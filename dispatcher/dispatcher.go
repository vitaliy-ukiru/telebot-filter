package dispatcher

import (
	"errors"

	"github.com/vitaliy-ukiru/telebot-filter/internal"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

type HandlerContainer interface {
	Use(mw ...tb.MiddlewareFunc)
	Handle(endpoint any, h tb.HandlerFunc, mw ...tb.MiddlewareFunc)
}

// Dispatcher is base object for handling updates.
// It gates between raw telebot bot and filters support.
//
// Handler routing methods in [Router]. Dispatcher have root router.
type Dispatcher struct {
	*Router

	g                    *tb.Group
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
	dp.Router = newRouter(dp, nil)
	return dp
}

// NewHandler is shortcut for creating builder.
func (d *Dispatcher) NewHandler(endpoint any) *Builder {
	return NewBuilder(endpoint)
}

var errRouterNotMatch = errors.New("not find matching handler")

// UseOn registers middlewares on given endpoint like tb.OnText, etc.
func (d *Dispatcher) UseOn(onEndpoint any, mw ...tb.MiddlewareFunc) {
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

	d.g.Handle(endpoint, d.wrappedEndpointHandler(endpoint))
	d.wrapped.Add(endpoint)
}

func (d *Dispatcher) addRoute(route tf.Route, router *Router) {
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
			fn = internal.ApplyMiddleware(fn, mw)
		}
		return fn(c)
	}
}
