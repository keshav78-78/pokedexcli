package pokecache

import "testing"

func TestNewCacheNotNil(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache()

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
