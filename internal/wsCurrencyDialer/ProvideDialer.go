package wsCurrencyDialer

import (
	"fmt"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocomms/common"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gocomms/netDial"
	"go.uber.org/fx"
)

func ProvideDialer(
	serviceIdentifier model.ServiceIdentifier_,
	serviceDependentOn model.ServiceIdentifier_,
	options ...IDialerSetting) fx.Option {
	settings := &DialerSettings{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option.apply(settings)
	}
	crfName := "goPolygon-io.Dialer.CRF"
	return fx.Provide(
		fx.Annotated{
			Group: "Apps",
			Target: func(params struct {
				fx.In
				TickersService                    TickersService.ITickersService `name:"Polygon"`
				ApiKey                            string                         `name:"Polygon-io.API.Key"`
				FxCurrencyRegistration            string                         `name:"Polygon-io.WS.FX.Registration.C"`
				FxCurrencyAggregationRegistration string                         `name:"Polygon-io.WS.FX.Registration.CA"`
				NetAppFuncInParams                common.NetAppFuncInParams
			}) messages.CreateAppCallback {
				fxOptions := fx.Options(
					fx.Provide(fx.Annotated{Name: "Polygon", Target: func() TickersService.ITickersService { return params.TickersService }}),
					fx.Provide(fx.Annotated{Name: "Polygon-io.API.Key", Target: model.CreateStringContext(params.ApiKey)}),
					fx.Provide(fx.Annotated{Name: "Polygon-io.WS.FX.Registration.C", Target: model.CreateStringContext(params.FxCurrencyAggregationRegistration)}),
					fx.Provide(fx.Annotated{Name: "Polygon-io.WS.FX.Registration.CA", Target: model.CreateStringContext(params.FxCurrencyAggregationRegistration)}),
					fx.Provide(
						fx.Annotated{
							Target: func(params struct {
								fx.In
								TickersService                    TickersService.ITickersService `name:"Polygon"`
								ApiKey                            string                         `name:"Polygon-io.API.Key"`
								FxCurrencyRegistration            string                         `name:"Polygon-io.WS.FX.Registration.C"`
								FxCurrencyAggregationRegistration string                         `name:"Polygon-io.WS.FX.Registration.CA"`
							}) intf.ConnectionReactorFactoryCallback {
								return func() (intf.IConnectionReactorFactory, error) {
									return NewConnectionReactorFactory(
										crfName,
										params.ApiKey,
										params.FxCurrencyRegistration,
										params.FxCurrencyAggregationRegistration,
										params.TickersService), nil
								}
							},
						}))
				f := netDial.NewNetDialApp(
					fmt.Sprintf("goPolygon-io Dialer"), serviceIdentifier, serviceDependentOn, fxOptions,
					fmt.Sprintf("goPolygon-io Dialer"),
					fmt.Sprintf("wss://socket.polygon.io:443/forex"),
					common.WebSocketName,
					func() (intf.IConnectionReactorFactory, error) {
						return nil, fmt.Errorf("asdffdf")
					},
					netDial.MaxConnectionsSetting(1),
					//netDial.UserContextValue(option),
					netDial.CanDial(settings.canDial...))
				return f(params.NetAppFuncInParams)
			},
		})
}
