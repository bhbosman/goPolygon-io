package TickerDetailsService_test

import (
	"github.com/bhbosman/goPolygon-io/internal"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickerDetailsService"
	resthttp "github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestName(t *testing.T) {

	targets := struct {
		fx.In
		Tickers *TickerDetailsService.TickerDetailsService `name:"Polygon"`
	}{}

	app := fxtest.New(
		t,
		internal.ProvidePolygonKeys(),
		resthttp.Provide(),
		TickerDetailsService.Provide(),
		fx.Populate(&targets),
	)
	assert.NoError(t, app.Err())
	app.RequireStart()
	defer app.StartTimeout()

	detailsResponse, err := targets.Tickers.TickerDetails("AAPL")
	assert.NoError(t, err)
	assert.NotNil(t, detailsResponse)
	println(detailsResponse.String())
}
