package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

type UserMagicFilter struct {
	getter ItemGetter[*tele.User]
}

func (u UserMagicFilter) ID() NumberFilter[int64] {
	return NumberFilter[int64]{
		getter: joinGetter(u.getter, func(user *tele.User) int64 {
			return user.ID
		}),
	}
}
func (u UserMagicFilter) FirstName() *StringPipeline {
	return newStringPipeline(joinGetter(u.getter, func(user *tele.User) string {
		return user.FirstName
	}))
}
func (u UserMagicFilter) LastName() *StringPipeline {
	return newStringPipeline(joinGetter(u.getter, func(user *tele.User) string {
		return user.LastName
	}))
}
func (u UserMagicFilter) Username() *StringPipeline {
	return newStringPipeline(joinGetter(u.getter, func(user *tele.User) string {
		return user.Username
	}))
}
func (u UserMagicFilter) LanguageCode() *StringPipeline {
	return newStringPipeline(joinGetter(u.getter, func(user *tele.User) string {
		return user.LanguageCode
	}))
}
func (u UserMagicFilter) CustomEmojiStatus() *StringPipeline {
	return newStringPipeline(joinGetter(u.getter, func(user *tele.User) string {
		return user.CustomEmojiStatus
	}))
}

func (u UserMagicFilter) IsForum(status bool) tf.Filter {
	return newPredicate(
		joinGetter(u.getter, func(user *tele.User) bool {
			return user.IsForum
		}),
		boolFilter(status),
	)
}
func (u UserMagicFilter) IsBot(status bool) tf.Filter {
	return newPredicate(
		joinGetter(u.getter, func(user *tele.User) bool {
			return user.IsBot
		}),
		boolFilter(status),
	)
}
func (u UserMagicFilter) IsPremium(status bool) tf.Filter {
	return newPredicate(
		joinGetter(u.getter, func(user *tele.User) bool {
			return user.IsPremium
		}),
		boolFilter(status),
	)
}
func (u UserMagicFilter) AddedToMenu(status bool) tf.Filter {
	return newPredicate(
		joinGetter(u.getter, func(user *tele.User) bool {
			return user.AddedToMenu
		}),
		boolFilter(status),
	)
}

func boolFilter(want bool) ItemFilter[bool] {
	return func(got bool) bool {
		return got == want
	}
}

//func (s UserMagicFilter) Usernames()         []string

func joinGetter[T, U any](a ItemGetter[T], b func(T) U) ItemGetter[U] {
	return func(ctx tele.Context) (U, bool) {
		t, ok := a(ctx)
		if !ok {
			var zero U
			return zero, false
		}
		return b(t), true
	}
}
