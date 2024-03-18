package internal

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tb "gopkg.in/telebot.v3"
)

type MiddlewareList = container.List[tb.MiddlewareFunc]

// ApplyMiddleware is copy of tele.applyMiddleware.
// For support middlewares in handlers we need packs middlewares independently
// of telebot.
func ApplyMiddleware(h tb.HandlerFunc, m *MiddlewareList) tb.HandlerFunc {
	for e := m.Back(); e != nil; e = e.Prev() {
		h = e.Value(h)
	}
	return h
}
