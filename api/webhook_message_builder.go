package api

import "fmt"

// WebhookMessageBuilder helper to build WebhookMessage(s) easier
type WebhookMessageBuilder struct {
	WebhookMessageCreate
}

// NewWebhookMessageBuilder creates a new WebhookMessageBuilder to be built later
func NewWebhookMessageBuilder() *WebhookMessageBuilder {
	return &WebhookMessageBuilder{
		WebhookMessageCreate: WebhookMessageCreate{
			AllowedMentions: &DefaultAllowedMentions,
		},
	}
}

func NewWebhookMessageBuilderWithEmbeds(embed *Embed, embeds ...*Embed) *WebhookMessageBuilder {
	return NewWebhookMessageBuilder().AddEmbeds(embed, embeds...)
}

func NewWebhookMessageBuilderWithContent(content string) *WebhookMessageBuilder {
	return NewWebhookMessageBuilder().SetContent(content)
}

// SetContent sets content of the WebhookMessage
func (b *WebhookMessageBuilder) SetContent(content string) *WebhookMessageBuilder {
	b.Content = content
	return b
}

// SetContentf sets content of the WebhookMessage
func (b *WebhookMessageBuilder) SetContentf(content string, a ...interface{}) *WebhookMessageBuilder {
	b.Content = fmt.Sprintf(content, a...)
	return b
}

// SetTTS sets the text to speech of the WebhookMessage
func (b *WebhookMessageBuilder) SetTTS(tts bool) *WebhookMessageBuilder {
	b.TTS = tts
	return b
}

// SetEmbeds sets the embeds of the WebhookMessage
func (b *WebhookMessageBuilder) SetEmbeds(embeds ...*Embed) *WebhookMessageBuilder {
	b.Embeds = embeds
	return b
}

// AddEmbeds adds multiple embeds to the WebhookMessage
func (b *WebhookMessageBuilder) AddEmbeds(embed *Embed, embeds ...*Embed) *WebhookMessageBuilder {
	b.Embeds = append(append(b.Embeds, embed), embeds...)
	return b
}

// ClearEmbeds removes all of the embeds from the WebhookMessage
func (b *WebhookMessageBuilder) ClearEmbeds() *WebhookMessageBuilder {
	b.Embeds = []*Embed{}
	return b
}

// RemoveEmbed removes an embed from the WebhookMessage
func (b *WebhookMessageBuilder) RemoveEmbed(i int) *WebhookMessageBuilder {
	if len(b.Embeds) > i {
		b.Embeds = append(b.Embeds[:i], b.Embeds[i+1:]...)
	}
	return b
}

// SetAllowedMentions sets the AllowedMentions of the WebhookMessage
func (b *WebhookMessageBuilder) SetAllowedMentions(allowedMentions *AllowedMentions) *WebhookMessageBuilder {
	b.AllowedMentions = allowedMentions
	return b
}

// ClearAllowedMentions clears the allowed mentions of the WebhookMessage
func (b *WebhookMessageBuilder) ClearAllowedMentions() *WebhookMessageBuilder {
	return b.SetAllowedMentions(&AllowedMentions{})
}

// Build builds the WebhookMessageBuilder to a MessageCreate struct
func (b *WebhookMessageBuilder) Build() *WebhookMessageCreate {
	return &b.WebhookMessageCreate
}
