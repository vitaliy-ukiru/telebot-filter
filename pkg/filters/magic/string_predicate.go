package magic

import (
	"slices"
	"strings"

	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
)

func (s *StringPipeline) Contains(substr string) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.Contains(s, substr)
	})
}
func (s *StringPipeline) ContainsRune(r rune) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.ContainsRune(s, r)
	})
}
func (s *StringPipeline) ContainsAny(chars string) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.ContainsAny(s, chars)
	})
}
func (s *StringPipeline) ContainsFunc(f func(rune) bool) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.ContainsFunc(s, f)
	})
}
func (s *StringPipeline) EqualFold(t string) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.EqualFold(s, t)
	})
}

func (s *StringPipeline) NotEqual(t string) tf.Filter {
	return s.predicate(func(s string) bool {
		return s != t
	})
}
func (s *StringPipeline) HasPrefix(prefix string) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.HasPrefix(s, prefix)
	})
}
func (s *StringPipeline) HasSuffix(suffix string) tf.Filter {
	return s.predicate(func(s string) bool {
		return strings.HasSuffix(s, suffix)
	})
}

func (s *StringPipeline) Equal(t string) tf.Filter {
	return s.predicate(func(s string) bool {
		return s == t
	})
}

func (s *StringPipeline) In(values ...string) tf.Filter {
	return s.predicate(func(s string) bool {
		return slices.Contains(values, s)
	})
}
