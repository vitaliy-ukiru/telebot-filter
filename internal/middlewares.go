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

type Yield = func(mw tb.MiddlewareFunc) (next bool)

type Iter = func(yield Yield)

func IterateApply(h tb.HandlerFunc, iterator Iter) tb.HandlerFunc {
	iterator(func(mw tb.MiddlewareFunc) (next bool) {
		h = mw(h)
		return true
	})
	return h
}
