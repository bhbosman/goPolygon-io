package TickersService_test

import (
	"github.com/bhbosman/goPolygon-io/internal"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	resthttp "github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestTickersService(t *testing.T) {

	targets := struct {
		fx.In
		Tickers *TickersService.TickersService `name:"Polygon"`
	}{}

	app := fxtest.New(
		t,
		internal.ProvidePolygonKeys(),
		resthttp.Provide(),
		TickersService.Provide(),
		fx.Populate(&targets),
	)
	assert.NoError(t, app.Err())
	app.RequireStart()
	defer app.StartTimeout()
	t.Run("Ask For Apple", func(t *testing.T) {
		tickersResponse, err := targets.Tickers.Tickers(
			TickersService.TickersOptionActive(true),
			TickersService.TickersOptionTicker("AAPL"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, tickersResponse)
		assert.Len(t, tickersResponse.Results, 1)
		assert.Equal(t, "AAPL", tickersResponse.Results[0].Ticker)
		assert.Equal(t, "CS", tickersResponse.Results[0].Type)
	})
	t.Run("Ask For Fx", func(t *testing.T) {
		tickersResponse, err := targets.Tickers.Tickers(
			TickersService.TickersOptionActive(true),
			TickersService.TickersOptionLimit(100),
			TickersService.TickersOptionMarket("fx"),
		)
		assert.NoError(t, err)
		assert.NotNil(t, tickersResponse)
		assert.Len(t, tickersResponse.Results, 100)
	})
}
