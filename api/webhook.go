package api

import "io"

type WebhookType int

const (
	WebhookTypeIncoming WebhookType = iota + 1
	WebhookTypeChannelFollower
	WebhookTypeApplication
)

type Webhook struct {
	ID            Snowflake   `json:"id"`
	Type          WebhookType `json:"type"`
	GuildID       *Snowflake  `json:"guild_id"`
	ChannelID     *Snowflake  `json:"channel_id"`
	Name          string      `json:"name"`
	Avatar        string      `json:"avatar"`
	Token         *string     `json:"token"`
	ApplicationID Snowflake   `json:"application_id"`
}

type WebhookUpdate struct {
	Name   *string   `json:"name,omitempty"`
	Avatar io.Reader `json:"avatar,omitempty"`
}
