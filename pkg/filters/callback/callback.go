package callback

// On returns unique callback that telebot
// can resolve as endpoint for inline buttons.
//
// This feature works over telebot unique and data button parts.
// When creating a new button, you can pass a payload there.
//
//	// create button
//	menu := new(telebot.ReplyMarkup)
//	btn := menu.Data("Text", "unique", dataParts...)
//
// You can get this payload in handlers like
//
//	// Both methods return payload without unique part
//	payloadParts := c.Args()
//	rawPayload := c.Data()
func On(unique string) string {
	return "\f" + unique
}
