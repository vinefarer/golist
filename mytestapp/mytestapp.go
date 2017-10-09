package main

import (
	"fmt"

	"github.com/golang/golist"
)

func main() {
	list := golist.CreateNewNode()
	for i := 0; i < 10; i++ {
		list.InsertNode(int32(i))
		list.PrintList()
	}

	list.FindNode(5).ChangeNode(13)
	list.PrintList()
	fmt.Println("-----------------")
	for i := 0; i < 10; i++ {
		list.DeleteNode()
		list.PrintList()
	}

	list = nil
	fmt.Println("=========LinkedList End=========")

	queue := golist.NewQueue()
	for i := 'a'; i < 'k'; i++ {
		queue.InsertElement(i)
		queue.PrintQueue()
	}
	queue.ModifyElement(queue.SearchElement('d'), 'z')
	queue.PrintQueue()
	fmt.Println("-----------------")

	for i := 'a'; i < 'k'; i++ {
		queue.RemoveElement()
		queue.PrintQueue()
	}
	queue = nil
	fmt.Println("=========Queue End=========")

	stack := golist.InitStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
		stack.Print()
	}
	stack.Modify(stack.Search(5), 8)
	stack.Print()
	fmt.Println("-----------------")

	for i := 0; i < 10; i++ {
		stack.Pop()
		stack.Print()
	}
	stack = nil
	fmt.Println("=========Stack End=========")
}
