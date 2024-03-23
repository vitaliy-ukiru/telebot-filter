package dispatcher

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

type handlerRoute struct {
	tf.Route
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
	endpoint = internal.ExtractRawEndpoint(r.Endpoint)

	l := hs[endpoint]
	if l == nil {
		l = new(container.List[handlerRoute])
		hs[endpoint] = l
	}
	l.Insert(r)
	return
}

// overAllMiddlewares returns iterator with middlewares
// from all parent routers (including current) and handler
// middlewares.
// It saves correct order of execution middlewares.
//
// For example
//
//	root router with middleware [A, B, C]
//	r2 is child of root with middlewares [X, Y, Z]
//	handler's middlewares is [1, 2]
//	execution chain will [A -> B -> C -> X -> Y -> Z -> 1 -> 2]
func (hr handlerRoute) overAllMiddlewares(yield internal.Yield) {
	for i := len(hr.Middlewares) - 1; i >= 0; i-- {
		if !yield(hr.Middlewares[i]) {
			return
		}
	}

	for r := hr.router; r != nil; r = r.parent {
		// don't use List.IterateBackward, because iter.Seq
		// must be without return values.
		// But without return value we can't catch stop signal from yield.
		for e := r.mw.Back(); e != nil; e = e.Prev() {
			if !yield(e.Value) {
				return
			}
		}

	}
}

func (hr handlerRoute) run(c tb.Context) error {
	callback := internal.ApplyMiddleware(
		hr.Route.Handler.Execute,
		hr.overAllMiddlewares,
	)

	return callback(c)
}
