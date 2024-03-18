package message

import (
	"strings"

	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

func Contains(part string) tf.Filter {
	return func(c tb.Context) bool {
		return strings.Contains(c.Text(), part)
	}
}

func EqualFold(s string) tf.Filter {
	return func(c tb.Context) bool {
		return strings.EqualFold(c.Text(), s)
	}
}
