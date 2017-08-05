package main

import (
	"math/rand"
	"time"

	"github.com/lwhile/gocache"

	"strconv"
)

const (
	max = 99999
	cap = 10000
)

var rd *rand.Rand

func init() {
	rd = rand.New(rand.NewSource(time.Now().Unix()))
}

func getData(i int) string {
	time.Sleep(time.Microsecond)
	return strconv.Itoa(i)
}

// func main() {
// 	cache := gocache.NewCache(max * 2)

// 	// No use cache
// 	start := time.Now()
// 	for i := 0; i < max; i++ {
// 		key := rd.Intn(max)
// 		getData(key)
// 	}
// 	fmt.Println("No cache use time:", time.Since(start))

// 	// Use cache
// 	start = time.Now()
// 	for i := 0; i < max; i++ {
// 		key := rd.Intn(max)
// 		var v interface{}
// 		var err error
// 		if v, err = cache.Get(key); err != nil {
// 			v = getData(key)
// 			vv, _ := strconv.Atoi(v.(string))
// 			cache.Set(v, vv)
// 		}
// 	}
// 	fmt.Println("Use cache use time:", time.Since(start))
// }

func main() {
	max := 999999
	cache := gocache.NewCache(max)
	for i := 0; i < max; i++ {
		cache.SetWithTTL(strconv.Itoa(i), i, 1)
	}

	time.Sleep(time.Second * 10)

}
