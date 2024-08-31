package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := "i love eating burger"
	search1 := "burg"
	test2 := "this problem is an easy problem"
	search2 := "pro"
	test3 := "i am tired"
	search3 := "you"
	fmt.Printf("%d\n", isPrefixOfWord(test1, search1))
	fmt.Printf("%d\n", isPrefixOfWord(test2, search2))
	fmt.Printf("%d\n", isPrefixOfWord(test3, search3))
}

func isPrefixOfWord(sentence string, searchWord string) int {
	// left to right
	parts := strings.Split(sentence, " ")
	for i := 0; i < len(parts); i++ {
		if strings.HasPrefix(parts[i], searchWord) {
			return i + 1
		}
	}
	return -1
}
