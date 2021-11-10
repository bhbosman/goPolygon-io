package TickerNewsService_test

import (
	"github.com/bhbosman/goPolygon-io/internal"
	"github.com/bhbosman/goPolygon-io/internal/rest/ReferenceApi/TickerNewsService"
	resthttp "github.com/bhbosman/goPolygon-io/internal/rest/http"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestName(t *testing.T) {
	targets := struct {
		fx.In
		Sut *TickerNewsService.TickerNewsService `name:"Polygon"`
	}{}

	app := fxtest.New(
		t,
		internal.ProvidePolygonKeys(),
		resthttp.Provide(),
		TickerNewsService.Provide(),
		fx.Populate(&targets),
	)
	assert.NoError(t, app.Err())
	app.RequireStart()
	defer app.StartTimeout()

	detailsResponse, err := targets.Sut.TickerNews()
	assert.NoError(t, err)
	assert.NotNil(t, detailsResponse)
	println(detailsResponse.String())
}
