package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	"golang.org/x/exp/constraints"
)

type IntFilter[T constraints.Integer] struct {
	getter ItemGetter[T]
}

func (i IntFilter[T]) predicate(filter ItemFilter[T]) tf.Filter {
	return newPredicate(i.getter, filter)
}

func (i IntFilter[T]) Equal(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t == n
	})
}

func (i IntFilter[T]) GreatThan(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t > n
	})
}

func (i IntFilter[T]) LessThan(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t > n
	})
}

func (i IntFilter[T]) GreatOrEqual(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t >= n
	})
}

func (i IntFilter[T]) LessOrEqual(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t <= n
	})
}

func (i IntFilter[T]) NotEqual(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t != n
	})
}

func (i IntFilter[T]) On(filter ItemFilter[T]) tf.Filter {
	return i.predicate(filter)
}
