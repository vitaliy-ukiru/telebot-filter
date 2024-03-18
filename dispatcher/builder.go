package dispatcher

import (
	"github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

// Builder for handler.
//
// For setup built handler use [Router.Bind]
// or this (Bind) method from dispatcher instance.
type Builder struct {
	h  telefilter.RawHandler
	mw []tb.MiddlewareFunc
}

func NewBuilder(endpoint any) *Builder {
	return &Builder{h: telefilter.RawHandler{Endpoint: endpoint}}
}

func (b *Builder) Use(mw ...tb.MiddlewareFunc) *Builder {
	b.mw = append(b.mw, mw...)
	return b
}

func (b *Builder) On(e any) *Builder {
	b.h.Endpoint = e
	return b
}

func (b *Builder) Filter(filters ...telefilter.Filter) *Builder {
	b.h.Filters = append(b.h.Filters, filters...)
	return b
}

func (b *Builder) Do(h tb.HandlerFunc) *Builder {
	b.h.Callback = h
	return b
}

func (b *Builder) Build() telefilter.Route {
	return telefilter.Route{
		b.h,
		b.mw,
	}
}
