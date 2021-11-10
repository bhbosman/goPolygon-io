package TickerTypesService

import "fmt"

type tickerTypesOption struct {
	key   string
	value string
}

func TickerTypesOption(key string, value string) *tickerTypesOption {
	return &tickerTypesOption{key: key, value: value}
}

func (self *tickerTypesOption) applyTickerTypesOption() string {
	return fmt.Sprintf("%v=%v", self.key, self.value)
}

type tickerTypesOptionAssetClass struct {
	tickerTypesOption
}

type AssetClassType int

const (
	AssetClassTypeStocks = iota
	AssetClassTypeOptions
	AssetClassTypeCrypto
	AssetClassTypeFx
)

func TickerTypesOptionAssetClass(value AssetClassType) *tickerTypesOptionAssetClass {
	return &tickerTypesOptionAssetClass{
		tickerTypesOption: tickerTypesOption{
			key: "asset_class",
			value: func() string {
				switch value {
				case AssetClassTypeStocks:
					return "stocks"
				case AssetClassTypeOptions:
					return "options"
				case AssetClassTypeCrypto:
					return "crypto"
				case AssetClassTypeFx:
					return "fx"
				default:
					return ""
				}
			}(),
		},
	}
}

type tickerTypesOptionLocale struct {
	tickerTypesOption
}

type LocaleType int

const (
	LocaleTypeUs = iota
	LocaleTypeGlobal
)

func TickerTypesOptionLocale(value LocaleType) *tickerTypesOptionLocale {
	return &tickerTypesOptionLocale{
		tickerTypesOption: tickerTypesOption{
			key: "locale",
			value: func() string {
				switch value {
				case LocaleTypeUs:
					return "us"
				case LocaleTypeGlobal:
					return "global"
				default:
					return ""
				}
			}(),
		},
	}
}
