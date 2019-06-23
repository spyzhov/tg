package telego

// Update
// https://core.telegram.org/bots/api#update
//
// This object represents an incoming update.At most one of the optional parameters can be present in
// any given update.
//
//    Field                Type                 Description
//    update_id            Integer              The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
//    message              Message              Optional. New incoming message of any kind — text, photo, sticker, etc.
//    edited_message       Message              Optional. New version of a message that is known to the bot and was edited
//    channel_post         Message              Optional. New incoming channel post of any kind — text, photo, sticker, etc.
//    edited_channel_post  Message              Optional. New version of a channel post that is known to the bot and was edited
//    inline_query         InlineQuery          Optional. New incoming inline query
//    chosen_inline_result ChosenInlineResult   Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
//    callback_query       CallbackQuery        Optional. New incoming callback query
//    shipping_query       ShippingQuery        Optional. New incoming shipping query. Only for invoices with flexible price
//    pre_checkout_query   PreCheckoutQuery     Optional. New incoming pre-checkout query. Contains full information about checkout
//    poll                 Poll                 Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
//
type Update struct {
	// The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateId int `json:"update_id"`
	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`
	// Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`
	// Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`
	// Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	// Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	// Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`
}

// WebhookInfo
// https://core.telegram.org/bots/api#webhookinfo
//
// Contains information about the current status of a webhook.
//
//    Field                Type                 Description
//    url                  String               Webhook URL, may be empty if webhook is not set up
//    has_custom_certificate Boolean              True, if a custom certificate was provided for webhook certificate checks
//    pending_update_count Integer              Number of updates awaiting delivery
//    last_error_date      Integer              Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
//    last_error_message   String               Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
//    max_connections      Integer              Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
//    allowed_updates      Array of String      Optional. A list of update types the bot is subscribed to. Defaults to all update types
//
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	Url string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`
	// Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorDate int `json:"last_error_date,omitempty"`
	// Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections int `json:"max_connections,omitempty"`
	// Optional. A list of update types the bot is subscribed to. Defaults to all update types
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// User
// https://core.telegram.org/bots/api#user
//
// This object represents a Telegram user or bot.
//
//    Field                Type                 Description
//    id                   Integer              Unique identifier for this user or bot
//    is_bot               Boolean              True, if this user is a bot
//    first_name           String               User‘s or bot’s first name
//    last_name            String               Optional. User‘s or bot’s last name
//    username             String               Optional. User‘s or bot’s username
//    language_code        String               Optional. IETF language tag of the user's language
//
type User struct {
	// Unique identifier for this user or bot
	Id int `json:"id"`
	// True, if this user is a bot
	IsBot bool `json:"is_bot"`
	// User‘s or bot’s first name
	FirstName string `json:"first_name"`
	// Optional. User‘s or bot’s last name
	LastName string `json:"last_name,omitempty"`
	// Optional. User‘s or bot’s username
	Username string `json:"username,omitempty"`
	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`
}

// Chat
// https://core.telegram.org/bots/api#chat
//
// This object represents a chat.
//
//    Field                Type                 Description
//    id                   Integer              Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
//    type                 String               Type of chat, can be either “private”, “group”, “supergroup” or “channel”
//    title                String               Optional. Title, for supergroups, channels and group chats
//    username             String               Optional. Username, for private chats, supergroups and channels if available
//    first_name           String               Optional. First name of the other party in a private chat
//    last_name            String               Optional. Last name of the other party in a private chat
//    all_members_are_administrators Boolean              Optional. True if a group has ‘All Members Are Admins’ enabled.
//    photo                ChatPhoto            Optional. Chat photo. Returned only in getChat.
//    description          String               Optional. Description, for supergroups and channel chats. Returned only in getChat.
//    invite_link          String               Optional. Chat invite link, for supergroups and channel chats. Each administrator in a chat generates their own invite links, so the bot must first generate the link using exportChatInviteLink. Returned only in getChat.
//    pinned_message       Message              Optional. Pinned message, for groups, supergroups and channels. Returned only in getChat.
//    sticker_set_name     String               Optional. For supergroups, name of group sticker set. Returned only in getChat.
//    can_set_sticker_set  Boolean              Optional. True, if the bot can change the group sticker set. Returned only in getChat.
//
type Chat struct {
	// Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Id int `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`
	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`
	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`
	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`
	// Optional. True if a group has ‘All Members Are Admins’ enabled.
	AllMembersAreAdministrators bool `json:"all_members_are_administrators,omitempty"`
	// Optional. Chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo,omitempty"`
	// Optional. Description, for supergroups and channel chats. Returned only in getChat.
	Description string `json:"description,omitempty"`
	// Optional. Chat invite link, for supergroups and channel chats. Each administrator in a chat generates their own invite links, so the bot must first generate the link using exportChatInviteLink. Returned only in getChat.
	InviteLink string `json:"invite_link,omitempty"`
	// Optional. Pinned message, for groups, supergroups and channels. Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. For supergroups, name of group sticker set. Returned only in getChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`
	// Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
}

