package internal

import (
	"github.com/bhbosman/goPolygon-io/internal/wsCurrencyDialer"
	"github.com/bhbosman/gocommon/app"
	"github.com/bhbosman/gocomms/connectionManager"
	"github.com/bhbosman/gocomms/connectionManager/endpoints"
	"github.com/bhbosman/gocomms/connectionManager/view"
	"github.com/bhbosman/gocomms/netDial"
	"github.com/bhbosman/gocomms/provide"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type App struct {
	FxApp      *fx.App
	ShutDowner fx.Shutdowner
}

func NewDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewDevelopment(options ...zap.Option) (*zap.Logger, error) {
	return NewDevelopmentConfig().Build(options...)
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
		fx.Provide(func() (*zap.Logger, error) { return NewDevelopment() }),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: logger} }),
		ProvidePolygonKeys(),
		wsCurrencyDialer.ProvideDialer(wsCurrencyDialer.MaxConnections(1), wsCurrencyDialer.CanDial(ConsumerCounter)),
		app.RegisterRootContext(),
		connectionManager.RegisterDefaultConnectionManager(),
		provide.RegisterHttpHandler("http://127.0.0.1:8084"),
		endpoints.RegisterConnectionManagerEndpoint(),
		view.RegisterConnectionsHtmlTemplate(),

		fx.Populate(&shutDowner),
		//fx.Populate(&dd),
		app.InvokeApps(),
	)
	return &App{
		FxApp:      fxApp,
		ShutDowner: shutDowner,
	}
}
