package dispatcher

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

// Router is main handler setup type.
// All routers must be linked to [Dispatcher]
// and you can legal create only with Dispatch.NewRouter or [Router.NewRouter]
type Router struct {
	dp     *Dispatcher
	parent *Router
	mw     *internal.MiddlewareList
}

func newRouter(dp *Dispatcher, parent *Router) *Router {
	return &Router{dp: dp, parent: parent, mw: new(internal.MiddlewareList)}
}

// Bind builds and add handler from builder.
func (r *Router) Bind(b *Builder) {
	r.Dispatch(b.Build())
}

// Handle manually adds handler with middlewares
// almost like telebot.
func (r *Router) Handle(endpoint any, handler tf.Handler, mw ...tb.MiddlewareFunc) {
	route := tf.NewRoute(endpoint, handler, mw...)
	r.Dispatch(route)
}

// Dispatch most low-level api accessible from outside the module.
// I think it better way to use in addons for this module.
func (r *Router) Dispatch(route tf.Route) {
	r.dp.addRoute(route, r)
}

// Use middlewares for handlers of this router and other child routers.
func (r *Router) Use(mw ...tb.MiddlewareFunc) {
	r.mw.ExtendSlice(mw)
}

// NewRouter create child router.
func (r *Router) NewRouter() *Router {
	return newRouter(r.dp, r)
}
