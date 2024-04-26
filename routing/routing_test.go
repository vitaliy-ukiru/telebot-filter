package routing

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

func newMiddleware(buff *strings.Builder, id string) tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			buff.WriteString(id)
			return next(c)
		}
	}
}

func newHandler(buff *strings.Builder, id string) tb.HandlerFunc {
	return func(_ tb.Context) error {
		buff.WriteString(id)
		return nil
	}
}

func TestNew(t *testing.T) {
	buff := new(strings.Builder)

	unreachedHandler := tf.NewRawHandler(
		func(_ tb.Context) error {
			t.Errorf("unreached code")
			return nil
		},
		func(_ tb.Context) bool { return false },
	)

	tests := []struct {
		name string
		args []tf.Handler
		mw   []tb.MiddlewareFunc
		want string
	}{
		{
			name: "no middlewares",
			args: []tf.Handler{
				unreachedHandler,
				tf.NewRawHandler(newHandler(buff, "main")),
				tf.NewRoute(nil, unreachedHandler, newMiddleware(buff, "unreached")),
			},
			mw: []tb.MiddlewareFunc{
				newMiddleware(buff, "tb_"),
			},
			want: "tb_main",
		},
		{
			name: "with middlewares",
			args: []tf.Handler{
				unreachedHandler,
				tf.NewRoute(nil, unreachedHandler, newMiddleware(buff, "unreached")),
				tf.NewRoute(
					nil,
					tf.NewRawHandler(newHandler(buff, "main")),
					newMiddleware(buff, "1"),
					newMiddleware(buff, "2"),
				),
			},
			mw: []tb.MiddlewareFunc{
				newMiddleware(buff, "tb_"),
			},
			want: "tb_12main",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buff.Reset()
			h := New(tt.args...)
			executeHandler(t, h, tt.mw)
			got := buff.String()

			assert.Equal(t, tt.want, got, "New()")
		})
	}
}

func executeHandler(t *testing.T, h tb.HandlerFunc, mw []tb.MiddlewareFunc) {
	const endpoint = "test"

	bot, err := tb.NewBot(tb.Settings{
		Offline:     true,
		Synchronous: true,
	})
	assert.NoError(t, err, "create bot")

	bot.Handle(endpoint, h, mw...)
	bot.ProcessUpdate(tb.Update{
		Message: &tb.Message{
			Text: endpoint,
		},
	})
}
