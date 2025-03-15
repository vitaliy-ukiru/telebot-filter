package telefilter

import (
	tb "gopkg.in/telebot.v4"
)

// Handler is filters and running update.
// You can use builtin implementation [RawHandler]
// for separated filtering and running or create
// custom types.
type Handler interface {
	// Check filters update.
	Check(c tb.Context) bool

	// Execute processing update.
	Execute(c tb.Context) error
}

// Route is unit of routing handler.
// In addition to the handler, it contains a middleware.
// Middlewares must storages separated of handler, because
// it's not their area of responsibility.
//
// Uses as transfer type and can use outside module for
// additional features.
type Route struct {
	Handler
	Endpoint    any
	Middlewares []tb.MiddlewareFunc
}

func NewRoute(endpoint any, handler Handler, middlewares ...tb.MiddlewareFunc) Route {
	return Route{Handler: handler, Endpoint: endpoint, Middlewares: middlewares}
}

// RawHandler is builtin handler with separated
// filters and callback.
type RawHandler struct {
	Filter   Filter
	Callback tb.HandlerFunc
}

func NewRawHandler(callback tb.HandlerFunc, filter Filter) RawHandler {
	return RawHandler{Filter: filter, Callback: callback}
}

func (h RawHandler) Check(c tb.Context) bool {
	if h.Filter == nil {
		return true
	}

	return h.Filter(c)
}

func (h RawHandler) Execute(c tb.Context) error {
	return h.Callback(c)
}
