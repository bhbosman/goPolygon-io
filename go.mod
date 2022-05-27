module github.com/bhbosman/goPolygon-io

go 1.15

require (
	github.com/bhbosman/goMessages v0.0.0-20210414134625-4d7166d206a6
	github.com/bhbosman/gocommon v0.0.0-20220527203717-126c2c2c5ca7
	github.com/bhbosman/gocomms v0.0.0-20220527203916-c87f0e084175
	github.com/bhbosman/goerrors v0.0.0-20210201065523-bb3e832fa9ab
	github.com/bhbosman/gomessageblock v0.0.0-20210901070622-be36a3f8d303
	github.com/bhbosman/goprotoextra v0.0.2-0.20210817141206-117becbef7c7
	github.com/cskr/pubsub v1.0.2
	github.com/emirpasic/gods v1.12.0
	github.com/gdamore/tcell/v2 v2.5.1 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/kardianos/service v1.1.0
	go.uber.org/fx v1.17.1
	go.uber.org/zap v1.21.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/bhbosman/goMessages => ../goMessages

replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/bhbosman/gocommon => ../gocommon
