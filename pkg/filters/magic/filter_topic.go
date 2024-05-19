package magic

import tele "gopkg.in/telebot.v3"

type TopicFilter struct {
	getter ItemGetter[*tele.Topic]
}

func newTopicFilter(getter ItemGetter[*tele.Topic]) TopicFilter {
	return TopicFilter{getter: getter}
}

func (t TopicFilter) Name() *StringPipeline {
	return newStringPipeline(joinGetter(t.getter, func(t *tele.Topic) string {
		return t.Name
	}))
}

func (t TopicFilter) IconCustomEmojiID() *StringPipeline {
	return newStringPipeline(joinGetter(t.getter, func(t *tele.Topic) string {
		return t.IconCustomEmojiID
	}))
}

func (t TopicFilter) IconColor() NumberFilter[int] {
	return newNumberFilter(joinGetter(t.getter, func(t *tele.Topic) int {
		return t.IconColor
	}))
}
func (t TopicFilter) ThreadID() NumberFilter[int] {
	return newNumberFilter(joinGetter(t.getter, func(t *tele.Topic) int {
		return t.ThreadID
	}))
}
