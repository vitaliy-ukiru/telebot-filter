package filters

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

func Any(filters ...tf.Filter) tf.Filter {
	return func(c tb.Context) bool {
		for _, f := range filters {
			if f(c) {
				return true
			}
		}
		return false
	}
}

func Not(filter tf.Filter) tf.Filter {
	return func(c tb.Context) bool {
		return !filter(c)
	}
}

func All(filters ...tf.Filter) tf.Filter {
	return func(c tb.Context) bool {
		for _, f := range filters {
			if !f(c) {
				return false
			}
		}
		return true
	}
}
