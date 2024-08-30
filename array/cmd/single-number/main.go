package main

import (
	"fmt"
)

func main() {
	test1 := []int{1, 1, 2, 3, 4, 4, 3}
	fmt.Printf("%v", xor(test1)) // expect 2
}

// Since this one is a `non-empty` array,
// we assume that there is at least one `single number` in the array.
// and the check `if len(nums) == 0` is not necessary.
func singleNumber(nums []int) int {
	// We have to go through all numbers to know if they are `single` or not. O(n)
	// We can use a hash table to store the numbers and their counts. O(n)
	// Then we can go through the hash table to find the `single` number. O(n)
	// So the total time complexity is O(n)
	//
	numMaps := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		numMaps[nums[i]]++
	}
	for k, v := range numMaps {
		if v == 1 {
			return k
		}
	}
	return 0
}

// Based on the hint, we can use XOR
// By using XOR, we can remove the duplicated element in the array.
// So we can use XOR to find the `single` number.
func xor(nums []int) int {
	mask := 0
	for i := 0; i < len(nums); i++ {
		mask ^= nums[i]
	}
	return mask
}
