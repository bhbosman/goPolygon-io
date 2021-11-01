package internal

import "go.uber.org/fx"

func ProvidePolygonKeys() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Target: func() (*polygonKeys, error) {
					return ReadPolygonKeys()
				},
			}),
		fx.Provide(
			fx.Annotated{
				Name: "Polygon-io.API.Key",
				Target: func(data *polygonKeys) string {
					return "odC_XFlVPFMXEk8hxEKvLYsKnNP3jx6R"
				},
			}),
		fx.Provide(
			fx.Annotated{
				Name: "Polygon-io.WS.FX",
				Target: func(data *polygonKeys) string {
					return "wss://socket.polygon.io:443/forex"
				},
			}),
		fx.Provide(
			fx.Annotated{
				Name: "Polygon-io.WS.FX.Registration.C",
				Target: func(data *polygonKeys) string {
					return "C.*"
				},
			}),
		fx.Provide(
			fx.Annotated{
				Name: "Polygon-io.WS.FX.Registration.CA",
				Target: func(data *polygonKeys) string {
					return "CA.*"
				},
			}),
	)
}
