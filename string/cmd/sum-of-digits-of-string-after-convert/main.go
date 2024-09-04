package main

import (
	"fmt"
	"strconv"
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
		out = transform(convertedStr)
		convertedStr = strconv.Itoa(out)
	}
	return out
}

func convert(s string) string {
	out := ""
	for _, c := range s {
		out += strconv.Itoa(int(c - 'a' + 1))
	}
	return out
}

func transform(value string) int {
	out := 0
	for _, c := range value {
		num, _ := strconv.Atoi(string(c))
		out += num
	}
	return out
}
