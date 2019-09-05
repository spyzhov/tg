package main

import (
	"context"
	"fmt"
	. "github.com/spyzhov/tg"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	bot := New(os.Getenv("TOKEN"))
	bot.Log = log.Printf
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*600)

	go func() {
		timer := time.NewTicker(time.Second)
		last := 0
		defer timer.Stop()
		for {
			<-timer.C

			updates, err := bot.GetUpdates(ctx, &GetUpdatesRequest{
				Offset:         last + 1,
				Limit:          10,
				AllowedUpdates: []string{"message"},
			})
			check(err)

			for i, update := range updates {
				fmt.Printf("update [%03d]: %#v\n", i, update.Message)
				last = update.UpdateId
				if update.Message != nil {
					if update.Message.Sticker != nil {
						_, err = bot.SendMessage(ctx, &SendMessageRequest{
							ChatId:    update.Message.Chat.Id,
							Text:      "You send me a sticker: _" + update.Message.Sticker.FileId + "_ from *" + update.Message.Sticker.SetName + "*",
							ParseMode: "Markdown",
						})
					} else {
						_, err = bot.SendMessage(ctx, &SendMessageRequest{
							ChatId:    update.Message.Chat.Id,
							Text:      "You wrote me: _" + update.Message.Text + "_",
							ParseMode: "Markdown",
						})
					}
					check(err)
				}
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:
		cancel()
		fmt.Println("exiting...")
		time.Sleep(time.Second)
	case <-ctx.Done():
		fmt.Println("Done!")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