// Message
// https://core.telegram.org/bots/api#message
//
// This object represents a message.
//
//    Field                Type                 Description
//    message_id           Integer              Unique message identifier inside this chat
//    from                 User                 Optional. Sender, empty for messages sent to channels
//    date                 Integer              Date the message was sent in Unix time
//    chat                 Chat                 Conversation the message belongs to
//    forward_from         User                 Optional. For forwarded messages, sender of the original message
//    forward_from_chat    Chat                 Optional. For messages forwarded from channels, information about the original channel
//    forward_from_message_id Integer              Optional. For messages forwarded from channels, identifier of the original message in the channel
//    forward_signature    String               Optional. For messages forwarded from channels, signature of the post author if present
//    forward_sender_name  String               Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
//    forward_date         Integer              Optional. For forwarded messages, date the original message was sent in Unix time
//    reply_to_message     Message              Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
//    edit_date            Integer              Optional. Date the message was last edited in Unix time
//    media_group_id       String               Optional. The unique identifier of a media message group this message belongs to
//    author_signature     String               Optional. Signature of the post author for messages in channels
//    text                 String               Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
//    entities             Array of MessageEntity Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
//    caption_entities     Array of MessageEntity Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
//    audio                Audio                Optional. Message is an audio file, information about the file
//    document             Document             Optional. Message is a general file, information about the file
//    animation            Animation            Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
//    game                 Game                 Optional. Message is a game, information about the game. More about games »
//    photo                Array of PhotoSize   Optional. Message is a photo, available sizes of the photo
//    sticker              Sticker              Optional. Message is a sticker, information about the sticker
//    video                Video                Optional. Message is a video, information about the video
//    voice                Voice                Optional. Message is a voice message, information about the file
//    video_note           VideoNote            Optional. Message is a video note, information about the video message
//    caption              String               Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
//    contact              Contact              Optional. Message is a shared contact, information about the contact
//    location             Location             Optional. Message is a shared location, information about the location
//    venue                Venue                Optional. Message is a venue, information about the venue
//    poll                 Poll                 Optional. Message is a native poll, information about the poll
//    new_chat_members     Array of User        Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
//    left_chat_member     User                 Optional. A member was removed from the group, information about them (this member may be the bot itself)
//    new_chat_title       String               Optional. A chat title was changed to this value
//    new_chat_photo       Array of PhotoSize   Optional. A chat photo was change to this value
//    delete_chat_photo    True                 Optional. Service message: the chat photo was deleted
//    group_chat_created   True                 Optional. Service message: the group has been created
//    supergroup_chat_created True                 Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
//    channel_chat_created True                 Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
//    migrate_to_chat_id   Integer              Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
//    migrate_from_chat_id Integer              Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
//    pinned_message       Message              Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
//    invoice              Invoice              Optional. Message is an invoice for a payment, information about the invoice. More about payments »
//    successful_payment   SuccessfulPayment    Optional. Message is a service message about a successful payment, information about the payment. More about payments »
//    connected_website    String               Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
//    passport_data        PassportData         Optional. Telegram Passport data
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
//
type Message struct {
	// Unique message identifier inside this chat
	MessageId int `json:"message_id"`
	// Optional. Sender, empty for messages sent to channels
	From *User `json:"from,omitempty"`
	// Date the message was sent in Unix time
	Date int `json:"date"`
	// Conversation the message belongs to
	Chat *Chat `json:"chat"`
	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Optional. For messages forwarded from channels, information about the original channel
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardFromMessageId int `json:"forward_from_message_id,omitempty"`
	// Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSignature string `json:"forward_signature,omitempty"`
	// Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int `json:"forward_date,omitempty"`
	// Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date,omitempty"`
	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupId string `json:"media_group_id,omitempty"`
	// Optional. Signature of the post author for messages in channels
	AuthorSignature string `json:"author_signature,omitempty"`
	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text,omitempty"`
	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []*MessageEntity `json:"entities,omitempty"`
	// Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`
	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`
	// Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`
	// Optional. Message is a game, information about the game. More about games »
	Game *Game `json:"game,omitempty"`
	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`
	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`
	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`
	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`
	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`
	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`
	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`
	// Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	NewChatMembers []*User `json:"new_chat_members,omitempty"`
	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// Optional. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`
	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId int `json:"migrate_to_chat_id,omitempty"`
	// Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId int `json:"migrate_from_chat_id,omitempty"`
	// Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	Invoice *Invoice `json:"invoice,omitempty"`
	// Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	// Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`
	// Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// MessageEntity
// https://core.telegram.org/bots/api#messageentity
//
// This object represents one special entity in a text message. For example, hashtags, usernames,
// URLs, etc.
//
//    Field                Type                 Description
//    type                 String               Type of the entity. Can be mention (@username), hashtag, cashtag, bot_command, url, email, phone_number, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
//    offset               Integer              Offset in UTF-16 code units to the start of the entity
//    length               Integer              Length of the entity in UTF-16 code units
//    url                  String               Optional. For “text_link” only, url that will be opened after user taps on the text
//    user                 User                 Optional. For “text_mention” only, the mentioned user
//
type MessageEntity struct {
	// Type of the entity. Can be mention (@username), hashtag, cashtag, bot_command, url, email, phone_number, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// Optional. For “text_link” only, url that will be opened after user taps on the text
	Url string `json:"url,omitempty"`
	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`
}

// PhotoSize
// https://core.telegram.org/bots/api#photosize
//
// This object represents one size of a photo or a file / sticker thumbnail.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    width                Integer              Photo width
//    height               Integer              Photo height
//    file_size            Integer              Optional. File size
//
type PhotoSize struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Photo width
	Width int `json:"width"`
	// Photo height
	Height int `json:"height"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Audio
// https://core.telegram.org/bots/api#audio
//
// This object represents an audio file to be treated as music by the Telegram clients.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    duration             Integer              Duration of the audio in seconds as defined by sender
//    performer            String               Optional. Performer of the audio as defined by sender or by audio tags
//    title                String               Optional. Title of the audio as defined by sender or by audio tags
//    mime_type            String               Optional. MIME type of the file as defined by sender
//    file_size            Integer              Optional. File size
//    thumb                PhotoSize            Optional. Thumbnail of the album cover to which the music file belongs
//
type Audio struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Document
// https://core.telegram.org/bots/api#document
//
// This object represents a general file (as opposed to photos, voice messages and audio files).
//
//    Field                Type                 Description
//    file_id              String               Unique file identifier
//    thumb                PhotoSize            Optional. Document thumbnail as defined by sender
//    file_name            String               Optional. Original filename as defined by sender
//    mime_type            String               Optional. MIME type of the file as defined by sender
//    file_size            Integer              Optional. File size
//
type Document struct {
	// Unique file identifier
	FileId string `json:"file_id"`
	// Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Video
// https://core.telegram.org/bots/api#video
//
// This object represents a video file.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    width                Integer              Video width as defined by sender
//    height               Integer              Video height as defined by sender
//    duration             Integer              Duration of the video in seconds as defined by sender
//    thumb                PhotoSize            Optional. Video thumbnail
//    mime_type            String               Optional. Mime type of a file as defined by sender
//    file_size            Integer              Optional. File size
//
type Video struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Mime type of a file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Animation
// https://core.telegram.org/bots/api#animation
//
// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
//
//    Field                Type                 Description
//    file_id              String               Unique file identifier
//    width                Integer              Video width as defined by sender
//    height               Integer              Video height as defined by sender
//    duration             Integer              Duration of the video in seconds as defined by sender
//    thumb                PhotoSize            Optional. Animation thumbnail as defined by sender
//    file_name            String               Optional. Original animation filename as defined by sender
//    mime_type            String               Optional. MIME type of the file as defined by sender
//    file_size            Integer              Optional. File size
//
type Animation struct {
	// Unique file identifier
	FileId string `json:"file_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Animation thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Voice
// https://core.telegram.org/bots/api#voice
//
// This object represents a voice note.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    duration             Integer              Duration of the audio in seconds as defined by sender
//    mime_type            String               Optional. MIME type of the file as defined by sender
//    file_size            Integer              Optional. File size
//
type Voice struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// VideoNote
// https://core.telegram.org/bots/api#videonote
//
// This object represents a video message (available in Telegram apps as of v.4.0).
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    length               Integer              Video width and height (diameter of the video message) as defined by sender
//    duration             Integer              Duration of the video in seconds as defined by sender
//    thumb                PhotoSize            Optional. Video thumbnail
//    file_size            Integer              Optional. File size
//
type VideoNote struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// Contact
// https://core.telegram.org/bots/api#contact
//
// This object represents a phone contact.
//
//    Field                Type                 Description
//    phone_number         String               Contact's phone number
//    first_name           String               Contact's first name
//    last_name            String               Optional. Contact's last name
//    user_id              Integer              Optional. Contact's user identifier in Telegram
//    vcard                String               Optional. Additional data about the contact in the form of a vCard
//
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Contact's user identifier in Telegram
	UserId int `json:"user_id,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard,omitempty"`
}

// Location
// https://core.telegram.org/bots/api#location
//
// This object represents a point on the map.
//
//    Field                Type                 Description
//    longitude            Float                Longitude as defined by sender
//    latitude             Float                Latitude as defined by sender
//
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
}

// Venue
// https://core.telegram.org/bots/api#venue
//
// This object represents a venue.
//
//    Field                Type                 Description
//    location             Location             Venue location
//    title                String               Name of the venue
//    address              String               Address of the venue
//    foursquare_id        String               Optional. Foursquare identifier of the venue
//    foursquare_type      String               Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
//
type Venue struct {
	// Venue location
	Location *Location `json:"location"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue
	FoursquareId string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
}

// PollOption
// https://core.telegram.org/bots/api#polloption
//
// This object contains information about one answer option in a poll.
//
//    Field                Type                 Description
//    text                 String               Option text, 1-100 characters
//    voter_count          Integer              Number of users that voted for this option
//
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// Poll
// https://core.telegram.org/bots/api#poll
//
// This object contains information about a poll.
//
//    Field                Type                 Description
//    id                   String               Unique poll identifier
//    question             String               Poll question, 1-255 characters
//    options              Array of PollOption  List of poll options
//    is_closed            Boolean              True, if the poll is closed
//
type Poll struct {
	// Unique poll identifier
	Id string `json:"id"`
	// Poll question, 1-255 characters
	Question string `json:"question"`
	// List of poll options
	Options []*PollOption `json:"options"`
	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`
}

// UserProfilePhotos
// https://core.telegram.org/bots/api#userprofilephotos
//
// This object represent a user's profile pictures.
//
//    Field                Type                 Description
//    total_count          Integer              Total number of profile pictures the target user has
//    photos               Array of Array of PhotoSize Requested profile pictures (in up to 4 sizes each)
//
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}

// File
// https://core.telegram.org/bots/api#file
//
// This object represents a file ready to be downloaded. The file can be downloaded via the link
// https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for
// at least 1 hour. When the link expires, a new one can be requested by calling getFile.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    file_size            Integer              Optional. File size, if known
//    file_path            String               Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
//
type File struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Optional. File size, if known
	FileSize int `json:"file_size,omitempty"`
	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup
// https://core.telegram.org/bots/api#replykeyboardmarkup
//
// This object represents a custom keyboard with reply options (see Introduction to bots for details
// and examples).
//
//    Field                Type                 Description
//    keyboard             Array of Array of KeyboardButton Array of button rows, each represented by an Array of KeyboardButton objects
//    resize_keyboard      Boolean              Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
//    one_time_keyboard    Boolean              Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
//    selective            Boolean              Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
//
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]*KeyboardButton `json:"keyboard"`
	// Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// KeyboardButton
// https://core.telegram.org/bots/api#keyboardbutton
//
// This object represents one button of the reply keyboard. For simple text buttons String can be
// used instead of this object to specify text of the button. Optional fields are mutually exclusive.
//
//    Field                Type                 Description
//    text                 String               Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
//    request_contact      Boolean              Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
//    request_location     Boolean              Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
//
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	Text string `json:"text"`
	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`
	// Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`
}

// ReplyKeyboardRemove
// https://core.telegram.org/bots/api#replykeyboardremove
//
// Upon receiving a message with this object, Telegram clients will remove the current custom
// keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new
// keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately
// after the user presses a button (see ReplyKeyboardMarkup).
//
//    Field                Type                 Description
//    remove_keyboard      True                 Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
//    selective            Boolean              Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
//
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
//
// This object represents an inline keyboard that appears right next to the message it belongs to.
//
//    Field                Type                 Description
//    inline_keyboard      Array of Array of InlineKeyboardButton Array of button rows, each represented by an Array of InlineKeyboardButton objects
//
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton
// https://core.telegram.org/bots/api#inlinekeyboardbutton
//
// This object represents one button of an inline keyboard. You must use exactly one of the optional
// fields.
//
//    Field                Type                 Description
//    text                 String               Label text on the button
//    url                  String               Optional. HTTP or tg:// url to be opened when button is pressed
//    login_url            LoginUrl             Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
//    callback_data        String               Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
//    switch_inline_query  String               Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
//    switch_inline_query_current_chat String               Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
//    callback_game        CallbackGame         Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
//    pay                  Boolean              Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row.
//
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`
	// Optional. HTTP or tg:// url to be opened when button is pressed
	Url string `json:"url,omitempty"`
	// Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	LoginUrl *LoginUrl `json:"login_url,omitempty"`
	// Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`
	// Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
	// Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	// Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	// Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row.
	Pay bool `json:"pay,omitempty"`
}

// LoginUrl
// https://core.telegram.org/bots/api#loginurl
//
// This object represents a parameter of the inline keyboard button used to automatically authorize a
// user. Serves as a great replacement for the Telegram Login Widget when the user is coming from
// Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
//
//    Field                Type                 Description
//    url                  String               An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
//    forward_text         String               Optional. New text of the button in forwarded messages.
//    bot_username         String               Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
//    request_write_access Boolean              Optional. Pass True to request the permission for your bot to send messages to the user.
//
type LoginUrl struct {
	// An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	Url string `json:"url"`
	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`
	// Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	BotUsername string `json:"bot_username,omitempty"`
	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// CallbackQuery
// https://core.telegram.org/bots/api#callbackquery
//
// This object represents an incoming callback query from a callback button in an inline keyboard. If
// the button that originated the query was attached to a message sent by the bot, the field message
// will be present. If the button was attached to a message sent via the bot (in inline mode), the field
// inline_message_id will be present. Exactly one of the fields data or game_short_name will be
// present.
//
//    Field                Type                 Description
//    id                   String               Unique identifier for this query
//    from                 User                 Sender
//    message              Message              Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
//    inline_message_id    String               Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
//    chat_instance        String               Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
//    data                 String               Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
//    game_short_name      String               Optional. Short name of a Game to be returned, serves as the unique identifier for the game
//
type CallbackQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`
	// Sender
	From *User `json:"from"`
	// Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	Message *Message `json:"message,omitempty"`
	// Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`
	// Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data,omitempty"`
	// Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}

// ForceReply
// https://core.telegram.org/bots/api#forcereply
//
// Upon receiving a message with this object, Telegram clients will display a reply interface to the
// user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely
// useful if you want to create user-friendly step-by-step interfaces without having to sacrifice
// privacy mode.
//
//    Field                Type                 Description
//    force_reply          True                 Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
//    selective            Boolean              Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
//
type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	ForceReply bool `json:"force_reply"`
	// Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// ChatPhoto
// https://core.telegram.org/bots/api#chatphoto
//
// This object represents a chat photo.
//
//    Field                Type                 Description
//    small_file_id        String               Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
//    big_file_id          String               Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
//
type ChatPhoto struct {
	// Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
	SmallFileId string `json:"small_file_id"`
	// Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
	BigFileId string `json:"big_file_id"`
}

// ChatMember
// https://core.telegram.org/bots/api#chatmember
//
// This object contains information about one member of a chat.
//
//    Field                Type                 Description
//    user                 User                 Information about the user
//    status               String               The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
//    until_date           Integer              Optional. Restricted and kicked only. Date when restrictions will be lifted for this user, unix time
//    can_be_edited        Boolean              Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
//    can_change_info      Boolean              Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings
//    can_post_messages    Boolean              Optional. Administrators only. True, if the administrator can post in the channel, channels only
//    can_edit_messages    Boolean              Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only
//    can_delete_messages  Boolean              Optional. Administrators only. True, if the administrator can delete messages of other users
//    can_invite_users     Boolean              Optional. Administrators only. True, if the administrator can invite new users to the chat
//    can_restrict_members Boolean              Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
//    can_pin_messages     Boolean              Optional. Administrators only. True, if the administrator can pin messages, groups and supergroups only
//    can_promote_members  Boolean              Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
//    is_member            Boolean              Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
//    can_send_messages    Boolean              Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues
//    can_send_media_messages Boolean              Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
//    can_send_other_messages Boolean              Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
//    can_add_web_page_previews Boolean              Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
//
type ChatMember struct {
	// Information about the user
	User *User `json:"user"`
	// The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	Status string `json:"status"`
	// Optional. Restricted and kicked only. Date when restrictions will be lifted for this user, unix time
	UntilDate int `json:"until_date,omitempty"`
	// Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited,omitempty"`
	// Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// Optional. Administrators only. True, if the administrator can post in the channel, channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// Optional. Administrators only. True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// Optional. Administrators only. True, if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// Optional. Administrators only. True, if the administrator can pin messages, groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member,omitempty"`
	// Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
}

// ResponseParameters
// https://core.telegram.org/bots/api#responseparameters
//
// Contains information about why a request was unsuccessful.
//
//    Field                Type                 Description
//    migrate_to_chat_id   Integer              Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
//    retry_after          Integer              Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
//
type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId int `json:"migrate_to_chat_id,omitempty"`
	// Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter int `json:"retry_after,omitempty"`
}

// InputMedia
// https://core.telegram.org/bots/api#inputmedia
//
// This object represents the content of a media message to be sent. It should be one of
//
type InputMedia struct {
}

// InputMediaPhoto
// https://core.telegram.org/bots/api#inputmediaphoto
//
// Represents a photo to be sent.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be photo
//    media                String               File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
//    caption              String               Optional. Caption of the photo to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//
type InputMediaPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
}

// InputMediaVideo
// https://core.telegram.org/bots/api#inputmediavideo
//
// Represents a video to be sent.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be video
//    media                String               File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Caption of the video to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    width                Integer              Optional. Video width
//    height               Integer              Optional. Video height
//    duration             Integer              Optional. Video duration
//    supports_streaming   Boolean              Optional. Pass True, if the uploaded video is suitable for streaming
//
type InputMediaVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Caption of the video to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Video width
	Width int `json:"width,omitempty"`
	// Optional. Video height
	Height int `json:"height,omitempty"`
	// Optional. Video duration
	Duration int `json:"duration,omitempty"`
	// Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

// InputMediaAnimation
// https://core.telegram.org/bots/api#inputmediaanimation
//
// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be animation
//    media                String               File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Caption of the animation to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    width                Integer              Optional. Animation width
//    height               Integer              Optional. Animation height
//    duration             Integer              Optional. Animation duration
//
type InputMediaAnimation struct {
	// Type of the result, must be animation
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Caption of the animation to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Animation width
	Width int `json:"width,omitempty"`
	// Optional. Animation height
	Height int `json:"height,omitempty"`
	// Optional. Animation duration
	Duration int `json:"duration,omitempty"`
}

// InputMediaAudio
// https://core.telegram.org/bots/api#inputmediaaudio
//
// Represents an audio file to be treated as music to be sent.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be audio
//    media                String               File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Caption of the audio to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    duration             Integer              Optional. Duration of the audio in seconds
//    performer            String               Optional. Performer of the audio
//    title                String               Optional. Title of the audio
//
type InputMediaAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Caption of the audio to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Performer of the audio
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio
	Title string `json:"title,omitempty"`
}

