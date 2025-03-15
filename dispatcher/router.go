package dispatcher

import (
	"github.com/vitaliy-ukiru/telebot-filter/v2/internal"
	tf "github.com/vitaliy-ukiru/telebot-filter/v2/telefilter"
	tb "gopkg.in/telebot.v4"
)

// Router is main handler setup type.
//
// All routers must be linked to [Dispatcher]
// and you can legal create only with [Dispatcher.NewRouter] or [Router.NewRouter].
type Router struct {
	dp     *Dispatcher
	parent *Router
	mw     *internal.MiddlewareList
}

func newRouter(dp *Dispatcher, parent *Router) *Router {
	return &Router{dp: dp, parent: parent, mw: new(internal.MiddlewareList)}
}

// Handle is more telebot same method.
//
// The only difference is that you need to provide [tf.Handler],
// that contains filters and handler
func (r *Router) Handle(endpoint any, handler tf.Handler, mw ...tb.MiddlewareFunc) {
	route := tf.NewRoute(endpoint, handler, mw...)
	r.Dispatch(route)
}

// Dispatch is low-level API.
// This is not a user-friendly method because
// which does not support syntactic sugar and looks cumbersome.
//
// But it may find application in third-party modules or complex
// systems where it will be more convenient to use.
func (r *Router) Dispatch(route tf.Route) {
	r.dp.addRoute(route, r)
}

// Use adds middlewares for handlers of this router
// and children routers.
func (r *Router) Use(mw ...tb.MiddlewareFunc) {
	r.mw.ExtendSlice(mw)
}

// NewRouter create child router.
func (r *Router) NewRouter() *Router {
	return newRouter(r.dp, r)
}
