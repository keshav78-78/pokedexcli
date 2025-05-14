package pokecache

import (
	"testing"
	"time"
)

func TestNewCacheNotNil(t *testing.T) {
	cache := NewCache(time.Second) // Pass TTL duration
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(time.Second) // Long enough TTL

	cases := []struct {
		inputkey string
		inputval []byte
	}{
		{
			inputkey: "key1",
			inputval: []byte("val1"),
		},
		{
			inputkey: "key2",
			inputval: []byte("val2"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputkey, cas.inputval)
		actual, ok := cache.Get(cas.inputkey)
		if !ok {
			t.Errorf("Key '%s' not found", cas.inputkey)
			continue
		}
		if string(actual) != string(cas.inputval) {
			t.Errorf("For key '%s': expected '%s', got '%s'", cas.inputkey, string(cas.inputval), string(actual))
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond*5) // wait for reaping

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("Expected key '%s' to be reaped, but it was found", keyOne)
	}
}
