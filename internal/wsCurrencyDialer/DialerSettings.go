package wsCurrencyDialer

import "github.com/bhbosman/gocomms/netDial"

type DialerSettings struct {
	canDial        []netDial.ICanDial
	maxConnections int
}
