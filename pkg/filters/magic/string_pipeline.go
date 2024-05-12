package magic

import "strings"

func (s *StringPipeline) Map(mapping func(rune) rune) *StringPipeline {
	s.add(func(s string) string {
		return strings.Map(mapping, s)
	})
	return s
}

func (s *StringPipeline) Repeat(count int) *StringPipeline {
	s.add(func(s string) string {
		return strings.Repeat(s, count)
	})
	return s
}

func (s *StringPipeline) Replace(old, new string, n int) *StringPipeline {
	s.add(func(s string) string {
		return strings.Replace(s, old, new, n)
	})
	return s
}

func (s *StringPipeline) ReplaceAll(old, new string) *StringPipeline {
	s.add(func(s string) string {
		return strings.ReplaceAll(s, old, new)
	})
	return s
}

func (s *StringPipeline) Lower() *StringPipeline {
	s.add(strings.ToLower)
	return s
}

func (s *StringPipeline) Title() *StringPipeline {
	s.add(strings.ToTitle)
	return s
}

func (s *StringPipeline) Upper() *StringPipeline {
	s.add(strings.ToUpper)
	return s
}

func (s *StringPipeline) TrimSuffix(suffix string) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimSuffix(s, suffix)
	})
	return s
}

func (s *StringPipeline) TrimSpace() *StringPipeline {
	s.add(strings.TrimSpace)
	return s
}

func (s *StringPipeline) TrimRight(cutSet string) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimRight(s, cutSet)
	})
	return s
}

func (s *StringPipeline) TrimRightFunc(f func(rune) bool) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimRightFunc(s, f)
	})
	return s
}

func (s *StringPipeline) TrimPrefix(prefix string) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimPrefix(s, prefix)
	})
	return s
}

func (s *StringPipeline) TrimLeft(cutset string) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimLeft(s, cutset)
	})
	return s
}

func (s *StringPipeline) TrimLeftFunc(f func(rune) bool) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimLeftFunc(s, f)
	})
	return s
}

func (s *StringPipeline) Trim(cutset string) *StringPipeline {
	s.add(func(s string) string {
		return strings.Trim(s, cutset)
	})
	return s
}

func (s *StringPipeline) TrimFunc(f func(rune) bool) *StringPipeline {
	s.add(func(s string) string {
		return strings.TrimFunc(s, f)
	})
	return s
}
