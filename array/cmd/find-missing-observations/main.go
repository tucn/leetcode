package main

import "fmt"

func main() {
	fmt.Printf("%v\n", missingRolls([]int{3, 2, 4, 3}, 4, 2))
	fmt.Printf("%v\n", missingRolls([]int{1, 5, 6}, 3, 4))
	fmt.Printf("%v\n", missingRolls([]int{1, 2, 3, 4}, 6, 4))
}

// @WIP
// Find missing rolls so that the mean of the total rolls is equal to the given mean.
// The time complexity and space complexity will depending on the input n
func missingRolls(rolls []int, mean int, n int) []int {
	out := make([]int, n)
	// firstly, we check the mod of current out with means
	// if it is 0 then we just add means to the out or add some number that mod means to 0
	// For example if the [4,4,4,4] means 4 and n = 1 -> [4,4,4,4,4] -> [4,4,4,4,4] / 5 = 4
	// Another example: [3,2,4,3] means 4 and n = 2 -> 3+2+4+3 = 12 % 4 = 0 -> add double 4 or 6

	// if it is not 0 then we need to add the next n so that the mod is 0
	// Another example: [3,2,4,1] means 4 and n = 2 -> 3+2+4+1 = 10 % 4
	// = 2 -> next should be means - 2 + means = 6 -> [3,2,4,1,6] and then 4 -> [3,2,4,1,6,4] = 16 % 4 = 0
	//
	// Now, if the means - mod + means > 6 ? then we need to depends on the next n
	// to maximize the change, we add 6
	//
	// and then for the other we add means to the out
	total := 0
	for _, val := range rolls {
		total += val
	}

	// mod not equal 0
	div := mean
	for i := 0; i < n; i++ {
		mod := total % div
		if mod == 0 {
			for j := 0; j < n-i; j++ {
				out[j] = mean
			}
			return out
		}
		need := (mean - mod) + mean
		if need > 6 {
			total += 6
		} else {
			total += need
		}
		div += 1
	}

	return []int{}
}
