package wsCurrencyDialer

import "github.com/bhbosman/gocomms/netDial"

type canDialSetting struct {
	canDial netDial.ICanDial
}

func CanDial(canDial netDial.ICanDial) *canDialSetting {
	return &canDialSetting{
		canDial: canDial}
}

func (self canDialSetting) apply(settings *DialerSettings) {
	settings.canDial = append(settings.canDial, self.canDial)
}
