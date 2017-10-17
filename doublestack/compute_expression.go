package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vinefarer/golist"
)

var (
	inputReader *bufio.Reader
	inputValue  []byte
	err         error
	ops         rune
	val         byte
)

func main() {
	fmt.Println("Please input your compute expression :")
	inputReader = bufio.NewReader(os.Stdin)
	inputValue, err = inputReader.ReadBytes('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	opslist := golist.InitStack()
	vallist := golist.InitStack()

	for _, e := range inputValue {
		if e == '\n' || e == '\r' {
			break
		}
		fmt.Printf("e = %s\n", string(e))
		switch e {
		case ' ':
		case '(':
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			opslist.Push(rune(e))
		case ')':
			ops = opslist.Pop().(rune)
			val = vallist.Pop().(byte)
			switch ops {
			case '+':
				val = vallist.Pop().(byte) + val
			case '-':
				val = vallist.Pop().(byte) - val
			case '*':
				val = vallist.Pop().(byte) * val
			case '/':
				val = vallist.Pop().(byte) / val
			default:
				fmt.Println("Wrong operation!")
			}
			vallist.Push(val)
		default:
			vallist.Push(e)
		}
		fmt.Println("---------------")
		opslist.Print()
		vallist.Print()
		fmt.Println("---------------")
	}

	fmt.Printf("The result is : %s\n", string(val))
}
