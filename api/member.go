package api

import "time"

type Member struct {
	GuildID      string     `json:"guild_id"`
	User         *User      `json:"user"`
	Nick         *string    `json:"nick"`
	Roles        []string   `json:"roles,omitempty"`
	JoinedAt     time.Time  `json:"joined_at"`
	PremiumSince *time.Time `json:"premium_since,omitempty"`
	Deaf         *bool      `json:"deaf,omitempty"`
	Mute         *bool      `json:"mute,omitempty"`
	Pending      bool       `json:"pending"`
	//Permissions  *Permissions `json:"permissions,omitempty"`
}
