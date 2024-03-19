package dispatcher

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

// Builder for handler.
//
// For setup built handler use [Router.Bind]
// or this (Bind) method from dispatcher instance.
type Builder struct {
	endpoint any
	h        tf.RawHandler
	mw       []tb.MiddlewareFunc
}

func NewBuilder(endpoint any) *Builder {
	return &Builder{endpoint: endpoint}
}

func (b *Builder) Use(mw ...tb.MiddlewareFunc) *Builder {
	b.mw = append(b.mw, mw...)
	return b
}

func (b *Builder) On(e any) *Builder {
	b.endpoint = e
	return b
}

func (b *Builder) Filter(filters ...tf.Filter) *Builder {
	b.h.Filters = append(b.h.Filters, filters...)
	return b
}

func (b *Builder) Do(h tb.HandlerFunc) *Builder {
	b.h.Callback = h
	return b
}

func (b *Builder) Build() tf.Route {
	return tf.NewRoute(b.endpoint, b.h, b.mw...)
}
