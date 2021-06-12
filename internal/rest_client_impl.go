package internal

import (
	"net/http"

	"github.com/DisgoOrg/disgohook/api"

	"github.com/DisgoOrg/restclient"
)

func newRestClientImpl(httpClient *http.Client, webhook api.WebhookClient) api.RestClient {
	return &restClientImpl{
		RestClient:    restclient.NewRestClient(httpClient, webhook.Logger(), "DisgoHook", nil),
		webhookClient: webhook,
	}
}

type restClientImpl struct {
	restclient.RestClient
	webhookClient api.WebhookClient
}

func (r *restClientImpl) WebhookClient() api.WebhookClient {
	return r.webhookClient
}

func (r *restClientImpl) GetWebhook(webhookID api.Snowflake, webhookToken string) (webhook *api.Webhook, err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.GetWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, nil, &webhook)
	return
}

func (r *restClientImpl) UpdateWebhook(webhookID api.Snowflake, webhookToken string, webhookUpdate api.WebhookUpdate) (webhook *api.Webhook, err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.UpdateWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, webhookUpdate, &webhook)
	return
}

func (r *restClientImpl) DeleteWebhook(webhookID api.Snowflake, webhookToken string) (err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.DeleteWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, nil, nil)
	return
}

func (r *restClientImpl) CreateWebhookMessage(webhookID api.Snowflake, webhookToken string, messageCreate api.WebhookMessageCreate, wait bool, threadID api.Snowflake) (message *api.WebhookMessage, err error) {
	params := map[string]interface{}{}
	if wait {
		params["wait"] = true
	}
	if threadID != "" {
		params["thread_id"] = threadID
	}
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.CreateWebhookMessage.Compile(params, webhookID, webhookToken)
	if err != nil {
		return
	}

	body, err := messageCreate.ToBody()
	if err != nil {
		return nil, err
	}

	var fullMessage *api.FullWebhookMessage
	if wait {
		err = r.Do(compiledRoute, body, &fullMessage)
	} else {
		err = r.Do(compiledRoute, body, nil)
	}
	if fullMessage != nil {
		message = r.createMessage(fullMessage)
	}
	return
}

func (r *restClientImpl) UpdateWebhookMessage(webhookID api.Snowflake, webhookToken string, messageID api.Snowflake, messageUpdate api.WebhookMessageUpdate) (message *api.WebhookMessage, err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.UpdateWebhookMessage.Compile(nil, webhookID, webhookToken, messageID)
	if err != nil {
		return
	}

	body, err := messageUpdate.ToBody()
	if err != nil {
		return nil, err
	}

	var fullMessage *api.FullWebhookMessage
	err = r.Do(compiledRoute, body, &fullMessage)
	if fullMessage != nil {
		message = r.createMessage(fullMessage)
	}
	return
}

func (r *restClientImpl) DeleteWebhookMessage(webhookID api.Snowflake, webhookToken string, messageID api.Snowflake) (err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.DeleteWebhookMessage.Compile(nil, webhookID, webhookToken, messageID)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, nil, nil)
	return
}

func (r *restClientImpl) createComponent(unmarshalComponent api.UnmarshalComponent) api.Component {
	switch unmarshalComponent.ComponentType {
	case api.ComponentTypeActionRow:
		components := make([]api.Component, len(unmarshalComponent.Components))
		for i, unmarshalC := range unmarshalComponent.Components {
			components[i] = r.createComponent(unmarshalC)
		}
		return &api.ActionRow{
			ComponentImpl: api.ComponentImpl{
				ComponentType: api.ComponentTypeActionRow,
			},
			Components: components,
		}

	case api.ComponentTypeButton:
		button := &api.Button{
			ComponentImpl: api.ComponentImpl{
				ComponentType: api.ComponentTypeButton,
			},
			Style:    unmarshalComponent.Style,
			Label:    unmarshalComponent.Label,
			CustomID: unmarshalComponent.CustomID,
			URL:      unmarshalComponent.URL,
			Disabled: unmarshalComponent.Disabled,
			Emoji:    unmarshalComponent.Emoji,
		}
		return button

	default:
		r.Logger().Errorf("unexpected component type %d received", unmarshalComponent.ComponentType)
		return nil
	}
}

func (r *restClientImpl) createMessage(fullMessage *api.FullWebhookMessage) *api.WebhookMessage {
	message := fullMessage.WebhookMessage
	message.Webhook = r.webhookClient
	if fullMessage.UnmarshalComponents != nil {
		for _, component := range fullMessage.UnmarshalComponents {
			message.Components = append(message.Components, r.createComponent(component))
		}
	}

	return message
}
