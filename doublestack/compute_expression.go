package main

import (
	"github.com/vinefarer/golist"
	"os"

)

func evaluate() interface{} {
	ops := golist.InitStack()
	val := golist.InitStack()
	for i := 0; i < len(os.Args); i++ {
		op := os.Args[i]

	}

	return val.Pop()
}

func main() {

}
