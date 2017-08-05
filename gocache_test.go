package gocache

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/lwhile/gocache"
)

func genRandomSeq(size int) []int {
	seq := make([]int, size)
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < size; i++ {
		seq[i] = rd.Intn(9999999)
	}
	return seq
}

func genNewTestCache(ttl int) LRU {
	max := 999999
	cache := gocache.NewCache(max)
	seq := genRandomSeq(max)
	for i := 0; i < max; i++ {
		cache.Set(strconv.Itoa(i), seq[i])
	}
	return cache
}

func genNewTestCacheWithTTL(ttl int) LRU {
	max := 999999
	cache := gocache.NewCache(max)
	seq := genRandomSeq(max)
	for i := 0; i < max; i++ {
		cache.SetWithTTL(strconv.Itoa(i), seq[i], ttl)
	}
	return cache
}

// func TestNewCache(t *testing.T) {
// 	max := 999999
// 	cache := gocache.NewCache(max)
// 	seq := genRandomSeq(max)
// 	for i := 0; i < max; i++ {
// 		cache.Set(strconv.Itoa(i), seq[i])
// 	}

// 	if cache.Len() != max {
// 		t.Fatalf("cache size %d != %d\n", cache.Len(), max)
// 	}

// 	for i := 0; i < max; i++ {
// 		v, err := cache.Get(strconv.Itoa(i))
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		if v.(int) != seq[i] {
// 			es := fmt.Sprintf("%v != %d\n", v.(int), seq[i])
// 			t.Fatal(es)
// 		}
// 	}
// }

// func TestCache_Del(t *testing.T) {
// 	// insert something
// 	max := 10
// 	cache := gocache.NewCache(max)

// 	for i := 0; i < max; i++ {
// 		cache.Set(strconv.Itoa(i), i)
// 	}

// 	for i := 0; i < max; i++ {
// 		cache.Del(strconv.Itoa(i))
// 		v, err := cache.Get(strconv.Itoa(i))
// 		if err == nil {
// 			t.Fatalf("delete key %v fail", v)
// 		}

// 		if cache.Len() != max-i-1 {
// 			t.Fatalf("%d != %d\n", cache.Len(), max)
// 		}
// 	}

// 	if cache.Len() != 0 {
// 		t.Fatalf("%d !=0\n", cache.Len())
// 	}

// }

func TestCache_clean(t *testing.T) {
	cache := genNewTestCacheWithTTL(2)

	time.Sleep(time.Second * 10)

	if cache.Len() != 0 {
		t.Fatalf("cache clean data fail: %d != 0\n", cache.Len())
	}
}
