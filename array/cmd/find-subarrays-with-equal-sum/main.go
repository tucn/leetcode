package main

import "fmt"

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", findSubarrays(test1))
	test2 := []int{4, 2, 4}
	fmt.Printf("%v\n", findSubarrays(test2))
}

func findSubarrays(nums []int) bool {
	sumMaps := make(map[int]int, 0)
	for k, v := range nums {
		if k < len(nums)-1 {
			// check if the sum of the current element and the next element is in the map
			nextVal := nums[k+1]
			if _, ok := sumMaps[v+nextVal]; ok {
				return true
			}
			// Add to the map if the sum is not there
			sumMaps[v+nextVal]++
		}
	}
	return false
}
