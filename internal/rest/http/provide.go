package http

import (
	"go.uber.org/fx"
	"net/http"
)

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Name: "Default",
				Target: func() (http.RoundTripper, error) {
					return http.DefaultTransport, nil
				},
			}),
		fx.Provide(
			fx.Annotated{
				Name: "Polygon",
				Target: func(params struct {
					fx.In
					Transport http.RoundTripper `name:"Polygon"`
				}) (*http.Client, error) {
					return &http.Client{
						Transport:     params.Transport,
						CheckRedirect: nil,
						Jar:           nil,
						Timeout:       0,
					}, nil
				},
			}),
		fx.Provide(
			fx.Annotated{
				Name: "Polygon",
				Target: func(params struct {
					fx.In
					DefaultTransport http.RoundTripper `name:"Default"`
					ApiKey           string            `name:"Polygon-io.API.Key"`
				}) (http.RoundTripper, error) {
					return NewTransport(
						params.ApiKey,
						params.DefaultTransport), nil
				},
			}),
	)
}
