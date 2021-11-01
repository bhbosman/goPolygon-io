package TickersService_test

import (
	"github.com/bhbosman/goPolygon-io/internal"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickersService"
	resthttp "github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {

	targets := struct {
		fx.In
		Client  *http.Client                   `name:"Polygon"`
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

	tickersResponse, err := targets.Tickers.Tickers(
		TickersService.TickersOptionActive(true),
		TickersService.TickersOptionLimit(1000),
		TickersService.TickersOptionMarket("fx"),
		//TickersService.TickerOptionTicker("AAPL"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, tickersResponse)
	for {
		for _, result := range tickersResponse.Results {
			println(result.String())

		}
		if tickersResponse.NextUrl == "" {
			break
		}
		tickersResponse, err = targets.Tickers.TickersNext(tickersResponse.NextUrl)
	}
}
