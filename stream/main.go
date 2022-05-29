package main

import (
	"github.com/bhbosman/goPolygon-io/internal"
)

func main() {
	app := internal.NewApp()

	if app.FxApp.Err() != nil {
		println(app.FxApp.Err().Error())
		return
	}
	app.RunTerminalApp()

}
