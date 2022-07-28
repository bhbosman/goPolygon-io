package wsCurrencyDialer

import (
	"fmt"
	"github.com/bhbosman/goCommsDefinitions"
	"github.com/bhbosman/goCommsNetDialer"
	"github.com/bhbosman/goCommsStacks/bottom"
	"github.com/bhbosman/goCommsStacks/topStack"
	"github.com/bhbosman/goCommsStacks/websocket"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/fx/PubSub"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocomms/common"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"net/url"
)

func ProvideDialer(
	options ...IDialerSetting) fx.Option {
	settings := &DialerSettings{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option.apply(settings)
	}
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
					PubSub                            *pubsub.PubSub `name:"Application"`
					GoFunctionCounter                 GoFunctionCounter.IService
				},
			) (messages.CreateAppCallback, error) {
				u, e := url.Parse("wss://socket.polygon.io:443/forex")
				if e != nil {
					return messages.CreateAppCallback{}, e
				}
				f := goCommsNetDialer.NewSingleNetDialApp(
					fmt.Sprintf("goPolygon-io Dialer"),
					common.MoreOptions(
						goCommsDefinitions.ProvideUrl("ConnectionUrl", u),
						goCommsDefinitions.ProvideUrl("ProxyUrl", nil),
						goCommsDefinitions.ProvideBool("UseProxy", false),
					), common.NewConnectionInstanceOptions(
						goCommsDefinitions.ProvideTransportFactoryForWebSocketName(
							topStack.ProvideTopStack(),
							bottom.Provide(),
							websocket.ProvideWebsocketStacks(),
						),
						goCommsDefinitions.ProvideStringContext("Polygon-io.API.Key", params.ApiKey),
						goCommsDefinitions.ProvideStringContext("Polygon-io.WS.FX.Registration.C", params.FxCurrencyRegistration),
						goCommsDefinitions.ProvideStringContext("Polygon-io.WS.FX.Registration.CA", params.FxCurrencyAggregationRegistration),
						PubSub.ProvidePubSubInstance("Application", params.PubSub),
						fx.Provide(
							fx.Annotated{
								Target: func() GoFunctionCounter.IService {
									return params.GoFunctionCounter
								},
							},
						),
						fx.Provide(
							fx.Annotated{
								Name: "Polygon",
								Target: func() TickersService.ITickersService {
									return params.TickersService
								},
							},
						),
						ProvideConnectionReactor(),
					),
					common.MaxConnectionsSetting(1),
					goCommsNetDialer.CanDial(settings.canDial...))
				return f(params.NetAppFuncInParams), nil
			},
		},
	)
}
