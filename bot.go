package tg

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//go:generate go run gen.go

// Bot structure as provider for Bot-API
type Bot struct {
	Host  string
	Log   Logger
	Debug bool
	token string
}

// Response for default message
type Response struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

// Error wrapper
type Error struct {
	Response *Response
	Basic    error
}

// Logger function for debugging
type Logger func(format string, v ...interface{})

const (
	Host = "https://api.telegram.org"
)

var (
	InvalidStatusCode = errors.New("invalid status code")
	WrongResponse     = errors.New("invalid response")
)

func New(token string) *Bot {
	return &Bot{
		Host:  Host,
		token: token,
		Log:   func(string, ...interface{}) {},
		Debug: false,
	}
}

func (b *Bot) Request(ctx context.Context, action string, body interface{}) (result []byte, err error) {
	defer func(start time.Time) {
		b.debug("[STOP ] POST request: %s [%0.5fs]", action, float64(time.Since(start))/float64(time.Second))
	}(time.Now())
	b.debug("[START] POST request: %s", action)
	response, err := b.post(ctx, action, body)
	if response != nil {
		defer b.closer(response.Body, "response body")
	}
	if err != nil {
		return
	}
	return result, b.parse(response, &result)
}

func (b *Bot) post(ctx context.Context, action string, body interface{}) (*http.Response, error) {
	target, err := url.Parse(fmt.Sprintf("%s/bot%s/%s", b.Host, b.token, action))
	if err != nil {
		return nil, err
	}

	var data []byte
	buffer := bytes.NewBuffer(nil)
	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(http.MethodPost, target.String(), buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	if b.Debug {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			return &http.Response{}, err
		}
		b.Log("DUMP\n%s", string(dump))
	}
	return new(http.Client).Do(req)
}

func (b *Bot) closer(closer io.Closer, scope string) {
	if closer != nil {
		if err := closer.Close(); err != nil {
			b.Log("closer [%s] error: %s", scope, err)
		}
	}
}

func (b *Bot) postResult(ctx context.Context, action string, body interface{}, result interface{}) error {
	defer func(start time.Time) {
		b.debug("[STOP ] POST request: %s [%0.5fs]", action, float64(time.Since(start))/float64(time.Second))
	}(time.Now())
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

func (b *Bot) parse(response *http.Response, object interface{}) error {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		b.Log("invalid status code: %d // %s", response.StatusCode, string(data))
		return InvalidStatusCode
	}
	b.debug("response: %s", data)
	result := new(Response)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	if !result.OK {
		return &Error{
			Basic:    WrongResponse,
			Response: result,
		}
	}
	if err = json.Unmarshal(result.Result, &object); err != nil {
		return &Error{
			Basic:    err,
			Response: result,
		}
	}
	return nil
}

func (b *Bot) debug(format string, v ...interface{}) {
	b.Log(format, v...)
}

func (e Error) Error() string {
	return e.Basic.Error()
}
