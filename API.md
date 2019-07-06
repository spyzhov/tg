# Telegram Bot API

> Generated: **06 July 2019**

Please feel free to reed an original [Telegram Bot API](https://core.telegram.org/bots/api) documentation. 

API for Bot was generated automatically using original documentation and may has the difference. 
If this has happened, please [inform the author](https://github.com/spyzhov/tg/issues/new).


## Table of contents

* [Available methods](#available-methods)
  * [addStickerToSet](#addstickertoset)
  * [answerCallbackQuery](#answercallbackquery)
  * [answerInlineQuery](#answerinlinequery)
  * [answerPreCheckoutQuery](#answerprecheckoutquery)
  * [answerShippingQuery](#answershippingquery)
  * [createNewStickerSet](#createnewstickerset)
  * [deleteChatPhoto](#deletechatphoto)
  * [deleteChatStickerSet](#deletechatstickerset)
  * [deleteMessage](#deletemessage)
  * [deleteStickerFromSet](#deletestickerfromset)
  * [deleteWebhook](#deletewebhook)
  * [editMessageCaption](#editmessagecaption)
  * [editMessageLiveLocation](#editmessagelivelocation)
  * [editMessageMedia](#editmessagemedia)
  * [editMessageReplyMarkup](#editmessagereplymarkup)
  * [editMessageText](#editmessagetext)
  * [exportChatInviteLink](#exportchatinvitelink)
  * [forwardMessage](#forwardmessage)
  * [getChat](#getchat)
  * [getChatAdministrators](#getchatadministrators)
  * [getChatMember](#getchatmember)
  * [getChatMembersCount](#getchatmemberscount)
  * [getFile](#getfile)
  * [getGameHighScores](#getgamehighscores)
  * [getMe](#getme)
  * [getStickerSet](#getstickerset)
  * [getUpdates](#getupdates)
  * [getUserProfilePhotos](#getuserprofilephotos)
  * [getWebhookInfo](#getwebhookinfo)
  * [kickChatMember](#kickchatmember)
  * [leaveChat](#leavechat)
  * [pinChatMessage](#pinchatmessage)
  * [promoteChatMember](#promotechatmember)
  * [restrictChatMember](#restrictchatmember)
  * [sendAnimation](#sendanimation)
  * [sendAudio](#sendaudio)
  * [sendChatAction](#sendchataction)
  * [sendContact](#sendcontact)
  * [sendDocument](#senddocument)
  * [sendGame](#sendgame)
  * [sendInvoice](#sendinvoice)
  * [sendLocation](#sendlocation)
  * [sendMediaGroup](#sendmediagroup)
  * [sendMessage](#sendmessage)
  * [sendPhoto](#sendphoto)
  * [sendPoll](#sendpoll)
  * [sendSticker](#sendsticker)
  * [sendVenue](#sendvenue)
  * [sendVideo](#sendvideo)
  * [sendVideoNote](#sendvideonote)
  * [sendVoice](#sendvoice)
  * [setChatDescription](#setchatdescription)
  * [setChatPhoto](#setchatphoto)
  * [setChatStickerSet](#setchatstickerset)
  * [setChatTitle](#setchattitle)
  * [setGameScore](#setgamescore)
  * [setPassportDataErrors](#setpassportdataerrors)
  * [setStickerPositionInSet](#setstickerpositioninset)
  * [setWebhook](#setwebhook)
  * [stopMessageLiveLocation](#stopmessagelivelocation)
  * [stopPoll](#stoppoll)
  * [unbanChatMember](#unbanchatmember)
  * [unpinChatMessage](#unpinchatmessage)
  * [uploadStickerFile](#uploadstickerfile)
* [Available types](#available-types)
  * [Animation](#animation)
  * [Audio](#audio)
  * [CallbackGame](#callbackgame)
  * [CallbackQuery](#callbackquery)
  * [Chat](#chat)
  * [ChatMember](#chatmember)
  * [ChatPhoto](#chatphoto)
  * [ChosenInlineResult](#choseninlineresult)
  * [Contact](#contact)
  * [Document](#document)
  * [EncryptedCredentials](#encryptedcredentials)
  * [EncryptedPassportElement](#encryptedpassportelement)
  * [File](#file)
  * [ForceReply](#forcereply)
  * [Game](#game)
  * [GameHighScore](#gamehighscore)
  * [InlineKeyboardButton](#inlinekeyboardbutton)
  * [InlineKeyboardMarkup](#inlinekeyboardmarkup)
  * [InlineQuery](#inlinequery)
  * [InlineQueryResult](#inlinequeryresult)
  * [InlineQueryResultArticle](#inlinequeryresultarticle)
  * [InlineQueryResultAudio](#inlinequeryresultaudio)
  * [InlineQueryResultCachedAudio](#inlinequeryresultcachedaudio)
  * [InlineQueryResultCachedDocument](#inlinequeryresultcacheddocument)
  * [InlineQueryResultCachedGif](#inlinequeryresultcachedgif)
  * [InlineQueryResultCachedMpeg4Gif](#inlinequeryresultcachedmpeg4gif)
  * [InlineQueryResultCachedPhoto](#inlinequeryresultcachedphoto)
  * [InlineQueryResultCachedSticker](#inlinequeryresultcachedsticker)
  * [InlineQueryResultCachedVideo](#inlinequeryresultcachedvideo)
  * [InlineQueryResultCachedVoice](#inlinequeryresultcachedvoice)
  * [InlineQueryResultContact](#inlinequeryresultcontact)
  * [InlineQueryResultDocument](#inlinequeryresultdocument)
  * [InlineQueryResultGame](#inlinequeryresultgame)
  * [InlineQueryResultGif](#inlinequeryresultgif)
  * [InlineQueryResultLocation](#inlinequeryresultlocation)
  * [InlineQueryResultMpeg4Gif](#inlinequeryresultmpeg4gif)
  * [InlineQueryResultPhoto](#inlinequeryresultphoto)
  * [InlineQueryResultVenue](#inlinequeryresultvenue)
  * [InlineQueryResultVideo](#inlinequeryresultvideo)
  * [InlineQueryResultVoice](#inlinequeryresultvoice)
  * [InputContactMessageContent](#inputcontactmessagecontent)
  * [InputFile](#inputfile)
  * [InputLocationMessageContent](#inputlocationmessagecontent)
  * [InputMedia](#inputmedia)
  * [InputMediaAnimation](#inputmediaanimation)
  * [InputMediaAudio](#inputmediaaudio)
  * [InputMediaDocument](#inputmediadocument)
  * [InputMediaPhoto](#inputmediaphoto)
  * [InputMediaVideo](#inputmediavideo)
  * [InputMessageContent](#inputmessagecontent)
  * [InputTextMessageContent](#inputtextmessagecontent)
  * [InputVenueMessageContent](#inputvenuemessagecontent)
  * [Invoice](#invoice)
  * [KeyboardButton](#keyboardbutton)
  * [LabeledPrice](#labeledprice)
  * [Location](#location)
  * [LoginUrl](#loginurl)
  * [MaskPosition](#maskposition)
  * [Message](#message)
  * [MessageEntity](#messageentity)
  * [OrderInfo](#orderinfo)
  * [PassportData](#passportdata)
  * [PassportElementError](#passportelementerror)
  * [PassportElementErrorDataField](#passportelementerrordatafield)
  * [PassportElementErrorFile](#passportelementerrorfile)
  * [PassportElementErrorFiles](#passportelementerrorfiles)
  * [PassportElementErrorFrontSide](#passportelementerrorfrontside)
  * [PassportElementErrorReverseSide](#passportelementerrorreverseside)
  * [PassportElementErrorSelfie](#passportelementerrorselfie)
  * [PassportElementErrorTranslationFile](#passportelementerrortranslationfile)
  * [PassportElementErrorTranslationFiles](#passportelementerrortranslationfiles)
  * [PassportElementErrorUnspecified](#passportelementerrorunspecified)
  * [PassportFile](#passportfile)
  * [PhotoSize](#photosize)
  * [Poll](#poll)
  * [PollOption](#polloption)
  * [PreCheckoutQuery](#precheckoutquery)
  * [ReplyKeyboardMarkup](#replykeyboardmarkup)
  * [ReplyKeyboardRemove](#replykeyboardremove)
  * [ResponseParameters](#responseparameters)
  * [ShippingAddress](#shippingaddress)
  * [ShippingOption](#shippingoption)
  * [ShippingQuery](#shippingquery)
  * [Sticker](#sticker)
  * [StickerSet](#stickerset)
  * [SuccessfulPayment](#successfulpayment)
  * [Update](#update)
  * [User](#user)
  * [UserProfilePhotos](#userprofilephotos)
  * [Venue](#venue)
  * [Video](#video)
  * [VideoNote](#videonote)
  * [Voice](#voice)
  * [WebhookInfo](#webhookinfo)

## Available methods

### addStickerToSet

**[Official documentation](https://core.telegram.org/bots/api#addstickertoset)**

#### Description

Use this method to add a new sticker to a set created by the bot. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | User identifier of sticker set owner |
| name | String | Sticker set name |
| png_sticker | InputFile or String | Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » |
| emojis | String | One or more emoji corresponding to the sticker |
| mask_position | MaskPosition | Optional. A JSON-serialized object for position where the mask should be placed on faces |

#### Interface

```go
func (b *Bot) AddStickerToSet(ctx context.Context, request *AddStickerToSetRequest) (result bool, err error)
```

#### Request

```go
type AddStickerToSetRequest struct {
	// User identifier of sticker set owner
	UserId int `json:"user_id"`
	// Sticker set name
	Name string `json:"name"`
	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	PngSticker interface{} `json:"png_sticker"`
	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`
	// Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}
```

### answerCallbackQuery

**[Official documentation](https://core.telegram.org/bots/api#answercallbackquery)**

#### Description

Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| callback_query_id | String | Unique identifier for the query to be answered |
| text | String | Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters |
| show_alert | Boolean | Optional. If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false. |
| url | String | Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter. |
| cache_time | Integer | Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0. |

#### Interface

```go
func (b *Bot) AnswerCallbackQuery(ctx context.Context, request *AnswerCallbackQueryRequest) (result bool, err error)
```

#### Request

```go
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
```

### answerInlineQuery

**[Official documentation](https://core.telegram.org/bots/api#answerinlinequery)**

#### Description

Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| inline_query_id | String | Unique identifier for the answered query |
| results | Array of InlineQueryResult | A JSON-serialized array of results for the inline query |
| cache_time | Integer | Optional. The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300. |
| is_personal | Boolean | Optional. Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query |
| next_offset | String | Optional. Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes. |
| switch_pm_text | String | Optional. If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter |
| switch_pm_parameter | String | Optional. Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities. |

#### Interface

```go
func (b *Bot) AnswerInlineQuery(ctx context.Context, request *AnswerInlineQueryRequest) (result bool, err error)
```

#### Request

```go
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
```

### answerPreCheckoutQuery

**[Official documentation](https://core.telegram.org/bots/api#answerprecheckoutquery)**

#### Description

Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| pre_checkout_query_id | String | Unique identifier for the query to be answered |
| ok | Boolean | Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems. |
| error_message | String | Optional. Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user. |

#### Interface

```go
func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, request *AnswerPreCheckoutQueryRequest) (result bool, err error)
```

#### Request

```go
type AnswerPreCheckoutQueryRequest struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	Ok bool `json:"ok"`
	// Optional. Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}
```

### answerShippingQuery

**[Official documentation](https://core.telegram.org/bots/api#answershippingquery)**

#### Description

If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| shipping_query_id | String | Unique identifier for the query to be answered |
| ok | Boolean | Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible) |
| shipping_options | Array of ShippingOption | Optional. Required if ok is True. A JSON-serialized array of available shipping options. |
| error_message | String | Optional. Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user. |

#### Interface

```go
func (b *Bot) AnswerShippingQuery(ctx context.Context, request *AnswerShippingQueryRequest) (result bool, err error)
```

#### Request

```go
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
```

### createNewStickerSet

**[Official documentation](https://core.telegram.org/bots/api#createnewstickerset)**

#### Description

Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | User identifier of created sticker set owner |
| name | String | Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters. |
| title | String | Sticker set title, 1-64 characters |
| png_sticker | InputFile or String | Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » |
| emojis | String | One or more emoji corresponding to the sticker |
| contains_masks | Boolean | Optional. Pass True, if a set of mask stickers should be created |
| mask_position | MaskPosition | Optional. A JSON-serialized object for position where the mask should be placed on faces |

#### Interface

```go
func (b *Bot) CreateNewStickerSet(ctx context.Context, request *CreateNewStickerSetRequest) (result bool, err error)
```

#### Request

```go
type CreateNewStickerSetRequest struct {
	// User identifier of created sticker set owner
	UserId int `json:"user_id"`
	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`
	// Sticker set title, 1-64 characters
	Title string `json:"title"`
	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	PngSticker interface{} `json:"png_sticker"`
	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`
	// Optional. Pass True, if a set of mask stickers should be created
	ContainsMasks bool `json:"contains_masks,omitempty"`
	// Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}
```

### deleteChatPhoto

**[Official documentation](https://core.telegram.org/bots/api#deletechatphoto)**

#### Description

Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) DeleteChatPhoto(ctx context.Context, request *DeleteChatPhotoRequest) (result bool, err error)
```

#### Request

```go
type DeleteChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### deleteChatStickerSet

**[Official documentation](https://core.telegram.org/bots/api#deletechatstickerset)**

#### Description

Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername) |

#### Interface

```go
func (b *Bot) DeleteChatStickerSet(ctx context.Context, request *DeleteChatStickerSetRequest) (result bool, err error)
```

#### Request

```go
type DeleteChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId int `json:"chat_id"`
}
```

### deleteMessage

**[Official documentation](https://core.telegram.org/bots/api#deletemessage)**

#### Description

Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Identifier of the message to delete |

#### Interface

```go
func (b *Bot) DeleteMessage(ctx context.Context, request *DeleteMessageRequest) (result bool, err error)
```

#### Request

```go
type DeleteMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Identifier of the message to delete
	MessageId int `json:"message_id"`
}
```

### deleteStickerFromSet

**[Official documentation](https://core.telegram.org/bots/api#deletestickerfromset)**

#### Description

Use this method to delete a sticker from a set created by the bot. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| sticker | String | File identifier of the sticker |

#### Interface

```go
func (b *Bot) DeleteStickerFromSet(ctx context.Context, request *DeleteStickerFromSetRequest) (result bool, err error)
```

#### Request

```go
type DeleteStickerFromSetRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker"`
}
```

### deleteWebhook

**[Official documentation](https://core.telegram.org/bots/api#deletewebhook)**

#### Description

Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
#### Interface

```go
func (b *Bot) DeleteWebhook(ctx context.Context) (result bool, err error)
```
### editMessageCaption

**[Official documentation](https://core.telegram.org/bots/api#editmessagecaption)**

#### Description

Use this method to edit captions of messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the message to edit |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |
| caption | String | Optional. New caption of the message |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for an inline keyboard. |

#### Interface

```go
func (b *Bot) EditMessageCaption(ctx context.Context, request *EditMessageCaptionRequest) (result *Message, err error)
```

#### Request

```go
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
```

### editMessageLiveLocation

**[Official documentation](https://core.telegram.org/bots/api#editmessagelivelocation)**

#### Description

Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the message to edit |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |
| latitude | Float number | Latitude of new location |
| longitude | Float number | Longitude of new location |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for a new inline keyboard. |

#### Interface

```go
func (b *Bot) EditMessageLiveLocation(ctx context.Context, request *EditMessageLiveLocationRequest) (result *Message, err error)
```

#### Request

```go
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
```

### editMessageMedia

**[Official documentation](https://core.telegram.org/bots/api#editmessagemedia)**

#### Description

Use this method to edit animation, audio, document, photo, or video messages. If a message is a part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the message to edit |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |
| media | InputMedia | A JSON-serialized object for a new media content of the message |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for a new inline keyboard. |

#### Interface

```go
func (b *Bot) EditMessageMedia(ctx context.Context, request *EditMessageMediaRequest) (result *Message, err error)
```

#### Request

```go
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
```

### editMessageReplyMarkup

**[Official documentation](https://core.telegram.org/bots/api#editmessagereplymarkup)**

#### Description

Use this method to edit only the reply markup of messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the message to edit |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for an inline keyboard. |

#### Interface

```go
func (b *Bot) EditMessageReplyMarkup(ctx context.Context, request *EditMessageReplyMarkupRequest) (result *Message, err error)
```

#### Request

```go
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
```

### editMessageText

**[Official documentation](https://core.telegram.org/bots/api#editmessagetext)**

#### Description

Use this method to edit text and game messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the message to edit |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |
| text | String | New text of the message |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message. |
| disable_web_page_preview | Boolean | Optional. Disables link previews for links in this message |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for an inline keyboard. |

#### Interface

```go
func (b *Bot) EditMessageText(ctx context.Context, request *EditMessageTextRequest) (result *Message, err error)
```

#### Request

```go
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
```

### exportChatInviteLink

**[Official documentation](https://core.telegram.org/bots/api#exportchatinvitelink)**

#### Description

Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) ExportChatInviteLink(ctx context.Context, request *ExportChatInviteLinkRequest) (result string, err error)
```

#### Request

```go
type ExportChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### forwardMessage

**[Official documentation](https://core.telegram.org/bots/api#forwardmessage)**

#### Description

Use this method to forward messages of any kind. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| from_chat_id | Integer or String | Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername) |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| message_id | Integer | Message identifier in the chat specified in from_chat_id |

#### Interface

```go
func (b *Bot) ForwardMessage(ctx context.Context, request *ForwardMessageRequest) (result *Message, err error)
```

#### Request

```go
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
```

### getChat

**[Official documentation](https://core.telegram.org/bots/api#getchat)**

#### Description

Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) GetChat(ctx context.Context, request *GetChatRequest) (result *Chat, err error)
```

#### Request

```go
type GetChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### getChatAdministrators

**[Official documentation](https://core.telegram.org/bots/api#getchatadministrators)**

#### Description

Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) GetChatAdministrators(ctx context.Context, request *GetChatAdministratorsRequest) (result []*ChatMember, err error)
```

#### Request

```go
type GetChatAdministratorsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### getChatMember

**[Official documentation](https://core.telegram.org/bots/api#getchatmember)**

#### Description

Use this method to get information about a member of a chat. Returns a ChatMember object on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername) |
| user_id | Integer | Unique identifier of the target user |

#### Interface

```go
func (b *Bot) GetChatMember(ctx context.Context, request *GetChatMemberRequest) (result *ChatMember, err error)
```

#### Request

```go
type GetChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}
```

### getChatMembersCount

**[Official documentation](https://core.telegram.org/bots/api#getchatmemberscount)**

#### Description

Use this method to get the number of members in a chat. Returns Int on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) GetChatMembersCount(ctx context.Context, request *GetChatMembersCountRequest) (result int, err error)
```

#### Request

```go
type GetChatMembersCountRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### getFile

**[Official documentation](https://core.telegram.org/bots/api#getfile)**

#### Description

Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | File identifier to get info about |

#### Interface

```go
func (b *Bot) GetFile(ctx context.Context, request *GetFileRequest) (result *File, err error)
```

#### Request

```go
type GetFileRequest struct {
	// File identifier to get info about
	FileId string `json:"file_id"`
}
```

### getGameHighScores

**[Official documentation](https://core.telegram.org/bots/api#getgamehighscores)**

#### Description

Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | Target user id |
| chat_id | Integer | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the sent message |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |

#### Interface

```go
func (b *Bot) GetGameHighScores(ctx context.Context, request *GetGameHighScoresRequest) (result []*GameHighScore, err error)
```

#### Request

```go
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
```

### getMe

**[Official documentation](https://core.telegram.org/bots/api#getme)**

#### Description

A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
#### Interface

```go
func (b *Bot) GetMe(ctx context.Context) (result *User, err error)
```
### getStickerSet

**[Official documentation](https://core.telegram.org/bots/api#getstickerset)**

#### Description

Use this method to get a sticker set. On success, a StickerSet object is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| name | String | Name of the sticker set |

#### Interface

```go
func (b *Bot) GetStickerSet(ctx context.Context, request *GetStickerSetRequest) (result *StickerSet, err error)
```

#### Request

```go
type GetStickerSetRequest struct {
	// Name of the sticker set
	Name string `json:"name"`
}
```

### getUpdates

**[Official documentation](https://core.telegram.org/bots/api#getupdates)**

#### Description

Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| offset | Integer | Optional. Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten. |
| limit | Integer | Optional. Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100. |
| timeout | Integer | Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only. |
| allowed_updates | Array of String | Optional. List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time. |

#### Interface

```go
func (b *Bot) GetUpdates(ctx context.Context, request *GetUpdatesRequest) (result []*Update, err error)
```

#### Request

```go
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
```

### getUserProfilePhotos

**[Official documentation](https://core.telegram.org/bots/api#getuserprofilephotos)**

#### Description

Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | Unique identifier of the target user |
| offset | Integer | Optional. Sequential number of the first photo to be returned. By default, all photos are returned. |
| limit | Integer | Optional. Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100. |

#### Interface

```go
func (b *Bot) GetUserProfilePhotos(ctx context.Context, request *GetUserProfilePhotosRequest) (result interface{}, err error)
```

#### Request

```go
type GetUserProfilePhotosRequest struct {
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Optional. Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset int `json:"offset,omitempty"`
	// Optional. Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}
```

### getWebhookInfo

**[Official documentation](https://core.telegram.org/bots/api#getwebhookinfo)**

#### Description

Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
#### Interface

```go
func (b *Bot) GetWebhookInfo(ctx context.Context) (result *WebhookInfo, err error)
```
### kickChatMember

**[Official documentation](https://core.telegram.org/bots/api#kickchatmember)**

#### Description

Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername) |
| user_id | Integer | Unique identifier of the target user |
| until_date | Integer | Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever |

#### Interface

```go
func (b *Bot) KickChatMember(ctx context.Context, request *KickChatMemberRequest) (result bool, err error)
```

#### Request

```go
type KickChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Optional. Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
	UntilDate int `json:"until_date,omitempty"`
}
```

### leaveChat

**[Official documentation](https://core.telegram.org/bots/api#leavechat)**

#### Description

Use this method for your bot to leave a group, supergroup or channel. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) LeaveChat(ctx context.Context, request *LeaveChatRequest) (result bool, err error)
```

#### Request

```go
type LeaveChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### pinChatMessage

**[Official documentation](https://core.telegram.org/bots/api#pinchatmessage)**

#### Description

Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Identifier of a message to pin |
| disable_notification | Boolean | Optional. Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels. |

#### Interface

```go
func (b *Bot) PinChatMessage(ctx context.Context, request *PinChatMessageRequest) (result bool, err error)
```

#### Request

```go
type PinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Identifier of a message to pin
	MessageId int `json:"message_id"`
	// Optional. Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
	DisableNotification bool `json:"disable_notification,omitempty"`
}
```

### promoteChatMember

**[Official documentation](https://core.telegram.org/bots/api#promotechatmember)**

#### Description

Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| user_id | Integer | Unique identifier of the target user |
| can_change_info | Boolean | Optional. Pass True, if the administrator can change chat title, photo and other settings |
| can_post_messages | Boolean | Optional. Pass True, if the administrator can create channel posts, channels only |
| can_edit_messages | Boolean | Optional. Pass True, if the administrator can edit messages of other users and can pin messages, channels only |
| can_delete_messages | Boolean | Optional. Pass True, if the administrator can delete messages of other users |
| can_invite_users | Boolean | Optional. Pass True, if the administrator can invite new users to the chat |
| can_restrict_members | Boolean | Optional. Pass True, if the administrator can restrict, ban or unban chat members |
| can_pin_messages | Boolean | Optional. Pass True, if the administrator can pin messages, supergroups only |
| can_promote_members | Boolean | Optional. Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him) |

#### Interface

```go
func (b *Bot) PromoteChatMember(ctx context.Context, request *PromoteChatMemberRequest) (result bool, err error)
```

#### Request

```go
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
```

### restrictChatMember

**[Official documentation](https://core.telegram.org/bots/api#restrictchatmember)**

#### Description

Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername) |
| user_id | Integer | Unique identifier of the target user |
| until_date | Integer | Optional. Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever |
| can_send_messages | Boolean | Optional. Pass True, if the user can send text messages, contacts, locations and venues |
| can_send_media_messages | Boolean | Optional. Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages |
| can_send_other_messages | Boolean | Optional. Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages |
| can_add_web_page_previews | Boolean | Optional. Pass True, if the user may add web page previews to their messages, implies can_send_media_messages |

#### Interface

```go
func (b *Bot) RestrictChatMember(ctx context.Context, request *RestrictChatMemberRequest) (result bool, err error)
```

#### Request

```go
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
```

### sendAnimation

**[Official documentation](https://core.telegram.org/bots/api#sendanimation)**

#### Description

Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| animation | InputFile or String | Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files » |
| duration | Integer | Optional. Duration of sent animation in seconds |
| width | Integer | Optional. Animation width |
| height | Integer | Optional. Animation height |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Animation caption (may also be used when resending animation by file_id), 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendAnimation(ctx context.Context, request *SendAnimationRequest) (result *Message, err error)
```

#### Request

```go
type SendAnimationRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
	Animation interface{} `json:"animation"`
	// Optional. Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Animation width
	Width int `json:"width,omitempty"`
	// Optional. Animation height
	Height int `json:"height,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
	// Optional. Animation caption (may also be used when resending animation by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendAudio

**[Official documentation](https://core.telegram.org/bots/api#sendaudio)**

#### Description

Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| audio | InputFile or String | Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » |
| caption | String | Optional. Audio caption, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| duration | Integer | Optional. Duration of the audio in seconds |
| performer | String | Optional. Performer |
| title | String | Optional. Track name |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendAudio(ctx context.Context, request *SendAudioRequest) (result *Message, err error)
```

#### Request

```go
type SendAudioRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Audio interface{} `json:"audio"`
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
	Thumb interface{} `json:"thumb,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendChatAction

**[Official documentation](https://core.telegram.org/bots/api#sendchataction)**

#### Description

Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| action | String | Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes. |

#### Interface

```go
func (b *Bot) SendChatAction(ctx context.Context, request *SendChatActionRequest) (result bool, err error)
```

#### Request

```go
type SendChatActionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
	Action string `json:"action"`
}
```

### sendContact

**[Official documentation](https://core.telegram.org/bots/api#sendcontact)**

#### Description

Use this method to send phone contacts. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| phone_number | String | Contact's phone number |
| first_name | String | Contact's first name |
| last_name | String | Optional. Contact's last name |
| vcard | String | Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendContact(ctx context.Context, request *SendContactRequest) (result *Message, err error)
```

#### Request

```go
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendDocument

**[Official documentation](https://core.telegram.org/bots/api#senddocument)**

#### Description

Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| document | InputFile or String | File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Document caption (may also be used when resending documents by file_id), 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendDocument(ctx context.Context, request *SendDocumentRequest) (result *Message, err error)
```

#### Request

```go
type SendDocumentRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Document interface{} `json:"document"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
	// Optional. Document caption (may also be used when resending documents by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendGame

**[Official documentation](https://core.telegram.org/bots/api#sendgame)**

#### Description

Use this method to send a game. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer | Unique identifier for the target chat |
| game_short_name | String | Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather. |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game. |

#### Interface

```go
func (b *Bot) SendGame(ctx context.Context, request *SendGameRequest) (result *Message, err error)
```

#### Request

```go
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
```

### sendInvoice

**[Official documentation](https://core.telegram.org/bots/api#sendinvoice)**

#### Description

Use this method to send invoices. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer | Unique identifier for the target private chat |
| title | String | Product name, 1-32 characters |
| description | String | Product description, 1-255 characters |
| payload | String | Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes. |
| provider_token | String | Payments provider token, obtained via Botfather |
| start_parameter | String | Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter |
| currency | String | Three-letter ISO 4217 currency code, see more on currencies |
| prices | Array of LabeledPrice | Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.) |
| provider_data | String | Optional. JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider. |
| photo_url | String | Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for. |
| photo_size | Integer | Optional. Photo size |
| photo_width | Integer | Optional. Photo width |
| photo_height | Integer | Optional. Photo height |
| need_name | Boolean | Optional. Pass True, if you require the user's full name to complete the order |
| need_phone_number | Boolean | Optional. Pass True, if you require the user's phone number to complete the order |
| need_email | Boolean | Optional. Pass True, if you require the user's email address to complete the order |
| need_shipping_address | Boolean | Optional. Pass True, if you require the user's shipping address to complete the order |
| send_phone_number_to_provider | Boolean | Optional. Pass True, if user's phone number should be sent to provider |
| send_email_to_provider | Boolean | Optional. Pass True, if user's email address should be sent to provider |
| is_flexible | Boolean | Optional. Pass True, if the final price depends on the shipping method |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button. |

#### Interface

```go
func (b *Bot) SendInvoice(ctx context.Context, request *SendInvoiceRequest) (result *Message, err error)
```

#### Request

```go
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
```

### sendLocation

**[Official documentation](https://core.telegram.org/bots/api#sendlocation)**

#### Description

Use this method to send point on the map. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| latitude | Float number | Latitude of the location |
| longitude | Float number | Longitude of the location |
| live_period | Integer | Optional. Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400. |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendLocation(ctx context.Context, request *SendLocationRequest) (result *Message, err error)
```

#### Request

```go
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendMediaGroup

**[Official documentation](https://core.telegram.org/bots/api#sendmediagroup)**

#### Description

Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| media | Array of InputMediaPhoto and InputMediaVideo | A JSON-serialized array describing photos and videos to be sent, must include 2–10 items |
| disable_notification | Boolean | Optional. Sends the messages silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the messages are a reply, ID of the original message |

#### Interface

```go
func (b *Bot) SendMediaGroup(ctx context.Context, request *SendMediaGroupRequest) (result []*Message, err error)
```

#### Request

```go
type SendMediaGroupRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
	Media interface{} `json:"media"`
	// Optional. Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the messages are a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
}
```

### sendMessage

**[Official documentation](https://core.telegram.org/bots/api#sendmessage)**

#### Description

Use this method to send text messages. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| text | String | Text of the message to be sent |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message. |
| disable_web_page_preview | Boolean | Optional. Disables link previews for links in this message |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendMessage(ctx context.Context, request *SendMessageRequest) (result *Message, err error)
```

#### Request

```go
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendPhoto

**[Official documentation](https://core.telegram.org/bots/api#sendphoto)**

#### Description

Use this method to send photos. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| photo | InputFile or String | Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files » |
| caption | String | Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendPhoto(ctx context.Context, request *SendPhotoRequest) (result *Message, err error)
```

#### Request

```go
type SendPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files »
	Photo interface{} `json:"photo"`
	// Optional. Photo caption (may also be used when resending photos by file_id), 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendPoll

**[Official documentation](https://core.telegram.org/bots/api#sendpoll)**

#### Description

Use this method to send a native poll. A native poll can't be sent to a private chat. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername). A native poll can't be sent to a private chat. |
| question | String | Poll question, 1-255 characters |
| options | Array of String | List of answer options, 2-10 strings 1-100 characters each |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendPoll(ctx context.Context, request *SendPollRequest) (result *Message, err error)
```

#### Request

```go
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendSticker

**[Official documentation](https://core.telegram.org/bots/api#sendsticker)**

#### Description

Use this method to send .webp stickers. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| sticker | InputFile or String | Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendSticker(ctx context.Context, request *SendStickerRequest) (result *Message, err error)
```

#### Request

```go
type SendStickerRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Sticker interface{} `json:"sticker"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendVenue

**[Official documentation](https://core.telegram.org/bots/api#sendvenue)**

#### Description

Use this method to send information about a venue. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| latitude | Float number | Latitude of the venue |
| longitude | Float number | Longitude of the venue |
| title | String | Name of the venue |
| address | String | Address of the venue |
| foursquare_id | String | Optional. Foursquare identifier of the venue |
| foursquare_type | String | Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.) |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendVenue(ctx context.Context, request *SendVenueRequest) (result *Message, err error)
```

#### Request

```go
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendVideo

**[Official documentation](https://core.telegram.org/bots/api#sendvideo)**

#### Description

Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| video | InputFile or String | Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files » |
| duration | Integer | Optional. Duration of sent video in seconds |
| width | Integer | Optional. Video width |
| height | Integer | Optional. Video height |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Video caption (may also be used when resending videos by file_id), 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| supports_streaming | Boolean | Optional. Pass True, if the uploaded video is suitable for streaming |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendVideo(ctx context.Context, request *SendVideoRequest) (result *Message, err error)
```

#### Request

```go
type SendVideoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
	Video interface{} `json:"video"`
	// Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Video width
	Width int `json:"width,omitempty"`
	// Optional. Video height
	Height int `json:"height,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendVideoNote

**[Official documentation](https://core.telegram.org/bots/api#sendvideonote)**

#### Description

As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| video_note | InputFile or String | Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported |
| duration | Integer | Optional. Duration of sent video in seconds |
| length | Integer | Optional. Video width and height, i.e. diameter of the video message |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendVideoNote(ctx context.Context, request *SendVideoNoteRequest) (result *Message, err error)
```

#### Request

```go
type SendVideoNoteRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	VideoNote interface{} `json:"video_note"`
	// Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Video width and height, i.e. diameter of the video message
	Length int `json:"length,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### sendVoice

**[Official documentation](https://core.telegram.org/bots/api#sendvoice)**

#### Description

Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| voice | InputFile or String | Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files » |
| caption | String | Optional. Voice message caption, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| duration | Integer | Optional. Duration of the voice message in seconds |
| disable_notification | Boolean | Optional. Sends the message silently. Users will receive a notification with no sound. |
| reply_to_message_id | Integer | Optional. If the message is a reply, ID of the original message |
| reply_markup | InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply | Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user. |

#### Interface

```go
func (b *Bot) SendVoice(ctx context.Context, request *SendVoiceRequest) (result *Message, err error)
```

#### Request

```go
type SendVoiceRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Voice interface{} `json:"voice"`
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
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
```

### setChatDescription

**[Official documentation](https://core.telegram.org/bots/api#setchatdescription)**

#### Description

Use this method to change the description of a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| description | String | Optional. New chat description, 0-255 characters |

#### Interface

```go
func (b *Bot) SetChatDescription(ctx context.Context, request *SetChatDescriptionRequest) (result bool, err error)
```

#### Request

```go
type SetChatDescriptionRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Optional. New chat description, 0-255 characters
	Description string `json:"description,omitempty"`
}
```

### setChatPhoto

**[Official documentation](https://core.telegram.org/bots/api#setchatphoto)**

#### Description

Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| photo | InputFile | New chat photo, uploaded using multipart/form-data |

#### Interface

```go
func (b *Bot) SetChatPhoto(ctx context.Context, request *SetChatPhotoRequest) (result bool, err error)
```

#### Request

```go
type SetChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// New chat photo, uploaded using multipart/form-data
	Photo *InputFile `json:"photo"`
}
```

### setChatStickerSet

**[Official documentation](https://core.telegram.org/bots/api#setchatstickerset)**

#### Description

Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername) |
| sticker_set_name | String | Name of the sticker set to be set as the group sticker set |

#### Interface

```go
func (b *Bot) SetChatStickerSet(ctx context.Context, request *SetChatStickerSetRequest) (result bool, err error)
```

#### Request

```go
type SetChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId int `json:"chat_id"`
	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}
```

### setChatTitle

**[Official documentation](https://core.telegram.org/bots/api#setchattitle)**

#### Description

Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| title | String | New chat title, 1-255 characters |

#### Interface

```go
func (b *Bot) SetChatTitle(ctx context.Context, request *SetChatTitleRequest) (result bool, err error)
```

#### Request

```go
type SetChatTitleRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// New chat title, 1-255 characters
	Title string `json:"title"`
}
```

### setGameScore

**[Official documentation](https://core.telegram.org/bots/api#setgamescore)**

#### Description

Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | User identifier |
| score | Integer | New score, must be non-negative |
| force | Boolean | Optional. Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters |
| disable_edit_message | Boolean | Optional. Pass True, if the game message should not be automatically edited to include the current scoreboard |
| chat_id | Integer | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the sent message |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |

#### Interface

```go
func (b *Bot) SetGameScore(ctx context.Context, request *SetGameScoreRequest) (result *Message, err error)
```

#### Request

```go
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
```

### setPassportDataErrors

**[Official documentation](https://core.telegram.org/bots/api#setpassportdataerrors)**

#### Description

Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | User identifier |
| errors | Array of PassportElementError | A JSON-serialized array describing the errors |

#### Interface

```go
func (b *Bot) SetPassportDataErrors(ctx context.Context, request *SetPassportDataErrorsRequest) (result bool, err error)
```

#### Request

```go
type SetPassportDataErrorsRequest struct {
	// User identifier
	UserId int `json:"user_id"`
	// A JSON-serialized array describing the errors
	Errors []*PassportElementError `json:"errors"`
}
```

### setStickerPositionInSet

**[Official documentation](https://core.telegram.org/bots/api#setstickerpositioninset)**

#### Description

Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| sticker | String | File identifier of the sticker |
| position | Integer | New sticker position in the set, zero-based |

#### Interface

```go
func (b *Bot) SetStickerPositionInSet(ctx context.Context, request *SetStickerPositionInSetRequest) (result bool, err error)
```

#### Request

```go
type SetStickerPositionInSetRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker"`
	// New sticker position in the set, zero-based
	Position int `json:"position"`
}
```

### setWebhook

**[Official documentation](https://core.telegram.org/bots/api#setwebhook)**

#### Description

Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| url | String | HTTPS url to send updates to. Use an empty string to remove webhook integration |
| certificate | InputFile | Optional. Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details. |
| max_connections | Integer | Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput. |
| allowed_updates | Array of String | Optional. List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time. |

#### Interface

```go
func (b *Bot) SetWebhook(ctx context.Context, request *SetWebhookRequest) (result bool, err error)
```

#### Request

```go
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
```

### stopMessageLiveLocation

**[Official documentation](https://core.telegram.org/bots/api#stopmessagelivelocation)**

#### Description

Use this method to stop updating a live location message before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Optional. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Optional. Required if inline_message_id is not specified. Identifier of the message with live location to stop |
| inline_message_id | String | Optional. Required if chat_id and message_id are not specified. Identifier of the inline message |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for a new inline keyboard. |

#### Interface

```go
func (b *Bot) StopMessageLiveLocation(ctx context.Context, request *StopMessageLiveLocationRequest) (result *Message, err error)
```

#### Request

```go
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
```

### stopPoll

**[Official documentation](https://core.telegram.org/bots/api#stoppoll)**

#### Description

Use this method to stop a poll which was sent by the bot. On success, the stopped Poll with the final results is returned.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |
| message_id | Integer | Identifier of the original message with the poll |
| reply_markup | InlineKeyboardMarkup | Optional. A JSON-serialized object for a new message inline keyboard. |

#### Interface

```go
func (b *Bot) StopPoll(ctx context.Context, request *StopPollRequest) (result *Poll, err error)
```

#### Request

```go
type StopPollRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Identifier of the original message with the poll
	MessageId int `json:"message_id"`
	// Optional. A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
```

### unbanChatMember

**[Official documentation](https://core.telegram.org/bots/api#unbanchatmember)**

#### Description

Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target group or username of the target supergroup or channel (in the format @username) |
| user_id | Integer | Unique identifier of the target user |

#### Interface

```go
func (b *Bot) UnbanChatMember(ctx context.Context, request *UnbanChatMemberRequest) (result bool, err error)
```

#### Request

```go
type UnbanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	ChatId int `json:"chat_id"`
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}
```

### unpinChatMessage

**[Official documentation](https://core.telegram.org/bots/api#unpinchatmessage)**

#### Description

Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| chat_id | Integer or String | Unique identifier for the target chat or username of the target channel (in the format @channelusername) |

#### Interface

```go
func (b *Bot) UnpinChatMessage(ctx context.Context, request *UnpinChatMessageRequest) (result bool, err error)
```

#### Request

```go
type UnpinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
}
```

### uploadStickerFile

**[Official documentation](https://core.telegram.org/bots/api#uploadstickerfile)**

#### Description

Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.

#### Parameters

| Field | Type | Description |
| ----- | ---- | ----------- |
| user_id | Integer | User identifier of sticker file owner |
| png_sticker | InputFile | Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files » |

#### Interface

```go
func (b *Bot) UploadStickerFile(ctx context.Context, request *UploadStickerFileRequest) (result *File, err error)
```

#### Request

```go
type UploadStickerFileRequest struct {
	// User identifier of sticker file owner
	UserId int `json:"user_id"`
	// Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
	PngSticker *InputFile `json:"png_sticker"`
}
```


## Available types

### Animation

**[Official documentation](https://core.telegram.org/bots/api#animation)**

#### Description

This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique file identifier |
| width | Integer | Video width as defined by sender |
| height | Integer | Video height as defined by sender |
| duration | Integer | Duration of the video in seconds as defined by sender |
| thumb | PhotoSize | Optional. Animation thumbnail as defined by sender |
| file_name | String | Optional. Original animation filename as defined by sender |
| mime_type | String | Optional. MIME type of the file as defined by sender |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### Audio

**[Official documentation](https://core.telegram.org/bots/api#audio)**

#### Description

This object represents an audio file to be treated as music by the Telegram clients.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| duration | Integer | Duration of the audio in seconds as defined by sender |
| performer | String | Optional. Performer of the audio as defined by sender or by audio tags |
| title | String | Optional. Title of the audio as defined by sender or by audio tags |
| mime_type | String | Optional. MIME type of the file as defined by sender |
| file_size | Integer | Optional. File size |
| thumb | PhotoSize | Optional. Thumbnail of the album cover to which the music file belongs |

#### Interface

```go
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
```

### CallbackGame

**[Official documentation](https://core.telegram.org/bots/api#callbackgame)**

#### Description

A placeholder, currently holds no information. Use BotFather to set up your game.
#### Interface

```go
type CallbackGame struct {
}
```

### CallbackQuery

**[Official documentation](https://core.telegram.org/bots/api#callbackquery)**

#### Description

This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | String | Unique identifier for this query |
| from | User | Sender |
| message | Message | Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old |
| inline_message_id | String | Optional. Identifier of the message sent via the bot in inline mode, that originated the query. |
| chat_instance | String | Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games. |
| data | String | Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field. |
| game_short_name | String | Optional. Short name of a Game to be returned, serves as the unique identifier for the game |

#### Interface

```go
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
```

### Chat

**[Official documentation](https://core.telegram.org/bots/api#chat)**

#### Description

This object represents a chat.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | Integer | Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. |
| type | String | Type of chat, can be either “private”, “group”, “supergroup” or “channel” |
| title | String | Optional. Title, for supergroups, channels and group chats |
| username | String | Optional. Username, for private chats, supergroups and channels if available |
| first_name | String | Optional. First name of the other party in a private chat |
| last_name | String | Optional. Last name of the other party in a private chat |
| all_members_are_administrators | Boolean | Optional. True if a group has ‘All Members Are Admins’ enabled. |
| photo | ChatPhoto | Optional. Chat photo. Returned only in getChat. |
| description | String | Optional. Description, for supergroups and channel chats. Returned only in getChat. |
| invite_link | String | Optional. Chat invite link, for supergroups and channel chats. Each administrator in a chat generates their own invite links, so the bot must first generate the link using exportChatInviteLink. Returned only in getChat. |
| pinned_message | Message | Optional. Pinned message, for groups, supergroups and channels. Returned only in getChat. |
| sticker_set_name | String | Optional. For supergroups, name of group sticker set. Returned only in getChat. |
| can_set_sticker_set | Boolean | Optional. True, if the bot can change the group sticker set. Returned only in getChat. |

#### Interface

```go
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
```

### ChatMember

**[Official documentation](https://core.telegram.org/bots/api#chatmember)**

#### Description

This object contains information about one member of a chat.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| user | User | Information about the user |
| status | String | The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked” |
| until_date | Integer | Optional. Restricted and kicked only. Date when restrictions will be lifted for this user, unix time |
| can_be_edited | Boolean | Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user |
| can_change_info | Boolean | Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings |
| can_post_messages | Boolean | Optional. Administrators only. True, if the administrator can post in the channel, channels only |
| can_edit_messages | Boolean | Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only |
| can_delete_messages | Boolean | Optional. Administrators only. True, if the administrator can delete messages of other users |
| can_invite_users | Boolean | Optional. Administrators only. True, if the administrator can invite new users to the chat |
| can_restrict_members | Boolean | Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members |
| can_pin_messages | Boolean | Optional. Administrators only. True, if the administrator can pin messages, groups and supergroups only |
| can_promote_members | Boolean | Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user) |
| is_member | Boolean | Optional. Restricted only. True, if the user is a member of the chat at the moment of the request |
| can_send_messages | Boolean | Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues |
| can_send_media_messages | Boolean | Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages |
| can_send_other_messages | Boolean | Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages |
| can_add_web_page_previews | Boolean | Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages |

#### Interface

```go
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
```

### ChatPhoto

**[Official documentation](https://core.telegram.org/bots/api#chatphoto)**

#### Description

This object represents a chat photo.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| small_file_id | String | Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download. |
| big_file_id | String | Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download. |

#### Interface

```go
type ChatPhoto struct {
	// Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
	SmallFileId string `json:"small_file_id"`
	// Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
	BigFileId string `json:"big_file_id"`
}
```

### ChosenInlineResult

**[Official documentation](https://core.telegram.org/bots/api#choseninlineresult)**

#### Description

Represents a result of an inline query that was chosen by the user and sent to their chat partner.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| result_id | String | The unique identifier for the result that was chosen |
| from | User | The user that chose the result |
| location | Location | Optional. Sender location, only for bots that require user location |
| inline_message_id | String | Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message. |
| query | String | The query that was used to obtain the result |

#### Interface

```go
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
```

### Contact

**[Official documentation](https://core.telegram.org/bots/api#contact)**

#### Description

This object represents a phone contact.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| phone_number | String | Contact's phone number |
| first_name | String | Contact's first name |
| last_name | String | Optional. Contact's last name |
| user_id | Integer | Optional. Contact's user identifier in Telegram |
| vcard | String | Optional. Additional data about the contact in the form of a vCard |

#### Interface

```go
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
```

### Document

**[Official documentation](https://core.telegram.org/bots/api#document)**

#### Description

This object represents a general file (as opposed to photos, voice messages and audio files).

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique file identifier |
| thumb | PhotoSize | Optional. Document thumbnail as defined by sender |
| file_name | String | Optional. Original filename as defined by sender |
| mime_type | String | Optional. MIME type of the file as defined by sender |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### EncryptedCredentials

**[Official documentation](https://core.telegram.org/bots/api#encryptedcredentials)**

#### Description

Contains data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| data | String | Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication |
| hash | String | Base64-encoded data hash for data authentication |
| secret | String | Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption |

#### Interface

```go
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}
```

### EncryptedPassportElement

**[Official documentation](https://core.telegram.org/bots/api#encryptedpassportelement)**

#### Description

Contains information about documents or other Telegram Passport elements shared with the bot by the user.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”. |
| data | String | Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials. |
| phone_number | String | Optional. User's verified phone number, available only for “phone_number” type |
| email | String | Optional. User's verified email address, available only for “email” type |
| files | Array of PassportFile | Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials. |
| front_side | PassportFile | Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials. |
| reverse_side | PassportFile | Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials. |
| selfie | PassportFile | Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials. |
| translation | Array of PassportFile | Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials. |
| hash | String | Base64-encoded element hash for using in PassportElementErrorUnspecified |

#### Interface

```go
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
```

### File

**[Official documentation](https://core.telegram.org/bots/api#file)**

#### Description

This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| file_size | Integer | Optional. File size, if known |
| file_path | String | Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file. |

#### Interface

```go
type File struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Optional. File size, if known
	FileSize int `json:"file_size,omitempty"`
	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`
}
```

### ForceReply

**[Official documentation](https://core.telegram.org/bots/api#forcereply)**

#### Description

Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| force_reply | True | Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply' |
| selective | Boolean | Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. |

#### Interface

```go
type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	ForceReply bool `json:"force_reply"`
	// Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}
```

### Game

**[Official documentation](https://core.telegram.org/bots/api#game)**

#### Description

This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| title | String | Title of the game |
| description | String | Description of the game |
| photo | Array of PhotoSize | Photo that will be displayed in the game message in chats. |
| text | String | Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters. |
| text_entities | Array of MessageEntity | Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc. |
| animation | Animation | Optional. Animation that will be displayed in the game message in chats. Upload via BotFather |

#### Interface

```go
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
```

### GameHighScore

**[Official documentation](https://core.telegram.org/bots/api#gamehighscore)**

#### Description

This object represents one row of the high scores table for a game.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| position | Integer | Position in high score table for the game |
| user | User | User |
| score | Integer | Score |

#### Interface

```go
type GameHighScore struct {
	// Position in high score table for the game
	Position int `json:"position"`
	// User
	User *User `json:"user"`
	// Score
	Score int `json:"score"`
}
```

### InlineKeyboardButton

**[Official documentation](https://core.telegram.org/bots/api#inlinekeyboardbutton)**

#### Description

This object represents one button of an inline keyboard. You must use exactly one of the optional fields.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| text | String | Label text on the button |
| url | String | Optional. HTTP or tg:// url to be opened when button is pressed |
| login_url | LoginUrl | Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget. |
| callback_data | String | Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes |
| switch_inline_query | String | Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen. |
| switch_inline_query_current_chat | String | Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options. |
| callback_game | CallbackGame | Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row. |
| pay | Boolean | Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row. |

#### Interface

```go
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
```

### InlineKeyboardMarkup

**[Official documentation](https://core.telegram.org/bots/api#inlinekeyboardmarkup)**

#### Description

This object represents an inline keyboard that appears right next to the message it belongs to.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| inline_keyboard | Array of Array of InlineKeyboardButton | Array of button rows, each represented by an Array of InlineKeyboardButton objects |

#### Interface

```go
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}
```

### InlineQuery

**[Official documentation](https://core.telegram.org/bots/api#inlinequery)**

#### Description

This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | String | Unique identifier for this query |
| from | User | Sender |
| location | Location | Optional. Sender location, only for bots that request user location |
| query | String | Text of the query (up to 512 characters) |
| offset | String | Offset of the results to be returned, can be controlled by the bot |

#### Interface

```go
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
```

### InlineQueryResult

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresult)**

#### Description

This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
#### Interface

```go
type InlineQueryResult struct {
}
```

### InlineQueryResultArticle

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultarticle)**

#### Description

Represents a link to an article or web page.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be article |
| id | String | Unique identifier for this result, 1-64 Bytes |
| title | String | Title of the result |
| input_message_content | InputMessageContent | Content of the message to be sent |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| url | String | Optional. URL of the result |
| hide_url | Boolean | Optional. Pass True, if you don't want the URL to be shown in the message |
| description | String | Optional. Short description of the result |
| thumb_url | String | Optional. Url of the thumbnail for the result |
| thumb_width | Integer | Optional. Thumbnail width |
| thumb_height | Integer | Optional. Thumbnail height |

#### Interface

```go
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
```

### InlineQueryResultAudio

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultaudio)**

#### Description

Represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be audio |
| id | String | Unique identifier for this result, 1-64 bytes |
| audio_url | String | A valid URL for the audio file |
| title | String | Title |
| caption | String | Optional. Caption, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| performer | String | Optional. Performer |
| audio_duration | Integer | Optional. Audio duration in seconds |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the audio |

#### Interface

```go
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
```

### InlineQueryResultCachedAudio

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedaudio)**

#### Description

Represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be audio |
| id | String | Unique identifier for this result, 1-64 bytes |
| audio_file_id | String | A valid file identifier for the audio file |
| caption | String | Optional. Caption, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the audio |

#### Interface

```go
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
```

### InlineQueryResultCachedDocument

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcacheddocument)**

#### Description

Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be document |
| id | String | Unique identifier for this result, 1-64 bytes |
| title | String | Title for the result |
| document_file_id | String | A valid file identifier for the file |
| description | String | Optional. Short description of the result |
| caption | String | Optional. Caption of the document to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the file |

#### Interface

```go
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
```

### InlineQueryResultCachedGif

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedgif)**

#### Description

Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be gif |
| id | String | Unique identifier for this result, 1-64 bytes |
| gif_file_id | String | A valid file identifier for the GIF file |
| title | String | Optional. Title for the result |
| caption | String | Optional. Caption of the GIF file to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the GIF animation |

#### Interface

```go
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
```

### InlineQueryResultCachedMpeg4Gif

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif)**

#### Description

Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be mpeg4_gif |
| id | String | Unique identifier for this result, 1-64 bytes |
| mpeg4_file_id | String | A valid file identifier for the MP4 file |
| title | String | Optional. Title for the result |
| caption | String | Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the video animation |

#### Interface

```go
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
```

### InlineQueryResultCachedPhoto

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedphoto)**

#### Description

Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be photo |
| id | String | Unique identifier for this result, 1-64 bytes |
| photo_file_id | String | A valid file identifier of the photo |
| title | String | Optional. Title for the result |
| description | String | Optional. Short description of the result |
| caption | String | Optional. Caption of the photo to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the photo |

#### Interface

```go
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
```

### InlineQueryResultCachedSticker

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedsticker)**

#### Description

Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be sticker |
| id | String | Unique identifier for this result, 1-64 bytes |
| sticker_file_id | String | A valid file identifier of the sticker |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the sticker |

#### Interface

```go
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
```

### InlineQueryResultCachedVideo

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedvideo)**

#### Description

Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be video |
| id | String | Unique identifier for this result, 1-64 bytes |
| video_file_id | String | A valid file identifier for the video file |
| title | String | Title for the result |
| description | String | Optional. Short description of the result |
| caption | String | Optional. Caption of the video to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the video |

#### Interface

```go
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
```

### InlineQueryResultCachedVoice

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcachedvoice)**

#### Description

Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be voice |
| id | String | Unique identifier for this result, 1-64 bytes |
| voice_file_id | String | A valid file identifier for the voice message |
| title | String | Voice message title |
| caption | String | Optional. Caption, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the voice message |

#### Interface

```go
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
```

### InlineQueryResultContact

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultcontact)**

#### Description

Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be contact |
| id | String | Unique identifier for this result, 1-64 Bytes |
| phone_number | String | Contact's phone number |
| first_name | String | Contact's first name |
| last_name | String | Optional. Contact's last name |
| vcard | String | Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the contact |
| thumb_url | String | Optional. Url of the thumbnail for the result |
| thumb_width | Integer | Optional. Thumbnail width |
| thumb_height | Integer | Optional. Thumbnail height |

#### Interface

```go
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
```

### InlineQueryResultDocument

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultdocument)**

#### Description

Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be document |
| id | String | Unique identifier for this result, 1-64 bytes |
| title | String | Title for the result |
| caption | String | Optional. Caption of the document to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| document_url | String | A valid URL for the file |
| mime_type | String | Mime type of the content of the file, either “application/pdf” or “application/zip” |
| description | String | Optional. Short description of the result |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the file |
| thumb_url | String | Optional. URL of the thumbnail (jpeg only) for the file |
| thumb_width | Integer | Optional. Thumbnail width |
| thumb_height | Integer | Optional. Thumbnail height |

#### Interface

```go
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
```

### InlineQueryResultGame

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultgame)**

#### Description

Represents a Game.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be game |
| id | String | Unique identifier for this result, 1-64 bytes |
| game_short_name | String | Short name of the game |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |

#### Interface

```go
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
```

### InlineQueryResultGif

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultgif)**

#### Description

Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be gif |
| id | String | Unique identifier for this result, 1-64 bytes |
| gif_url | String | A valid URL for the GIF file. File size must not exceed 1MB |
| gif_width | Integer | Optional. Width of the GIF |
| gif_height | Integer | Optional. Height of the GIF |
| gif_duration | Integer | Optional. Duration of the GIF |
| thumb_url | String | URL of the static thumbnail for the result (jpeg or gif) |
| title | String | Optional. Title for the result |
| caption | String | Optional. Caption of the GIF file to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the GIF animation |

#### Interface

```go
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
```

### InlineQueryResultLocation

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultlocation)**

#### Description

Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be location |
| id | String | Unique identifier for this result, 1-64 Bytes |
| latitude | Float number | Location latitude in degrees |
| longitude | Float number | Location longitude in degrees |
| title | String | Location title |
| live_period | Integer | Optional. Period in seconds for which the location can be updated, should be between 60 and 86400. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the location |
| thumb_url | String | Optional. Url of the thumbnail for the result |
| thumb_width | Integer | Optional. Thumbnail width |
| thumb_height | Integer | Optional. Thumbnail height |

#### Interface

```go
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
```

### InlineQueryResultMpeg4Gif

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif)**

#### Description

Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be mpeg4_gif |
| id | String | Unique identifier for this result, 1-64 bytes |
| mpeg4_url | String | A valid URL for the MP4 file. File size must not exceed 1MB |
| mpeg4_width | Integer | Optional. Video width |
| mpeg4_height | Integer | Optional. Video height |
| mpeg4_duration | Integer | Optional. Video duration |
| thumb_url | String | URL of the static thumbnail (jpeg or gif) for the result |
| title | String | Optional. Title for the result |
| caption | String | Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the video animation |

#### Interface

```go
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
```

### InlineQueryResultPhoto

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultphoto)**

#### Description

Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be photo |
| id | String | Unique identifier for this result, 1-64 bytes |
| photo_url | String | A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB |
| thumb_url | String | URL of the thumbnail for the photo |
| photo_width | Integer | Optional. Width of the photo |
| photo_height | Integer | Optional. Height of the photo |
| title | String | Optional. Title for the result |
| description | String | Optional. Short description of the result |
| caption | String | Optional. Caption of the photo to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the photo |

#### Interface

```go
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
```

### InlineQueryResultVenue

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultvenue)**

#### Description

Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be venue |
| id | String | Unique identifier for this result, 1-64 Bytes |
| latitude | Float | Latitude of the venue location in degrees |
| longitude | Float | Longitude of the venue location in degrees |
| title | String | Title of the venue |
| address | String | Address of the venue |
| foursquare_id | String | Optional. Foursquare identifier of the venue if known |
| foursquare_type | String | Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.) |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the venue |
| thumb_url | String | Optional. Url of the thumbnail for the result |
| thumb_width | Integer | Optional. Thumbnail width |
| thumb_height | Integer | Optional. Thumbnail height |

#### Interface

```go
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
```

### InlineQueryResultVideo

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultvideo)**

#### Description

Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be video |
| id | String | Unique identifier for this result, 1-64 bytes |
| video_url | String | A valid URL for the embedded video player or video file |
| mime_type | String | Mime type of the content of video url, “text/html” or “video/mp4” |
| thumb_url | String | URL of the thumbnail (jpeg only) for the video |
| title | String | Title for the result |
| caption | String | Optional. Caption of the video to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| video_width | Integer | Optional. Video width |
| video_height | Integer | Optional. Video height |
| video_duration | Integer | Optional. Video duration in seconds |
| description | String | Optional. Short description of the result |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video). |

#### Interface

```go
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
```

### InlineQueryResultVoice

**[Official documentation](https://core.telegram.org/bots/api#inlinequeryresultvoice)**

#### Description

Represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be voice |
| id | String | Unique identifier for this result, 1-64 bytes |
| voice_url | String | A valid URL for the voice recording |
| title | String | Recording title |
| caption | String | Optional. Caption, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| voice_duration | Integer | Optional. Recording duration in seconds |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message |
| input_message_content | InputMessageContent | Optional. Content of the message to be sent instead of the voice recording |

#### Interface

```go
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
```

### InputContactMessageContent

**[Official documentation](https://core.telegram.org/bots/api#inputcontactmessagecontent)**

#### Description

Represents the content of a contact message to be sent as the result of an inline query.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| phone_number | String | Contact's phone number |
| first_name | String | Contact's first name |
| last_name | String | Optional. Contact's last name |
| vcard | String | Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes |

#### Interface

```go
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
```

### InputFile

**[Official documentation](https://core.telegram.org/bots/api#inputfile)**

#### Description

This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
#### Interface

```go
type InputFile struct {
}
```

### InputLocationMessageContent

**[Official documentation](https://core.telegram.org/bots/api#inputlocationmessagecontent)**

#### Description

Represents the content of a location message to be sent as the result of an inline query.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| latitude | Float | Latitude of the location in degrees |
| longitude | Float | Longitude of the location in degrees |
| live_period | Integer | Optional. Period in seconds for which the location can be updated, should be between 60 and 86400. |

#### Interface

```go
type InputLocationMessageContent struct {
	// Latitude of the location in degrees
	Latitude float64 `json:"latitude"`
	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`
	// Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
}
```

### InputMedia

**[Official documentation](https://core.telegram.org/bots/api#inputmedia)**

#### Description

This object represents the content of a media message to be sent. It should be one of
#### Interface

```go
type InputMedia struct {
}
```

### InputMediaAnimation

**[Official documentation](https://core.telegram.org/bots/api#inputmediaanimation)**

#### Description

Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be animation |
| media | String | File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Caption of the animation to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| width | Integer | Optional. Animation width |
| height | Integer | Optional. Animation height |
| duration | Integer | Optional. Animation duration |

#### Interface

```go
type InputMediaAnimation struct {
	// Type of the result, must be animation
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
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
```

### InputMediaAudio

**[Official documentation](https://core.telegram.org/bots/api#inputmediaaudio)**

#### Description

Represents an audio file to be treated as music to be sent.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be audio |
| media | String | File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Caption of the audio to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| duration | Integer | Optional. Duration of the audio in seconds |
| performer | String | Optional. Performer of the audio |
| title | String | Optional. Title of the audio |

#### Interface

```go
type InputMediaAudio struct {
	// Type of the result, must be audio
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
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
```

### InputMediaDocument

**[Official documentation](https://core.telegram.org/bots/api#inputmediadocument)**

#### Description

Represents a general file to be sent.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be document |
| media | String | File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Caption of the document to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |

#### Interface

```go
type InputMediaDocument struct {
	// Type of the result, must be document
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
	// Optional. Caption of the document to be sent, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ParseMode string `json:"parse_mode,omitempty"`
}
```

### InputMediaPhoto

**[Official documentation](https://core.telegram.org/bots/api#inputmediaphoto)**

#### Description

Represents a photo to be sent.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be photo |
| media | String | File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » |
| caption | String | Optional. Caption of the photo to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |

#### Interface

```go
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
```

### InputMediaVideo

**[Official documentation](https://core.telegram.org/bots/api#inputmediavideo)**

#### Description

Represents a video to be sent.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the result, must be video |
| media | String | File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files » |
| thumb | InputFile or String | Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files » |
| caption | String | Optional. Caption of the video to be sent, 0-1024 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption. |
| width | Integer | Optional. Video width |
| height | Integer | Optional. Video height |
| duration | Integer | Optional. Video duration |
| supports_streaming | Boolean | Optional. Pass True, if the uploaded video is suitable for streaming |

#### Interface

```go
type InputMediaVideo struct {
	// Type of the result, must be video
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Media string `json:"media"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Thumb interface{} `json:"thumb,omitempty"`
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
```

### InputMessageContent

**[Official documentation](https://core.telegram.org/bots/api#inputmessagecontent)**

#### Description

This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 4 types:
#### Interface

```go
type InputMessageContent struct {
}
```

### InputTextMessageContent

**[Official documentation](https://core.telegram.org/bots/api#inputtextmessagecontent)**

#### Description

Represents the content of a text message to be sent as the result of an inline query.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| message_text | String | Text of the message to be sent, 1-4096 characters |
| parse_mode | String | Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message. |
| disable_web_page_preview | Boolean | Optional. Disables link previews for links in the sent message |

#### Interface

```go
type InputTextMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	ParseMode string `json:"parse_mode,omitempty"`
	// Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}
```

### InputVenueMessageContent

**[Official documentation](https://core.telegram.org/bots/api#inputvenuemessagecontent)**

#### Description

Represents the content of a venue message to be sent as the result of an inline query.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| latitude | Float | Latitude of the venue in degrees |
| longitude | Float | Longitude of the venue in degrees |
| title | String | Name of the venue |
| address | String | Address of the venue |
| foursquare_id | String | Optional. Foursquare identifier of the venue, if known |
| foursquare_type | String | Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.) |

#### Interface

```go
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
```

### Invoice

**[Official documentation](https://core.telegram.org/bots/api#invoice)**

#### Description

This object contains basic information about an invoice.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| title | String | Product name |
| description | String | Product description |
| start_parameter | String | Unique bot deep-linking parameter that can be used to generate this invoice |
| currency | String | Three-letter ISO 4217 currency code |
| total_amount | Integer | Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). |

#### Interface

```go
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
```

### KeyboardButton

**[Official documentation](https://core.telegram.org/bots/api#keyboardbutton)**

#### Description

This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| text | String | Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed |
| request_contact | Boolean | Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only |
| request_location | Boolean | Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only |

#### Interface

```go
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	Text string `json:"text"`
	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`
	// Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`
}
```

### LabeledPrice

**[Official documentation](https://core.telegram.org/bots/api#labeledprice)**

#### Description

This object represents a portion of the price for goods or services.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| label | String | Portion label |
| amount | Integer | Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). |

#### Interface

```go
type LabeledPrice struct {
	// Portion label
	Label string `json:"label"`
	// Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}
```

### Location

**[Official documentation](https://core.telegram.org/bots/api#location)**

#### Description

This object represents a point on the map.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| longitude | Float | Longitude as defined by sender |
| latitude | Float | Latitude as defined by sender |

#### Interface

```go
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
}
```

### LoginUrl

**[Official documentation](https://core.telegram.org/bots/api#loginurl)**

#### Description

This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| url | String | An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization. |
| forward_text | String | Optional. New text of the button in forwarded messages. |
| bot_username | String | Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details. |
| request_write_access | Boolean | Optional. Pass True to request the permission for your bot to send messages to the user. |

#### Interface

```go
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
```

### MaskPosition

**[Official documentation](https://core.telegram.org/bots/api#maskposition)**

#### Description

This object describes the position on faces where a mask should be placed by default.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| point | String | The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”. |
| x_shift | Float number | Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position. |
| y_shift | Float number | Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position. |
| scale | Float number | Mask scaling coefficient. For example, 2.0 means double size. |

#### Interface

```go
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
```

### Message

**[Official documentation](https://core.telegram.org/bots/api#message)**

#### Description

This object represents a message.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| message_id | Integer | Unique message identifier inside this chat |
| from | User | Optional. Sender, empty for messages sent to channels |
| date | Integer | Date the message was sent in Unix time |
| chat | Chat | Conversation the message belongs to |
| forward_from | User | Optional. For forwarded messages, sender of the original message |
| forward_from_chat | Chat | Optional. For messages forwarded from channels, information about the original channel |
| forward_from_message_id | Integer | Optional. For messages forwarded from channels, identifier of the original message in the channel |
| forward_signature | String | Optional. For messages forwarded from channels, signature of the post author if present |
| forward_sender_name | String | Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages |
| forward_date | Integer | Optional. For forwarded messages, date the original message was sent in Unix time |
| reply_to_message | Message | Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply. |
| edit_date | Integer | Optional. Date the message was last edited in Unix time |
| media_group_id | String | Optional. The unique identifier of a media message group this message belongs to |
| author_signature | String | Optional. Signature of the post author for messages in channels |
| text | String | Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters. |
| entities | Array of MessageEntity | Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text |
| caption_entities | Array of MessageEntity | Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption |
| audio | Audio | Optional. Message is an audio file, information about the file |
| document | Document | Optional. Message is a general file, information about the file |
| animation | Animation | Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set |
| game | Game | Optional. Message is a game, information about the game. More about games » |
| photo | Array of PhotoSize | Optional. Message is a photo, available sizes of the photo |
| sticker | Sticker | Optional. Message is a sticker, information about the sticker |
| video | Video | Optional. Message is a video, information about the video |
| voice | Voice | Optional. Message is a voice message, information about the file |
| video_note | VideoNote | Optional. Message is a video note, information about the video message |
| caption | String | Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters |
| contact | Contact | Optional. Message is a shared contact, information about the contact |
| location | Location | Optional. Message is a shared location, information about the location |
| venue | Venue | Optional. Message is a venue, information about the venue |
| poll | Poll | Optional. Message is a native poll, information about the poll |
| new_chat_members | Array of User | Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members) |
| left_chat_member | User | Optional. A member was removed from the group, information about them (this member may be the bot itself) |
| new_chat_title | String | Optional. A chat title was changed to this value |
| new_chat_photo | Array of PhotoSize | Optional. A chat photo was change to this value |
| delete_chat_photo | True | Optional. Service message: the chat photo was deleted |
| group_chat_created | True | Optional. Service message: the group has been created |
| supergroup_chat_created | True | Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup. |
| channel_chat_created | True | Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel. |
| migrate_to_chat_id | Integer | Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. |
| migrate_from_chat_id | Integer | Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. |
| pinned_message | Message | Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply. |
| invoice | Invoice | Optional. Message is an invoice for a payment, information about the invoice. More about payments » |
| successful_payment | SuccessfulPayment | Optional. Message is a service message about a successful payment, information about the payment. More about payments » |
| connected_website | String | Optional. The domain name of the website on which the user has logged in. More about Telegram Login » |
| passport_data | PassportData | Optional. Telegram Passport data |
| reply_markup | InlineKeyboardMarkup | Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons. |

#### Interface

```go
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
```

### MessageEntity

**[Official documentation](https://core.telegram.org/bots/api#messageentity)**

#### Description

This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| type | String | Type of the entity. Can be mention (@username), hashtag, cashtag, bot_command, url, email, phone_number, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames) |
| offset | Integer | Offset in UTF-16 code units to the start of the entity |
| length | Integer | Length of the entity in UTF-16 code units |
| url | String | Optional. For “text_link” only, url that will be opened after user taps on the text |
| user | User | Optional. For “text_mention” only, the mentioned user |

#### Interface

```go
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
```

### OrderInfo

**[Official documentation](https://core.telegram.org/bots/api#orderinfo)**

#### Description

This object represents information about an order.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| name | String | Optional. User name |
| phone_number | String | Optional. User's phone number |
| email | String | Optional. User email |
| shipping_address | ShippingAddress | Optional. User shipping address |

#### Interface

```go
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
```

### PassportData

**[Official documentation](https://core.telegram.org/bots/api#passportdata)**

#### Description

Contains information about Telegram Passport data shared with the bot by the user.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| data | Array of EncryptedPassportElement | Array with information about documents and other Telegram Passport elements that was shared with the bot |
| credentials | EncryptedCredentials | Encrypted credentials required to decrypt the data |

#### Interface

```go
type PassportData struct {
	// Array with information about documents and other Telegram Passport elements that was shared with the bot
	Data []*EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}
```

### PassportElementError

**[Official documentation](https://core.telegram.org/bots/api#passportelementerror)**

#### Description

This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
#### Interface

```go
type PassportElementError struct {
}
```

### PassportElementErrorDataField

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrordatafield)**

#### Description

Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be data |
| type | String | The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address” |
| field_name | String | Name of the data field which has the error |
| data_hash | String | Base64-encoded data hash |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorFile

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrorfile)**

#### Description

Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be file |
| type | String | The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration” |
| file_hash | String | Base64-encoded file hash |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorFiles

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrorfiles)**

#### Description

Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be files |
| type | String | The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration” |
| file_hashes | Array of String | List of base64-encoded file hashes |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorFrontSide

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrorfrontside)**

#### Description

Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be front_side |
| type | String | The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport” |
| file_hash | String | Base64-encoded hash of the file with the front side of the document |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorReverseSide

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrorreverseside)**

#### Description

Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be reverse_side |
| type | String | The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card” |
| file_hash | String | Base64-encoded hash of the file with the reverse side of the document |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorSelfie

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrorselfie)**

#### Description

Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be selfie |
| type | String | The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport” |
| file_hash | String | Base64-encoded hash of the file with the selfie |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorTranslationFile

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrortranslationfile)**

#### Description

Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be translation_file |
| type | String | Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration” |
| file_hash | String | Base64-encoded file hash |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorTranslationFiles

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrortranslationfiles)**

#### Description

Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be translation_files |
| type | String | Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration” |
| file_hashes | Array of String | List of base64-encoded file hashes |
| message | String | Error message |

#### Interface

```go
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
```

### PassportElementErrorUnspecified

**[Official documentation](https://core.telegram.org/bots/api#passportelementerrorunspecified)**

#### Description

Represents an issue in an unspecified place. The error is considered resolved when new data is added.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| source | String | Error source, must be unspecified |
| type | String | Type of element of the user's Telegram Passport which has the issue |
| element_hash | String | Base64-encoded element hash |
| message | String | Error message |

#### Interface

```go
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
```

### PassportFile

**[Official documentation](https://core.telegram.org/bots/api#passportfile)**

#### Description

This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| file_size | Integer | File size |
| file_date | Integer | Unix time when the file was uploaded |

#### Interface

```go
type PassportFile struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// File size
	FileSize int `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate int `json:"file_date"`
}
```

### PhotoSize

**[Official documentation](https://core.telegram.org/bots/api#photosize)**

#### Description

This object represents one size of a photo or a file / sticker thumbnail.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| width | Integer | Photo width |
| height | Integer | Photo height |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### Poll

**[Official documentation](https://core.telegram.org/bots/api#poll)**

#### Description

This object contains information about a poll.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | String | Unique poll identifier |
| question | String | Poll question, 1-255 characters |
| options | Array of PollOption | List of poll options |
| is_closed | Boolean | True, if the poll is closed |

#### Interface

```go
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
```

### PollOption

**[Official documentation](https://core.telegram.org/bots/api#polloption)**

#### Description

This object contains information about one answer option in a poll.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| text | String | Option text, 1-100 characters |
| voter_count | Integer | Number of users that voted for this option |

#### Interface

```go
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}
```

### PreCheckoutQuery

**[Official documentation](https://core.telegram.org/bots/api#precheckoutquery)**

#### Description

This object contains information about an incoming pre-checkout query.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | String | Unique query identifier |
| from | User | User who sent the query |
| currency | String | Three-letter ISO 4217 currency code |
| total_amount | Integer | Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). |
| invoice_payload | String | Bot specified invoice payload |
| shipping_option_id | String | Optional. Identifier of the shipping option chosen by the user |
| order_info | OrderInfo | Optional. Order info provided by the user |

#### Interface

```go
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
```

### ReplyKeyboardMarkup

**[Official documentation](https://core.telegram.org/bots/api#replykeyboardmarkup)**

#### Description

This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| keyboard | Array of Array of KeyboardButton | Array of button rows, each represented by an Array of KeyboardButton objects |
| resize_keyboard | Boolean | Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard. |
| one_time_keyboard | Boolean | Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false. |
| selective | Boolean | Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard. |

#### Interface

```go
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
```

### ReplyKeyboardRemove

**[Official documentation](https://core.telegram.org/bots/api#replykeyboardremove)**

#### Description

Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| remove_keyboard | True | Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup) |
| selective | Boolean | Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet. |

#### Interface

```go
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}
```

### ResponseParameters

**[Official documentation](https://core.telegram.org/bots/api#responseparameters)**

#### Description

Contains information about why a request was unsuccessful.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| migrate_to_chat_id | Integer | Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. |
| retry_after | Integer | Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated |

#### Interface

```go
type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId int `json:"migrate_to_chat_id,omitempty"`
	// Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
	RetryAfter int `json:"retry_after,omitempty"`
}
```

### ShippingAddress

**[Official documentation](https://core.telegram.org/bots/api#shippingaddress)**

#### Description

This object represents a shipping address.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| country_code | String | ISO 3166-1 alpha-2 country code |
| state | String | State, if applicable |
| city | String | City |
| street_line1 | String | First line for the address |
| street_line2 | String | Second line for the address |
| post_code | String | Address post code |

#### Interface

```go
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
```

### ShippingOption

**[Official documentation](https://core.telegram.org/bots/api#shippingoption)**

#### Description

This object represents one shipping option.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | String | Shipping option identifier |
| title | String | Option title |
| prices | Array of LabeledPrice | List of price portions |

#### Interface

```go
type ShippingOption struct {
	// Shipping option identifier
	Id string `json:"id"`
	// Option title
	Title string `json:"title"`
	// List of price portions
	Prices []*LabeledPrice `json:"prices"`
}
```

### ShippingQuery

**[Official documentation](https://core.telegram.org/bots/api#shippingquery)**

#### Description

This object contains information about an incoming shipping query.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | String | Unique query identifier |
| from | User | User who sent the query |
| invoice_payload | String | Bot specified invoice payload |
| shipping_address | ShippingAddress | User specified shipping address |

#### Interface

```go
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
```

### Sticker

**[Official documentation](https://core.telegram.org/bots/api#sticker)**

#### Description

This object represents a sticker.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| width | Integer | Sticker width |
| height | Integer | Sticker height |
| thumb | PhotoSize | Optional. Sticker thumbnail in the .webp or .jpg format |
| emoji | String | Optional. Emoji associated with the sticker |
| set_name | String | Optional. Name of the sticker set to which the sticker belongs |
| mask_position | MaskPosition | Optional. For mask stickers, the position where the mask should be placed |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### StickerSet

**[Official documentation](https://core.telegram.org/bots/api#stickerset)**

#### Description

This object represents a sticker set.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| name | String | Sticker set name |
| title | String | Sticker set title |
| contains_masks | Boolean | True, if the sticker set contains masks |
| stickers | Array of Sticker | List of all set stickers |

#### Interface

```go
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
```

### SuccessfulPayment

**[Official documentation](https://core.telegram.org/bots/api#successfulpayment)**

#### Description

This object contains basic information about a successful payment.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| currency | String | Three-letter ISO 4217 currency code |
| total_amount | Integer | Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). |
| invoice_payload | String | Bot specified invoice payload |
| shipping_option_id | String | Optional. Identifier of the shipping option chosen by the user |
| order_info | OrderInfo | Optional. Order info provided by the user |
| telegram_payment_charge_id | String | Telegram payment identifier |
| provider_payment_charge_id | String | Provider payment identifier |

#### Interface

```go
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
```

### Update

**[Official documentation](https://core.telegram.org/bots/api#update)**

#### Description

This object represents an incoming update.At most one of the optional parameters can be present in any given update.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| update_id | Integer | The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially. |
| message | Message | Optional. New incoming message of any kind — text, photo, sticker, etc. |
| edited_message | Message | Optional. New version of a message that is known to the bot and was edited |
| channel_post | Message | Optional. New incoming channel post of any kind — text, photo, sticker, etc. |
| edited_channel_post | Message | Optional. New version of a channel post that is known to the bot and was edited |
| inline_query | InlineQuery | Optional. New incoming inline query |
| chosen_inline_result | ChosenInlineResult | Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot. |
| callback_query | CallbackQuery | Optional. New incoming callback query |
| shipping_query | ShippingQuery | Optional. New incoming shipping query. Only for invoices with flexible price |
| pre_checkout_query | PreCheckoutQuery | Optional. New incoming pre-checkout query. Contains full information about checkout |
| poll | Poll | Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot |

#### Interface

```go
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
```

### User

**[Official documentation](https://core.telegram.org/bots/api#user)**

#### Description

This object represents a Telegram user or bot.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| id | Integer | Unique identifier for this user or bot |
| is_bot | Boolean | True, if this user is a bot |
| first_name | String | User‘s or bot’s first name |
| last_name | String | Optional. User‘s or bot’s last name |
| username | String | Optional. User‘s or bot’s username |
| language_code | String | Optional. IETF language tag of the user's language |

#### Interface

```go
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
```

### UserProfilePhotos

**[Official documentation](https://core.telegram.org/bots/api#userprofilephotos)**

#### Description

This object represent a user's profile pictures.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| total_count | Integer | Total number of profile pictures the target user has |
| photos | Array of Array of PhotoSize | Requested profile pictures (in up to 4 sizes each) |

#### Interface

```go
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}
```

### Venue

**[Official documentation](https://core.telegram.org/bots/api#venue)**

#### Description

This object represents a venue.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| location | Location | Venue location |
| title | String | Name of the venue |
| address | String | Address of the venue |
| foursquare_id | String | Optional. Foursquare identifier of the venue |
| foursquare_type | String | Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.) |

#### Interface

```go
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
```

### Video

**[Official documentation](https://core.telegram.org/bots/api#video)**

#### Description

This object represents a video file.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| width | Integer | Video width as defined by sender |
| height | Integer | Video height as defined by sender |
| duration | Integer | Duration of the video in seconds as defined by sender |
| thumb | PhotoSize | Optional. Video thumbnail |
| mime_type | String | Optional. Mime type of a file as defined by sender |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### VideoNote

**[Official documentation](https://core.telegram.org/bots/api#videonote)**

#### Description

This object represents a video message (available in Telegram apps as of v.4.0).

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| length | Integer | Video width and height (diameter of the video message) as defined by sender |
| duration | Integer | Duration of the video in seconds as defined by sender |
| thumb | PhotoSize | Optional. Video thumbnail |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### Voice

**[Official documentation](https://core.telegram.org/bots/api#voice)**

#### Description

This object represents a voice note.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| file_id | String | Unique identifier for this file |
| duration | Integer | Duration of the audio in seconds as defined by sender |
| mime_type | String | Optional. MIME type of the file as defined by sender |
| file_size | Integer | Optional. File size |

#### Interface

```go
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
```

### WebhookInfo

**[Official documentation](https://core.telegram.org/bots/api#webhookinfo)**

#### Description

Contains information about the current status of a webhook.

#### Properties

| Field | Type | Description |
| ----- | ---- | ----------- |
| url | String | Webhook URL, may be empty if webhook is not set up |
| has_custom_certificate | Boolean | True, if a custom certificate was provided for webhook certificate checks |
| pending_update_count | Integer | Number of updates awaiting delivery |
| last_error_date | Integer | Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook |
| last_error_message | String | Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook |
| max_connections | Integer | Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery |
| allowed_updates | Array of String | Optional. A list of update types the bot is subscribed to. Defaults to all update types |

#### Interface

```go
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
```

