package routing_test

import (
	"github.com/vitaliy-ukiru/telebot-filter/routing"
	"github.com/vitaliy-ukiru/telebot-filter/telefilter"
	"gopkg.in/telebot.v3"
)

var bot telebot.Bot

var (
	handleHi     telebot.HandlerFunc
	filterHiText telefilter.Filter

	handleWakeUpInGroup telebot.HandlerFunc
	filterGroup         telefilter.Filter
	filterWakeUpText    telefilter.Filter
)

func ExampleNew() {
	bot.Handle(telebot.OnText, routing.New(
		telefilter.NewRawHandler(
			handleHi,
			filterHiText,
		),
		telefilter.NewRawHandler(
			handleWakeUpInGroup,
			filterGroup,
			filterWakeUpText,
		),
	))
}

func ExampleRoute_Add() {
	var route routing.Route
	route.Add(telefilter.NewRawHandler(handleHi, filterHiText))
	bot.Handle(telebot.OnText, route.Handler)
}
