package lru

import (
	"fmt"
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
func newDoubleLinkList(size int) *doubleLinkList {
	head := new(Node)
	head.ID = 0

	tail := new(Node)
	tail.ID = size - 1

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
	memory    map[interface{}]*Node
}

// Get a value from cache
func (c *Cache) Get(ifce interface{}) (interface{}, error) {
	node, ok := c.memory[ifce]

	// 存在cache中, 将结点移到链表的头部,然后返回值
	if ok {
		// TODO: 移动结点
		c.Set(node.Key, node.Value)
		return node.Value, nil
	}

	// 不在cache里
	return nil, fmt.Errorf("Cache no contain this key: %s", ifce)
}

// Set a value to cache
func (c *Cache) Set(key, value interface{}) {
	movedNode := c.Container.Head.Next

	newNode := new(Node)
	c.memory[key] = newNode
	newNode.Key = key
	newNode.Value = value

	newNode.Next = movedNode
	newNode.Last = c.Container.Head

	movedNode.Last = newNode

	c.Container.Head.Next = newNode
}

func (c *Cache) Show() {
	n := c.Container.Head
	var count int
	for n.Next != nil {
		fmt.Print(n.ID, n.Key, n.Value, n)
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
func NewCache(size int) LRU {
	return &Cache{
		Container: newDoubleLinkList(size),
		size:      size,
		memory:    make(map[interface{}]*Node),
	}
}
