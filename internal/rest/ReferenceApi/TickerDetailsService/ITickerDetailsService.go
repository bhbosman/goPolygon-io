package TickerDetailsService

import "github.com/bhbosman/goPolygon-io/internal/stream"

type ITickerDetailsService interface {
	TickerDetails(stocksTicker string) (*stream.ReferenceDataTickerDetails, error)
}
