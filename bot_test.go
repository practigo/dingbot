package dingbot_test

import (
	"fmt"

	"github.com/practigo/dingbot"
)

func ExampleWebhook() {
	var testToken = "YOUR_TOKEN_HERE"

	bot := dingbot.NewWebhook(testToken)
	msg := dingbot.DingMessage{
		Msgtype: dingbot.MsgTypeText,
		Text: &dingbot.TextMsg{
			Content: "hello world",
		},
	}
	fmt.Println(bot.Send(&msg))
	// Output: Error 300001: token is not exist
}
