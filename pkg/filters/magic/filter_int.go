package magic

import (
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type NumberFilter[T Number] struct {
	getter ItemGetter[T]
}

func newNumberFilter[T Number](getter ItemGetter[T]) NumberFilter[T] {
	return NumberFilter[T]{getter: getter}
}

func (i NumberFilter[T]) predicate(filter ItemFilter[T]) tf.Filter {
	return newPredicate(i.getter, filter)
}

func (i NumberFilter[T]) Equal(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t == n
	})
}

func (i NumberFilter[T]) GreatThan(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t > n
	})
}

func (i NumberFilter[T]) LessThan(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t > n
	})
}

func (i NumberFilter[T]) GreatOrEqual(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t >= n
	})
}

func (i NumberFilter[T]) LessOrEqual(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t <= n
	})
}

func (i NumberFilter[T]) NotEqual(n T) tf.Filter {
	return i.predicate(func(t T) bool {
		return t != n
	})
}

func (i NumberFilter[T]) On(filter ItemFilter[T]) tf.Filter {
	return i.predicate(filter)
}

func (i NumberFilter[T]) All(factories ...FilterFactory[NumberFilter[T]]) tf.Filter {
	return logicBranch(i, And, factories)
}

func (i NumberFilter[T]) Any(factories ...FilterFactory[NumberFilter[T]]) tf.Filter {
	return logicBranch(i, Or, factories)
}
