package wsCurrencyDialer

import (
	"fmt"
	"github.com/bhbosman/gocomms/impl"
	"github.com/bhbosman/gocomms/intf"
	"github.com/bhbosman/gocomms/netDial"
	"go.uber.org/fx"
)

func ProvideDialer(options ...IDialerSetting) fx.Option {
	settings := &DialerSettings{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option.apply(settings)
	}
	crfName := "goPolygon-io.Dialer.CRF"
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Group: "CFR",
				Target: func(params struct {
					fx.In
					ApiKey                    string `name:"Polygon-io.API.Key"`
					FxRegistration            string `name:"Polygon-io.WS.FX.Registration.C"`
					FxAggregationRegistration string `name:"Polygon-io.WS.FX.Registration.CA"`
				}) (intf.IConnectionReactorFactory, error) {
					return NewConnectionReactorFactory(
						crfName,
						params.ApiKey,
						params.FxRegistration,
						params.FxAggregationRegistration), nil
				},
			}),
		fx.Provide(
			fx.Annotated{
				Group: "Apps",
				Target: netDial.NewNetDialApp(
					fmt.Sprintf("goPolygon-io Dialer"),
					fmt.Sprintf("wss://socket.polygon.io:443/forex"),
					impl.WebSocketName,
					crfName,
					netDial.MaxConnectionsSetting(1),
					//netDial.UserContextValue(option),
					netDial.CanDial(settings.canDial...)),
			}))
}
