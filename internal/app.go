package internal

import (
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/bhbosman/goPolygon-io/internal/wsCurrencyDialer"
	"github.com/bhbosman/gocommon/Services/implementations"
	"github.com/bhbosman/gocommon/app"
	"github.com/bhbosman/gocommon/logSettings"
	"github.com/bhbosman/gocomms/connectionManager"
	"github.com/bhbosman/gocomms/connectionManager/endpoints"
	"github.com/bhbosman/gocomms/connectionManager/view"
	"github.com/bhbosman/gocomms/netDial"
	"github.com/bhbosman/gocomms/provide"
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

	fxApp := fx.New(
		ProvidePolygonKeys(),
		logSettings.ProvideZapConfig(),
		wsCurrencyDialer.ProvideDialer(wsCurrencyDialer.MaxConnections(1), wsCurrencyDialer.CanDial(ConsumerCounter)),
		app.RegisterRootContext(),
		connectionManager.RegisterDefaultConnectionManager(),
		provide.RegisterHttpHandler("http://127.0.0.1:8084"),
		endpoints.RegisterConnectionManagerEndpoint(),
		view.RegisterConnectionsHtmlTemplate(),
		implementations.ProvideNewUniqueReferenceService(),
		implementations.ProvideUniqueSessionNumber(),
		TickersService.Provide(),
		http.Provide(),

		fx.Populate(&shutDowner),
		//fx.Populate(&dd),
		app.InvokeApps(),
	)
	return &App{
		FxApp:      fxApp,
		ShutDowner: shutDowner,
	}
}
