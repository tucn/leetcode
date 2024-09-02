package main

import (
	"fmt"
)

// @WIP
func main() {
	//nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]
	nums1 := [][]int{
		{1, 2},
		{2, 3},
		{4, 5},
	}
	nums2 := [][]int{
		{1, 4},
		{3, 2},
		{4, 1},
	}
	fmt.Printf("%v\n", mergeArrays(nums1, nums2))
}

// Each array contains unique ids and is sorted in ascending order by id.
// => no duplicated ids and can be iterated in order
//
// output[dup_id] = sum(nums1[dup_id], nums2[dup_id])
// Return the resulting array. The returned array must be sorted in ascending order by id.
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	// don't need to check for empty arrays based on the constraints
	out := make([][]int, 0)

	// should iterate based on the longest array
	// swap the arrays if nums1 is longer
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	mapOut := make(map[int]int, len(nums1))
	i := 0
	for i < len(nums1) {
		id1 := nums1[i][0]
		val1 := nums1[i][1]
		id2 := nums2[i][0]
		val2 := nums2[i][1]
		if id1 > len(mapOut) {
			mapOut[id1] += val1
		}
		if id2 > len(mapOut) {
			mapOut[id2] += val2
		}
		if len(nums1) == len(mapOut) {
			break
		}
		i++
	}
	// convert map to array
	for k, v := range mapOut {
		out = append(out, []int{k, v})
	}
	return out
}
