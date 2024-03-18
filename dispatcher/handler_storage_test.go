package dispatcher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitaliy-ukiru/telebot-filter/internal"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	"github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

func TestA(t *testing.T) {
	assert.True(t, true)
}

func newMiddleware(buff *strings.Builder, id string) tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			buff.WriteString(id)
			return next(c)
		}
	}
}

func newList(m ...tb.MiddlewareFunc) *internal.MiddlewareList {
	return container.NewListFromSlice(m)
}

func newHandler(buff *strings.Builder, id string) telefilter.Handler {
	return telefilter.RawHandler{
		Callback: func(_ tb.Context) error {
			buff.WriteString(id)
			return nil
		},
	}
}

func Test_handlerRoute_run(t *testing.T) {
	type fields struct {
		router  *Router
		Handler telefilter.Handler
		mw      []tb.MiddlewareFunc
	}
	var c tb.Context = nil
	buff := new(strings.Builder)

	tests := []struct {
		name      string
		fields    fields
		wantStack string
	}{
		{
			name: "top level",
			fields: fields{
				router: &Router{
					mw: newList(
						newMiddleware(buff, "a"),
						newMiddleware(buff, "b"),
						newMiddleware(buff, "c"),
					),
				},
				Handler: newHandler(buff, "_callback"),
				mw: []tb.MiddlewareFunc{
					newMiddleware(buff, "1"),
					newMiddleware(buff, "2"),
				},
			},
			wantStack: "abc12_callback",
		},
		{
			name: "multi level (3)",
			fields: fields{
				router: &Router{
					parent: &Router{
						parent: &Router{
							mw: newList(
								newMiddleware(buff, "x"),
								newMiddleware(buff, "y"),
								newMiddleware(buff, "z"),
							),
						},
						mw: newList(newMiddleware(buff, "_")),
					},
					mw: newList(
						newMiddleware(buff, "a"),
						newMiddleware(buff, "b"),
						newMiddleware(buff, "c"),
					),
				},
				Handler: newHandler(buff, "_callback"),
				mw: []tb.MiddlewareFunc{
					newMiddleware(buff, "1"),
					newMiddleware(buff, "2"),
				},
			},
			wantStack: "xyz_abc12_callback",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buff.Reset()
			hr := handlerRoute{
				router: tt.fields.router,
				Route: telefilter.Route{
					Handler:     tt.fields.Handler,
					Middlewares: tt.fields.mw,
				},
			}
			_ = hr.run(c)
			stack := buff.String()
			assert.Equal(t, tt.wantStack, stack, "run()")
		})
	}
}
