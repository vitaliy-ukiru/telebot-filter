package magic

import (
	"github.com/vitaliy-ukiru/telebot-filter/pkg/filters"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
)

func And(predicates ...tf.Filter) tf.Filter {
	return filters.All(predicates...)
}

func Or(predicates ...tf.Filter) tf.Filter {
	return filters.Any(predicates...)
}

func Not(filter tf.Filter) tf.Filter {
	return filters.Not(filter)
}
