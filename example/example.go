package main

import (
	"github.com/lwhile/lru"
)

func main() {
	cache := lru.NewCache(10)

	for i := 0; i < 100; i++ {
		cache.Set(i, i)
	}

	cache.Show()
}
