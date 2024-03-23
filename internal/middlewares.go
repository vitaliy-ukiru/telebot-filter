package internal

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tb "gopkg.in/telebot.v3"
)

type MiddlewareList = container.List[tb.MiddlewareFunc]

type Yield = func(mw tb.MiddlewareFunc) (next bool)

// Iter is copy iter.Seq[tb.MiddlewareFunc] from go1.22.
// Now is module's go version in 1.21.
// But the version increases to 1.22, the new syntax will be used.
type Iter = func(yield Yield)

func ApplyMiddleware(h tb.HandlerFunc, iterator Iter) tb.HandlerFunc {
	iterator(func(mw tb.MiddlewareFunc) (next bool) {
		h = mw(h)
		return true
	})
	return h
}
