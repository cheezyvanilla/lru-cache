package main

import "fmt"

type LruCache struct {
	Head     *Node
	Tail     *Node
	Capacity int
	Items    map[int]*Node
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

func NewLruCache(capacity int) LruCache {
	return LruCache{
		Capacity: capacity,
		Items:    make(map[int]*Node),
	}
}

func (lru *LruCache) Get(key int) int {
	node, ok := lru.Items[key]
	if ok {
		lru.removeNode(node)
		lru.addNode(node)
		return node.Value
	}
	return -1
}

func (lru *LruCache) Put(key int, value int) {
	// exist check
	node, ok := lru.Items[key]
	if ok {
		node.Value = value
		lru.removeNode(node)
		lru.addNode(node)
		return
	} else {
		newNode := Node{Key: key, Value: value}
		lru.Items[key] = &newNode
		lru.addNode(&newNode)
		fmt.Printf("tail while put %d: %v\n", key, lru.Tail.Key)

	}

	// the bug is here, this new list, the list is [3,1] but 1 is deleted because outcapacity
	if len(lru.Items) > lru.Capacity {
		delete(lru.Items, lru.Tail.Key)
		lru.removeNode(lru.Tail)
	}
}

func (lru *LruCache) addNode(node *Node) {
	node.Prev = nil
	node.Next = lru.Head
	if lru.Head != nil { // handle nil error
		lru.Head.Prev = node
	}
	lru.Head = node

	if lru.Tail == nil {
		lru.Tail = node
	}
}

func (lru *LruCache) removeNode(node *Node) {
	fmt.Println("removing node: ", node.Key)
	if node == lru.Head {
		fmt.Println("removing head")
		lru.Head = lru.Head.Next
	} else if node == lru.Tail {
		fmt.Println("removing tail")
		lru.Tail = lru.Tail.Prev
	} else {
		fmt.Println("removing middle")
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
	}
}
