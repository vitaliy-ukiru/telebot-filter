package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

const nameKey = "name"

func handlerFactory(name string) tb.HandlerFunc {
	return func(context tb.Context) error {
		context.Set(nameKey, name)
		return nil
	}
}

func trueFilter(tb.Context) bool  { return true }
func falseFilter(tb.Context) bool { return false }

type testCase struct {
	name     string
	args     []tf.Handler
	want     string
	notMatch bool
}

var testingRoutes = []testCase{
	{
		name: "first handler",
		args: []tf.Handler{
			tf.NewRawHandler(handlerFactory("first"), trueFilter),
			tf.NewRawHandler(handlerFactory("2"), trueFilter),
		},
		want: "first",
	},
	{
		name: "middle",
		args: []tf.Handler{
			tf.NewRawHandler(handlerFactory("top"), falseFilter),
			tf.NewRawHandler(handlerFactory("middle")),
			tf.NewRawHandler(handlerFactory("bottom")),
		},
		want: "middle",
	},
	{
		name: "not match",
		args: []tf.Handler{
			tf.NewRawHandler(handlerFactory("not1"), falseFilter),
			tf.NewRawHandler(handlerFactory("not2"), falseFilter),
		},
		notMatch: true,
	},
	{
		name:     "empty args",
		notMatch: true,
	},
}

func testExecute(
	t *testing.T,
	bot *tb.Bot,
	h tb.HandlerFunc,
	want string,
	notMatch bool,
) {
	ctx := bot.NewContext(tb.Update{})
	err := h(ctx)
	assert.NoError(t, err, "execute handler")

	got, match := ctx.Get(nameKey).(string)

	assert.NotEqual(t, notMatch, match, "match status")

	assert.EqualValues(
		t,
		want,
		got,
		"executed handler name",
	)
}

func TestNew(t *testing.T) {

	bot, err := tb.NewBot(tb.Settings{
		Offline:     true,
		Synchronous: true,
	})

	assert.NoError(t, err, "tb.NewBot")

	for _, tt := range testingRoutes {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.args...)
			testExecute(t, bot, h, tt.want, tt.notMatch)
		})
	}
}

func TestRoute_Add(t *testing.T) {
	type fields struct {
		handlers *container.List[tf.Handler]
	}
	tests := []struct {
		name   string
		fields fields
		args   tf.Handler
		panics bool
	}{
		{
			name: "self initialization",
			args: new(tf.RawHandler),
		},
		{
			name: "with pre init list",
			fields: fields{
				handlers: container.NewListFromSlice([]tf.Handler{
					tf.Route{},
					tf.RawHandler{},
				}),
			},
			args: &tf.RawHandler{Callback: func(_ tb.Context) error {
				return nil
			}},
		},
		{
			name:   "panics on nil handler",
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Route{
				handlers: tt.fields.handlers,
			}
			fn := assert.NotPanicsf
			if tt.panics {
				fn = assert.Panicsf
			}

			fn(t, func() {
				r.Add(tt.args)
			}, "r.Add(%v)", tt.args)

			assert.NotNil(t, r.handlers, "list of handlers")

			if !tt.panics {
				tail := r.handlers.Back().Value
				assert.EqualValues(t, tt.args, tail, "tail of list")
			}
		})
	}
}

func TestRoute_Handler(t *testing.T) {

	bot, err := tb.NewBot(tb.Settings{
		Offline:     true,
		Synchronous: true,
	})

	assert.NoError(t, err, "tb.NewBot")

	for _, tt := range testingRoutes {
		t.Run(tt.name, func(t *testing.T) {
			//r := &Route{
			//	handlers: tt.fields.handlers,
			//}
			//tt.wantErr(t, r.Handler(tt.args.c), fmt.Sprintf("Handler(%v)", tt.args.c))
			//
			r := &Route{}
			if len(tt.args) > 0 {
				r.handlers = container.NewListFromSlice(tt.args)
			}
			testExecute(t, bot, r.Handler, tt.want, tt.notMatch)
		})
	}
}
