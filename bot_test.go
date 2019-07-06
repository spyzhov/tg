package tg

import (
	"context"
	"fmt"
)

func ExampleBot_GetMe() {
	bot := New("TOKEN")
	user, _ := bot.GetMe(context.Background())
	fmt.Printf("%#v\n", user)
}
