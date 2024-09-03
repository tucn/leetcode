package main

import (
	"fmt"
)

// @WIP
func main() {
	fmt.Printf("%d\n", getLucky("iiii", 1))
	fmt.Printf("%d\n", getLucky("leetcode", 2))
}

func getLucky(s string, k int) int {
	convertedStr := convert(s)
	out := 0
	for i := 0; i < k; i++ {
		out += transform(convertedStr)
	}
	return out
}

func convert(s string) int {
	out := 0
	for _, c := range s {
		out += int(c - '0')
	}
	return out
}

func transform(value int) int {
	out := 0
	for {
		if value == 0 {
			break
		}
		out += value % 10
		value /= 10
	}
	return out
}
