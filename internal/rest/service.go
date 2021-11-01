package rest

import "github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"

type IReferenceService interface {
	TickersService.ITickersService
}
type ReferenceService struct {
}
