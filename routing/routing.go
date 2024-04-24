// Package routing is simple package for use filters.
// It fulls compatibility with other telebot handlers.
//
// But all handlers for endpoint must be installed in one bot.Handle call.
// Else it will override like in default telebot.
// Middlewares also will work like in telebot.
package routing

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

// New creates generic handler in place.
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
