package http

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

type CustomClient struct {
	Client *http.Client
}

func New(client *http.Client) *CustomClient {
	return &CustomClient{
		Client: client,
	}
}

type IHttpClient interface {
	NewRequest(method, url string, data interface{}) (*http.Response, error)
}

// NewClient - creates new http client
func NewClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	return client
}

func (c *CustomClient) NewRequest(method, url string, data interface{}) (*http.Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
