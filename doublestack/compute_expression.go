package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/vinefarer/golist"
)

var (
	inputReader *bufio.Reader
	inputValue byte
	err error
	res interface{}
)

//func evaluate(str string) interface{} {
//	ops := golist.InitStack()
//	val := golist.InitStack()
//
//	return str
//}

func main() {
	fmt.Println("Please input your expression : ")
	inputReader = bufio.NewReader(os.Stdin)
	ops := golist.InitStack()
	val := golist.InitStack()
	for {
		inputValue, err = inputReader.ReadByte()
		if err != nil {
			fmt.Println("There were errors reading, exiting program.")
			break
		}
		switch inputValue {
		case '(':
		case '+': fallthrough
		case '-': fallthrough
		case '*': fallthrough
		case '/': ops.Push(inputValue)
		case ')': {
			op := ops.Pop()
			res := val.Pop()
			switch op {
			case '+': res = val.Pop() + res
			case '-': res = val.Pop() - res
			case '*': res = val.Pop() * res
			case '/': res = val.Pop() / res
			}
		}
		default: val.Push(inputValue)
		}
		if inputValue == '\n' {
			break
		}
		fmt.Printf("The result is : %s\n", string(inputValue))
	}
}
