package main

import (
	"fmt"
)

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
// Return the outulting array. The returned array must be sorted in ascending order by id.
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var out [][]int
	// two pointers to iterate over nums1 and nums2
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		// if one of the arrays is exhausted, append the out of the other array
		if i == len(nums1) && j < len(nums2) {
			out = append(out, nums2[j:]...)
			break
		}
		if j == len(nums2) && i < len(nums1) {
			out = append(out, nums1[i:]...)
			break
		}
		// if the ids are the same, add the sum of the values
		if nums1[i][0] == nums2[j][0] {
			out = append(out, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
			continue
		}
		// if nums1[i][0] < nums2[j][0], append nums1[i]
		if nums1[i][0] < nums2[j][0] {
			out = append(out, nums1[i])
			i++
			continue
		}
		// if nums1[i][0] > nums2[j][0], append nums2[j]
		out = append(out, nums2[j])
		j++
	}
	return out
}
