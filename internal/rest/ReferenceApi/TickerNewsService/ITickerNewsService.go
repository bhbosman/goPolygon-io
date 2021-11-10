package TickerNewsService

import "github.com/bhbosman/goPolygon-io/internal/stream"

type ITickerNewsService interface {
	TickerNews(options ...ITickerDetailsServiceOption) (*stream.ReferenceDataTickerNewsResponse, error)
	TickerNewsNext(nextUrl string) (*stream.ReferenceDataTickerNewsResponse, error)
}
