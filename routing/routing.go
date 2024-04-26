// Package routing is simple package for use filters.
// It fulls compatibility with other telebot handlers.
//
// But all handlers for endpoint must be installed in one bot.Handle call.
// Else it will override like in default telebot.
//
// You can add middlewares to current handler via using [tf.Route].
// These middlewares will execute after handler test, but before Handler.Execute
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

// Route is container for handler.
// Difference with [New] function is that it can
// add handlers in runtime.
//
// You can install route.Handler to telebot and after
// add handlers.
type Route struct {
	handlers *container.List[tf.Handler]
}

func (r *Route) Add(handler tf.Handler) {
	if r.handlers == nil {
		r.handlers = new(container.List[tf.Handler])
	}

	if handler == nil {
		panic("routing: handler must be not nil")
	}

	r.handlers.Insert(handler)
}

func (r *Route) Handler(c tb.Context) error {
	if r.handlers == nil {
		return nil
	}

	for e := r.handlers.Front(); e != nil; e = e.Next() {
		h := e.Value
		if !h.Check(c) {
			continue
		}

		return h.Execute(c)
	}
	return nil
}
