package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/DisgoOrg/disgohook/api/endpoints"
)

type RestClientImpl struct {

}

// Request makes a new rest request to discords api with the specific endpoints.APIRoute
func (r RestClientImpl) Request(route endpoints.CompiledAPIRoute, rqBody interface{}, rsBody interface{}) error {
	var reader io.Reader
	var rqJSON []byte
	if rqBody != nil {
		rqJSON, err := json.Marshal(rqBody)
		if err != nil {
			return err
		}
		r.Disgo().Logger().Debugf("request json: \"%s\"", string(rqJSON))
		reader = bytes.NewBuffer(rqJSON)
	} else {
		reader = nil
	}

	rq, err := http.NewRequest(route.Method().String(), route.Route(), reader)
	if err != nil {
		return err
	}

	rq.Header.Set("User-Agent", r.UserAgent())
	rq.Header.Set("Authorization", string("Bot "+r.disgo.Token()))
	rq.Header.Set("Content-Type", "application/json")

	rs, err := r.Client.Do(rq)
	if err != nil {
		return err
	}

	defer func() {
		err = rs.Body.Close()
		if err != nil {
			r.Disgo().Logger().Error("error closing response body", err.Error())
		}
	}()

	rawRsBody, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		r.Disgo().Logger().Errorf("error reading from response body: %s", err)
		return err
	}

	r.Disgo().Logger().Debugf("code: %d, response: %s", rs.StatusCode, string(rawRsBody))

	r.Disgo().EventManager().Dispatch(events.HTTPRequestEvent{
		GenericEvent: events.NewEvent(r.Disgo(), 0),
		Request:      rq,
		Response:     rs,
	})

	switch rs.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		if rsBody != nil {
			if err = json.Unmarshal(rawRsBody, rsBody); err != nil {
				r.Disgo().Logger().Errorf("error unmarshalling response. error: %s", err)
				return err
			}
		}
		return nil

	case http.StatusTooManyRequests:
		limit := rs.Header.Get("X-RateLimit-Limit")
		remaining := rs.Header.Get("X-RateLimit-Limit")
		reset := rs.Header.Get("X-RateLimit-Limit")
		bucket := rs.Header.Get("X-RateLimit-Limit")
		r.Disgo().Logger().Errorf("too many requests. limit: %s, remaining: %s, reset: %s,bucket: %s", limit, remaining, reset, bucket)
		return api.ErrRatelimited

	case http.StatusBadGateway:
		r.Disgo().Logger().Error(api.ErrBadGateway)
		return api.ErrBadGateway

	case http.StatusBadRequest:
		r.Disgo().Logger().Errorf("bad request request: \"%s\", response: \"%s\"", string(rqJSON), string(rawRsBody))
		return api.ErrBadRequest

	case http.StatusUnauthorized:
		r.Disgo().Logger().Error(api.ErrUnauthorized)
		return api.ErrUnauthorized

	default:
		var errorRs api.ErrorResponse
		if err = json.Unmarshal(rawRsBody, &errorRs); err != nil {
			r.Disgo().Logger().Errorf("error unmarshalling error response. code: %d, error: %s", rs.StatusCode, err)
			return err
		}
		return fmt.Errorf("request to %s failed. statuscode: %d, errorcode: %d, message_events: %s", rq.URL, rs.StatusCode, errorRs.Code, errorRs.Message)
	}
}
