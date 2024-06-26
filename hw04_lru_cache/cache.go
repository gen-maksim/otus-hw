package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type CacheEl struct {
	Key Key
	Val interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	v, ok := l.GetListItem(key)

	if !ok {
		ni := l.queue.PushFront(CacheEl{Key: key, Val: value})
		l.items[key] = ni

		if l.capacity < len(l.items) {
			remove := l.queue.Back()
			s := (remove.Value).(CacheEl)
			delete(l.items, s.Key)
			l.queue.Remove(remove)
		}
	} else {
		s := (v.Value).(CacheEl)
		s.Val = value
		v.Value = s
	}

	return ok
}

func (l *lruCache) GetListItem(key Key) (*ListItem, bool) {
	if value, ok := l.items[key]; ok {
		l.queue.MoveToFront(value)

		return value, true
	}

	return nil, false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if value, ok := l.GetListItem(key); ok {
		if s, ok := (value.Value).(CacheEl); ok {
			return s.Val, true
		}

		return value.Value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
