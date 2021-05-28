package api

import "time"

// WebhookMessage represents a message created by a Webhook
type WebhookMessage struct {
	Webhook         WebhookClient
	ID              Snowflake            `json:"id"`
	WebhookID       Snowflake            `json:"webhook_id"`
	ChannelID       Snowflake            `json:"channel_id"`
	GuildID         Snowflake            `json:"guild_id"`
	TTS             bool                 `json:"tts"`
	Author          *User                `json:"author"`
	CreatedAt       time.Time            `json:"timestamp"`
	EditedAt        *time.Time           `json:"edited_timestamp"`
	Content         *string              `json:"content,omitempty"`
	Embeds          []*Embed             `json:"embeds,omitempty"`
	Attachments     []*WebhookAttachment `json:"attachments,omitempty"`
	Mentions        []interface{}        `json:"mentions"`
	MentionEveryone bool                 `json:"mention_everyone"`
	MentionRoles    []string             `json:"mention_roles"`
}

// Edit allows you to edit an existing WebhookMessage sent by you
func (m *WebhookMessage) Edit(message *WebhookMessageUpdate) (*WebhookMessage, error) {
	return m.Webhook.EditMessage(m.ID, message)
}

// Delete allows you to delete an existing WebhookMessage sent by you
func (m *WebhookMessage) Delete() error {
	return m.Webhook.DeleteMessage(m.ID)
}

//WebhookAttachment is used for files sent in a Message
type WebhookAttachment struct {
	ID       string `json:"id,omitempty"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   *int   `json:"height"`
	Width    *int   `json:"width"`
}

// WebhookMessageUpdate is used to edit a Message
type WebhookMessageUpdate struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
}

// WebhookMessageCreate is the struct to create a new WebhookMessage
type WebhookMessageCreate struct {
	Content         string           `json:"content,omitempty"`
	Username        string           `json:"username,omitempty"`
	AvatarURL       string           `json:"avatar_url,omitempty"`
	TTS             bool             `json:"tts,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
}
