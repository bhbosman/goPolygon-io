package TickersService

import (
	"fmt"
	"github.com/bhbosman/goPolygon-io/internal/stream"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"net/http"
	"strings"
)

type TickersService struct {
	client *http.Client
}

func NewTickers(client *http.Client) *TickersService {
	return &TickersService{
		client: client,
	}
}

func (self *TickersService) Tickers(options ...ITickersServiceOption) (*stream.ReferenceDataTickersResponse, error) {
	var sa []string
	for _, option := range options {
		if option == nil {
			continue
		}
		s := option.applyTickersOption()
		if s == "" {
			continue
		}
		sa = append(sa, s)
	}
	params := strings.Join(sa, "&")

	url := func(params string) string {
		url := "https://api.polygon.io/v3/reference/tickers"
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
	return readReferenceDataTickersResponse(response.Body)
}

func (self *TickersService) TickersNext(nextUrl string) (*stream.ReferenceDataTickersResponse, error) {
	response, err := self.client.Get(nextUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	return readReferenceDataTickersResponse(response.Body)
}

func readReferenceDataTickersResponse(body io.Reader) (*stream.ReferenceDataTickersResponse, error) {
	result := &stream.ReferenceDataTickersResponse{}
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
