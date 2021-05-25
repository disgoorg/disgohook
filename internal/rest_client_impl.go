package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/DisgoOrg/disgohook/api"

	"github.com/DisgoOrg/disgohook/api/endpoints"
)

func newRestClientImpl(client *http.Client, webhook api.WebhookClient) api.RestClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &RestClientImpl{
		client:        client,
		webhookClient: webhook,
	}
}

type RestClientImpl struct {
	client        *http.Client
	webhookClient api.WebhookClient
}

func (r *RestClientImpl) UserAgent() string {
	return "DisgoHook"
}

// Request makes a new rest request to discords api with the specific endpoints.APIRoute
func (r *RestClientImpl) Request(route *endpoints.CompiledAPIRoute, rqBody interface{}, rsBody interface{}) error {
	var reader io.Reader
	var rqJSON []byte
	if rqBody != nil {
		rqJSON, err := json.Marshal(rqBody)
		if err != nil {
			return err
		}
		r.webhookClient.Logger().Debugf("request json: \"%s\"", string(rqJSON))
		reader = bytes.NewBuffer(rqJSON)
	} else {
		reader = nil
	}

	rq, err := http.NewRequest(route.Method().String(), route.Route(), reader)
	if err != nil {
		return err
	}

	rq.Header.Set("User-Agent", r.UserAgent())
	rq.Header.Set("Content-Type", "application/json")

	rs, err := r.client.Do(rq)
	if err != nil {
		return err
	}

	defer func() {
		err = rs.Body.Close()
		if err != nil {
			r.webhookClient.Logger().Error("error closing response body", err.Error())
		}
	}()

	rawRsBody, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		r.webhookClient.Logger().Errorf("error reading from response body: %s", err)
		return err
	}

	r.webhookClient.Logger().Debugf("code: %d, response: %s", rs.StatusCode, string(rawRsBody))

	switch rs.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		if rsBody != nil {
			if err = json.Unmarshal(rawRsBody, rsBody); err != nil {
				r.webhookClient.Logger().Errorf("error unmarshalling response. error: %s", err)
				return err
			}
		}
		return nil

	case http.StatusTooManyRequests:
		limit := rs.Header.Get("X-RateLimit-Limit")
		remaining := rs.Header.Get("X-RateLimit-Limit")
		reset := rs.Header.Get("X-RateLimit-Limit")
		bucket := rs.Header.Get("X-RateLimit-Limit")
		r.webhookClient.Logger().Errorf("too many requests. limit: %s, remaining: %s, reset: %s,bucket: %s", limit, remaining, reset, bucket)
		return api.ErrRatelimited

	case http.StatusBadGateway:
		r.webhookClient.Logger().Error(api.ErrBadGateway)
		return api.ErrBadGateway

	case http.StatusBadRequest:
		r.webhookClient.Logger().Errorf("bad request request: \"%s\", response: \"%s\"", string(rqJSON), string(rawRsBody))
		return api.ErrBadRequest

	case http.StatusUnauthorized:
		r.webhookClient.Logger().Error(api.ErrUnauthorized)
		return api.ErrUnauthorized

	default:
		var errorRs api.ErrorResponse
		if err = json.Unmarshal(rawRsBody, &errorRs); err != nil {
			r.webhookClient.Logger().Errorf("error unmarshalling error response. code: %d, error: %s", rs.StatusCode, err)
			return err
		}
		return fmt.Errorf("request to %s failed. statuscode: %d, errorcode: %d, message_events: %s", rq.URL, rs.StatusCode, errorRs.Code, errorRs.Message)
	}
}

func (r *RestClientImpl) GetWebhook(webhookID api.Snowflake, webhookToken string) (webhook *api.Webhook, err error) {
	var compiledRoute *endpoints.CompiledAPIRoute
	compiledRoute, err = endpoints.GetWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Request(compiledRoute, nil, &webhook)
	return
}

func (r *RestClientImpl) UpdateWebhook(webhookID api.Snowflake, webhookToken string, webhookUpdate *api.WebhookUpdate) (webhook *api.Webhook, err error) {
	var compiledRoute *endpoints.CompiledAPIRoute
	compiledRoute, err = endpoints.UpdateWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Request(compiledRoute, webhookUpdate, &webhook)
	return
}

func (r *RestClientImpl) DeleteWebhook(webhookID api.Snowflake, webhookToken string) (err error) {
	var compiledRoute *endpoints.CompiledAPIRoute
	compiledRoute, err = endpoints.DeleteWebhook.Compile(nil, webhookID, webhookToken)
	if err != nil {
		return
	}
	err = r.Request(compiledRoute, nil, nil)
	return
}

func (r *RestClientImpl) CreateWebhookMessage(webhookID api.Snowflake, webhookToken string, messageCreate *api.WebhookMessageCreate, wait bool, threadID api.Snowflake) (message *api.WebhookMessage, err error) {
	params := map[string]interface{}{}
	if wait {
		params["wait"] = true
	}
	if threadID != "" {
		params["thread_id"] = threadID
	}
	var compiledRoute *endpoints.CompiledAPIRoute
	compiledRoute, err = endpoints.CreateWebhookMessage.Compile(params, webhookID, webhookToken)
	if err != nil {
		return
	}
	if wait {
		err = r.Request(compiledRoute, messageCreate, &message)
	} else {
		err = r.Request(compiledRoute, messageCreate, nil)
	}
	if message != nil {
		message.Webhook = r.webhookClient
	}
	return
}

func (r *RestClientImpl) UpdateWebhookMessage(webhookID api.Snowflake, webhookToken string, messageID api.Snowflake, messageUpdate *api.WebhookMessageUpdate) (message *api.WebhookMessage, err error) {
	var compiledRoute *endpoints.CompiledAPIRoute
	compiledRoute, err = endpoints.UpdateWebhookMessage.Compile(nil, webhookID, webhookToken, messageID)
	if err != nil {
		return
	}
	err = r.Request(compiledRoute, messageUpdate, &message)
	if message != nil {
		message.Webhook = r.webhookClient
	}
	return
}

func (r *RestClientImpl) DeleteWebhookMessage(webhookID api.Snowflake, webhookToken string, messageID api.Snowflake) (err error) {
	var compiledRoute *endpoints.CompiledAPIRoute
	compiledRoute, err = endpoints.DeleteWebhookMessage.Compile(nil, webhookID, webhookToken, messageID)
	if err != nil {
		return
	}
	err = r.Request(compiledRoute, nil, nil)
	return
}
