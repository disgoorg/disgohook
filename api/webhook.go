package api

type WebhookType int

const (
	WebhookTypeIncoming WebhookType = iota
	WebhookTypeChannelFollower
)

type Webhook struct {
	DisgoHook     DisgoHook
	ID            string      `json:"id"`
	Type          WebhookType `json:"type"`
	GuildID       *string     `json:"guild_id,omitempty"`
	ChannelID     string      `json:"channel_id"`
	User          *User       `json:"user,omitempty"`
	Name          *string     `json:"name"`
	Avatar        *string     `json:"avatar"`
	Token         *string     `json:"token"`
	ApplicationID *string     `json:"application_id"`
	Guild         *Guild      `json:"guild"`
	Channel       *Channel    `json:"channel"`
	URL           *string
}

func (w *Webhook) Update(webhookUpdate UpdateWebhook) error {
	webhook, err := w.DisgoHook.RestClient().UpdateWebhook(w.ID, *w.Token, webhookUpdate)
}

type UpdateWebhook struct {
}
