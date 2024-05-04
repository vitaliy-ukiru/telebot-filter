package routing_test

import (
	"github.com/vitaliy-ukiru/telebot-filter/routing"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

var bot tb.Bot

var (
	handleHi     tb.HandlerFunc
	filterHiText tf.Filter

	handleWakeUpInGroup tb.HandlerFunc
	filterGroup         tf.Filter
	filterWakeUpText    tf.Filter
)

func ExampleNew() {
	bot.Handle(tb.OnText, routing.New(
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

func ExampleRoute_Add() {
	var route routing.Route
	route.Add(tf.NewRawHandler(handleHi, filterHiText))
	bot.Handle(tb.OnText, route.Handler)
}
