package wsCurrencyDialer

import (
	"github.com/bhbosman/goCommsNetDialer"
)

type canDialSetting struct {
	canDial goCommsNetDialer.ICanDial
}

func CanDial(canDial goCommsNetDialer.ICanDial) *canDialSetting {
	return &canDialSetting{
		canDial: canDial}
}

func (self canDialSetting) apply(settings *DialerSettings) {
	settings.canDial = append(settings.canDial, self.canDial)
}
