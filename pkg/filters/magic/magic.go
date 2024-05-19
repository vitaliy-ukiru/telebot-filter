package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

func Text() *StringPipeline {
	return &StringPipeline{start: func(c tele.Context) (string, bool) {
		text := c.Text()
		return text, text != ""
	}}
}

func Data() *StringPipeline {
	return &StringPipeline{start: func(c tele.Context) (string, bool) {
		data := c.Data()
		return data, data != ""
	}}
}

func Sender() UserMagicFilter {
	return UserMagicFilter{
		getter: func(ctx tele.Context) (*tele.User, bool) {
			u := ctx.Sender()
			return u, u != nil
		},
	}
}

func Chat() ChatMagicFilter {
	return ChatMagicFilter{
		getter: func(ctx tele.Context) (*tele.Chat, bool) {
			chat := ctx.Chat()
			return chat, chat != nil
		},
	}
}

func Update(fn ItemFilter[tele.Update]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (tele.Update, bool) {
			x := ctx.Update()
			return x, true
		},
		fn,
	)
}

func Message() MessageFilter {
	return newMessageFilter(
		func(ctx tele.Context) (*tele.Message, bool) {
			x := ctx.Message()
			return x, x != nil
		},
	)
}

func Callback(fn ItemFilter[*tele.Callback]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.Callback, bool) {
			x := ctx.Callback()
			return x, x != nil
		},
		fn,
	)
}

func Query(fn ItemFilter[*tele.Query]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.Query, bool) {
			x := ctx.Query()
			return x, x != nil
		},
		fn,
	)
}

func InlineResult(fn ItemFilter[*tele.InlineResult]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.InlineResult, bool) {
			x := ctx.InlineResult()
			return x, x != nil
		},
		fn,
	)
}

func ShippingQuery(fn ItemFilter[*tele.ShippingQuery]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.ShippingQuery, bool) {
			x := ctx.ShippingQuery()
			return x, x != nil
		},
		fn,
	)
}

func PreCheckoutQuery(fn ItemFilter[*tele.PreCheckoutQuery]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.PreCheckoutQuery, bool) {
			x := ctx.PreCheckoutQuery()
			return x, x != nil
		},
		fn,
	)
}

func Poll(fn ItemFilter[*tele.Poll]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.Poll, bool) {
			x := ctx.Poll()
			return x, x != nil
		},
		fn,
	)
}

func PollAnswer(fn ItemFilter[*tele.PollAnswer]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.PollAnswer, bool) {
			x := ctx.PollAnswer()
			return x, x != nil
		},
		fn,
	)
}

func ChatMember(fn ItemFilter[*tele.ChatMemberUpdate]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.ChatMemberUpdate, bool) {
			x := ctx.ChatMember()
			return x, x != nil
		},
		fn,
	)
}

func ChatJoinRequest(fn ItemFilter[*tele.ChatJoinRequest]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.ChatJoinRequest, bool) {
			x := ctx.ChatJoinRequest()
			return x, x != nil
		},
		fn,
	)
}

func Topic(fn ItemFilter[*tele.Topic]) tf.Filter {
	return newPredicate(
		func(ctx tele.Context) (*tele.Topic, bool) {
			x := ctx.Topic()
			return x, x != nil
		},
		fn,
	)
}

func Entities() EntitiesMagicFilter {
	return newEntitiesFilter(func(ctx tele.Context) (tele.Entities, bool) {
		e := ctx.Entities()
		return e, len(e) > 0
	})
}
