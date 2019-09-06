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
				AllowedUpdates: []string{"message", "callback_query"},
			})
			check(err)

			for i, update := range updates {
				fmt.Printf("update [%03d]: %#v\n", i, update.Message)
				last = update.UpdateId
				if update.Message != nil {
					_, err = bot.SendMessage(ctx, &SendMessageRequest{
						ChatId:    update.Message.Chat.Id,
						Text:      "Push the button!",
						ParseMode: "Markdown",
						ReplyMarkup: &InlineKeyboardMarkup{
							InlineKeyboard: [][]*InlineKeyboardButton{
								{
									{Text: "Push!", CallbackData: "Push!"},
								},
							},
						},
					})
					check(err)
				} else if update.CallbackQuery != nil {
					var markup *InlineKeyboardMarkup
					if update.CallbackQuery.Data == "Push!" {
						markup = &InlineKeyboardMarkup{
							InlineKeyboard: [][]*InlineKeyboardButton{
								{
									{Text: "Push the button!", CallbackData: "Push the button!"},
								},
							},
						}
					}
					_, err = bot.EditMessageText(ctx, &EditMessageTextRequest{
						Text:        update.CallbackQuery.Message.Text + " " + update.CallbackQuery.Data,
						MessageId:   update.CallbackQuery.Message.MessageId,
						ChatId:      update.CallbackQuery.Message.Chat.Id,
						ReplyMarkup: markup,
					})
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
