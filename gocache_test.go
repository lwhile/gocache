package gocache

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/lwhile/gocache"
)

var max = 9999

func genNewTestCache() LRU {
	cache := gocache.NewCache(max)
	for i := 0; i < max; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
	return cache
}

func genNewTestCacheWithTTL(ttl int) LRU {
	cache := gocache.NewCache(max)
	for i := 0; i < max; i++ {
		cache.SetWithTTL(strconv.Itoa(i), i, ttl)
	}
	return cache
}

func TestNewCache(t *testing.T) {
	cache := genNewTestCache()

	if cache.Len() != max {
		t.Fatalf("cache size %d != %d\n", cache.Len(), max)
	}

	for i := 0; i < max; i++ {
		v, err := cache.Get(strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		if v.(int) != i {
			es := fmt.Sprintf("%v != %d\n", v.(int), i)
			t.Fatal(es)
		}
	}
}

func TestCache_Del(t *testing.T) {
	// insert something
	cache := genNewTestCache()

	for i := 0; i < max; i++ {
		cache.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < max; i++ {
		cache.Del(strconv.Itoa(i))
		v, err := cache.Get(strconv.Itoa(i))
		if err == nil {
			t.Fatalf("delete key %v fail", v)
		}

		if cache.Len() != max-i-1 {
			t.Fatalf("%d != %d\n", cache.Len(), max)
		}
	}

	if cache.Len() != 0 {
		t.Fatalf("%d !=0\n", cache.Len())
	}

}

func TestCache_clean(t *testing.T) {
	cache := genNewTestCacheWithTTL(2)

	time.Sleep(time.Second * 10)

	if cache.Len() != 0 {
		t.Fatalf("cache clean data fail: %d != 0\n", cache.Len())
	}
}
