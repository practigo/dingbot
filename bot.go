// Package dingbot provides Bot related APIs for dingtalk open platform.
package dingbot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// constants
const (
	WebhookURL = "https://oapi.dingtalk.com/robot/send"
	DingCodeOk = 0
)

// DingResponse is the response for sending msg.
type DingResponse struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
}

func (d DingResponse) Error() string {
	return fmt.Sprintf("Error %d: %s", d.Errcode, d.Errmsg)
}

// Sign signs the timestamp with a secret.
func Sign(secret, timestamp string) string {
	toSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(toSign)) // sha1 Write() returns no error
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Sender sends messages of various types to a Dingtalk group.
type Sender interface {
	// Send sends a DingMessage.
	Send(*DingMessage) error
}

// Webhook sends the message via webhook API.
// See https://open-doc.dingtalk.com/microapp/serverapi2/qf2nxq.
type Webhook struct {
	token string
	cl    *http.Client
	// new security feature
	secret string
}

// WithSecret sets the webhook serect for passing the security check.
func (b *Webhook) WithSecret(s string) {
	b.secret = s
}

// Send sends a DingMessage. The underlying HTTP client
// has a 5s timeout and will not retry when fail.
func (b *Webhook) Send(msg *DingMessage) (err error) {
	data, err := json.Marshal(*msg)
	if err != nil {
		return errors.Wrap(err, "form message")
	}

	v := url.Values{}
	v.Add("access_token", b.token)

	if b.secret != "" {
		t := fmt.Sprintf("%d", time.Now().Unix()*1000)
		signature := Sign(b.secret, t)
		v.Add("timestamp", t)
		v.Add("sign", signature)
	}

	uri := fmt.Sprintf("%s?%s", WebhookURL, v.Encode())
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return errors.Wrap(err, "form request")
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := b.cl.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return errors.Wrap(err, "do request")
	}

	// status := resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read reponse")
	}

	var ret DingResponse
	if err = json.Unmarshal(body, &ret); err != nil {
		return errors.Wrap(err, "parse resp")
	}

	if ret.Errcode != DingCodeOk {
		return ret
	}

	return nil
}

// NewWebhook returns a Webhook with the provided token.
func NewWebhook(token string) *Webhook {
	b := &Webhook{
		token: token,
		cl: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	return b
}
