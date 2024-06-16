package magic

import (
	"slices"

	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

type ItemFilter[T any] func(T) bool

type ItemGetter[T any] func(ctx tele.Context) (T, bool)

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

func (c CompareFilter[T]) All(factories ...FilterFactory[CompareFilter[T]]) tf.Filter {
	return logicBranch(c, And, factories)
}

func (c CompareFilter[T]) Any(factories ...FilterFactory[CompareFilter[T]]) tf.Filter {
	return logicBranch(c, Or, factories)
}

type FilterFactory[T any] func(T) tf.Filter

func logicBranch[T any](
	base T,
	operator func(...tf.Filter) tf.Filter,
	factories []FilterFactory[T],
) tf.Filter {
	filters := make([]tf.Filter, 0, len(factories))
	for _, factory := range factories {
		filters = append(filters, factory(base))
	}
	return operator(filters...)
}

func newBoolFilter(getter ItemGetter[bool], want bool) tf.Filter {
	return newPredicate(getter, boolFilter(want))
}

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

func joinNotNil[T any, U *K, K any](a ItemGetter[T], b func(T) U) ItemGetter[U] {
	return func(ctx tele.Context) (U, bool) {
		t, ok := a(ctx)
		if !ok {
			var zero U
			return zero, false
		}
		v := b(t)
		return v, v != nil
	}
}

func boolFilter(want bool) ItemFilter[bool] {
	return func(got bool) bool {
		return got == want
	}
}
