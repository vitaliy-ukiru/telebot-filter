package filters

import (
	"github.com/vitaliy-ukiru/filter-telebot/dispatcher"
	tb "gopkg.in/telebot.v3"
)

func Any(filters ...dispatcher.Filter) dispatcher.Filter {
	return func(c tb.Context) bool {
		for _, f := range filters {
			if f(c) {
				return true
			}
		}
		return false
	}
}

func Not(filter dispatcher.Filter) dispatcher.Filter {
	return func(c tb.Context) bool {
		return !filter(c)
	}
}

func All(filters ...dispatcher.Filter) dispatcher.Filter {
	return func(c tb.Context) bool {
		for _, f := range filters {
			if !f(c) {
				return false
			}
		}
		return true
	}
}
