package gocache

import "time"

// 删除策略可以有三种:定时删除, 被动删除, 定期删除
// 定时删除对内存友好而对 CPU 不友好
// 被动删除与定时删除相反
// 定期删除介于两者之间
// 这里和redis一样, 采用被动删除和定期删除同时作为删除策略

const (
	// 10 mins * 60 seconds
	clearInterval = 10 * 60
)

// ttlHandler deal with ttl login
type ttlHandler struct {
	interval int
	cache    *Cache
}

func newTTLHandler(c *Cache) *ttlHandler {
	t := &ttlHandler{
		interval: clearInterval,
		cache:    c,
	}

	go t.clean()
	return t
}

func (t *ttlHandler) clean() {
	for {
		for k, v := range t.cache.memory {
			if v.isExpire() {
				t.cache.Del(k)
			}
		}
		time.Sleep(time.Second * time.Duration(t.interval))
	}
}
