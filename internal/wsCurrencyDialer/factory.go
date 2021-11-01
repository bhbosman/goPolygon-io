package wsCurrencyDialer

import (
	"context"
	"github.com/bhbosman/gocomms/common"
	"github.com/bhbosman/gocomms/intf"
	"go.uber.org/zap"
)

type factory struct {
	crfName                   string
	apiKey                    string
	fxRegistration            string
	fxAggregationRegistration string
}

func NewConnectionReactorFactory(
	crfName string,
	apiKey string,
	fxRegistration string,
	fxAggregationRegistration string) *factory {
	return &factory{
		crfName:                   crfName,
		apiKey:                    apiKey,
		fxRegistration:            fxRegistration,
		fxAggregationRegistration: fxAggregationRegistration,
	}
}

func (self *factory) Create(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	connectionCancelFunc common.ConnectionCancelFunc,
	logger *zap.Logger,
	userContext interface{}) intf.IConnectionReactor {
	result := NewConnectionReactor(
		logger,
		cancelCtx,
		cancelFunc,
		connectionCancelFunc,
		self.apiKey,
		self.fxRegistration,
		self.fxAggregationRegistration,
		userContext)
	return result
}

func (self *factory) Values(inputValues map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	return result, nil
}

func (self *factory) Name() string {
	return self.crfName
}
