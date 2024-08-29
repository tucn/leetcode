package main

import (
	"fmt"
	"math"

	"container/list"
)

func main() {
	test1 := []int{10, 4, 8, 3}
	fmt.Printf("%v", justWork_leftRightDifference(test1))
}

// How to get left sums?
//
//	index 0: 0
//	index 1: 0 + 10
//	index 2: 0 + 10 + 4
//	index 3: 0 + 10 + 4 + 8
//
// How to get right sums?
//
//	index 3: 0
//	index 2: 0 + 3
//	index 1: 0 + 3 + 8
//	index 0: 0 + 3 + 8 + 4
//
// left sums =  [0,10,14,22]
// right sums = [15,11,3,0]
// output =     [15,1,11,22]
// How to get output? (|left sums[i] - right sums[i]|)
//
//	index 0: 15 - 0 = 15
//	index 1: 11 - 10 = 1
//	index 2: 22 - 14 = 8
//	index 3: 0 - 22 = -22
//
// First solution, we have to 2*O(n) loops, which is not efficient.
// However, just try the simple way first
func justWork_leftRightDifference(nums []int) []int {
	output := make([]int, len(nums))
	// declare leftSums and rightSums as FIFO
	leftSums := list.New()
	// declare output as First In Last Out (FILO)
	rightSums := list.New()
	//
	// -> back -> []
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			leftSums.PushBack(0)
			fmt.Printf("Left: %d\n", leftSums.Back().Value)
			continue
		}
		nextVal := leftSums.Back().Value.(int) + nums[i-1]
		leftSums.PushBack(nextVal)
		fmt.Printf("Left: %d\n", leftSums.Back().Value)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightSums.PushBack(0)
			fmt.Printf("Right: %d\n", rightSums.Back().Value)
			continue
		}
		nextVal := rightSums.Back().Value.(int) + nums[i+1]
		rightSums.PushBack(nextVal)
		fmt.Printf("Right: %d\n", rightSums.Back().Value)
	}
	for i := 0; i < len(nums); i++ {
		output[i] = int(math.Abs(float64(leftSums.Front().Value.(int) - rightSums.Back().Value.(int))))
		leftSums.Remove(leftSums.Front())
		rightSums.Remove(rightSums.Back())
	}
	return output
}
