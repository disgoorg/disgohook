package internal

import (
	"net/http"

	"github.com/DisgoOrg/disgohook/api"

	"github.com/DisgoOrg/restclient"
)

func newRestClientImpl(httpClient *http.Client, webhook api.WebhookClient) api.RestClient {
	return &RestClientImpl{
		RestClient:    restclient.NewRestClient(httpClient, webhook.Logger(), "DisgoHook", nil),
		webhookClient: webhook,
	}
}

type RestClientImpl struct {
	restclient.RestClient
	webhookClient api.WebhookClient
}

func (r *RestClientImpl) WebhookClient() api.WebhookClient {
	return r.webhookClient
}

func (r *RestClientImpl) GetWebhook(webhookID api.Snowflake, webhookToken string) (webhook *api.Webhook, err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.GetWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, nil, &webhook)
	return
}

func (r *RestClientImpl) UpdateWebhook(webhookID api.Snowflake, webhookToken string, webhookUpdate api.WebhookUpdate) (webhook *api.Webhook, err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.UpdateWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, webhookUpdate, &webhook)
	return
}

func (r *RestClientImpl) DeleteWebhook(webhookID api.Snowflake, webhookToken string) (err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.DeleteWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, nil, nil)
	return
}

func (r *RestClientImpl) CreateWebhookMessage(webhookID api.Snowflake, webhookToken string, messageCreate api.WebhookMessageCreate, wait bool, threadID api.Snowflake) (message *api.WebhookMessage, err error) {
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
	if wait {
		err = r.Do(compiledRoute, messageCreate, &message)
	} else {
		err = r.Do(compiledRoute, messageCreate, nil)
	}
	if message != nil {
		message.Webhook = r.webhookClient
	}
	return
}

func (r *RestClientImpl) UpdateWebhookMessage(webhookID api.Snowflake, webhookToken string, messageID api.Snowflake, messageUpdate api.WebhookMessageUpdate) (message *api.WebhookMessage, err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.UpdateWebhookMessage.Compile(nil, webhookID, webhookToken, messageID)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, messageUpdate, &message)
	if message != nil {
		message.Webhook = r.webhookClient
	}
	return
}

func (r *RestClientImpl) DeleteWebhookMessage(webhookID api.Snowflake, webhookToken string, messageID api.Snowflake) (err error) {
	var compiledRoute *restclient.CompiledAPIRoute
	compiledRoute, err = restclient.DeleteWebhookMessage.Compile(nil, webhookID, webhookToken, messageID)
	if err != nil {
		return
	}
	err = r.Do(compiledRoute, nil, nil)
	return
}
