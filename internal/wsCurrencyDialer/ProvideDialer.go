package wsCurrencyDialer

import (
	"fmt"
	"github.com/bhbosman/goCommsDefinitions"
	"github.com/bhbosman/goCommsNetDialer"
	"github.com/bhbosman/goCommsStacks/bottom"
	"github.com/bhbosman/goCommsStacks/top"
	"github.com/bhbosman/goCommsStacks/websocket"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocomms/common"
	"github.com/bhbosman/gocomms/intf"
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
			Target: func(
				params struct {
					fx.In
					TickersService                    TickersService.ITickersService `name:"Polygon"`
					ApiKey                            string                         `name:"Polygon-io.API.Key"`
					FxCurrencyRegistration            string                         `name:"Polygon-io.WS.FX.Registration.C"`
					FxCurrencyAggregationRegistration string                         `name:"Polygon-io.WS.FX.Registration.CA"`
					NetAppFuncInParams                common.NetAppFuncInParams
				},
			) (messages.CreateAppCallback, error) {
				f := goCommsNetDialer.NewNetDialApp(
					fmt.Sprintf("goPolygon-io Dialer"),
					serviceIdentifier,
					serviceDependentOn,
					fmt.Sprintf("goPolygon-io Dialer"),
					fmt.Sprintf("wss://socket.polygon.io:443/forex"),
					goCommsDefinitions.WebSocketName,
					common.NewConnectionInstanceOptions(
						goCommsDefinitions.ProvideTransportFactoryForWebSocketName(
							top.ProvideTopStack(),
							bottom.ProvideBottomStack(),
							websocket.ProvideWebsocketStacks(),
						),
						fx.Provide(
							fx.Annotated{
								Target: func() (intf.IConnectionReactorFactory, error) {
									return NewConnectionReactorFactory(
										crfName,
										params.ApiKey,
										params.FxCurrencyRegistration,
										params.FxCurrencyAggregationRegistration,
										params.TickersService), nil
								},
							},
						),
					),
					common.MaxConnectionsSetting(1),
					goCommsNetDialer.CanDial(settings.canDial...))
				return f(params.NetAppFuncInParams), nil
			},
		})
}
