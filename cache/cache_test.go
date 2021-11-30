package cache

import "testing"

type MyString string
func (d MyString) Len() int {
	return len(d)
}
func TestCache_Get(t *testing.T) {
	lru := New(0, nil)

	lru.Add("key1", MyString("1234"))

	if v, ok := lru.Get("key1"); !ok || string(v.(MyString)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}
