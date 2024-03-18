package container

type Element[T any] struct {
	Value T
	prev  *Element[T]
	next  *Element[T]
}

type List[T any] struct {
	head *Element[T]
	tail *Element[T]
	len  int
}

func NewListFromSlice[T any](s []T) *List[T] {
	l := new(List[T])
	l.ExtendSlice(s)
	return l
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) Insert(d T) *Element[T] {
	e := &Element[T]{Value: d}
	if l.head == nil {
		l.head = e
	}
	if l.tail != nil {
		l.tail.next = e
		e.prev = l.tail
	}
	l.tail = e
	l.len++
	return e
}

func (l *List[T]) Extend(other *List[T]) {
	if other == nil || other.len == 0 {
		return
	}
	for e := other.head; e != nil; e = e.next {
		l.Insert(e.Value)
	}
}

func (l *List[T]) ExtendSlice(values []T) {
	if len(values) == 0 {
		return
	}
	for _, v := range values {
		l.Insert(v)
	}
}

func (l *List[T]) Front() *Element[T] {
	return l.head
}

func (l *List[T]) Back() *Element[T] {
	return l.tail
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}
