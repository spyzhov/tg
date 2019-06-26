# Telegram BOT

[![GoDoc](https://godoc.org/github.com/spyzhov/telego?status.svg)](https://godoc.org/github.com/spyzhov/telego)

Implement golang interface for [Telegram Bot API](https://core.telegram.org/bots/api).

Implemented all methods up to: May 31, 2019 Bot API 4.3

# API

All documentation about Bot.API you can find at [API.md](API.md)

# Examples

All examples are at: [example](example/) dir.

# Generator

Generates API using [`generator.py`](generator.py): [requirements](requirements.txt) listed at file, Python3.7 required.

## Simple use

```go
package main

import (
	"context"
	"fmt"
	. "github.com/spyzhov/telego"
	"os"
	"strconv"
)

func main() {
	bot := New(os.Getenv("TOKEN"))
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

```