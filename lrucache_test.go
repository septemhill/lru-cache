package lrucache

import (
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	cache := NewLRUCache(5)
	if cache == nil {
		t.Fatalf("cache should not be nil")
	}
}

func TestSet(t *testing.T) {
	cache := NewLRUCache(5)
	data := []int{123, 434, 534}

	for i := 0; i < len(data); i++ {
		cache.Set(i, data[i])
		elem := cache.Get(i)
		if elem == nil || elem != data[i] {
			t.Fatalf("got value: %v, expected value: %v", elem, data[i])
		}
	}
}

func TestGet(t *testing.T) {
}

func TestDelete(t *testing.T) {

}
