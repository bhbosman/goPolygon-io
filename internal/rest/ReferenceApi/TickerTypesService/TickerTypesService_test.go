package TickerTypesService_test

import (
	"github.com/bhbosman/goPolygon-io/internal"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickerTypesService"
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
		Client  *http.Client                           `name:"Polygon"`
		Tickers *TickerTypesService.TickerTypesService `name:"Polygon"`
	}{}

	app := fxtest.New(
		t,
		internal.ProvidePolygonKeys(),
		resthttp.Provide(),
		TickerTypesService.Provide(),
		fx.Populate(&targets),
	)
	assert.NoError(t, app.Err())
	app.RequireStart()
	defer app.StartTimeout()

	tickersResponse, err := targets.Tickers.TickerTypes(
		TickerTypesService.TickerTypesOptionAssetClass(TickerTypesService.AssetClassTypeFx))
	assert.NoError(t, err)
	assert.NotNil(t, tickersResponse)
	for _, result := range tickersResponse.Results {
		println(result.String())

	}
}
