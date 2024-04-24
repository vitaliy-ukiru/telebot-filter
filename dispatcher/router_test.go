package dispatcher_test

import (
	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	"gopkg.in/telebot.v3"
)

var (
	router      *dispatcher.Router
	handleFunc  telebot.HandlerFunc
	filters     []tf.Filter
	middlewares []telebot.MiddlewareFunc
)

func ExampleRouter_Handle() {
	router.Handle(
		"/start",
		tf.NewRawHandler(handleFunc, filters...),
		middlewares...,
	)
}

func ExampleRouter_Bind() {
	router.Bind(
		dispatcher.
			NewBuilder("/start").
			Filter(filters...).
			Do(handleFunc).
			Use(middlewares...),
	)
}

func ExampleRouter_Dispatch() {
	router.Dispatch(
		tf.NewRoute(
			"/start",
			tf.NewRawHandler(handleFunc, filters...),
			middlewares...,
		),
	)
}

func ExampleRouter_Dispatch_withoutConstructor() {
	router.Dispatch(
		tf.Route{
			Endpoint:    "/start",
			Handler:     tf.NewRawHandler(handleFunc, filters...),
			Middlewares: middlewares,
		},
	)
}
