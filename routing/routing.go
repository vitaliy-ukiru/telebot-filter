package routing

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

func New(routes ...tf.Handler) tb.HandlerFunc {
	return func(c tb.Context) error {
		for _, route := range routes {
			if !route.Check(c) {
				continue
			}

			return route.Execute(c)
		}
		return nil
	}
}
