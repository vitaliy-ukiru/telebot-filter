// Package message provides filters for checks message.
package message

import (
	"regexp"
	"strings"

	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tb "gopkg.in/telebot.v3"
)

func Contains(part string) tf.Filter {
	return func(c tb.Context) bool {
		return strings.Contains(c.Text(), part)
	}
}

func EqualFold(s string) tf.Filter {
	return func(c tb.Context) bool {
		return strings.EqualFold(c.Text(), s)
	}
}

func Regexp(pattern regexp.Regexp) tf.Filter {
	return func(c tb.Context) bool {
		return pattern.MatchString(c.Text())
	}
}

// EntityFunc checks message's entities at returns true if
// at least one matches predicate.
func EntityFunc(predicate func(entity tb.MessageEntity) bool) tf.Filter {
	return func(c tb.Context) bool {
		for _, entity := range c.Entities() {
			if predicate(entity) {
				return true
			}
		}
		return false
	}
}

func HaveEntities(c tb.Context) bool {
	return len(c.Entities()) > 0
}

func HaveEntity(entityType tb.EntityType) tf.Filter {
	return EntityFunc(func(entity tb.MessageEntity) bool {
		return entity.Type == entityType
	})
}
