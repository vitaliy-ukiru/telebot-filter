package dispatcher

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

// Builder helps for connect handler function, filters,
// endpoint and middlewares.
//
// For create Builder use [Dispatcher.B] or [NewBuilder].
type Builder struct {
	endpoint any
	filters  []tf.Filter
	callback tb.HandlerFunc
	mw       []tb.MiddlewareFunc
}

// NewBuilder creates new builder with endpoint.
// You can override endpoint with [Builder.On] method.
func NewBuilder(endpoint any) *Builder {
	return &Builder{endpoint: endpoint}
}

// Use adds middlewares for handler.
func (b *Builder) Use(mw ...tb.MiddlewareFunc) *Builder {
	b.mw = append(b.mw, mw...)
	return b
}

// On sets endpoint.
func (b *Builder) On(e any) *Builder {
	b.endpoint = e
	return b
}

// Filter appends filters to builder.
func (b *Builder) Filter(filters ...tf.Filter) *Builder {
	b.filters = append(b.filters, filters...)
	return b
}

// Do sets callback.
func (b *Builder) Do(h tb.HandlerFunc) *Builder {
	b.callback = h
	return b
}

func (b *Builder) Build() tf.Route {
	if b.callback == nil {
		panic("telebot-filter: builder: callback must be not nil")
	}
	return tf.NewRoute(
		b.endpoint,
		tf.NewRawHandler(b.callback, b.filters...),
		b.mw...,
	)
}
