package http

import (
	"fmt"
	"net/http"
)

type Transport struct {
	ApiKey string
	base   http.RoundTripper
}

func (self *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", self.ApiKey))
	return self.base.RoundTrip(request)
}

func NewTransport(apiKey string, base http.RoundTripper) *Transport {
	return &Transport{
		ApiKey: apiKey,
		base:   base,
	}
}
