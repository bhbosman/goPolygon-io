package TickersService

import (
	"fmt"
	"strconv"
)

type tickersOption struct {
	key   string
	value string
}

func TickersOption(key string, value string) *tickersOption {
	return &tickersOption{key: key, value: value}
}

func (self *tickersOption) applyTickersOption() string {
	return fmt.Sprintf("%v=%v", self.key, self.value)
}

type tickersOptionTicker struct {
	tickersOption
}

func TickersOptionTicker(value string) *tickersOptionTicker {
	return &tickersOptionTicker{
		tickersOption{
			key:   "ticker",
			value: value,
		},
	}
}

type tickersOptionType struct {
	tickersOption
}

func TickersOptionType(value string) *tickersOptionType {
	return &tickersOptionType{
		tickersOption{
			key:   "type",
			value: value,
		},
	}
}

type tickersOptionMarket struct {
	tickersOption
}

func TickersOptionMarket(value string) *tickersOptionMarket {
	return &tickersOptionMarket{
		tickersOption{
			key:   "market",
			value: value,
		},
	}
}

type tickersOptionExchange struct {
	tickersOption
}

func TickersOptionExchange(value string) *tickersOptionExchange {
	return &tickersOptionExchange{
		tickersOption{
			key:   "exchange",
			value: value,
		},
	}
}

type tickersOptionCusip struct {
	tickersOption
}

func TickersOptionCusip(value string) *tickersOptionCusip {
	return &tickersOptionCusip{
		tickersOption{
			key:   "cusip",
			value: value,
		},
	}
}

type tickersOptionCik struct {
	tickersOption
}

func TickersOptionCik(value string) *tickersOptionCik {
	return &tickersOptionCik{
		tickersOption{
			key:   "cik",
			value: value,
		},
	}
}

type tickersOptionDate struct {
	tickersOption
}

func TickersOptionDate(value string) *tickersOptionDate {
	return &tickersOptionDate{
		tickersOption{
			key:   "date",
			value: value,
		},
	}
}

type tickersOptionSearch struct {
	tickersOption
}

func TickersOptionSearch(value string) *tickersOptionSearch {
	return &tickersOptionSearch{
		tickersOption{
			key:   "search",
			value: value,
		},
	}
}

type tickersOptionActive struct {
	tickersOption
}

func TickersOptionActive(value bool) *tickersOptionActive {
	return &tickersOptionActive{
		tickersOption{
			key: "active",
			value: func() string {
				switch value {
				case true:
					return "true"
				default:
					return "false"
				}
			}(),
		},
	}
}

type tickersOptionSort struct {
	tickersOption
}

func TickersOptionSort(value string) *tickersOptionSort {
	return &tickersOptionSort{
		tickersOption{
			key:   "sort",
			value: value,
		},
	}
}

type tickersOptionOrder struct {
	tickersOption
}

func TickersOptionOrder(value string) *tickersOptionOrder {
	return &tickersOptionOrder{
		tickersOption{
			key:   "order",
			value: value,
		},
	}
}

type tickersOptionLimit struct {
	tickersOption
}

func TickersOptionLimit(value int) *tickersOptionLimit {
	return &tickersOptionLimit{
		tickersOption{
			key:   "limit",
			value: strconv.Itoa(value),
		},
	}
}
