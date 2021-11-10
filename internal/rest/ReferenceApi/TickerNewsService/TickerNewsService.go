package TickerNewsService

import (
	"fmt"
	"github.com/bhbosman/goPolygon-io/internal/stream"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"net/http"
	"strings"
)

type TickerNewsService struct {
	client *http.Client
}

func NewTickerNewsService(client *http.Client) *TickerNewsService {
	return &TickerNewsService{client: client}
}

func (self *TickerNewsService) TickerNews(options ...ITickerDetailsServiceOption) (*stream.ReferenceDataTickerNewsResponse, error) {
	var sa []string
	for _, option := range options {
		if option == nil {
			continue
		}
		s := option.applyTickerDetailsServiceOption()
		if s == "" {
			continue
		}
		sa = append(sa, s)
	}
	params := strings.Join(sa, "&")

	url := func(params string) string {
		url := "https://api.polygon.io/v2/reference/news"
		switch {
		case params == "":
			return url
		default:
			return fmt.Sprintf("%v?%v", url, params)
		}
	}(params)
	var request *http.Request
	var err error
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = self.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	return readReferenceDataTickerNewsResponse(response.Body)
}

func (self *TickerNewsService) TickerNewsNext(nextUrl string) (*stream.ReferenceDataTickerNewsResponse, error) {
	response, err := self.client.Get(nextUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	return readReferenceDataTickerNewsResponse(response.Body)
}

func readReferenceDataTickerNewsResponse(body io.Reader) (*stream.ReferenceDataTickerNewsResponse, error) {
	result := &stream.ReferenceDataTickerNewsResponse{}
	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
		AnyResolver:        nil,
	}
	err := unmarshaler.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
