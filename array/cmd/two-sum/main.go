package main

import "fmt"

func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v", twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	// complement
	compMaps := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := compMaps[nums[i]]; ok {
			return []int{compMaps[nums[i]], nums[i]}
		}
		compMaps[complement] = nums[i]
	}
	return []int{}
}
