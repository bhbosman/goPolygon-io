package TickerTypesService

import (
	"fmt"
	"github.com/bhbosman/goPolygon-io/internal/stream"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"net/http"
	"strings"
)

type TickerTypesService struct {
	client *http.Client
}

func NewTickerTypesService(client *http.Client) *TickerTypesService {
	return &TickerTypesService{client: client}
}

func (self *TickerTypesService) TickerTypes(options ...ITickerTypesServiceOption) (*stream.ReferenceDataTickerTypesResponse, error) {
	var sa []string
	for _, option := range options {
		if option == nil {
			continue
		}
		s := option.applyTickerTypesOption()
		if s == "" {
			continue
		}
		sa = append(sa, s)
	}
	params := strings.Join(sa, "&")

	url := func(params string) string {
		url := "https://api.polygon.io/v3/reference/tickers/types"
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
	return self.readReferenceDataTickerTypesResponse(response.Body)

}

func (self *TickerTypesService) readReferenceDataTickerTypesResponse(body io.Reader) (*stream.ReferenceDataTickerTypesResponse, error) {
	result := &stream.ReferenceDataTickerTypesResponse{}
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
