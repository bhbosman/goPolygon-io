package TickerNewsService

import (
	"go.uber.org/fx"
	"net/http"
)

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Name: "Polygon",
				Target: func(params struct {
					fx.In
					Client *http.Client `name:"Polygon"`
				}) (*TickerNewsService, error) {
					return NewTickerNewsService(params.Client), nil
				},
			},
		),
	)
}
