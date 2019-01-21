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
	cache := NewLRUCache(5)
	data := []float64{123.22, 4323.333, 7554.2349}

	for i := 0; i < len(data); i++ {
		cache.Set(i, data[i])
		elem := cache.Get(i)
		if elem == nil || elem != data[i] {
			t.Fatalf("got value: %v, expected value: %v", elem, data[i])
		}
	}
}

func TestDelete(t *testing.T) {
	cache := NewLRUCache(5)
	data := []string{"Septem", "Nicole", "Asolia"}

	for i := 0; i < len(data); i++ {
		cache.Set(i, data[i])
	}

	for i := 0; i < len(data); i++ {
		cache.Delete(i)
		if elem := cache.Get(i); elem != nil {
			t.Fatalf("got value: %v, expected value: nil", elem)
		}
	}
}

func TestPurge(t *testing.T) {
	cache := NewLRUCache(100)

	for i := 0; i < 1000; i++ {
		cache.Set(i, i+332)
	}
	len := cache.Len()

	if len != 100 {
		t.Fatalf("got value: %v, expected value: %v", len, 100)
	}

	cache.Purge()
	len = cache.Len()

	if len != 0 {
		t.Fatalf("got value: %v, expected value: %v", len, 0)
	}
}
