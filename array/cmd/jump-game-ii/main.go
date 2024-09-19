package main

import "fmt"

func main() {
	fmt.Printf("%v\n", jump([]int{2, 3, 1, 1, 4, 1, 2, 1, 1, 2, 3}))
}

// - Greedy approach
// + The time complexity of this solution is O(n),
// where n is the length of the input array nums.
// This is because we iterate through the array once,
// making constant-time operations at each step.
//
// + The space complexity of this solution is O(1).
// We only use a fixed amount of extra space for variables like jumps,
// currentEnd, and farthest, regardless of the input size.
func jump(nums []int) int {
	// Quick check
	if len(nums) <= 1 {
		return 0
	}
	// output jumps
	jumps := 0
	// We need to use 1 pointer for
	// determine which is the target jumps
	currentEnd := 0
	// farthest destination to reach
	farthest := 0
	for index, val := range nums {
		// Choose the farthest point from the index point to index+val
		farthest = max(farthest, index+val)
		// We shouldn't increase jump
		// if index == last index
		if index == currentEnd && index < len(nums)-1 {
			jumps++
			currentEnd = farthest
		}
	}
	return jumps
}
