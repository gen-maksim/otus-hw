package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
	index int
}

type list struct {
	count int
	first *ListItem
	last  *ListItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	nl := ListItem{Value: v, index: l.count}
	l.count++

	if l.count != 1 {
		prevL := l.first
		prevL.Prev = &nl
		nl.Next = prevL
		nl.Prev = nil
	} else {
		l.last = &nl
	}
	l.first = &nl

	return &nl
}

func (l *list) PushBack(v interface{}) *ListItem {
	nl := ListItem{Value: v, index: l.count}
	l.count++

	if l.count != 1 {
		nextL := l.last
		nextL.Next = &nl
		nl.Prev = nextL
		nl.Next = nil
	} else {
		l.first = &nl
	}
	l.last = &nl

	return &nl
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil: // first
		i.Next.Prev = nil
		l.first = i.Next
	case i.Next == nil: // last
		i.Prev.Next = nil
		l.last = i.Prev
	default: // in the middle
		i.Next.Prev, i.Prev.Next = i.Prev, i.Next
	}

	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.count == 1 {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
	i = nil
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) Len() int {
	return l.count
}

func NewList() List {
	return new(list)
}
