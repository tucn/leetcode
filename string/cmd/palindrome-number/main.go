package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v", isPalindrome(121))
	fmt.Printf("%v", isPalindrome(-121))
	fmt.Printf("%v", isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// Revert the number and then compare, if same then is palindrome
	// If not same then not palindrome
	stack := list.New()
	for x != 0 {
		stack.PushBack(x % 10)
		x = x / 10
	}
	// print stack
	for e := stack.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v", e.Value)
	}
	return false
}
