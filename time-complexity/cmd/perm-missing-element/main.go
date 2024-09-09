package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", Solution([]int{2, 3, 1, 5}))
}

func Solution(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			return i + 1
		}
	}
	return len(A) + 1
}
