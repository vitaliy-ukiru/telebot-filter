package dispatcher_test

import (
	"io"

	"github.com/vitaliy-ukiru/telebot-filter/dispatcher"
	tb "gopkg.in/telebot.v3"
)

func ExampleNewBuilder() {
	h := func(c tb.Context) error {
		doc := c.Message().Document.File
		file, err := c.Bot().File(&doc)
		if err != nil {
			return err
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			return c.Send("fail read file: " + err.Error())
		}
		return c.Send(string(data))
	}

	filterDoc := func(c tb.Context) bool {
		doc := c.Message().Document
		return doc.MIME == "text/plain"
	}
	_ = dispatcher.NewBuilder(tb.OnDocument).
		Filter(filterDoc).
		Do(h)
}
