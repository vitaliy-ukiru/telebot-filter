package magic

import (
	"slices"

	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

type ItemFilter[T any] func(T) bool

type ItemGetter[T any] func(ctx tele.Context) (T, bool)

type ItemPredicate[T any] struct {
	getter    ItemGetter[T]
	predicate ItemFilter[T]
}

func newPredicate[T any](getter ItemGetter[T], predicate ItemFilter[T]) tf.Filter {
	return func(c tele.Context) bool {
		val, ok := getter(c)
		if !ok {
			return false
		}
		return predicate(val)
	}
}

type CompareFilter[T comparable] struct {
	getter ItemGetter[T]
}

func newCompareFilter[T comparable](getter ItemGetter[T]) CompareFilter[T] {
	return CompareFilter[T]{getter: getter}
}

func (c CompareFilter[T]) Equal(value T) tf.Filter {
	return newPredicate(c.getter, func(got T) bool {
		return value == got
	})
}

func (c CompareFilter[T]) NotEqual(value T) tf.Filter {
	return newPredicate(c.getter, func(got T) bool {
		return value != got
	})
}

func (c CompareFilter[T]) OneOf(values ...T) tf.Filter {
	return newPredicate(c.getter, func(got T) bool {
		return slices.Contains(values, got)
	})
}

func (c CompareFilter[T]) On(filter ItemFilter[T]) tf.Filter {
	return newPredicate(c.getter, filter)
}
