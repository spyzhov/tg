package main

import (
	"context"
	"fmt"
	"github.com/spyzhov/telego"
	"os"
)

func main() {
	bot := telego.New(os.Getenv("TOKEN"))
	user, err := bot.GetMe(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", user)
}
