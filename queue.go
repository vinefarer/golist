package golist

import (
	"fmt"
)

// Element is an element of a queue
type Element struct {
	prev, next *Element
	queue      *Queue
	value      interface{}
}

// Prev returns the previous element
func (e *Element) Prev() *Element {
	if p := e.prev; e != nil && p.queue.root != p {
		return p
	}
	return nil
}

// Next returns the next element
func (e *Element) Next() *Element {
	if n := e.next; e != nil && n.queue.root != n {
		return n
	}
	return nil
}

// InitElement returns an element
func (e *Element) InitElement(v interface{}) *Element {
	e.prev = e
	e.next = e
	e.queue = nil
	e.value = v

	return e
}

// NewElement creates a new element of a queue
func NewElement(v interface{}) *Element {
	return new(Element).InitElement(v)
}

// Queue is a queue
type Queue struct {
	root *Element
	len  int
}

// InitQueueNode inits a queue head
func (q *Queue) InitQueueNode() *Queue {
	q.root = NewElement(nil)
	q.len = 0

	return q
}

// NewQueue creates a new queue
func NewQueue() *Queue {
	return new(Queue).InitQueueNode()
}

// Len returns the length of the queue
func (q *Queue) Len() int {
	return q.len
}

// InsertElement inserts a new element at the tail of the queue
func (q *Queue) InsertElement(v interface{}) *Queue {
	e := NewElement(v)
	e.queue = q
	e.next = q.root
	e.prev = q.root.prev
	q.root.prev.next = e
	q.root.prev = e
	q.len++

	return q
}

// RemoveElement removes an element at the tail of the queue
func (q *Queue) RemoveElement() interface{} {
	e := q.root.next
	e.next.prev = q.root
	q.root.next = e.next
	e.queue = nil
	q.len--

	return e.value
}

// SearchElement returns the pointer of the element what we need
func (q *Queue) SearchElement(v interface{}) *Element {
	e := q.root
	for e.next != q.root {
		if e.next.value == v {
			return e.next
		}
		e = e.next
	}

	return nil
}

// ModifyElement modify value of the element
func (q *Queue) ModifyElement(e *Element, v interface{}) {
	e.value = v
}

// PrintQueue prints the value of the queue
func (q *Queue) PrintQueue() {
	e := q.root
	for e.next != q.root {
		fmt.Printf("%c\t", e.next.value)
		e = e.next
	}
	fmt.Printf("\n")
}
