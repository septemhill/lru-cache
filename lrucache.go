package lrucache

import (
	"container/list"
	"fmt"
)

type data struct {
	key   interface{}
	value interface{}
}

//LRUCache
type LRUCache struct {
	l        *list.List
	m        map[interface{}]*list.Element
	capacity int
}

//NewLRUCache creates a lru cache with specified size
func NewLRUCache(cap int) *LRUCache {
	leastCap := 30
	if cap > 0 {
		leastCap = cap
	}

	return &LRUCache{
		l:        list.New(),
		m:        make(map[interface{}]*list.Element),
		capacity: leastCap,
	}
}

func (l *LRUCache) removeLast() {
	elem := l.l.Back()
	l.l.Remove(elem)
	key := elem.Value.(*data).key
	delete(l.m, key)
}

//Set set data to lru cache
func (l *LRUCache) Set(key, value interface{}) {
	if elem, ok := l.m[key]; ok {
		l.l.MoveToFront(elem)
		elem.Value.(*data).value = value
		//l.Trace()
		return
	}

	elem := l.l.PushFront(&data{key: key, value: value})
	l.m[key] = elem

	if len(l.m) > l.capacity {
		l.removeLast()
	}

	//l.Trace()
}

//Get gets data from lru cache by key
func (l *LRUCache) Get(key interface{}) interface{} {
	if len(l.m) > 0 {
		if elem, ok := l.m[key]; ok {
			l.l.MoveToFront(elem)
			return elem.Value.(*data).value
		}
		return nil
	}
	return nil
}

//Delete removes data from lru cache by key
func (l *LRUCache) Delete(key interface{}) {
	if len(l.m) > 0 {
		if elem, ok := l.m[key]; ok {
			l.l.Remove(elem)
			delete(l.m, key)
		}
	}
}

func (l *LRUCache) Trace() {
	fmt.Println("list:")
	for iter := l.l.Front(); iter != nil; iter = iter.Next() {
		fmt.Println(iter.Value)
	}
	fmt.Println("map:")
	for key, val := range l.m {
		fmt.Println(key, val)
	}
}

