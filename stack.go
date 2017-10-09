package golist

import (
	"fmt"
)

type element struct {
	prev, next *element
	stack      *Stack
	value      interface{}
}

func (e *element) Init(v interface{}) *element {
	e.prev = e
	e.next = e
	e.stack = nil
	e.value = v
	return e
}

func InitElement(v interface{}) *element {
	return new(element).Init(v)
}

func (e *element) Prev() *element {
	if p := e.prev; e != nil && e.stack.root != e {
		return p
	}
	return nil
}

func (e *element) Next() *element {
	if n := e.next; e != nil && e.stack.root != e {
		return n
	}
	return nil
}

type Stack struct {
	root *element
	len  int
}

func InitStack() *Stack {
	return &Stack{
		root: InitElement(nil),
		len:  0,
	}
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) GetBottomVal() interface{} {
	return s.root.next.value
}

func (s *Stack) GetTopVal() interface{} {
	return s.root.prev.value
}

func (s *Stack) Push(v interface{}) *Stack {
	e := InitElement(v)
	e.next = s.root.next
	e.prev = s.root
	e.stack = s
	s.root.next.prev = e
	s.root.next = e
	return s
}

func (s *Stack) Pop() interface{} {
	e := s.root.next
	s.root.next = e.next
	e.next.prev = s.root
	e.stack = nil
	e.next = nil
	e.prev = nil
	return e.value
}

func (s *Stack) Search(v interface{}) *element {
	e := s.root
	for e.next != s.root {
		if e.next.value == v {
			return e.next
		}
		e = e.next
	}
	return nil
}

func (s *Stack) Modify(e *element, v interface{}) {
	e.value = v
}

func (s *Stack) Print() {
	e := s.root
	for e.next != s.root {
		fmt.Printf("%d\t", e.next.value)
		e = e.next
	}
	fmt.Printf("\n")
}
