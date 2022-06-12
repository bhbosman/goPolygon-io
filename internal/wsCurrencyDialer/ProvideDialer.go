package wsCurrencyDialer

import (
	"fmt"
	"github.com/bhbosman/goCommsDefinitions"
	"github.com/bhbosman/goCommsNetDialer"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocomms/common"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gocomms/netDial"
	"go.uber.org/fx"
)

func ProvideDialer(
	serviceIdentifier model.ServiceIdentifier,
	serviceDependentOn model.ServiceIdentifier,
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
				f := goCommsNetDialer.NewNetDialApp(
					fmt.Sprintf("goPolygon-io Dialer"),
					serviceIdentifier,
					serviceDependentOn,
					fmt.Sprintf("goPolygon-io Dialer"),
					fmt.Sprintf("wss://socket.polygon.io:443/forex"),
					goCommsDefinitions.WebSocketName,
					func() (intf.IConnectionReactorFactory, error) {
						return NewConnectionReactorFactory(
							crfName,
							params.ApiKey,
							params.FxCurrencyRegistration,
							params.FxCurrencyAggregationRegistration,
							params.TickersService), nil
					},
					netDial.MaxConnectionsSetting(1),
					//netDial.UserContextValue(option),
					goCommsNetDialer.CanDial(settings.canDial...))
				return f(params.NetAppFuncInParams)
			},
		})
}
