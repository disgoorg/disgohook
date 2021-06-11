package api

import "time"

// WebhookMessage represents a message created by a Webhook
type WebhookMessage struct {
	Webhook         WebhookClient
	ID              Snowflake     `json:"id"`
	WebhookID       Snowflake     `json:"webhook_id"`
	ChannelID       Snowflake     `json:"channel_id"`
	GuildID         Snowflake     `json:"guild_id"`
	TTS             bool          `json:"tts"`
	Author          *User         `json:"author"`
	CreatedAt       time.Time     `json:"timestamp"`
	EditedAt        *time.Time    `json:"edited_timestamp"`
	Content         *string       `json:"content,omitempty"`
	Embeds          []Embed       `json:"embeds,omitempty"`
	Components      []Component   `json:"components,omitempty"`
	Attachments     []Attachment  `json:"attachments,omitempty"`
	Mentions        []interface{} `json:"mentions"`
	MentionEveryone bool          `json:"mention_everyone"`
	MentionRoles    []string      `json:"mention_roles"`
}

// FullWebhookMessage is used for easier unmarshalling of Component(s) in WebhookMessage(s)
type FullWebhookMessage struct {
	*WebhookMessage
	UnmarshalComponents []UnmarshalComponent `json:"components,omitempty"`
}

// Edit allows you to edit an existing WebhookMessage sent by you
func (m *WebhookMessage) Edit(message WebhookMessageUpdate) (*WebhookMessage, error) {
	return m.Webhook.EditMessage(m.ID, message)
}

// Delete allows you to delete an existing WebhookMessage sent by you
func (m *WebhookMessage) Delete() error {
	return m.Webhook.DeleteMessage(m.ID)
}

//Attachment is used for files sent in a Message
type Attachment struct {
	ID       Snowflake `json:"id,omitempty"`
	Filename string    `json:"filename"`
	Size     int       `json:"size"`
	URL      string    `json:"url"`
	ProxyURL string    `json:"proxy_url"`
	Height   *int      `json:"height"`
	Width    *int      `json:"width"`
}
