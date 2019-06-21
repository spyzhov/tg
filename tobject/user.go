package tobject

/**
User
This object represents a Telegram user or bot.

Field			Type	Description
id				Integer	Unique identifier for this user or bot
is_bot			Boolean	True, if this user is a bot
first_name		String	User‘s or bot’s first name
last_name		String	Optional. User‘s or bot’s last name
username		String	Optional. User‘s or bot’s username
language_code	String	Optional. IETF language tag of the user's language
*/
// https://core.telegram.org/bots/api#user
type User struct {
	// Unique identifier for this user or bot
	ID int `json:"id"`
	// True, if this user is a bot
	IsBot bool `json:"is_bot"`
	// User‘s or bot’s first name
	FirstName string `json:"first_name"`
	// Optional. User‘s or bot’s last name
	LastName string `json:"last_name,omitempty"`
	// Optional. User‘s or bot’s username
	Username string `json:"username,omitempty"`
	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code"`
}
