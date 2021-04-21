package api

import (
	"strconv"
	"strings"

	"github.com/DisgoOrg/disgohook/api/endpoints"
)

type User struct {
	ID            string  `json:"id"`
	Discriminator string  `json:"discriminator"`
	Bot           *bool   `json:"bot,omitempty"`
	Username      string  `json:"username"`
	Avatar        *string `json:"avatar"`
}

// Mention returns the user as a mention
func (u User) Mention() string {
	return "<@" + u.ID + ">"
}

// Tag returns the user's Username and Discriminator
func (u User) Tag() string {
	return u.Username + "#" + u.Discriminator
}

// String returns
func (u User) String() string {
	return u.Mention()
}

// AvatarURL returns the Icon of a User
func (g *User) AvatarURL() string {
	if g.Avatar == nil {
		discrim, _ := strconv.Atoi(g.Discriminator)
		route, err := endpoints.DefaultUserAvatar.Compile(endpoints.PNG, discrim%5)
		if err != nil {
			return ""
		}
		return route.Route()
	}
	format := endpoints.PNG
	if strings.HasPrefix(*g.Avatar, "a_") {
		format = endpoints.GIF
	}
	route, err := endpoints.UserAvatar.Compile(format, g.ID, *g.Avatar)
	if err != nil {
		return ""
	}
	return route.Route()
}
