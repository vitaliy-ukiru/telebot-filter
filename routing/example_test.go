package routing_test

import (
	"github.com/vitaliy-ukiru/telebot-filter/routing"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

var bot tele.Bot

var (
	handleHi     tele.HandlerFunc
	filterHiText tf.Filter

	handleWakeUpInGroup tele.HandlerFunc
	filterGroup         tf.Filter
	filterWakeUpText    tf.Filter

	middlewareStatistic tele.MiddlewareFunc
)

func ExampleNew() {
	bot.Handle(tele.OnText, routing.New(
		tf.NewRawHandler(
			handleHi,
			filterHiText,
		),
		tf.NewRawHandler(
			handleWakeUpInGroup,
			filterGroup,
			filterWakeUpText,
		),
	))
}

func ExampleNew_middlewares() {
	bot.Handle(tele.OnText, routing.New(
		tf.NewRoute(
			nil, // passing any value, it doesn't matter
			tf.NewRawHandler(handleHi, filterHiText),
			middlewareStatistic, // will execute only after filters
		),

		tf.NewRawHandler(
			handleWakeUpInGroup,
			filterGroup,
			filterWakeUpText,
		),
	))
}

func ExampleRoute_Add() {
	var route routing.Route
	route.Add(tf.NewRawHandler(handleHi, filterHiText))
	bot.Handle(tele.OnText, route.Handler)
}
