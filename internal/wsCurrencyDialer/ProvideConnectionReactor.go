package wsCurrencyDialer

import (
	"context"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/interfaces"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocomms/intf"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideConnectionReactor() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
					fx.In
					CancelCtx                         context.Context
					CancelFunc                        context.CancelFunc
					ConnectionCancelFunc              model.ConnectionCancelFunc
					Logger                            *zap.Logger
					TickersService                    TickersService.ITickersService `name:"Polygon"`
					ApiKey                            string                         `name:"Polygon-io.API.Key"`
					FxCurrencyRegistration            string                         `name:"Polygon-io.WS.FX.Registration.C"`
					FxCurrencyAggregationRegistration string                         `name:"Polygon-io.WS.FX.Registration.CA"`
					UniqueReferenceService            interfaces.IUniqueReferenceService
					PubSub                            *pubsub.PubSub
					GoFunctionCounter                 GoFunctionCounter.IService
				},
				) (intf.IConnectionReactor, error) {
					result := NewConnectionReactor(
						params.Logger,
						params.CancelCtx,
						params.CancelFunc,
						params.ConnectionCancelFunc,
						params.ApiKey,
						params.FxCurrencyRegistration,
						params.FxCurrencyAggregationRegistration,
						params.TickersService,
						params.UniqueReferenceService,
						params.PubSub,
						params.GoFunctionCounter,
					)
					return result, nil
				},
			},
		),
	)
}
