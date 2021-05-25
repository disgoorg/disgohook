package endpoints

// Discord Endpoint Constants
const (
	APIVersion = "8"
	Base       = "https://discord.com/"
	CDN        = "https://cdn.discordapp.com"
	API        = Base + "api/v" + APIVersion
)

// Webhooks
var (
	GetWebhook    = NewAPIRoute(GET, "/webhooks/{webhook.id}/{token}")
	UpdateWebhook  = NewAPIRoute(PATCH, "/webhooks/{webhook.id}/{token}")
	DeleteWebhook = NewAPIRoute(DELETE, "/webhooks/{webhook.id}/{token}")

	CreateWebhookMessage = NewAPIRoute(POST, "/webhooks/{webhook.id}/{webhook.token}", "wait", "thread_id")
	UpdateWebhookMessage = NewAPIRoute(POST, "/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}")
	DeleteWebhookMessage = NewAPIRoute(POST, "/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}")
)

// CDN
var (
	Emote             = NewCDNRoute("/emojis/{emote.id}.", PNG, GIF)
	DefaultUserAvatar = NewCDNRoute("/embed/avatars/{user.discriminator%5}.", PNG)
	UserAvatar        = NewCDNRoute("/avatars/{user.id}/user.avatar.", PNG, JPEG, WEBP, GIF)
	Attachments       = NewCDNRoute("/attachments/{channel.id}/{attachment.id}/{file.name}", BLANK)
)
