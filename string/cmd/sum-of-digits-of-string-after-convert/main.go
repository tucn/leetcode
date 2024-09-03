package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

// @WIP
func main() {
	fmt.Printf("%d\n", getLucky("iiii", 1))
	fmt.Printf("%d\n", getLucky("leetcode", 2))
}
func getLucky(s string, k int) int {
	numStack := list.New()
	for index, _ := range s {
		num, err := strconv.Atoi(string(s[index]))
		if err != nil {
			return 0
		}
		if num < 10 {
			numStack.PushBack(num)
			continue
		}
		if num >= 10 {
			numStack.PushBack(num % 10)
			numStack.PushBack(math.Round(float64(num) / 10))
			continue
		}
	}
	count := 0
	out := 0
	for {
		if numStack.Len() == 0 {
			break
		}
		if count == k {
			break
		}
		count += 1
		out += transform(numStack)
	}
	return transform(numStack)
}

func transform(stack *list.List) int {
	out := 0
	for {
		num := stack.Back().Value.(int)
		out += num
		stack.Remove(stack.Back())
		if stack.Len() == 0 {
			break
		}
	}
	return out
}
