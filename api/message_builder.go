package api

// WebhookMessageBuilder helper to build Message(s) easier
type WebhookMessageBuilder struct {
	WebhookMessageCreate
}

func NewWebhookMessageWithEmbeds(embed *Embed, embeds ...*Embed) *WebhookMessageBuilder {
	return NewWebhookMessageBuilder().AddEmbeds(embed).AddEmbeds(embeds...)
}

func NewWebhookMessageWithContent(content string) *WebhookMessageBuilder {
	return NewWebhookMessageBuilder().SetContent(content)
}

// NewWebhookMessageBuilder creates a new WebhookMessageBuilder to be built later
func NewWebhookMessageBuilder() *WebhookMessageBuilder {
	return &WebhookMessageBuilder{
		WebhookMessageCreate: WebhookMessageCreate{
			AllowedMentions: &DefaultAllowedMentions,
		},
	}
}

// SetContent sets content of the Message
func (b *WebhookMessageBuilder) SetContent(content string) *WebhookMessageBuilder {
	b.Content = content
	return b
}

// SetTTS sets the text to speech of the Message
func (b *WebhookMessageBuilder) SetTTS(tts bool) *WebhookMessageBuilder {
	b.TTS = tts
	return b
}

// SetEmbeds sets the embeds of the InteractionResponse
func (b *WebhookMessageBuilder) SetEmbeds(embeds ...*Embed) *WebhookMessageBuilder {
	b.Embeds = embeds
	return b
}

// AddEmbeds adds multiple embeds to the InteractionResponse
func (b *WebhookMessageBuilder) AddEmbeds(embeds ...*Embed) *WebhookMessageBuilder {
	b.Embeds = append(b.Embeds, embeds...)
	return b
}

// ClearEmbeds removes all of the embeds from the InteractionResponse
func (b *WebhookMessageBuilder) ClearEmbeds() *WebhookMessageBuilder {
	if b.Embeds != nil {
		b.Embeds = []*Embed{}
	}
	return b
}

// RemoveEmbed removes an embed from the InteractionResponse
func (b *WebhookMessageBuilder) RemoveEmbed(index int) *WebhookMessageBuilder {
	if len(b.Embeds) > index {
		b.Embeds = append(b.Embeds[:index], b.Embeds[index+1:]...)
	}
	return b
}

// SetAllowedMentions sets the AllowedMentions of the Message
func (b *WebhookMessageBuilder) SetAllowedMentions(allowedMentions *AllowedMentions) *WebhookMessageBuilder {
	b.AllowedMentions = allowedMentions
	return b
}

// SetAllowedMentionsEmpty sets the allowed mentions of the Message to nothing
func (b *WebhookMessageBuilder) SetAllowedMentionsEmpty() *WebhookMessageBuilder {
	return b.SetAllowedMentions(&AllowedMentions{})
}


// Build builds the WebhookMessageBuilder to a MessageCreate struct
func (b *WebhookMessageBuilder) Build() WebhookMessageCreate {
	return b.WebhookMessageCreate
}