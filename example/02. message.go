package main

import (
	"context"
	"fmt"
	. "github.com/spyzhov/telego"
	"os"
)

const CHAT_ID = -1

func main() {
	bot := New(os.Getenv("TOKEN"))
	user, err := bot.GetMe(context.Background())
	if err != nil {
		panic(err)
	}
	result, err := bot.SendMessage(context.Background(), &SendMessageRequest{
		ChatId: CHAT_ID,
		Text:   "Hello, my name is " + user.Username,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)
}
