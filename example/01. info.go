package main

import (
	"context"
	"fmt"
	"github.com/spyzhov/tg"
	"log"
	"os"
)

func main() {
	bot := tg.New(os.Getenv("TOKEN"))
	bot.Log = log.Printf
	user, err := bot.GetMe(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", user)
}
