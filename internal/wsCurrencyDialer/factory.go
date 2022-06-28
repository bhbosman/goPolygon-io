package wsCurrencyDialer

import (
	"context"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocomms/intf"
	"go.uber.org/zap"
)

type factory struct {
	apiKey                    string
	fxRegistration            string
	fxAggregationRegistration string
	tickersService            TickersService.ITickersService
}

func NewConnectionReactorFactory(
	apiKey string,
	fxRegistration string,
	fxAggregationRegistration string,
	TickersService TickersService.ITickersService) *factory {
	return &factory{
		apiKey:                    apiKey,
		fxRegistration:            fxRegistration,
		fxAggregationRegistration: fxAggregationRegistration,
		tickersService:            TickersService,
	}
}

func (self *factory) Create(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	connectionCancelFunc model.ConnectionCancelFunc,
	logger *zap.Logger,
	userContext interface{},
) (intf.IConnectionReactor, error) {
	result := NewConnectionReactor(
		logger,
		cancelCtx,
		cancelFunc,
		connectionCancelFunc,
		self.apiKey,
		self.fxRegistration,
		self.fxAggregationRegistration,
		userContext,
		self.tickersService)
	return result, nil
}
