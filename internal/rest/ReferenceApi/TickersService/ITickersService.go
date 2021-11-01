package TickersService

import (
	"github.com/bhbosman/goPolygon-io/internal/stream"
)

type ITickersService interface {
	Tickers(option ...ITickersServiceOption) (*stream.ReferenceDataTickersResponse, error)
	TickersNext(nextUrl string) (*stream.ReferenceDataTickersResponse, error)
}
