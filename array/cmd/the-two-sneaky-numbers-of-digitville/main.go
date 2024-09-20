package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", getSneakyNumbers([]int{0, 3, 2, 1, 3, 2}))
	fmt.Printf("%v\n", getSneakyNumbers([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}

func getSneakyNumbers(nums []int) []int {
	// Sort in ascending order
	sort.Sort(sort.IntSlice(nums))
	occurred := make(map[int]bool, len(nums))
	out := make([]int, 0, 2)
	for _, v := range nums {
		if occurred[v] {
			out = append(out, v)
			continue
		}
		occurred[v] = true
	}
	return out
}
