package internal

import (
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/bhbosman/goPolygon-io/internal/wsCurrencyDialer"
	"github.com/bhbosman/gocommon/FxWrappers"
	"github.com/bhbosman/gocommon/Providers"
	"github.com/bhbosman/gocomms/connectionManager/CMImpl"
	"github.com/bhbosman/gocomms/connectionManager/endpoints"
	"github.com/bhbosman/gocomms/connectionManager/view"
	"github.com/bhbosman/gocomms/netDial"
	"go.uber.org/fx"
)

type App struct {
	FxApp      *fx.App
	ShutDowner fx.Shutdowner
}

func NewApp(setting ...IAppSettings) *App {
	settingInstance := &settings{}
	for _, s := range setting {
		if s == nil {
			continue
		}
		s.apply(settingInstance)
	}

	var shutDowner fx.Shutdowner
	ConsumerCounter := netDial.NewCanDialDefaultImpl()
	//var dd *gocommon.RunTimeManager

	fxApp := FxWrappers.NewFxMainApplicationServices(
		"PolygonIo",
		false,
		ProvidePolygonKeys(),
		Providers.RegisterRunTimeManager(),

		CMImpl.RegisterDefaultConnectionManager(),
		wsCurrencyDialer.ProvideDialer(wsCurrencyDialer.MaxConnections(1), wsCurrencyDialer.CanDial(ConsumerCounter)),
		endpoints.RegisterConnectionManagerEndpoint(),
		view.RegisterConnectionsHtmlTemplate(),
		TickersService.Provide(),
		http.Provide(),

		fx.Populate(&shutDowner),
	)
	return &App{
		FxApp:      fxApp,
		ShutDowner: shutDowner,
	}
}
