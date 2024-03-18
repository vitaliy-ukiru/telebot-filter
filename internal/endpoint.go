package internal

import tb "gopkg.in/telebot.v3"

func ExtractRawEndpoint(e any) string {
	switch end := e.(type) {
	case string:
		return end
	case tb.CallbackEndpoint:
		return end.CallbackUnique()
	default:
		panic("fsm: telebot: unsupported endpoint")
	}
}
