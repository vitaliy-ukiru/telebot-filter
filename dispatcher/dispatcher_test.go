package dispatcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tb "gopkg.in/telebot.v3"
)

type _mockBotHandlerCounter struct {
	handlers map[string]int
}

func (*_mockBotHandlerCounter) Use(...tb.MiddlewareFunc) {}

func (m *_mockBotHandlerCounter) Handle(endpoint any, _ tb.HandlerFunc, _ ...tb.MiddlewareFunc) {
	e := endpoint.(string)
	if m.handlers == nil {
		m.handlers = make(map[string]int)
	}
	m.handlers[e]++
}

func (m *_mockBotHandlerCounter) get(e string) int { return m.handlers[e] }

func TestDispatcher_wrapEndpoint(t *testing.T) {
	bot := new(_mockBotHandlerCounter)

	d := &Dispatcher{
		bot:     bot,
		wrapped: make(container.Set[string]),
	}
	cases := []string{"test", "/test", tb.OnPhoto, "test"}
	for _, e := range cases {
		d.wrapEndpoint(e)
	}

	for _, s := range cases {
		assert.Equalf(t, 1, bot.get(s), "call bot.Handle for %s", s)
	}
}
