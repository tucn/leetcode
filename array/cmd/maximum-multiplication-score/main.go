package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxScore(
		[]int{3, 2, 5, 6},
		[]int{2, -6, 4, -5, -3, 2, -7},
	))
}

func maxScore(a []int, b []int) int64 {
	// sliding windows
	max := 0
	return int64(max)
}
