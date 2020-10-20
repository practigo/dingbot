package dingbot_test

import (
	"fmt"
	"testing"

	"github.com/practigo/dingbot"
)

func ExampleWebhook() {
	var testToken = "YOUR_TOKEN_HERE"

	bot := dingbot.NewWebhook(testToken)

	var testSecret = "" // if you have secret set, fill here
	bot.WithSecret(testSecret)

	msg := dingbot.DingMessage{
		Msgtype: dingbot.MsgTypeText,
		Text: &dingbot.TextMsg{
			Content: "hello world",
		},
	}
	fmt.Println(bot.Send(&msg))
	// Output: Error 300001: token is not exist
}

// TestSign refs python script
// #python 2.7 from https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq/uKPlK
// import time
// import hmac
// import hashlib
// import base64
// import urllib
// import struct
//
// timestamp = "1601187240000"
// secret = 'this is a secret'
// secret_enc = bytes(secret).encode('utf-8')
// string_to_sign = '{}\n{}'.format(timestamp, secret)
// string_to_sign_enc = bytes(string_to_sign).encode('utf-8')
// hmac_code = hmac.new(secret_enc, string_to_sign_enc, digestmod=hashlib.sha256).digest()
// sign = base64.b64encode(hmac_code)
// print(sign)
func TestSign(t *testing.T) {
	sign := dingbot.Sign("this is a secret", "1601187240000")
	if sign != "7gmCSzcAc2XzfB14K9+cRSM1hqBng7kT+N5k61qXXz0=" {
		t.Error("wrong sign")
	}
}
