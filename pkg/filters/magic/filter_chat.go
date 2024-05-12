package magic

import tele "gopkg.in/telebot.v3"

type ChatMagicFilter struct {
	getter ItemGetter[*tele.Chat]
}

func (c ChatMagicFilter) ID() NumberFilter[int64] {
	return NumberFilter[int64]{
		getter: joinGetter(c.getter, func(chat *tele.Chat) int64 {
			return chat.ID
		}),
	}
}
func (c ChatMagicFilter) FirstName() *StringPipeline {
	return newStringPipeline(joinGetter(c.getter, func(chat *tele.Chat) string {
		return chat.FirstName
	}))
}
func (c ChatMagicFilter) LastName() *StringPipeline {
	return newStringPipeline(joinGetter(c.getter, func(chat *tele.Chat) string {
		return chat.LastName
	}))
}
func (c ChatMagicFilter) Username() *StringPipeline {
	return newStringPipeline(joinGetter(c.getter, func(chat *tele.Chat) string {
		return chat.Username
	}))
}

func (c ChatMagicFilter) Title() *StringPipeline {
	return newStringPipeline(joinGetter(c.getter, func(chat *tele.Chat) string {
		return chat.Title
	}))
}

func (c ChatMagicFilter) Type() CompareFilter[tele.ChatType] {
	return newCompareFilter(joinGetter(c.getter, func(chat *tele.Chat) tele.ChatType {
		return chat.Type
	}))
}
