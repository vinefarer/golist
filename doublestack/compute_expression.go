package main

import (
	"fmt"
	"os"
	"bufio"
	//"github.com/vinefarer/golist"
)

var (
	inputReader *bufio.Reader
	inputValue string
	err error
	res interface{}
)

func evaluate(str string) interface{} {
	//ops := golist.InitStack()
	//val := golist.InitStack()

	return str
}

func main() {
	fmt.Println("Please input your expression : ")

	inputReader = bufio.NewReader(os.Stdin)
	inputValue, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	res = evaluate(inputValue)
	fmt.Println("The result is : ", res)
}
