package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", isPalindrome(121))
	fmt.Printf("%v\n", isPalindrome(-121))
	fmt.Printf("%v\n", isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for x != 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return reversed == original
}
