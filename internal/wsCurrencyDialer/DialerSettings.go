package wsCurrencyDialer

import (
	"github.com/bhbosman/goCommsNetDialer"
)

type DialerSettings struct {
	canDial        []goCommsNetDialer.ICanDial
	maxConnections int
}
