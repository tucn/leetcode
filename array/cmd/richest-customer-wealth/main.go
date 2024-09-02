package main

import (
	"fmt"
)

func main() {
	test1 := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Printf("%d", maximumWealth(test1))
	test2 := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Printf("%d", maximumWealth(test2))
}

// @WIP
func maximumWealth(accounts [][]int) int {
	return 0
}
