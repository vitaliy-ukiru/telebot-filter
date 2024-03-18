package telefilter

import "gopkg.in/telebot.v3"

// Filter is functions for filters updates.
type Filter func(c telebot.Context) bool
