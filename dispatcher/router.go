package dispatcher

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal"
	"github.com/vitaliy-ukiru/telebot-filter/telefilter"
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
func (r *Router) Bind(b *Builder) (endpoint string) {
	return r.Dispatch(b.Build())
}

// Handle manually adds handler with middlewares
// almost like telebot.
func (r *Router) Handle(handler telefilter.Handler, mw ...tb.MiddlewareFunc) (endpoint string) {
	h := telefilter.Route{
		Handler:     handler,
		Middlewares: mw,
	}
	return r.Dispatch(h)
}

// Dispatch most low-level api accessible from outside the module.
// I think it better way to use in addons for this module.
func (r *Router) Dispatch(route telefilter.Route) (endpoint string) {
	endpoint = internal.ExtractRawEndpoint(route.HandlerEndpoint())

	r.dp.addRoute(route, r)
	return
}

// Use middlewares for handlers of this router and other child routers.
func (r *Router) Use(mw ...tb.MiddlewareFunc) {
	r.mw.ExtendSlice(mw)
}

// NewRouter create child router.
func (r *Router) NewRouter() *Router {
	return newRouter(r.dp, r)
}
