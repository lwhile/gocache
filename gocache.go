package gocache

import (
	"fmt"
	"sync"
)

// LRU interface
type LRU interface {
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{})
	Show()
}

// Node type
type Node struct {
	Key   interface{}
	Value interface{}
	Next  *Node
	Last  *Node
	ID    int
}

// DoubleLinkList type
type doubleLinkList struct {
	Head *Node
	Tail *Node
}

// newDoubleLinkList return a doubleLinkList
func newDoubleLinkList(cap int) *doubleLinkList {
	head := new(Node)
	head.ID = 0

	tail := new(Node)
	tail.ID = cap - 1

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
	Container *doubleLinkList
	size      int
	cap       int
	memory    map[interface{}]*Node
	mux       sync.RWMutex
}

// Get a value from cache
func (c *Cache) Get(ifce interface{}) (interface{}, error) {
	node, ok := c.memory[ifce]

	// 存在cache中, 将结点移到链表的头部,然后返回值
	if ok {
		c.Set(node.Key, node.Value)
		return node.Value, nil
	}

	// 不在cache里
	return nil, fmt.Errorf("Cache no contain this key: %s", ifce)
}

// Set a value to cache
func (c *Cache) Set(key, value interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.size >= c.cap {
		removedNode := c.Container.Tail.Last
		removedNode.Last.Next = c.Container.Tail
		c.Container.Tail.Last = removedNode.Last
		c.size--
		delete(c.memory, key)
	} else {
		movedNode := c.Container.Head.Next

		newNode := new(Node)
		c.memory[key] = newNode
		newNode.Key = key
		newNode.Value = value
		newNode.ID = c.Size()
		newNode.Next = movedNode
		newNode.Last = c.Container.Head

		movedNode.Last = newNode

		c.Container.Head.Next = newNode

		c.size++
	}
}

// Show linklist
func (c *Cache) Show() {
	n := c.Container.Head
	var count int
	for n != nil {
		fmt.Print(n.ID, n.Key, n.Value)
		fmt.Print(" -> ")
		if count%10 == 0 {
			fmt.Println()
		}
		n = n.Next
	}
}

// Size return size of cache
func (c *Cache) Size() int {
	return c.size
}

// NewCache return a cache type
func NewCache(cap int) LRU {
	return &Cache{
		Container: newDoubleLinkList(cap),
		memory:    make(map[interface{}]*Node),
		cap:       cap,
	}
}
