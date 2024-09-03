package main

import (
	"fmt"
)

func main() {
	test1 := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Printf("%d\n", maximumWealth(test1))
	test2 := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Printf("%d\n", maximumWealth(test2))
	test3 := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Printf("%d\n", maximumWealth(test3))
}

// Find the maximum value of the multiple dimensional array
func maximumWealth(accounts [][]int) int {
	max := 0
	// O(m) * O(n)
	// But we know the length of the array is always the same
	// So we can do O(n)
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