// InputMediaDocument
// https://core.telegram.org/bots/api#inputmediadocument
//
// Represents a general file to be sent.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be document
//    media                String               File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Caption of the document to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//
type InputMediaDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Caption of the document to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
}

// InputFile
// https://core.telegram.org/bots/api#inputfile
//
// This object represents the contents of a file to be uploaded. Must be posted using
// multipart/form-data in the usual way that files are uploaded via the browser.
//
type InputFile struct {
}

// Sticker
// https://core.telegram.org/bots/api#sticker
//
// This object represents a sticker.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    width                Integer              Sticker width
//    height               Integer              Sticker height
//    thumb                PhotoSize            Optional. Sticker thumbnail in the .webp or .jpg format
//    emoji                String               Optional. Emoji associated with the sticker
//    set_name             String               Optional. Name of the sticker set to which the sticker belongs
//    mask_position        MaskPosition         Optional. For mask stickers, the position where the mask should be placed
//    file_size            Integer              Optional. File size
//
type Sticker struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Sticker width
	Width int `json:"width"`
	// Sticker height
	Height int `json:"height"`
	// Optional. Sticker thumbnail in the .webp or .jpg format
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`
	// Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`
	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// StickerSet
// https://core.telegram.org/bots/api#stickerset
//
// This object represents a sticker set.
//
//    Field                Type                 Description
//    name                 String               Sticker set name
//    title                String               Sticker set title
//    contains_masks       Boolean              True, if the sticker set contains masks
//    stickers             Array of Sticker     List of all set stickers
//
type StickerSet struct {
	// Sticker set name
	Name string `json:"name"`
	// Sticker set title
	Title string `json:"title"`
	// True, if the sticker set contains masks
	ContainsMasks bool `json:"contains_masks"`
	// List of all set stickers
	Stickers []*Sticker `json:"stickers"`
}

