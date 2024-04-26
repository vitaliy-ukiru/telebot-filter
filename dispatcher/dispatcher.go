package dispatcher

import (
	"errors"

	"github.com/vitaliy-ukiru/telebot-filter/internal"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

// HandlerContainer represent object for communicate with telebot.
// You can pass telebot.Group, telebot.Bot, or you custom type.
//
// For example, it can be multi bot abstraction (See pkg/multibot).
type HandlerContainer interface {
	Use(mw ...tb.MiddlewareFunc)
	Handle(endpoint any, h tb.HandlerFunc, mw ...tb.MiddlewareFunc)
}

// Dispatcher is base object for handling updates.
// It gates between raw telebot bot and filters support.
//
// You can use inject dispatcher with any
// object that implements [HandlerContainer],
// e.g. telebot.Bot, telebot.Group or you custom implementation.
//
// Dispatcher has root router and wrappers for router methods.
// And all routers will eventually be children of the root.
//
// You can add middlewares on endpoint.
// It will run only on selected endpoint like tele.OnText, etc...
//
// Also, dispatcher tries effective adds handlers to telebot.
// Because handler is closure function it'll make allocation.
type Dispatcher struct {
	router *Router

	bot                  HandlerContainer
	wrapped              container.Set[string]
	handlers             handlerStorage
	endpointsMiddlewares map[string]*internal.MiddlewareList
}

func NewDispatcher(bot HandlerContainer) *Dispatcher {
	dp := &Dispatcher{
		bot:                  bot,
		handlers:             make(handlerStorage),
		wrapped:              make(container.Set[string]),
		endpointsMiddlewares: make(map[string]*internal.MiddlewareList),
	}
	dp.router = newRouter(dp, nil)
	return dp
}

// B is shortcut for creating builder.
func (*Dispatcher) B(endpoint any) *Builder {
	return NewBuilder(endpoint)
}

var errRouterNotMatch = errors.New("not find matching handler")

// UseOuter registers middlewares on given endpoint like tb.OnText, etc.
//
// It will be execute middlewares before filters.
func (d *Dispatcher) UseOuter(onEndpoint any, mw ...tb.MiddlewareFunc) {
	endpoint := internal.ExtractRawEndpoint(onEndpoint)

	list := d.endpointsMiddlewares[endpoint]
	if list == nil {
		list = new(internal.MiddlewareList)
		d.endpointsMiddlewares[endpoint] = list
	}
	list.ExtendSlice(mw)
}

// wrapEndpoint add handler to telebot with "cache".
// We may not add to those endpoint that are already wrapped.
func (d *Dispatcher) wrapEndpoint(endpoint string) {
	if d.wrapped.Has(endpoint) {
		return
	}

	// don't provide endpoint middlewares here
	// because user can add middlewares later
	d.bot.Handle(endpoint, d.wrappedEndpointHandler(endpoint))
	d.wrapped.Add(endpoint)
}

func (d *Dispatcher) addRoute(route tf.Route, router *Router) {
	if route.Handler == nil {
		panic("telebot-filter: dispatcher: handler must be not nil")
	}

	endpoint := d.handlers.addRoute(handlerRoute{
		Route:  route,
		router: router,
	})
	d.wrapEndpoint(endpoint)
}

func (d *Dispatcher) wrappedEndpointHandler(endpoint string) tb.HandlerFunc {
	fn := func(teleCtx tb.Context) error {
		err := d.handlers.Process(endpoint, teleCtx)
		if errors.Is(err, errRouterNotMatch) {
			return nil
		}
		return err
	}
	return func(c tb.Context) error {
		mw := d.endpointsMiddlewares[endpoint]
		if mw != nil && mw.Len() > 0 {
			fn = internal.ApplyMiddleware(fn, mw.IterateBackward)
		}
		return fn(c)
	}
}

// Bind builds and add handler from builder to root router.
//
// See details in [Router.Bind].
func (d *Dispatcher) Bind(b *Builder) {
	d.router.Bind(b)
}

// Handle manually adds handler with middlewares to root router
// almost like telebot.
//
// See details in [Router.Handle].
func (d *Dispatcher) Handle(endpoint any, handler tf.Handler, mw ...tb.MiddlewareFunc) {
	d.router.Handle(endpoint, handler, mw...)
}

// Dispatch is low-level public API.
// See details in [Router.Dispatch].
func (d *Dispatcher) Dispatch(route tf.Route) {
	d.router.Dispatch(route)
}

// Use adds middlewares to all handler, because
// saves it to root router.
func (d *Dispatcher) Use(mw ...tb.MiddlewareFunc) {
	d.router.Use(mw...)
}

// NewRouter creates new router from root.
func (d *Dispatcher) NewRouter() *Router {
	return d.router.NewRouter()
}
