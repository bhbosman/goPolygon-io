package TickerDetailsService

import (
	"fmt"
	"github.com/bhbosman/goPolygon-io/internal/stream"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"net/http"
)

type TickerDetailsService struct {
	client *http.Client
}

func NewTickerDetailsService(client *http.Client) *TickerDetailsService {
	return &TickerDetailsService{client: client}
}

func (self *TickerDetailsService) TickerDetails(stocksTicker string) (*stream.ReferenceDataTickerDetails, error) {
	url := fmt.Sprintf("https://api.polygon.io/v1/meta/symbols/%v/company", stocksTicker)
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
	return self.readReferenceDataTickerDetails(response.Body)
}

func (self *TickerDetailsService) readReferenceDataTickerDetails(body io.Reader) (*stream.ReferenceDataTickerDetails, error) {
	result := &stream.ReferenceDataTickerDetails{}
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
