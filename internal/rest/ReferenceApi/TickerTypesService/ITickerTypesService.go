package TickerTypesService

import "github.com/bhbosman/goPolygon-io/internal/stream"

type ITickerTypesService interface {
	TickerTypes(options ...ITickerTypesServiceOption) (*stream.ReferenceDataTickerTypesResponse, error)
}
