package magic

import (
	"github.com/vitaliy-ukiru/telebot-filter/internal/container"
	tf "github.com/vitaliy-ukiru/telebot-filter/telefilter"
	tele "gopkg.in/telebot.v3"
)

type stringJob func(s string) string

type StringPipeline struct {
	start ItemGetter[string]
	list  *container.List[stringJob]
}

func newStringPipeline(start ItemGetter[string]) *StringPipeline {
	return &StringPipeline{start: start}
}

func (s *StringPipeline) add(job stringJob) {
	if s.list == nil {
		s.list = new(container.List[stringJob])
	}

	s.list.Insert(job)
}

func (s *StringPipeline) execute(ctx tele.Context) (string, bool) {
	value, ok := s.start(ctx)
	if !ok {
		return "", false
	}
	if s.list != nil {
		for e := s.list.Front(); e != nil; e = e.Next() {
			value = e.Value(value)
		}
	}
	return value, true
}

func (s *StringPipeline) predicate(f ItemFilter[string]) tf.Filter {
	return newPredicate(s.execute, f)
}

func (s *StringPipeline) Copy() *StringPipeline {
	return &StringPipeline{
		start: s.start,
		list:  s.list.Copy(),
	}
}
