package dispatcher_test

import (
	"github.com/vitaliy-ukiru/telebot-filter/v2/dispatcher"
	tf "github.com/vitaliy-ukiru/telebot-filter/v2/telefilter"
	tb "gopkg.in/telebot.v4"
)

var (
	router      *dispatcher.Router
	handleFunc  tb.HandlerFunc
	filter      tf.Filter
	middlewares []tb.MiddlewareFunc
)

func ExampleRouter_Handle() {
	router.Handle(
		"/start",
		tf.NewRawHandler(handleFunc, filter),
		middlewares...,
	)
}

func ExampleRouter_Dispatch() {
	router.Dispatch(
		tf.NewRoute(
			"/start",
			tf.NewRawHandler(handleFunc, filter),
			middlewares...,
		),
	)
}

func ExampleRouter_Dispatch_withoutConstructor() {
	router.Dispatch(
		tf.Route{
			Endpoint:    "/start",
			Handler:     tf.NewRawHandler(handleFunc, filter),
			Middlewares: middlewares,
		},
	)
}
