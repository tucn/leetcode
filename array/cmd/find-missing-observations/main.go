package main

import "fmt"

func main() {
	fmt.Printf("%v\n", missingRolls([]int{3, 2, 4, 3}, 4, 2))
	fmt.Printf("%v\n", missingRolls([]int{1, 5, 6}, 3, 4))
	fmt.Printf("%v\n", missingRolls([]int{1, 2, 3, 4}, 6, 4))
	fmt.Printf("%v\n", missingRolls([]int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}, 4, 40))
}

// Find missing rolls so that the mean of the total rolls is equal to the given mean.
// The time complexity and space complexity will depending on the input n
func missingRolls(rolls []int, mean int, n int) []int {
	out := make([]int, 0)
	totalSum := (len(rolls) + n) * mean
	// Collect total
	totalRolls := 0
	for _, v := range rolls {
		if mean == 6 && v < 6 {
			// There is no way you can make the means of 6
			// if there is one of the current element less than
			// 6
			return out
		}
		totalRolls += v
	}
	shouldFill := totalSum - totalRolls
	// shouldFill should not less then n
	if shouldFill < n {
		return out
	}
	// shouldFill should not greater than 6 * n
	if shouldFill > 6*n {
		return out
	}
	// Generate the Missing Rolls:
	// Go backward to make sure the last element is the largest
	for i := n - 1; i >= 0; i-- {
		// For example:
		// - We are assigning the value to result[0].
		// - There are i = 2 rolls left after this one.
		// - The formula becomes: min(6, 15 - 2) = min(6, 13) = 6.
		// - We assign 6 to result[0] and reduce missingSum to 15 - 6 = 9.
		fillin := min(6, shouldFill-i)

		out = append(out, fillin)
		shouldFill -= fillin
	}
	return out
}
