package telego

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spyzhov/telego/tobject"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Bot struct {
	Host  string
	Log   Logger
	token string
}

type Logger func(format string, v ...interface{})

type Action string

const (
	Host  = "https://api.telegram.org"
	debug = true
)

var (
	InvalidStatusCode = errors.New("invalid status code")
	WrongResponse     = errors.New("invalid response")
)

func New(token string) *Bot {
	return &Bot{
		Host:  Host,
		token: token,
		Log:   log.Printf,
	}
}

func (b *Bot) post(ctx context.Context, action Action, body interface{}) (*http.Response, error) {
	target, err := url.Parse(fmt.Sprintf("%s/bot%s/%s", b.Host, b.token, action))
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(http.MethodPost, target.String(), buffer)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	return new(http.Client).Do(req)
}

func (b *Bot) get(ctx context.Context, action Action, params map[string]string) (*http.Response, error) {
	target, err := url.Parse(fmt.Sprintf("%s/bot%s/%s", b.Host, b.token, action))
	if err != nil {
		return nil, err
	}

	for name, value := range params {
		target.Query().Set(name, value)
	}

	req, err := http.NewRequest(http.MethodGet, target.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	return new(http.Client).Do(req)
}

func (b *Bot) closer(closer io.Closer, scope string) {
	if closer != nil {
		if err := closer.Close(); err != nil {
			b.Log("closer [%s] error: %s", scope, err)
		}
	}
}

func (b *Bot) getResult(ctx context.Context, action Action, params map[string]string, result interface{}) error {
	defer b.debug("[STOP ] GET request: %s", action)
	b.debug("[START] GET request: %s", action)
	response, err := b.get(ctx, action, params)
	if response != nil {
		defer b.closer(response.Body, "response body")
	}
	if err != nil {
		return err
	}
	return b.parse(response, &result)
}

func (b *Bot) postResult(ctx context.Context, action Action, body interface{}, result interface{}) error {
	defer b.debug("[STOP ] POST request: %s", action)
	b.debug("[START] POST request: %s", action)
	response, err := b.post(ctx, action, body)
	if response != nil {
		defer b.closer(response.Body, "response body")
	}
	if err != nil {
		return err
	}
	return b.parse(response, &result)
}

func (b *Bot) parse(response *http.Response, object *interface{}) error {
	if response.StatusCode != http.StatusOK {
		b.Log("invalid status code: %d", response.StatusCode)
		return InvalidStatusCode
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	b.debug("response: %s", data)
	result := new(tobject.Response)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	if !result.OK {
		return WrongResponse
	}
	return json.Unmarshal(result.Result, &object)
}

func (b *Bot) debug(format string, v ...interface{}) {
	b.Log(format, v...)
}
