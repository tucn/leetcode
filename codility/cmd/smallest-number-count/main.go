package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", smallest(16))
	fmt.Printf("%v\n", smallest(50))
	fmt.Printf("%v\n", smallest2(16))
	fmt.Printf("%v\n", smallest2(50))
}

// Using slices instead of stack
func smallest2(N int) int {
	out := make([]int, 0)
	for N > 9 {
		out = append(out, 9)
		N -= 9
	}
	if N > 0 {
		out = append(out, N)
	}
	// Reverse the out
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	// Change to string
	outStr := ""
	for i := 0; i < len(out); i++ {
		outStr += fmt.Sprintf("%d", out[i])
	}
	outInt, err := strconv.Atoi(outStr)
	if err != nil {
		return 0
	}
	return outInt
}

// Using stack
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
