package main

import (
	"context"
	"fmt"
	. "github.com/spyzhov/telego"
	"log"
	"os"
	"strconv"
)

func main() {
	bot := New(os.Getenv("TOKEN"))
	bot.Log = log.Printf
	chatId, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		panic(err)
	}
	user, err := bot.GetMe(context.Background())
	if err != nil {
		panic(err)
	}
	result, err := bot.SendMessage(context.Background(), &SendMessageRequest{
		ChatId: chatId,
		Text:   "Hello, my name is " + user.Username,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)
	if result.From != nil {
		fmt.Printf("From: %#v\n", result.From)
	}
	fmt.Printf("Chat: %#v\n", result.Chat)
}
