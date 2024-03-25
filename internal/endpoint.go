package internal

import tb "gopkg.in/telebot.v3"

func ExtractRawEndpoint(e any) string {
	if e == nil {
		panic("telebot-filter: endpoint must be not nil")
	}

	switch end := e.(type) {
	case string:
		return end
	case tb.CallbackEndpoint:
		return end.CallbackUnique()
	default:
		panic("telebot-filter: unsupported endpoint")
	}
}
