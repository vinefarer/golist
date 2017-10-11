package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/vinefarer/golist"
)

var (
	inputReader *bufio.Reader
	inputValue string
	err error
	res interface{}
)

func evaluate(str string) interface{} {
	ops := golist.InitStack()
	val := golist.InitStack()

	return val.Pop()
}

func main() {
	fmt.Println("Please input your expression : ")
	inputReader = bufio.NewReader(os.Stdin)
	inputValue, err = inputReader.ReadString('\n')
	if nil == err {
		res = evaluate(inputValue)
		fmt.Println("The result is : ", res)
	}
}
