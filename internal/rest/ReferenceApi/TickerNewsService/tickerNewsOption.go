package TickerNewsService

import "fmt"

type tickerNewsOption struct {
	key   string
	value string
}

func TickerNewsOption(key string, value string) *tickerNewsOption {
	return &tickerNewsOption{
		key:   key,
		value: value,
	}
}

func (self *tickerNewsOption) applyTickerDetailsServiceOption() string {
	return fmt.Sprintf("%v=%v", self.key, self.value)
}

type tickerNewsOptionTicker struct {
	tickerNewsOption
}

func TickerNewsOptionTicker(value string) *tickerNewsOptionTicker {
	return &tickerNewsOptionTicker{
		tickerNewsOption{
			key:   "ticker",
			value: value,
		},
	}
}

type tickerNewsOptionPublishedUtc struct {
	tickerNewsOption
}

func TickerNewsOptionPublishedUtc(value string) *tickerNewsOptionPublishedUtc {
	return &tickerNewsOptionPublishedUtc{
		tickerNewsOption{
			key:   "published_utc",
			value: value,
		},
	}
}

type tickerNewsOptionOrder struct {
	tickerNewsOption
}

func TickerNewsOptionOrder(value string) *tickerNewsOptionOrder {
	return &tickerNewsOptionOrder{
		tickerNewsOption{
			key:   "order",
			value: value,
		},
	}
}

type tickerNewsOptionLimit struct {
	tickerNewsOption
}

func TickerNewsOptionLimit(value string) *tickerNewsOptionLimit {
	return &tickerNewsOptionLimit{
		tickerNewsOption{
			key:   "limit",
			value: value,
		},
	}
}

type tickerNewsOptionSort struct {
	tickerNewsOption
}

func TickerNewsOptionSort(value string) *tickerNewsOptionSort {
	return &tickerNewsOptionSort{
		tickerNewsOption{
			key:   "sort",
			value: value,
		},
	}
}
