package main

import (
	"fmt"
)

func main() {
	test1 := []int{8, 6, 1, 5, 3}
	fmt.Printf("%v\n", minimumSum(test1)) // 9 (1,5,3)
	test2 := []int{5, 4, 8, 7, 10, 2}
	fmt.Printf("%v\n", minimumSum(test2)) // 13 (4,7,2)
	test3 := []int{6, 5, 4, 3, 4, 5}
	fmt.Printf("%v\n", minimumSum(test3)) // -1

}

// Triplet index: increase order
// nums[left] < nums[center] and nums[right] < nums[center]
func minimumSum(nums []int) int {
	answer := -1
	// 3 pointers: left, center, right
	for left, leftVal := range nums {
		for center := left + 1; center < len(nums); center++ {
			centerVal := nums[center]
			if leftVal >= centerVal {
				continue
			}
			for right := center + 1; right < len(nums); right++ {
				rightVal := nums[right]
				if centerVal <= rightVal {
					continue
				}
				sum := leftVal + centerVal + rightVal
				if answer == -1 || answer > sum {
					answer = sum
				}
			}
		}
	}
	return answer
}
