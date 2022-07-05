package wsCurrencyDialer

import (
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocomms/intf"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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
						ClientContext                     interface{}                    `name:"UserContext"`
						TickersService                    TickersService.ITickersService `name:"Polygon"`
						ApiKey                            string                         `name:"Polygon-io.API.Key"`
						FxCurrencyRegistration            string                         `name:"Polygon-io.WS.FX.Registration.C"`
						FxCurrencyAggregationRegistration string                         `name:"Polygon-io.WS.FX.Registration.CA"`
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
						params.ClientContext,
						params.TickersService)
					return result, nil
				},
			},
		),
	)
}
