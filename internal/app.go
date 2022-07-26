package internal

import (
	"github.com/bhbosman/goFxApp"
	//"github.com/bhbosman/goFxApp/FxWrappers"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	"github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/bhbosman/goPolygon-io/internal/wsCurrencyDialer"
	"github.com/bhbosman/gocommon/Providers"
	//"github.com/bhbosman/gocomms/connectionManager/endpoints"
	//"github.com/bhbosman/gocomms/connectionManager/view"
	"go.uber.org/fx"
)

type App struct {
	FxApp      *fx.App
	ShutDowner fx.Shutdowner
}

func NewApp(
	setting ...IAppSettings,
) *goFxApp.TerminalAppUsingFxApp {
	settingInstance := &settings{}
	for _, s := range setting {
		if s == nil {
			continue
		}
		s.apply(settingInstance)
	}

	var shutDowner fx.Shutdowner
	//var dd *gocommon.RunTimeManager

	fxApp := goFxApp.NewFxMainApplicationServices(
		"PolygonIo",
		false,
		ProvidePolygonKeys(),
		Providers.RegisterRunTimeManager(),
		wsCurrencyDialer.ProvideDialer(
			wsCurrencyDialer.MaxConnections(1),
		),
		TickersService.Provide(),
		http.Provide(),

		fx.Populate(&shutDowner),
	)
	return fxApp
}
