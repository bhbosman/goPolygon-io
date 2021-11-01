package wsCurrencyDialer

type IDialerSetting interface {
	apply(*DialerSettings)
}
