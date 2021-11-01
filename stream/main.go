package main

import (
	"github.com/bhbosman/goPolygon-io/internal"
	"time"
)

func main() {
	app := internal.NewApp()

	if app.FxApp.Err() != nil {
		println(app.FxApp.Err().Error())
		return
	}
	app.FxApp.Run()
	// allow shutdown to complete
	time.Sleep(time.Second)

}
