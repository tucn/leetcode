package main

import "fmt"

func main() {
	// fmt.Printf("%v\n", removeDuplicates([]int{1, 1, 2}))
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Printf("k: %v\n", k)
}

// noted: remove in-place
func removeDuplicates(nums []int) int {
	chk := make(map[int]bool, len(nums))
	removed := 0
	for k := range len(nums) {
		v := nums[k-removed]
		if ok := chk[v]; ok {
			// duplicate found
			nums = removeIndex(nums, k-removed)
			removed++
			continue
		}
		chk[v] = true
	}
	return len(nums)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// enhance version
func enhanceRemoveDuplicates(nums []int) int {
	// Using two pointers
	left := 0
	// using a right pointer that increase index
	for right := 0; right < len(nums); right++ {
		// if the value in the left pointer is not same with right pointer
		// we increase the left pointer
		if nums[left] != nums[right] {
			// assign the left value with the right value
			nums[left+1] = nums[right]
			// increase
			left++
		}
	}
	return len(nums[:left+1])
}
