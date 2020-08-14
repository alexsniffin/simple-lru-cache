package simple_lru_cache

import (
	"testing"
)

func Test_Integration(t *testing.T) {
	lruCache := NewLRUCache(&LRUConfig{
		Limit: 3,
	})

	lruCache.Set("1", 1)
	lruCache.Set("2", 1)
	lruCache.Set("3", 1)

	// 1,2,3
	_, ok := lruCache.Get("1")
	if !ok {
		t.Error("item not found")
		t.FailNow()
	}
	// 2,3,1
	lruCache.Set("4", 1)

	// 3,1,4
	if _, ok := lruCache.items["1"]; !ok {
		t.Error("missing expected key: 1")
		t.FailNow()
	}
	if _, ok := lruCache.items["3"]; !ok {
		t.Error("missing expected key: 3")
		t.FailNow()
	}
	if _, ok := lruCache.items["4"]; !ok {
		t.Error("missing expected key: 4")
		t.FailNow()
	}
	if lruCache.first.key != "3" {
		t.Error("unexpected key for first")
		t.FailNow()
	}
	if lruCache.first.next.key != "1" {
		t.Error("unexpected key for first.next")
		t.FailNow()
	}
	if lruCache.first.next.next.key != "4" {
		t.Error("unexpected key for first.next.next")
		t.FailNow()
	}
	if lruCache.end.key != "4" {
		t.Error("unexpected key for end")
		t.FailNow()
	}
	if lruCache.end.prev.key != "1" {
		t.Error("unexpected key for end.prev")
		t.FailNow()
	}
	if lruCache.end.prev.prev.key != "3" {
		t.Error("unexpected key for end.prev.prev")
		t.FailNow()
	}
}
