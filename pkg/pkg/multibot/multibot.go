package multibot

import (
	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"
	tb "gopkg.in/telebot.v3"
)

type MultiBot []dispatcher.HandlerContainer

func NewMultiBot(bots ...dispatcher.HandlerContainer) MultiBot {
	return bots
}

func (m MultiBot) Use(mw ...tb.MiddlewareFunc) {
	for _, bot := range m {
		bot.Use(mw...)
	}
}

func (m MultiBot) Handle(endpoint any, h tb.HandlerFunc, mw ...tb.MiddlewareFunc) {
	for _, bot := range m {
		bot.Handle(endpoint, h, mw...)
	}
}
