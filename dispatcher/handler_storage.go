package dispatcher

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	"github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

type handlerRoute struct {
	telefilter.Route
	router *Router
}

type handlerStorage map[string]*container.List[handlerRoute]

func (hs handlerStorage) Process(endpoint string, c tb.Context) error {
	routes := hs[endpoint]
	for h := routes.Front(); h != nil; h = h.Next() {
		if !h.Value.Check(c) {
			continue
		}
		return h.Value.run(c)
	}

	return errRouterNotMatch
}

func (hs handlerStorage) addRoute(r handlerRoute) (endpoint string) {
	endpoint = internal.ExtractRawEndpoint(r.HandlerEndpoint())

	l := hs[endpoint]
	if l == nil {
		l = new(container.List[handlerRoute])
		hs[endpoint] = l
	}
	l.Insert(r)
	return
}

// getAllMiddlewares returns list with middlewares from all parent routers (including current)
// and handler middlewares. It saves correct structure of execution middlewares.
//
// For example
//
//	topRouter with middleware [A, B, C]
//	r2 is child of topRouter with middlewares [X, Y, Z]
//	handler routers is [1, 2]
//	execution chain will [A -> B -> C -> X -> Y -> Z -> 1 -> 2]
func (hr handlerRoute) getAllMiddlewares() *internal.MiddlewareList {
	l := new(internal.MiddlewareList)
	getMiddlewares(l, hr.router)
	l.ExtendSlice(hr.Middlewares)
	return l
}

// getMiddlewares recursively adds all middleware to the list.
func getMiddlewares(l *internal.MiddlewareList, r *Router) {
	// base case for stop.
	if r == nil {
		return
	}
	// fetch middlewares from top routers
	if r.parent != nil {
		getMiddlewares(l, r.parent)
	}
	l.Extend(r.mw)
}

func (hr handlerRoute) run(c tb.Context) error {
	callback := internal.ApplyMiddleware(
		hr.Execute,
		hr.getAllMiddlewares(),
	)

	return callback(c)
}
