package golist

import (
	"fmt"
	"sync"
)

type LinkedList struct {
	sync.RWMutex

	next *LinkedList
	// Point to next node
	prev *LinkedList
	// Point to preview node

	data int32
	// To store data
}

// Create a new node
func CreateNewNode() *LinkedList {
	// return new(node LinkedList)
	return &LinkedList{
		next: nil,
		prev: nil,
		data: 0,
	}
}

// Insert new data
func (node *LinkedList) InsertNode(data int32) *LinkedList {
	if node == nil {
		return nil
	}
	node.Lock()
	defer node.Unlock()

	newnode := CreateNewNode()
	newnode.data = data
	newnode.next = node.next
	node.next = newnode
	return newnode
}

// Delete old data
func (node *LinkedList) DeleteNode() *LinkedList {
	if node.next == nil || node == nil {
		return nil
	}
	node.Lock()
	defer node.Unlock()

	node.next = node.next.next
	return node
}

// Find node which you need
func (node *LinkedList) FindNode(data int32) *LinkedList {
	if node == nil || node.next == nil {
		return nil
	}
	node.Lock()
	defer node.Unlock()

	point := node
	for point.next != nil {
		if point.next.data == data {
			return point.next
		}
		point = point.next
	}
	return nil
}

// Change node
func (node *LinkedList) ChangeNode(data int32) {
	if node == nil || node.next == nil {
		return
	}
	node.Lock()
	defer node.Unlock()

	node.data = data
	return
}

// Print all nodes of list
func (node *LinkedList) PrintList() {
	if node == nil || node.next == nil {
		return
	}
	node.Lock()
	defer node.Unlock()

	point := node
	for point.next != nil {
		fmt.Printf("%d\t", point.next.data)
		point = point.next
	}
	fmt.Printf("\n")
	return
}
