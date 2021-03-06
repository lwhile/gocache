package gocache

import (
	"fmt"
	"time"

	"github.com/lwhile/utillib/safemap"
	"github.com/prometheus/common/log"
)

const (
	dftClearnInterval = 10 * 60 * time.Second
)

// LRU interface
type LRU interface {
	Get(string) (interface{}, error)
	Set(string, interface{})
	SetWithTTL(string, interface{}, int)
	Del(key string)
	Len() int
}

// Node type
type Node struct {
	Key   string
	Value interface{}
	Next  *Node
	Last  *Node

	// 过期时间戳
	// 0 表示无限
	TTL int64
}

func (node *Node) isExpire() bool {
	return time.Now().Unix() > node.TTL && node.TTL > 0
}

// DoubleLinkList type
type doubleLinkList struct {
	Head *Node
	Tail *Node
}

// newDoubleLinkList return a doubleLinkList
func newDoubleLinkList(cap int) *doubleLinkList {
	head := new(Node)
	tail := new(Node)

	head.Next = tail
	tail.Last = head

	dlst := doubleLinkList{
		Head: head,
		Tail: tail,
	}
	return &dlst
}

// Cache type
type Cache struct {
	Container     *doubleLinkList
	size          int
	cap           int
	memory        safemap.SafeMap
	clearInterval time.Duration
}

// Get a value from cache
func (c *Cache) Get(key string) (interface{}, error) {
	n, ok := c.memory.Get(key)
	// 存在cache中, 将结点移到链表的头部,然后返回值
	if ok {
		node := n.(*Node)
		// Check whether expire
		if node.isExpire() {
			c.Del(key)
			return nil, fmt.Errorf("key %s was expired", key)
		}
		c.Set(node.Key, node.Value)
		return node.Value, nil
	}

	// 不在cache里
	return nil, fmt.Errorf("Cache no contain this key: %s", key)
}

// Set a value to cache
func (c *Cache) Set(key string, value interface{}) {
	c.set(key, value, 0)
}

// SetWithTTL :
func (c *Cache) SetWithTTL(key string, value interface{}, ttl int) {
	c.set(key, value, time.Duration(ttl))
}

func (c *Cache) set(key string, value interface{}, ttl time.Duration) {
	if c.size >= c.cap {
		c.Del(c.Container.Tail.Last.Key)
	}
	movedNode := c.Container.Head.Next
	newNode := newNode(key, value, ttl)
	c.memory.Set(key, newNode)
	newNode.Next = movedNode
	newNode.Last = c.Container.Head
	movedNode.Last = newNode
	c.Container.Head.Next = newNode
	c.size++
}

// Del :
func (c *Cache) Del(key string) {
	dNode, ok := c.memory.Get(key)
	if !ok {
		return
	}

	delNode := dNode.(*Node)

	delNode.Last.Next = delNode.Next
	delNode.Next.Last = delNode.Last
	delNode.Next = nil
	delNode.Last = nil
	c.memory.Delete(key)
	c.size--
}

func (c *Cache) clean() {
	for {
		time.Sleep(c.clearInterval)
		log.Infoln("Begin clean cache...")
		allM := c.memory.GetAll()
		for k, v := range allM {
			//log.Infof("k:%s,v:%v\n", k, v)
			n := v.(*Node)
			if n.isExpire() {
				//log.Infof("Del %s\n", k)
				c.Del(k)
			}
		}
	}
}

func newNode(key string, value interface{}, ttl time.Duration) (node *Node) {
	node = &Node{
		Key:   key,
		Value: value,
	}
	if ttl <= 0 {
		node.TTL = 0
	} else {
		node.TTL = time.Now().Unix() + int64(ttl)
	}
	return
}

// Show linklist
func (c *Cache) Show() {
	n := c.Container.Head
	var count int
	for n != nil {
		fmt.Print(n.Key, n.Value)
		fmt.Print(" -> ")
		if count%10 == 0 {
			fmt.Println()
		}
		n = n.Next
	}
}

// Len return size of cache
func (c *Cache) Len() int {
	return c.size
}

// NewCache return a cache type
func NewCache(cap int) LRU {
	cache := Cache{
		Container:     newDoubleLinkList(cap),
		memory:        safemap.NewMap(),
		cap:           cap,
		clearInterval: dftClearnInterval,
	}
	go cache.clean()
	return &cache
}
