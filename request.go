package telego

// getUpdates
// https://core.telegram.org/bots/api#getupdates
//
// Use this method to receive incoming updates using long polling (wiki). An Array of Update objects
// is returned.
//
//    Field                Type                 Description
//    offset               Integer              Optional. Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
//    limit                Integer              Optional. Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
//    timeout              Integer              Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
//    allowed_updates      Array of String      Optional. List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
//
type GetUpdatesRequest struct {
	// Optional. Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Offset int `json:"offset,omitempty"`
	// Optional. Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
	// Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	Timeout int `json:"timeout,omitempty"`
	// Optional. List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// setWebhook
// https://core.telegram.org/bots/api#setwebhook
//
// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever
// there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a
// JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable
// amount of attempts. Returns True on success.
//
//    Field                Type                 Description
//    url                  String               HTTPS url to send updates to. Use an empty string to remove webhook integration
//    certificate          InputFile            Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
//    max_connections      Integer              Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
//    allowed_updates      Array of String      Optional. List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
//
type SetWebhookRequest struct {
	// HTTPS url to send updates to. Use an empty string to remove webhook integration
	Url string `json:"url"`
	// Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	Certificate *InputFile `json:"certificate,omitempty"`
	// Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
	MaxConnections int `json:"max_connections,omitempty"`
	// Optional. List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// sendMessage
// https://core.telegram.org/bots/api#sendmessage
//
// Use this method to send text messages. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    text                 String               Text of the message to be sent
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
//    disable_web_page_preview Boolean              Optional. Disables link previews for links in this message
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Text of the message to be sent
	Text string `json:"text"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// forwardMessage
// https://core.telegram.org/bots/api#forwardmessage
//
// Use this method to forward messages of any kind. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    from_chat_id         Integer or String    Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    message_id           Integer              Message identifier in the chat specified in from_chat_id
//
type ForwardMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	FromChatId int `json:"from_chat_id"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Message identifier in the chat specified in from_chat_id
	MessageId int `json:"message_id"`
}

// sendPhoto
// https://core.telegram.org/bots/api#sendphoto
//
// Use this method to send photos. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    photo                InputFile or String  Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files »
//    caption              String               Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files »
	Photo *InputFile `json:"photo"`
	// Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendAudio
// https://core.telegram.org/bots/api#sendaudio
//
// Use this method to send audio files, if you want Telegram clients to display them in the music
// player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can
// currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    audio                InputFile or String  Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
//    caption              String               Optional. Audio caption, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    duration             Integer              Optional. Duration of the audio in seconds
//    performer            String               Optional. Performer
//    title                String               Optional. Track name
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendAudioRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Audio *InputFile `json:"audio"`
	// Optional. Audio caption, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Performer
	Performer string `json:"performer,omitempty"`
	// Optional. Track name
	Title string `json:"title,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendDocument
// https://core.telegram.org/bots/api#senddocument
//
// Use this method to send general files. On success, the sent Message is returned. Bots can
// currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    document             InputFile or String  File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Document caption (may also be used when resending documents by file_id), 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendDocumentRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Document *InputFile `json:"document"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Document caption (may also be used when resending documents by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendVideo
// https://core.telegram.org/bots/api#sendvideo
//
// Use this method to send video files, Telegram clients support mp4 videos (other formats may be
// sent as Document). On success, the sent Message is returned. Bots can currently send video files of up
// to 50 MB in size, this limit may be changed in the future.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    video                InputFile or String  Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
//    duration             Integer              Optional. Duration of sent video in seconds
//    width                Integer              Optional. Video width
//    height               Integer              Optional. Video height
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    supports_streaming   Boolean              Optional. Pass True, if the uploaded video is suitable for streaming
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendVideoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
	Video *InputFile `json:"video"`
	// Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Video width
	Width int `json:"width,omitempty"`
	// Optional. Video height
	Height int `json:"height,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendAnimation
// https://core.telegram.org/bots/api#sendanimation
//
// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success,
// the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this
// limit may be changed in the future.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    animation            InputFile or String  Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
//    duration             Integer              Optional. Duration of sent animation in seconds
//    width                Integer              Optional. Animation width
//    height               Integer              Optional. Animation height
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    caption              String               Optional. Animation caption (may also be used when resending animation by file_id), 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendAnimationRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
	Animation *InputFile `json:"animation"`
	// Optional. Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Animation width
	Width int `json:"width,omitempty"`
	// Optional. Animation height
	Height int `json:"height,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Animation caption (may also be used when resending animation by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendVoice
// https://core.telegram.org/bots/api#sendvoice
//
// Use this method to send audio files, if you want Telegram clients to display the file as a
// playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other
// formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently
// send voice messages of up to 50 MB in size, this limit may be changed in the future.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    voice                InputFile or String  Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
//    caption              String               Optional. Voice message caption, 0-1024 characters
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    duration             Integer              Optional. Duration of the voice message in seconds
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendVoiceRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Voice *InputFile `json:"voice"`
	// Optional. Voice message caption, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Duration of the voice message in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendVideoNote
// https://core.telegram.org/bots/api#sendvideonote
//
// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this
// method to send video messages. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    video_note           InputFile or String  Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
//    duration             Integer              Optional. Duration of sent video in seconds
//    length               Integer              Optional. Video width and height, i.e. diameter of the video message
//    thumb                InputFile or String  Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendVideoNoteRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	VideoNote *InputFile `json:"video_note"`
	// Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Video width and height, i.e. diameter of the video message
	Length int `json:"length,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb *InputFile `json:"thumb,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendMediaGroup
// https://core.telegram.org/bots/api#sendmediagroup
//
// Use this method to send a group of photos or videos as an album. On success, an array of the sent
// Messages is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    media                Array of InputMediaPhoto and InputMediaVideo A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
//    disable_notification Boolean              Optional. Sends the messages silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the messages are a reply, ID of the original message
//
type SendMediaGroupRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
	Media []*InputMediaPhoto `json:"media"`
	// Optional. Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the messages are a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
}

// sendLocation
// https://core.telegram.org/bots/api#sendlocation
//
// Use this method to send point on the map. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    latitude             Float number         Latitude of the location
//    longitude            Float number         Longitude of the location
//    live_period          Integer              Optional. Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendLocationRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Latitude of the location
	Latitude float64 `json:"latitude"`
	// Longitude of the location
	Longitude float64 `json:"longitude"`
	// Optional. Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// editMessageLiveLocation
// https://core.telegram.org/bots/api#editmessagelivelocation
//
// Use this method to edit live location messages. A location can be edited until its live_period
// expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the
// edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the message to edit
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//    latitude             Float number         Latitude of new location
//    longitude            Float number         Longitude of new location
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for a new inline keyboard.
//
type EditMessageLiveLocationRequest struct {
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Latitude of new location
	Latitude float64 `json:"latitude"`
	// Longitude of new location
	Longitude float64 `json:"longitude"`
	// Optional. A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// stopMessageLiveLocation
// https://core.telegram.org/bots/api#stopmessagelivelocation
//
// Use this method to stop updating a live location message before live_period expires. On success,
// if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the message with live location to stop
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for a new inline keyboard.
//
type StopMessageLiveLocationRequest struct {
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message with live location to stop
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Optional. A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendVenue
// https://core.telegram.org/bots/api#sendvenue
//
// Use this method to send information about a venue. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    latitude             Float number         Latitude of the venue
//    longitude            Float number         Longitude of the venue
//    title                String               Name of the venue
//    address              String               Address of the venue
//    foursquare_id        String               Optional. Foursquare identifier of the venue
//    foursquare_type      String               Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendVenueRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Latitude of the venue
	Latitude float64 `json:"latitude"`
	// Longitude of the venue
	Longitude float64 `json:"longitude"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue
	FoursquareId string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendContact
// https://core.telegram.org/bots/api#sendcontact
//
// Use this method to send phone contacts. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    phone_number         String               Contact's phone number
//    first_name           String               Contact's first name
//    last_name            String               Optional. Contact's last name
//    vcard                String               Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
//
type SendContactRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard string `json:"vcard,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendPoll
// https://core.telegram.org/bots/api#sendpoll
//
// Use this method to send a native poll. A native poll can't be sent to a private chat. On success,
// the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername). A native poll can't be sent to a private chat.
//    question             String               Poll question, 1-255 characters
//    options              Array of String      List of answer options, 2-10 strings 1-100 characters each
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendPollRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername). A native poll can't be sent to a private chat.
	ChatId int `json:"chat_id"`
	// Poll question, 1-255 characters
	Question string `json:"question"`
	// List of answer options, 2-10 strings 1-100 characters each
	Options []string `json:"options"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// sendChatAction
// https://core.telegram.org/bots/api#sendchataction
//
// Use this method when you need to tell the user that something is happening on the bot's side. The
// status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear
// its typing status). Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    action               String               Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
//
type SendChatActionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
	Action string `json:"action"`
}

// getUserProfilePhotos
// https://core.telegram.org/bots/api#getuserprofilephotos
//
// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
//
//    Field                Type                 Description
//    user_id              Integer              Unique identifier of the target user
//    offset               Integer              Optional. Sequential number of the first photo to be returned. By default, all photos are returned.
//    limit                Integer              Optional. Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
//
type GetUserProfilePhotosRequest struct {
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Optional. Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset int `json:"offset,omitempty"`
	// Optional. Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

// getFile
// https://core.telegram.org/bots/api#getfile
//
// Use this method to get basic info about a file and prepare it for downloading. For the moment,
// bots can download files of up to 20MB in size. On success, a File object is returned. The file can then
// be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path>
// is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When
// the link expires, a new one can be requested by calling getFile again.
//
//    Field                Type                 Description
//    file_id              String               File identifier to get info about
//
type GetFileRequest struct {
	// File identifier to get info about
	FileId string `json:"file_id"`
}

// kickChatMember
// https://core.telegram.org/bots/api#kickchatmember
//
// Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups
// and channels, the user will not be able to return to the group on their own using invite links,
// etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
//    user_id              Integer              Unique identifier of the target user
//    until_date           Integer              Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
//
type KickChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
	UntilDate int `json:"until_date,omitempty"`
}

// unbanChatMember
// https://core.telegram.org/bots/api#unbanchatmember
//
// Use this method to unban a previously kicked user in a supergroup or channel. The user will not
// return to the group or channel automatically, but will be able to join via link, etc. The bot must be
// an administrator for this to work. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
//    user_id              Integer              Unique identifier of the target user
//
type UnbanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

// restrictChatMember
// https://core.telegram.org/bots/api#restrictchatmember
//
// Use this method to restrict a user in a supergroup. The bot must be an administrator in the
// supergroup for this to work and must have the appropriate admin rights. Pass True for all boolean
// parameters to lift restrictions from a user. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//    user_id              Integer              Unique identifier of the target user
//    until_date           Integer              Optional. Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
//    can_send_messages    Boolean              Optional. Pass True, if the user can send text messages, contacts, locations and venues
//    can_send_media_messages Boolean              Optional. Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
//    can_send_other_messages Boolean              Optional. Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
//    can_add_web_page_previews Boolean              Optional. Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
//
type RestrictChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Optional. Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
	UntilDate int `json:"until_date,omitempty"`
	// Optional. Pass True, if the user can send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// Optional. Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// Optional. Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// Optional. Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
}

// promoteChatMember
// https://core.telegram.org/bots/api#promotechatmember
//
// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all
// boolean parameters to demote a user. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    user_id              Integer              Unique identifier of the target user
//    can_change_info      Boolean              Optional. Pass True, if the administrator can change chat title, photo and other settings
//    can_post_messages    Boolean              Optional. Pass True, if the administrator can create channel posts, channels only
//    can_edit_messages    Boolean              Optional. Pass True, if the administrator can edit messages of other users and can pin messages, channels only
//    can_delete_messages  Boolean              Optional. Pass True, if the administrator can delete messages of other users
//    can_invite_users     Boolean              Optional. Pass True, if the administrator can invite new users to the chat
//    can_restrict_members Boolean              Optional. Pass True, if the administrator can restrict, ban or unban chat members
//    can_pin_messages     Boolean              Optional. Pass True, if the administrator can pin messages, supergroups only
//    can_promote_members  Boolean              Optional. Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
//
type PromoteChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Optional. Pass True, if the administrator can change chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// Optional. Pass True, if the administrator can create channel posts, channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// Optional. Pass True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// Optional. Pass True, if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. Pass True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// Optional. Pass True, if the administrator can pin messages, supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// Optional. Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
}

// exportChatInviteLink
// https://core.telegram.org/bots/api#exportchatinvitelink
//
// Use this method to generate a new invite link for a chat; any previously generated link is
// revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin
// rights. Returns the new invite link as String on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
type ExportChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// setChatPhoto
// https://core.telegram.org/bots/api#setchatphoto
//
// Use this method to set a new profile photo for the chat. Photos can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    photo                InputFile            New chat photo, uploaded using multipart/form-data
//
type SetChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// New chat photo, uploaded using multipart/form-data
	Photo *InputFile `json:"photo"`
}

// deleteChatPhoto
// https://core.telegram.org/bots/api#deletechatphoto
//
// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be
// an administrator in the chat for this to work and must have the appropriate admin rights. Returns
// True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
type DeleteChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// setChatTitle
// https://core.telegram.org/bots/api#setchattitle
//
// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot
// must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    title                String               New chat title, 1-255 characters
//
type SetChatTitleRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// New chat title, 1-255 characters
	Title string `json:"title"`
}

// setChatDescription
// https://core.telegram.org/bots/api#setchatdescription
//
// Use this method to change the description of a supergroup or a channel. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin rights. Returns True on
// success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    description          String               Optional. New chat description, 0-255 characters
//
type SetChatDescriptionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Optional. New chat description, 0-255 characters
	Description string `json:"description,omitempty"`
}

// pinChatMessage
// https://core.telegram.org/bots/api#pinchatmessage
//
// Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an
// administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the
// supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Identifier of a message to pin
//    disable_notification Boolean              Optional. Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
//
type PinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Identifier of a message to pin
	MessageId int `json:"message_id"`
	// Optional. Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// unpinChatMessage
// https://core.telegram.org/bots/api#unpinchatmessage
//
// Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an
// administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the
// supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
type UnpinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// leaveChat
// https://core.telegram.org/bots/api#leavechat
//
// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
//
type LeaveChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// getChat
// https://core.telegram.org/bots/api#getchat
//
// Use this method to get up to date information about the chat (current name of the user for
// one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on
// success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
//
type GetChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// getChatAdministrators
// https://core.telegram.org/bots/api#getchatadministrators
//
// Use this method to get a list of administrators in a chat. On success, returns an Array of
// ChatMember objects that contains information about all chat administrators except other bots. If the chat
// is a group or a supergroup and no administrators were appointed, only the creator will be returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
//
type GetChatAdministratorsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// getChatMembersCount
// https://core.telegram.org/bots/api#getchatmemberscount
//
// Use this method to get the number of members in a chat. Returns Int on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
//
type GetChatMembersCountRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}

// getChatMember
// https://core.telegram.org/bots/api#getchatmember
//
// Use this method to get information about a member of a chat. Returns a ChatMember object on
// success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
//    user_id              Integer              Unique identifier of the target user
//
type GetChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

// setChatStickerSet
// https://core.telegram.org/bots/api#setchatstickerset
//
// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True
// on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//    sticker_set_name     String               Name of the sticker set to be set as the group sticker set
//
type SetChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId int `json:"chat_id"`
	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

// deleteChatStickerSet
// https://core.telegram.org/bots/api#deletechatstickerset
//
// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True
// on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//
type DeleteChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId int `json:"chat_id"`
}

// answerCallbackQuery
// https://core.telegram.org/bots/api#answercallbackquery
//
// Use this method to send answers to callback queries sent from inline keyboards. The answer will be
// displayed to the user as a notification at the top of the chat screen or as an alert. On success,
// True is returned.
//
//    Field                Type                 Description
//    callback_query_id    String               Unique identifier for the query to be answered
//    text                 String               Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
//    show_alert           Boolean              Optional. If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
//    url                  String               Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
//    cache_time           Integer              Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
//
type AnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered
	CallbackQueryId string `json:"callback_query_id"`
	// Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text string `json:"text,omitempty"`
	// Optional. If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`
	// Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	Url string `json:"url,omitempty"`
	// Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}

// editMessageText
// https://core.telegram.org/bots/api#editmessagetext
//
// Use this method to edit text and game messages. On success, if edited message is sent by the bot,
// the edited Message is returned, otherwise True is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the message to edit
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//    text                 String               New text of the message
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
//    disable_web_page_preview Boolean              Optional. Disables link previews for links in this message
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for an inline keyboard.
//
type EditMessageTextRequest struct {
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// New text of the message
	Text string `json:"text"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// editMessageCaption
// https://core.telegram.org/bots/api#editmessagecaption
//
// Use this method to edit captions of messages. On success, if edited message is sent by the bot,
// the edited Message is returned, otherwise True is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the message to edit
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//    caption              String               Optional. New caption of the message
//    parse_mode           String               Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for an inline keyboard.
//
type EditMessageCaptionRequest struct {
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Optional. New caption of the message
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// editMessageMedia
// https://core.telegram.org/bots/api#editmessagemedia
//
// Use this method to edit animation, audio, document, photo, or video messages. If a message is a
// part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can
// be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously
// uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the
// bot, the edited Message is returned, otherwise True is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the message to edit
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//    media                InputMedia           A JSON-serialized object for a new media content of the message
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for a new inline keyboard.
//
type EditMessageMediaRequest struct {
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// A JSON-serialized object for a new media content of the message
	Media *InputMedia `json:"media"`
	// Optional. A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// editMessageReplyMarkup
// https://core.telegram.org/bots/api#editmessagereplymarkup
//
// Use this method to edit only the reply markup of messages. On success, if edited message is sent
// by the bot, the edited Message is returned, otherwise True is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the message to edit
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for an inline keyboard.
//
type EditMessageReplyMarkupRequest struct {
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// stopPoll
// https://core.telegram.org/bots/api#stoppoll
//
// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll with the
// final results is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Identifier of the original message with the poll
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for a new message inline keyboard.
//
type StopPollRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Identifier of the original message with the poll
	MessageId int `json:"message_id"`
	// Optional. A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// deleteMessage
// https://core.telegram.org/bots/api#deletemessage
//
// Use this method to delete a message, including service messages, with the following limitations:-
// A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing
// messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.-
// Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is
// an administrator of a group, it can delete any message there.- If the bot has can_delete_messages
// permission in a supergroup or a channel, it can delete any message there.Returns True on success.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    message_id           Integer              Identifier of the message to delete
//
type DeleteMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Identifier of the message to delete
	MessageId int `json:"message_id"`
}

// sendSticker
// https://core.telegram.org/bots/api#sendsticker
//
// Use this method to send .webp stickers. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer or String    Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//    sticker              InputFile or String  Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
//
type SendStickerRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Sticker *InputFile `json:"sticker"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// getStickerSet
// https://core.telegram.org/bots/api#getstickerset
//
// Use this method to get a sticker set. On success, a StickerSet object is returned.
//
//    Field                Type                 Description
//    name                 String               Name of the sticker set
//
type GetStickerSetRequest struct {
	// Name of the sticker set
	Name string `json:"name"`
}

// uploadStickerFile
// https://core.telegram.org/bots/api#uploadstickerfile
//
// Use this method to upload a .png file with a sticker for later use in createNewStickerSet and
// addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
//
//    Field                Type                 Description
//    user_id              Integer              User identifier of sticker file owner
//    png_sticker          InputFile            Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
//
type UploadStickerFileRequest struct {
	// User identifier of sticker file owner
	UserId int `json:"user_id"`
	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
	PngSticker *InputFile `json:"png_sticker"`
}

// createNewStickerSet
// https://core.telegram.org/bots/api#createnewstickerset
//
// Use this method to create new sticker set owned by a user. The bot will be able to edit the
// created sticker set. Returns True on success.
//
//    Field                Type                 Description
//    user_id              Integer              User identifier of created sticker set owner
//    name                 String               Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
//    title                String               Sticker set title, 1-64 characters
//    png_sticker          InputFile or String  Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
//    emojis               String               One or more emoji corresponding to the sticker
//    contains_masks       Boolean              Optional. Pass True, if a set of mask stickers should be created
//    mask_position        MaskPosition         Optional. A JSON-serialized object for position where the mask should be placed on faces
//
type CreateNewStickerSetRequest struct {
	// User identifier of created sticker set owner
	UserId int `json:"user_id"`
	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`
	// Sticker set title, 1-64 characters
	Title string `json:"title"`
	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	PngSticker *InputFile `json:"png_sticker"`
	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`
	// Optional. Pass True, if a set of mask stickers should be created
	ContainsMasks bool `json:"contains_masks,omitempty"`
	// Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// addStickerToSet
// https://core.telegram.org/bots/api#addstickertoset
//
// Use this method to add a new sticker to a set created by the bot. Returns True on success.
//
//    Field                Type                 Description
//    user_id              Integer              User identifier of sticker set owner
//    name                 String               Sticker set name
//    png_sticker          InputFile or String  Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
//    emojis               String               One or more emoji corresponding to the sticker
//    mask_position        MaskPosition         Optional. A JSON-serialized object for position where the mask should be placed on faces
//
type AddStickerToSetRequest struct {
	// User identifier of sticker set owner
	UserId int `json:"user_id"`
	// Sticker set name
	Name string `json:"name"`
	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	PngSticker *InputFile `json:"png_sticker"`
	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`
	// Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// setStickerPositionInSet
// https://core.telegram.org/bots/api#setstickerpositioninset
//
// Use this method to move a sticker in a set created by the bot to a specific position . Returns
// True on success.
//
//    Field                Type                 Description
//    sticker              String               File identifier of the sticker
//    position             Integer              New sticker position in the set, zero-based
//
type SetStickerPositionInSetRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker"`
	// New sticker position in the set, zero-based
	Position int `json:"position"`
}

// deleteStickerFromSet
// https://core.telegram.org/bots/api#deletestickerfromset
//
// Use this method to delete a sticker from a set created by the bot. Returns True on success.
//
//    Field                Type                 Description
//    sticker              String               File identifier of the sticker
//
type DeleteStickerFromSetRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker"`
}

// answerInlineQuery
// https://core.telegram.org/bots/api#answerinlinequery
//
// Use this method to send answers to an inline query. On success, True is returned.No more than 50
// results per query are allowed.
//
//    Field                Type                 Description
//    inline_query_id      String               Unique identifier for the answered query
//    results              Array of InlineQueryResult A JSON-serialized array of results for the inline query
//    cache_time           Integer              Optional. The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
//    is_personal          Boolean              Optional. Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
//    next_offset          String               Optional. Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
//    switch_pm_text       String               Optional. If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
//    switch_pm_parameter  String               Optional. Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
//
type AnswerInlineQueryRequest struct {
	// Unique identifier for the answered query
	InlineQueryId string `json:"inline_query_id"`
	// A JSON-serialized array of results for the inline query
	Results []*InlineQueryResult `json:"results"`
	// Optional. The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	CacheTime int `json:"cache_time,omitempty"`
	// Optional. Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	IsPersonal bool `json:"is_personal,omitempty"`
	// Optional. Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`
	// Optional. If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmText string `json:"switch_pm_text,omitempty"`
	// Optional. Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}

// sendInvoice
// https://core.telegram.org/bots/api#sendinvoice
//
// Use this method to send invoices. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer              Unique identifier for the target private chat
//    title                String               Product name, 1-32 characters
//    description          String               Product description, 1-255 characters
//    payload              String               Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
//    provider_token       String               Payments provider token, obtained via Botfather
//    start_parameter      String               Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
//    currency             String               Three-letter ISO 4217 currency code, see more on currencies
//    prices               Array of LabeledPrice Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
//    provider_data        String               Optional. JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
//    photo_url            String               Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
//    photo_size           Integer              Optional. Photo size
//    photo_width          Integer              Optional. Photo width
//    photo_height         Integer              Optional. Photo height
//    need_name            Boolean              Optional. Pass True, if you require the user's full name to complete the order
//    need_phone_number    Boolean              Optional. Pass True, if you require the user's phone number to complete the order
//    need_email           Boolean              Optional. Pass True, if you require the user's email address to complete the order
//    need_shipping_address Boolean              Optional. Pass True, if you require the user's shipping address to complete the order
//    send_phone_number_to_provider Boolean              Optional. Pass True, if user's phone number should be sent to provider
//    send_email_to_provider Boolean              Optional. Pass True, if user's email address should be sent to provider
//    is_flexible          Boolean              Optional. Pass True, if the final price depends on the shipping method
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
//
type SendInvoiceRequest struct {
	// Unique identifier for the target private chat
	ChatId int `json:"chat_id"`
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	Payload string `json:"payload"`
	// Payments provider token, obtained via Botfather
	ProviderToken string `json:"provider_token"`
	// Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code, see more on currencies
	Currency string `json:"currency"`
	// Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`
	// Optional. JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`
	// Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoUrl string `json:"photo_url,omitempty"`
	// Optional. Photo size
	PhotoSize int `json:"photo_size,omitempty"`
	// Optional. Photo width
	PhotoWidth int `json:"photo_width,omitempty"`
	// Optional. Photo height
	PhotoHeight int `json:"photo_height,omitempty"`
	// Optional. Pass True, if you require the user's full name to complete the order
	NeedName bool `json:"need_name,omitempty"`
	// Optional. Pass True, if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
	// Optional. Pass True, if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email,omitempty"`
	// Optional. Pass True, if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
	// Optional. Pass True, if user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	// Optional. Pass True, if user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
	// Optional. Pass True, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// answerShippingQuery
// https://core.telegram.org/bots/api#answershippingquery
//
// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified,
// the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to
// shipping queries. On success, True is returned.
//
//    Field                Type                 Description
//    shipping_query_id    String               Unique identifier for the query to be answered
//    ok                   Boolean              Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
//    shipping_options     Array of ShippingOption Optional. Required if ok is True. A JSON-serialized array of available shipping options.
//    error_message        String               Optional. Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
//
type AnswerShippingQueryRequest struct {
	// Unique identifier for the query to be answered
	ShippingQueryId string `json:"shipping_query_id"`
	// Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok"`
	// Optional. Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"`
	// Optional. Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// answerPreCheckoutQuery
// https://core.telegram.org/bots/api#answerprecheckoutquery
//
// Once the user has confirmed their payment and shipping details, the Bot API sends the final
// confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to
// such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer
// within 10 seconds after the pre-checkout query was sent.
//
//    Field                Type                 Description
//    pre_checkout_query_id String               Unique identifier for the query to be answered
//    ok                   Boolean              Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
//    error_message        String               Optional. Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
//
type AnswerPreCheckoutQueryRequest struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	Ok bool `json:"ok"`
	// Optional. Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// setPassportDataErrors
// https://core.telegram.org/bots/api#setpassportdataerrors
//
// Informs a user that some of the Telegram Passport elements they provided contains errors. The user
// will not be able to re-submit their Passport to you until the errors are fixed (the contents of the
// field for which you returned the error must change). Returns True on success.
//
//    Field                Type                 Description
//    user_id              Integer              User identifier
//    errors               Array of PassportElementError A JSON-serialized array describing the errors
//
type SetPassportDataErrorsRequest struct {
	// User identifier
	UserId int `json:"user_id"`
	// A JSON-serialized array describing the errors
	Errors []*PassportElementError `json:"errors"`
}

// sendGame
// https://core.telegram.org/bots/api#sendgame
//
// Use this method to send a game. On success, the sent Message is returned.
//
//    Field                Type                 Description
//    chat_id              Integer              Unique identifier for the target chat
//    game_short_name      String               Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
//    disable_notification Boolean              Optional. Sends the message silently. Users will receive a notification with no sound.
//    reply_to_message_id  Integer              Optional. If the message is a reply, ID of the original message
//    reply_markup         InlineKeyboardMarkup Optional. A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
//
type SendGameRequest struct {
	// Unique identifier for the target chat
	ChatId int `json:"chat_id"`
	// Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
	GameShortName string `json:"game_short_name"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// setGameScore
// https://core.telegram.org/bots/api#setgamescore
//
// Use this method to set the score of the specified user in a game. On success, if the message was
// sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new
// score is not greater than the user's current score in the chat and force is False.
//
//    Field                Type                 Description
//    user_id              Integer              User identifier
//    score                Integer              New score, must be non-negative
//    force                Boolean              Optional. Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
//    disable_edit_message Boolean              Optional. Pass True, if the game message should not be automatically edited to include the current scoreboard
//    chat_id              Integer              Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the sent message
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//
type SetGameScoreRequest struct {
	// User identifier
	UserId int `json:"user_id"`
	// New score, must be non-negative
	Score int `json:"score"`
	// Optional. Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	Force bool `json:"force,omitempty"`
	// Optional. Pass True, if the game message should not be automatically edited to include the current scoreboard
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the sent message
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
}

// getGameHighScores
// https://core.telegram.org/bots/api#getgamehighscores
//
// Use this method to get data for high score tables. Will return the score of the specified user and
// several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
//
//    Field                Type                 Description
//    user_id              Integer              Target user id
//    chat_id              Integer              Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
//    message_id           Integer              Optional. Required if inline_message_id is not specified. Identifier of the sent message
//    inline_message_id    String               Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
//
type GetGameHighScoresRequest struct {
	// Target user id
	UserId int `json:"user_id"`
	// Optional. Required if inline_message_id is not specified. Unique identifier for the target chat
	ChatId int `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of the sent message
	MessageId int `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId string `json:"inline_message_id,omitempty"`
}
