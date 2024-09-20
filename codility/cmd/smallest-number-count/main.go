package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", smallest(16))
	fmt.Printf("%v\n", smallest(50))
}

func smallest(N int) int {
	// Create a stack for storing number
	stack := list.New()
	for N > 9 {
		stack.PushBack(9)
		N -= 9
	}
	// N < 9
	if N > 0 {
		stack.PushBack(N)
	}
	// turn into a string
	outStr := ""
	for stack.Len() > 0 {
		ele := stack.Back()
		outStr += fmt.Sprintf("%v", ele.Value)
		stack.Remove(stack.Back())
	}
	outInt, err := strconv.Atoi(outStr)
	if err != nil {
		return 0
	}
	return outInt
}