// MaskPosition
// https://core.telegram.org/bots/api#maskposition
//
// This object describes the position on faces where a mask should be placed by default.
//
//    Field                Type                 Description
//    point                String               The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
//    x_shift              Float number         Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
//    y_shift              Float number         Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
//    scale                Float number         Mask scaling coefficient. For example, 2.0 means double size.
//
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float64 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	YShift float64 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

// InlineQuery
// https://core.telegram.org/bots/api#inlinequery
//
// This object represents an incoming inline query. When the user sends an empty query, your bot
// could return some default or trending results.
//
//    Field                Type                 Description
//    id                   String               Unique identifier for this query
//    from                 User                 Sender
//    location             Location             Optional. Sender location, only for bots that request user location
//    query                String               Text of the query (up to 512 characters)
//    offset               String               Offset of the results to be returned, can be controlled by the bot
//
type InlineQuery struct {
	// Unique identifier for this query
	Id string `json:"id"`
	// Sender
	From *User `json:"from"`
	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
	// Text of the query (up to 512 characters)
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
}

// InlineQueryResult
// https://core.telegram.org/bots/api#inlinequeryresult
//
// This object represents one result of an inline query. Telegram clients currently support results
// of the following 20 types:
//
type InlineQueryResult struct {
}

// InlineQueryResultArticle
// https://core.telegram.org/bots/api#inlinequeryresultarticle
//
// Represents a link to an article or web page.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be article
//    id                   String               Unique identifier for this result, 1-64 Bytes
//    title                String               Title of the result
//    input_message_content InputMessageContent  Content of the message to be sent
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    url                  String               Optional. URL of the result
//    hide_url             Boolean              Optional. Pass True, if you don't want the URL to be shown in the message
//    description          String               Optional. Short description of the result
//    thumb_url            String               Optional. Url of the thumbnail for the result
//    thumb_width          Integer              Optional. Thumbnail width
//    thumb_height         Integer              Optional. Thumbnail height
//
type InlineQueryResultArticle struct {
	// Type of the result, must be article
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Title of the result
	Title string `json:"title"`
	// Content of the message to be sent
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. URL of the result
	Url string `json:"url,omitempty"`
	// Optional. Pass True, if you don't want the URL to be shown in the message
	HideUrl bool `json:"hide_url,omitempty"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto
// https://core.telegram.org/bots/api#inlinequeryresultphoto
//
// Represents a link to a photo. By default, this photo will be sent by the user with optional
// caption. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the photo.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be photo
//    id                   String               Unique identifier for this result, 1-64 bytes
//    photo_url            String               A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
//    thumb_url            String               URL of the thumbnail for the photo
//    photo_width          Integer              Optional. Width of the photo
//    photo_height         Integer              Optional. Height of the photo
//    title                String               Optional. Title for the result
//    description          String               Optional. Short description of the result
//    caption              String               Optional. Caption of the photo to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the photo
//
type InlineQueryResultPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	PhotoUrl string `json:"photo_url"`
	// URL of the thumbnail for the photo
	ThumbUrl string `json:"thumb_url"`
	// Optional. Width of the photo
	PhotoWidth int `json:"photo_width,omitempty"`
	// Optional. Height of the photo
	PhotoHeight int `json:"photo_height,omitempty"`
	// Optional. Title for the result
	Title string `json:"title,omitempty"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Caption of the photo to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif
// https://core.telegram.org/bots/api#inlinequeryresultgif
//
// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the
// user with optional caption. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the animation.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be gif
//    id                   String               Unique identifier for this result, 1-64 bytes
//    gif_url              String               A valid URL for the GIF file. File size must not exceed 1MB
//    gif_width            Integer              Optional. Width of the GIF
//    gif_height           Integer              Optional. Height of the GIF
//    gif_duration         Integer              Optional. Duration of the GIF
//    thumb_url            String               URL of the static thumbnail for the result (jpeg or gif)
//    title                String               Optional. Title for the result
//    caption              String               Optional. Caption of the GIF file to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the GIF animation
//
type InlineQueryResultGif struct {
	// Type of the result, must be gif
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the GIF file. File size must not exceed 1MB
	GifUrl string `json:"gif_url"`
	// Optional. Width of the GIF
	GifWidth int `json:"gif_width,omitempty"`
	// Optional. Height of the GIF
	GifHeight int `json:"gif_height,omitempty"`
	// Optional. Duration of the GIF
	GifDuration int `json:"gif_duration,omitempty"`
	// URL of the static thumbnail for the result (jpeg or gif)
	ThumbUrl string `json:"thumb_url"`
	// Optional. Title for the result
	Title string `json:"title,omitempty"`
	// Optional. Caption of the GIF file to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
//
// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this
// animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the animation.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be mpeg4_gif
//    id                   String               Unique identifier for this result, 1-64 bytes
//    mpeg4_url            String               A valid URL for the MP4 file. File size must not exceed 1MB
//    mpeg4_width          Integer              Optional. Video width
//    mpeg4_height         Integer              Optional. Video height
//    mpeg4_duration       Integer              Optional. Video duration
//    thumb_url            String               URL of the static thumbnail (jpeg or gif) for the result
//    title                String               Optional. Title for the result
//    caption              String               Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the video animation
//
type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4Url string `json:"mpeg4_url"`
	// Optional. Video width
	Mpeg4Width int `json:"mpeg4_width,omitempty"`
	// Optional. Video height
	Mpeg4Height int `json:"mpeg4_height,omitempty"`
	// Optional. Video duration
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`
	// URL of the static thumbnail (jpeg or gif) for the result
	ThumbUrl string `json:"thumb_url"`
	// Optional. Title for the result
	Title string `json:"title,omitempty"`
	// Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo
// https://core.telegram.org/bots/api#inlinequeryresultvideo
//
// Represents a link to a page containing an embedded video player or a video file. By default, this
// video file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the video.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be video
//    id                   String               Unique identifier for this result, 1-64 bytes
//    video_url            String               A valid URL for the embedded video player or video file
//    mime_type            String               Mime type of the content of video url, “text/html” or “video/mp4”
//    thumb_url            String               URL of the thumbnail (jpeg only) for the video
//    title                String               Title for the result
//    caption              String               Optional. Caption of the video to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    video_width          Integer              Optional. Video width
//    video_height         Integer              Optional. Video height
//    video_duration       Integer              Optional. Video duration in seconds
//    description          String               Optional. Short description of the result
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
//
type InlineQueryResultVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the embedded video player or video file
	VideoUrl string `json:"video_url"`
	// Mime type of the content of video url, “text/html” or “video/mp4”
	MimeType string `json:"mime_type"`
	// URL of the thumbnail (jpeg only) for the video
	ThumbUrl string `json:"thumb_url"`
	// Title for the result
	Title string `json:"title"`
	// Optional. Caption of the video to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Video width
	VideoWidth int `json:"video_width,omitempty"`
	// Optional. Video height
	VideoHeight int `json:"video_height,omitempty"`
	// Optional. Video duration in seconds
	VideoDuration int `json:"video_duration,omitempty"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultAudio
// https://core.telegram.org/bots/api#inlinequeryresultaudio
//
// Represents a link to an mp3 audio file. By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of
// the audio.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be audio
//    id                   String               Unique identifier for this result, 1-64 bytes
//    audio_url            String               A valid URL for the audio file
//    title                String               Title
//    caption              String               Optional. Caption, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    performer            String               Optional. Performer
//    audio_duration       Integer              Optional. Audio duration in seconds
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the audio
//
type InlineQueryResultAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the audio file
	AudioUrl string `json:"audio_url"`
	// Title
	Title string `json:"title"`
	// Optional. Caption, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Performer
	Performer string `json:"performer,omitempty"`
	// Optional. Audio duration in seconds
	AudioDuration int `json:"audio_duration,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice
// https://core.telegram.org/bots/api#inlinequeryresultvoice
//
// Represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this
// voice recording will be sent by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the the voice message.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be voice
//    id                   String               Unique identifier for this result, 1-64 bytes
//    voice_url            String               A valid URL for the voice recording
//    title                String               Recording title
//    caption              String               Optional. Caption, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    voice_duration       Integer              Optional. Recording duration in seconds
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the voice recording
//
type InlineQueryResultVoice struct {
	// Type of the result, must be voice
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid URL for the voice recording
	VoiceUrl string `json:"voice_url"`
	// Recording title
	Title string `json:"title"`
	// Optional. Caption, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Recording duration in seconds
	VoiceDuration int `json:"voice_duration,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the voice recording
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument
// https://core.telegram.org/bots/api#inlinequeryresultdocument
//
// Represents a link to a file. By default, this file will be sent by the user with an optional
// caption. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be document
//    id                   String               Unique identifier for this result, 1-64 bytes
//    title                String               Title for the result
//    caption              String               Optional. Caption of the document to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    document_url         String               A valid URL for the file
//    mime_type            String               Mime type of the content of the file, either “application/pdf” or “application/zip”
//    description          String               Optional. Short description of the result
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the file
//    thumb_url            String               Optional. URL of the thumbnail (jpeg only) for the file
//    thumb_width          Integer              Optional. Thumbnail width
//    thumb_height         Integer              Optional. Thumbnail height
//
type InlineQueryResultDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Title for the result
	Title string `json:"title"`
	// Optional. Caption of the document to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// A valid URL for the file
	DocumentUrl string `json:"document_url"`
	// Mime type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// Optional. URL of the thumbnail (jpeg only) for the file
	ThumbUrl string `json:"thumb_url,omitempty"`
	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultLocation
// https://core.telegram.org/bots/api#inlinequeryresultlocation
//
// Represents a location on a map. By default, the location will be sent by the user. Alternatively,
// you can use input_message_content to send a message with the specified content instead of the
// location.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be location
//    id                   String               Unique identifier for this result, 1-64 Bytes
//    latitude             Float number         Location latitude in degrees
//    longitude            Float number         Location longitude in degrees
//    title                String               Location title
//    live_period          Integer              Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the location
//    thumb_url            String               Optional. Url of the thumbnail for the result
//    thumb_width          Integer              Optional. Thumbnail width
//    thumb_height         Integer              Optional. Thumbnail height
//
type InlineQueryResultLocation struct {
	// Type of the result, must be location
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Location latitude in degrees
	Latitude float64 `json:"latitude"`
	// Location longitude in degrees
	Longitude float64 `json:"longitude"`
	// Location title
	Title string `json:"title"`
	// Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the location
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// Optional. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultVenue
// https://core.telegram.org/bots/api#inlinequeryresultvenue
//
// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the venue.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be venue
//    id                   String               Unique identifier for this result, 1-64 Bytes
//    latitude             Float                Latitude of the venue location in degrees
//    longitude            Float                Longitude of the venue location in degrees
//    title                String               Title of the venue
//    address              String               Address of the venue
//    foursquare_id        String               Optional. Foursquare identifier of the venue if known
//    foursquare_type      String               Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the venue
//    thumb_url            String               Optional. Url of the thumbnail for the result
//    thumb_width          Integer              Optional. Thumbnail width
//    thumb_height         Integer              Optional. Thumbnail height
//
type InlineQueryResultVenue struct {
	// Type of the result, must be venue
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Latitude of the venue location in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the venue location in degrees
	Longitude float64 `json:"longitude"`
	// Title of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue if known
	FoursquareId string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the venue
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// Optional. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultContact
// https://core.telegram.org/bots/api#inlinequeryresultcontact
//
// Represents a contact with a phone number. By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of
// the contact.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be contact
//    id                   String               Unique identifier for this result, 1-64 Bytes
//    phone_number         String               Contact's phone number
//    first_name           String               Contact's first name
//    last_name            String               Optional. Contact's last name
//    vcard                String               Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the contact
//    thumb_url            String               Optional. Url of the thumbnail for the result
//    thumb_width          Integer              Optional. Thumbnail width
//    thumb_height         Integer              Optional. Thumbnail height
//
type InlineQueryResultContact struct {
	// Type of the result, must be contact
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	Id string `json:"id"`
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the contact
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// Optional. Url of the thumbnail for the result
	ThumbUrl string `json:"thumb_url,omitempty"`
	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// InlineQueryResultGame
// https://core.telegram.org/bots/api#inlinequeryresultgame
//
// Represents a Game.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be game
//    id                   String               Unique identifier for this result, 1-64 bytes
//    game_short_name      String               Short name of the game
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//
type InlineQueryResultGame struct {
	// Type of the result, must be game
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Short name of the game
	GameShortName string `json:"game_short_name"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// InlineQueryResultCachedPhoto
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
//
// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent
// by the user with an optional caption. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the photo.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be photo
//    id                   String               Unique identifier for this result, 1-64 bytes
//    photo_file_id        String               A valid file identifier of the photo
//    title                String               Optional. Title for the result
//    description          String               Optional. Short description of the result
//    caption              String               Optional. Caption of the photo to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the photo
//
type InlineQueryResultCachedPhoto struct {
	// Type of the result, must be photo
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier of the photo
	PhotoFileId string `json:"photo_file_id"`
	// Optional. Title for the result
	Title string `json:"title,omitempty"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Caption of the photo to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
//
// Represents a link to an animated GIF file stored on the Telegram servers. By default, this
// animated GIF file will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with specified content instead of the animation.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be gif
//    id                   String               Unique identifier for this result, 1-64 bytes
//    gif_file_id          String               A valid file identifier for the GIF file
//    title                String               Optional. Title for the result
//    caption              String               Optional. Caption of the GIF file to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the GIF animation
//
type InlineQueryResultCachedGif struct {
	// Type of the result, must be gif
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the GIF file
	GifFileId string `json:"gif_file_id"`
	// Optional. Title for the result
	Title string `json:"title,omitempty"`
	// Optional. Caption of the GIF file to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
//
// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the
// Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content
// instead of the animation.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be mpeg4_gif
//    id                   String               Unique identifier for this result, 1-64 bytes
//    mpeg4_file_id        String               A valid file identifier for the MP4 file
//    title                String               Optional. Title for the result
//    caption              String               Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the video animation
//
type InlineQueryResultCachedMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the MP4 file
	Mpeg4FileId string `json:"mpeg4_file_id"`
	// Optional. Title for the result
	Title string `json:"title,omitempty"`
	// Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedSticker
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
//
// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be
// sent by the user. Alternatively, you can use input_message_content to send a message with the
// specified content instead of the sticker.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be sticker
//    id                   String               Unique identifier for this result, 1-64 bytes
//    sticker_file_id      String               A valid file identifier of the sticker
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the sticker
//
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be sticker
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier of the sticker
	StickerFileId string `json:"sticker_file_id"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the sticker
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
//
// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by
// the user with an optional caption. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the file.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be document
//    id                   String               Unique identifier for this result, 1-64 bytes
//    title                String               Title for the result
//    document_file_id     String               A valid file identifier for the file
//    description          String               Optional. Short description of the result
//    caption              String               Optional. Caption of the document to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the file
//
type InlineQueryResultCachedDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// Title for the result
	Title string `json:"title"`
	// A valid file identifier for the file
	DocumentFileId string `json:"document_file_id"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Caption of the document to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the file
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
//
// Represents a link to a video file stored on the Telegram servers. By default, this video file will
// be sent by the user with an optional caption. Alternatively, you can use input_message_content to
// send a message with the specified content instead of the video.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be video
//    id                   String               Unique identifier for this result, 1-64 bytes
//    video_file_id        String               A valid file identifier for the video file
//    title                String               Title for the result
//    description          String               Optional. Short description of the result
//    caption              String               Optional. Caption of the video to be sent, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the video
//
type InlineQueryResultCachedVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the video file
	VideoFileId string `json:"video_file_id"`
	// Title for the result
	Title string `json:"title"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Caption of the video to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the video
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
//
// Represents a link to a voice message stored on the Telegram servers. By default, this voice
// message will be sent by the user. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the voice message.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be voice
//    id                   String               Unique identifier for this result, 1-64 bytes
//    voice_file_id        String               A valid file identifier for the voice message
//    title                String               Voice message title
//    caption              String               Optional. Caption, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the voice message
//
type InlineQueryResultCachedVoice struct {
	// Type of the result, must be voice
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the voice message
	VoiceFileId string `json:"voice_file_id"`
	// Voice message title
	Title string `json:"title"`
	// Optional. Caption, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the voice message
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
//
// Represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file
// will be sent by the user. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the audio.
//
//    Field                Type                 Description
//    type                 String               Type of the result, must be audio
//    id                   String               Unique identifier for this result, 1-64 bytes
//    audio_file_id        String               A valid file identifier for the audio file
//    caption              String               Optional. Caption, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. Inline keyboard attached to the message
//    input_message_content InputMessageContent  Optional. Content of the message to be sent instead of the audio
//
type InlineQueryResultCachedAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes
	Id string `json:"id"`
	// A valid file identifier for the audio file
	AudioFileId string `json:"audio_file_id"`
	// Optional. Caption, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Content of the message to be sent instead of the audio
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

// InputMessageContent
// https://core.telegram.org/bots/api#inputmessagecontent
//
// This object represents the content of a message to be sent as a result of an inline query.
// Telegram clients currently support the following 4 types:
//
type InputMessageContent struct {
}

// InputTextMessageContent
// https://core.telegram.org/bots/api#inputtextmessagecontent
//
// Represents the content of a text message to be sent as the result of an inline query.
//
//    Field                Type                 Description
//    message_text         String               Text of the message to be sent, 1-4096 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
//    disable_web_page_preview Boolean              Optional. Disables link previews for links in the sent message
//
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// InputLocationMessageContent
// https://core.telegram.org/bots/api#inputlocationmessagecontent
//
// Represents the content of a location message to be sent as the result of an inline query.
//
//    Field                Type                 Description
//    latitude             Float                Latitude of the location in degrees
//    longitude            Float                Longitude of the location in degrees
//    live_period          Integer              Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
//
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`
	// Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
}

// InputVenueMessageContent
// https://core.telegram.org/bots/api#inputvenuemessagecontent
//
// Represents the content of a venue message to be sent as the result of an inline query.
//
//    Field                Type                 Description
//    latitude             Float                Latitude of the venue in degrees
//    longitude            Float                Longitude of the venue in degrees
//    title                String               Name of the venue
//    address              String               Address of the venue
//    foursquare_id        String               Optional. Foursquare identifier of the venue, if known
//    foursquare_type      String               Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
//
type InputVenueMessageContent struct {
	// Latitude of the venue in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the venue in degrees
	Longitude float64 `json:"longitude"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue, if known
	FoursquareId string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
}

// InputContactMessageContent
// https://core.telegram.org/bots/api#inputcontactmessagecontent
//
// Represents the content of a contact message to be sent as the result of an inline query.
//
//    Field                Type                 Description
//    phone_number         String               Contact's phone number
//    first_name           String               Contact's first name
//    last_name            String               Optional. Contact's last name
//    vcard                String               Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
//
type InputContactMessageContent struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`
}

// ChosenInlineResult
// https://core.telegram.org/bots/api#choseninlineresult
//
// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
//
//    Field                Type                 Description
//    result_id            String               The unique identifier for the result that was chosen
//    from                 User                 The user that chose the result
//    location             Location             Optional. Sender location, only for bots that require user location
//    inline_message_id    String               Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
//    query                String               The query that was used to obtain the result
//
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultId string `json:"result_id"`
	// The user that chose the result
	From *User `json:"from"`
	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`
	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// The query that was used to obtain the result
	Query string `json:"query"`
}

// LabeledPrice
// https://core.telegram.org/bots/api#labeledprice
//
// This object represents a portion of the price for goods or services.
//
//    Field                Type                 Description
//    label                String               Portion label
//    amount               Integer              Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
//
type LabeledPrice struct {
	// Portion label
	Label string `json:"label"`
	// Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

// Invoice
// https://core.telegram.org/bots/api#invoice
//
// This object contains basic information about an invoice.
//
//    Field                Type                 Description
//    title                String               Product name
//    description          String               Product description
//    start_parameter      String               Unique bot deep-linking parameter that can be used to generate this invoice
//    currency             String               Three-letter ISO 4217 currency code
//    total_amount         Integer              Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
//
type Invoice struct {
	// Product name
	Title string `json:"title"`
	// Product description
	Description string `json:"description"`
	// Unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress
// https://core.telegram.org/bots/api#shippingaddress
//
// This object represents a shipping address.
//
//    Field                Type                 Description
//    country_code         String               ISO 3166-1 alpha-2 country code
//    state                String               State, if applicable
//    city                 String               City
//    street_line1         String               First line for the address
//    street_line2         String               Second line for the address
//    post_code            String               Address post code
//
type ShippingAddress struct {
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// State, if applicable
	State string `json:"state"`
	// City
	City string `json:"city"`
	// First line for the address
	StreetLine1 string `json:"street_line1"`
	// Second line for the address
	StreetLine2 string `json:"street_line2"`
	// Address post code
	PostCode string `json:"post_code"`
}

// OrderInfo
// https://core.telegram.org/bots/api#orderinfo
//
// This object represents information about an order.
//
//    Field                Type                 Description
//    name                 String               Optional. User name
//    phone_number         String               Optional. User's phone number
//    email                String               Optional. User email
//    shipping_address     ShippingAddress      Optional. User shipping address
//
type OrderInfo struct {
	// Optional. User name
	Name string `json:"name,omitempty"`
	// Optional. User's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// Optional. User email
	Email string `json:"email,omitempty"`
	// Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingOption
// https://core.telegram.org/bots/api#shippingoption
//
// This object represents one shipping option.
//
//    Field                Type                 Description
//    id                   String               Shipping option identifier
//    title                String               Option title
//    prices               Array of LabeledPrice List of price portions
//
type ShippingOption struct {
	// Shipping option identifier
	Id string `json:"id"`
	// Option title
	Title string `json:"title"`
	// List of price portions
	Prices []*LabeledPrice `json:"prices"`
}

// SuccessfulPayment
// https://core.telegram.org/bots/api#successfulpayment
//
// This object contains basic information about a successful payment.
//
//    Field                Type                 Description
//    currency             String               Three-letter ISO 4217 currency code
//    total_amount         Integer              Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
//    invoice_payload      String               Bot specified invoice payload
//    shipping_option_id   String               Optional. Identifier of the shipping option chosen by the user
//    order_info           OrderInfo            Optional. Order info provided by the user
//    telegram_payment_charge_id String               Telegram payment identifier
//    provider_payment_charge_id String               Provider payment identifier
//
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	// Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

// ShippingQuery
// https://core.telegram.org/bots/api#shippingquery
//
// This object contains information about an incoming shipping query.
//
//    Field                Type                 Description
//    id                   String               Unique query identifier
//    from                 User                 User who sent the query
//    invoice_payload      String               Bot specified invoice payload
//    shipping_address     ShippingAddress      User specified shipping address
//
type ShippingQuery struct {
	// Unique query identifier
	Id string `json:"id"`
	// User who sent the query
	From *User `json:"from"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery
// https://core.telegram.org/bots/api#precheckoutquery
//
// This object contains information about an incoming pre-checkout query.
//
//    Field                Type                 Description
//    id                   String               Unique query identifier
//    from                 User                 User who sent the query
//    currency             String               Three-letter ISO 4217 currency code
//    total_amount         Integer              Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
//    invoice_payload      String               Bot specified invoice payload
//    shipping_option_id   String               Optional. Identifier of the shipping option chosen by the user
//    order_info           OrderInfo            Optional. Order info provided by the user
//
type PreCheckoutQuery struct {
	// Unique query identifier
	Id string `json:"id"`
	// User who sent the query
	From *User `json:"from"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	// Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// PassportData
// https://core.telegram.org/bots/api#passportdata
//
// Contains information about Telegram Passport data shared with the bot by the user.
//
//    Field                Type                 Description
//    data                 Array of EncryptedPassportElement Array with information about documents and other Telegram Passport elements that was shared with the bot
//    credentials          EncryptedCredentials Encrypted credentials required to decrypt the data
//
type PassportData struct {
	// Array with information about documents and other Telegram Passport elements that was shared with the bot
	Data []*EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}

// PassportFile
// https://core.telegram.org/bots/api#passportfile
//
// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files
// are in JPEG format when decrypted and don't exceed 10MB.
//
//    Field                Type                 Description
//    file_id              String               Unique identifier for this file
//    file_size            Integer              File size
//    file_date            Integer              Unix time when the file was uploaded
//
type PassportFile struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// File size
	FileSize int `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate int `json:"file_date"`
}

// EncryptedPassportElement
// https://core.telegram.org/bots/api#encryptedpassportelement
//
// Contains information about documents or other Telegram Passport elements shared with the bot by
// the user.
//
//    Field                Type                 Description
//    type                 String               Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
//    data                 String               Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
//    phone_number         String               Optional. User's verified phone number, available only for “phone_number” type
//    email                String               Optional. User's verified email address, available only for “email” type
//    files                Array of PassportFile Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
//    front_side           PassportFile         Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
//    reverse_side         PassportFile         Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
//    selfie               PassportFile         Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
//    translation          Array of PassportFile Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
//    hash                 String               Base64-encoded element hash for using in PassportElementErrorUnspecified
//
type EncryptedPassportElement struct {
	// Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`
	// Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	Data string `json:"data,omitempty"`
	// Optional. User's verified phone number, available only for “phone_number” type
	PhoneNumber string `json:"phone_number,omitempty"`
	// Optional. User's verified email address, available only for “email” type
	Email string `json:"email,omitempty"`
	// Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Files []*PassportFile `json:"files,omitempty"`
	// Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`
	// Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	// Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`
	// Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []*PassportFile `json:"translation,omitempty"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

// EncryptedCredentials
// https://core.telegram.org/bots/api#encryptedcredentials
//
// Contains data required for decrypting and authenticating EncryptedPassportElement. See the
// Telegram Passport Documentation for a complete description of the data decryption and authentication
// processes.
//
//    Field                Type                 Description
//    data                 String               Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
//    hash                 String               Base64-encoded data hash for data authentication
//    secret               String               Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
//
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

// PassportElementError
// https://core.telegram.org/bots/api#passportelementerror
//
// This object represents an error in the Telegram Passport element which was submitted that should
// be resolved by the user. It should be one of:
//
type PassportElementError struct {
}

// PassportElementErrorDataField
// https://core.telegram.org/bots/api#passportelementerrordatafield
//
// Represents an issue in one of the data fields that was provided by the user. The error is
// considered resolved when the field's value changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be data
//    type                 String               The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
//    field_name           String               Name of the data field which has the error
//    data_hash            String               Base64-encoded data hash
//    message              String               Error message
//
type PassportElementErrorDataField struct {
	// Error source, must be data
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`
	// Name of the data field which has the error
	FieldName string `json:"field_name"`
	// Base64-encoded data hash
	DataHash string `json:"data_hash"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorFrontSide
// https://core.telegram.org/bots/api#passportelementerrorfrontside
//
// Represents an issue with the front side of a document. The error is considered resolved when the
// file with the front side of the document changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be front_side
//    type                 String               The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
//    file_hash            String               Base64-encoded hash of the file with the front side of the document
//    message              String               Error message
//
type PassportElementErrorFrontSide struct {
	// Error source, must be front_side
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorReverseSide
// https://core.telegram.org/bots/api#passportelementerrorreverseside
//
// Represents an issue with the reverse side of a document. The error is considered resolved when the
// file with reverse side of the document changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be reverse_side
//    type                 String               The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
//    file_hash            String               Base64-encoded hash of the file with the reverse side of the document
//    message              String               Error message
//
type PassportElementErrorReverseSide struct {
	// Error source, must be reverse_side
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorSelfie
// https://core.telegram.org/bots/api#passportelementerrorselfie
//
// Represents an issue with the selfie with a document. The error is considered resolved when the
// file with the selfie changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be selfie
//    type                 String               The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
//    file_hash            String               Base64-encoded hash of the file with the selfie
//    message              String               Error message
//
type PassportElementErrorSelfie struct {
	// Error source, must be selfie
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`
	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorFile
// https://core.telegram.org/bots/api#passportelementerrorfile
//
// Represents an issue with a document scan. The error is considered resolved when the file with the
// document scan changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be file
//    type                 String               The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
//    file_hash            String               Base64-encoded file hash
//    message              String               Error message
//
type PassportElementErrorFile struct {
	// Error source, must be file
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// Base64-encoded file hash
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorFiles
// https://core.telegram.org/bots/api#passportelementerrorfiles
//
// Represents an issue with a list of scans. The error is considered resolved when the list of files
// containing the scans changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be files
//    type                 String               The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
//    file_hashes          Array of String      List of base64-encoded file hashes
//    message              String               Error message
//
type PassportElementErrorFiles struct {
	// Error source, must be files
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorTranslationFile
// https://core.telegram.org/bots/api#passportelementerrortranslationfile
//
// Represents an issue with one of the files that constitute the translation of a document. The error
// is considered resolved when the file changes.
//
//    Field                Type                 Description
//    source               String               Error source, must be translation_file
//    type                 String               Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
//    file_hash            String               Base64-encoded file hash
//    message              String               Error message
//
type PassportElementErrorTranslationFile struct {
	// Error source, must be translation_file
	Source string `json:"source"`
	// Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// Base64-encoded file hash
	FileHash string `json:"file_hash"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorTranslationFiles
// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
//
// Represents an issue with the translated version of a document. The error is considered resolved
// when a file with the document translation change.
//
//    Field                Type                 Description
//    source               String               Error source, must be translation_files
//    type                 String               Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
//    file_hashes          Array of String      List of base64-encoded file hashes
//    message              String               Error message
//
type PassportElementErrorTranslationFiles struct {
	// Error source, must be translation_files
	Source string `json:"source"`
	// Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`
	// Error message
	Message string `json:"message"`
}

// PassportElementErrorUnspecified
// https://core.telegram.org/bots/api#passportelementerrorunspecified
//
// Represents an issue in an unspecified place. The error is considered resolved when new data is
// added.
//
//    Field                Type                 Description
//    source               String               Error source, must be unspecified
//    type                 String               Type of element of the user's Telegram Passport which has the issue
//    element_hash         String               Base64-encoded element hash
//    message              String               Error message
//
type PassportElementErrorUnspecified struct {
	// Error source, must be unspecified
	Source string `json:"source"`
	// Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type"`
	// Base64-encoded element hash
	ElementHash string `json:"element_hash"`
	// Error message
	Message string `json:"message"`
}

// Game
// https://core.telegram.org/bots/api#game
//
// This object represents a game. Use BotFather to create and edit games, their short names will act
// as unique identifiers.
//
//    Field                Type                 Description
//    title                String               Title of the game
//    description          String               Description of the game
//    photo                Array of PhotoSize   Photo that will be displayed in the game message in chats.
//    text                 String               Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
//    text_entities        Array of MessageEntity Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
//    animation            Animation            Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
//
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description string `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo"`
	// Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	Text string `json:"text,omitempty"`
	// Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	// Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
	Animation *Animation `json:"animation,omitempty"`
}

// CallbackGame
// https://core.telegram.org/bots/api#callbackgame
//
// A placeholder, currently holds no information. Use BotFather to set up your game.
//
type CallbackGame struct {
}

// GameHighScore
// https://core.telegram.org/bots/api#gamehighscore
//
// This object represents one row of the high scores table for a game.
//
//    Field                Type                 Description
//    position             Integer              Position in high score table for the game
//    user                 User                 User
//    score                Integer              Score
//
type GameHighScore struct {
	// Position in high score table for the game
	Position int `json:"position"`
	// User
	User *User `json:"user"`
	// Score
	Score int `json:"score"`
}
