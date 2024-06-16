package magic

import (
	"regexp"
	"slices"
	"strings"

	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
)

func (s *StringPipeline) Contains(substr string) tf.Filter {
	return s.On(func(s string) bool {
		return strings.Contains(s, substr)
	})
}
func (s *StringPipeline) ContainsRune(r rune) tf.Filter {
	return s.On(func(s string) bool {
		return strings.ContainsRune(s, r)
	})
}
func (s *StringPipeline) ContainsAny(chars string) tf.Filter {
	return s.On(func(s string) bool {
		return strings.ContainsAny(s, chars)
	})
}
func (s *StringPipeline) ContainsFunc(f func(rune) bool) tf.Filter {
	return s.On(func(s string) bool {
		return strings.ContainsFunc(s, f)
	})
}
func (s *StringPipeline) EqualFold(t string) tf.Filter {
	return s.On(func(s string) bool {
		return strings.EqualFold(s, t)
	})
}

func (s *StringPipeline) HasPrefix(prefix string) tf.Filter {
	return s.On(func(s string) bool {
		return strings.HasPrefix(s, prefix)
	})
}
func (s *StringPipeline) HasSuffix(suffix string) tf.Filter {
	return s.On(func(s string) bool {
		return strings.HasSuffix(s, suffix)
	})
}

func (s *StringPipeline) Equal(t string) tf.Filter {
	return s.On(func(s string) bool {
		return s == t
	})
}

func (s *StringPipeline) NotEqual(t string) tf.Filter {
	return s.On(func(s string) bool {
		return s != t
	})
}

func (s *StringPipeline) In(values ...string) tf.Filter {
	return s.On(func(s string) bool {
		return slices.Contains(values, s)
	})
}

func (s *StringPipeline) Regexp(pattern regexp.Regexp) tf.Filter {
	return s.On(pattern.MatchString)
}

func (s *StringPipeline) Len() NumberFilter[int] {
	return newNumberFilter(joinGetter(s.execute, func(s string) int {
		return len(s)
	}))
}

func (s *StringPipeline) All(filtersFactories ...func(s *StringPipeline) tf.Filter) tf.Filter {
	filters := make([]tf.Filter, 0, len(filtersFactories))
	for _, factory := range filtersFactories {
		filters = append(filters, factory(s.Copy()))
	}
	return And(filters...)
}

func (s *StringPipeline) Any(filtersFactories ...func(s *StringPipeline) tf.Filter) tf.Filter {
	filters := make([]tf.Filter, 0, len(filtersFactories))
	for _, factory := range filtersFactories {
		filters = append(filters, factory(s.Copy()))
	}
	return Or(filters...)
}
