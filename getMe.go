package telego

import (
	"context"
	"github.com/spyzhov/telego/tobject"
)

const getMeAction Action = "getMe"

// A simple method for testing your bot's auth token. Requires no parameters.
// Returns basic information about the bot in form of a User object.
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe(ctx context.Context) (*tobject.User, error) {
	result := new(tobject.User)
	return result, b.postResult(ctx, getMeAction, nil, &result)
}
