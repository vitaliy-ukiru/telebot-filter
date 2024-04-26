package telefilter

import tb "gopkg.in/telebot.v3"

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
