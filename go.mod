module github.com/bhbosman/goPolygon-io

go 1.18

require (
	github.com/bhbosman/goCommsDefinitions v0.0.0-20220801175552-c5aa68065af3
	github.com/bhbosman/goCommsNetDialer v0.0.0-20220623084808-6238b44079be
	github.com/bhbosman/goCommsStacks v0.0.0-20220628053835-8e5ac6a0c20b
	github.com/bhbosman/goFxApp v0.0.0-20220627212218-a29bbf40d76d
	github.com/bhbosman/gocommon v0.0.0-20230328150634-566a0f916878
	github.com/bhbosman/gocomms v0.0.0-20220628074707-e93417aaaed2
	github.com/bhbosman/goerrors v0.0.0-20220623084908-4d7bbcd178cf
	github.com/bhbosman/gomessageblock v0.0.0-20220617132215-32f430d7de62
	github.com/bhbosman/goprotoextra v0.0.2-0.20210817141206-117becbef7c7
	github.com/golang/protobuf v1.5.2
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/stretchr/testify v1.8.0
	go.uber.org/fx v1.19.2
	go.uber.org/zap v1.23.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/cskr/pubsub v1.0.2
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
)

require (
	github.com/bhbosman/goConnectionManager v0.0.0-20220802122816-d9db41f332dc // indirect
	github.com/bhbosman/goFxAppManager v0.0.0-20220730103022-02c299931769 // indirect
	github.com/bhbosman/goUi v0.0.0-20230328181044-49e31970d158
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.5.1 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/icza/gox v0.0.0-20220321141217-e2d488ab2fbc // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/tview v0.0.0-20220610163003-691f46d6f500 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gdamore/tcell/v2 => github.com/bhbosman/tcell/v2 v2.5.2-0.20220624055704-f9a9454fab5b

replace github.com/golang/mock => github.com/bhbosman/gomock v1.6.1-0.20230302060806-d02c40b7514e

replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/bhbosman/goFxAppManager => ../goFxAppManager

//replace github.com/bhbosman/gocommon => ../gocommon

replace github.com/bhbosman/goCommsStacks => ../goCommsStacks

replace github.com/bhbosman/goCommsNetDialer => ../goCommsNetDialer

replace github.com/bhbosman/goCommsNetListener => ../goCommsNetListener

replace github.com/bhbosman/goCommsDefinitions => ../goCommsDefinitions

replace github.com/bhbosman/goFxApp => ../goFxApp

//replace github.com/bhbosman/goUi => ../goUi

replace github.com/bhbosman/goerrors => ../goerrors

replace github.com/bhbosman/goConnectionManager => ../goConnectionManager

replace github.com/rivo/tview => ../tview

replace github.com/bhbosman/goprotoextra => ../goprotoextra

replace github.com/cskr/pubsub => github.com/bhbosman/pubsub v1.0.3-0.20220802200819-029949e8a8af
