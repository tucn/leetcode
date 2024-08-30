package main

import (
	"fmt"
)

// @WIP
func main() {
	test1 := []int{1, 1, 2, 3, 4, 4, 3}
	fmt.Printf("%v", singleNumber(test1)) // expect 2
}

// Since this one is a `non-empty` array,
// we assume that there is at least one `single number` in the array.
// and the check `if len(nums) == 0` is not necessary.
func singleNumber(nums []int) int {
	return 0
}
