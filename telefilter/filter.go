package telefilter

import tb "gopkg.in/telebot.v3"

// Filter is functions for filters updates.
type Filter func(c tb.Context) bool
