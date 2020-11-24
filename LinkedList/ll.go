package LinkedList

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Node struct {
	content Item
	next    *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (ll *LinkedList) Append(t Item) {
	ll.lock.Lock()

	node := Node{t, nil}

	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head

		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
	ll.lock.Unlock()
}

func (ll *LinkedList) Insert(i int, t Item) error {
	ll.lock.Lock()

	defer ll.lock.Unlock()

	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bounds")
	}

	addNode := Node{t, nil}

	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
		return nil
	}

	node := ll.head
	j := 0

	for j < i-2 {
		j++
		node = node.next
	}

	addNode.next = node.next
	node.next = &addNode
	ll.size++

	return nil
}
