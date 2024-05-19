package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
)

type SliceMagicFilter[S ~[]E, E any] struct {
	getter ItemGetter[S]
}

func newArrayFilter[S ~[]E, E any](getter ItemGetter[S]) SliceMagicFilter[S, E] {
	return SliceMagicFilter[S, E]{getter: getter}
}

func (s SliceMagicFilter[S, E]) Len() NumberFilter[int] {
	return newNumberFilter(joinGetter(s.getter, func(s S) int {
		return len(s)
	}))
}

func (s SliceMagicFilter[S, E]) AtLeastOne(predicate ItemFilter[E]) tf.Filter {
	return newPredicate(s.getter, func(s S) bool {
		for _, e := range s {
			if predicate(e) {
				return true
			}
		}
		return false
	})
}
