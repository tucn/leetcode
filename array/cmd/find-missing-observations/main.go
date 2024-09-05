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
	divider := len(rolls) + n
	// Collect total
	currentTotal := 0
	for _, v := range rolls {
		if mean == 6 && v < 6 {
			// There is no way you can make the means of 6
			// if there is one of the current element less than
			// 6
			return out
		}
		currentTotal += v
	}
	nRollsSum := n * 6
	found := 0
	// Should be start from n since there is no way we can have a 0 roll
	for i := n; i <= nRollsSum; i++ {
		total := currentTotal + i
		if total/divider == mean {
			fmt.Printf("Found: %d\n", i)
			found = i
			break
		}
	}
	// Return blank if not found
	if found == 0 {
		return []int{}
	}
	// found, just split all values with n
	remainer := found % n
	// if remainer > 6, we have to split the remainder through 6 or more random of n
	// make sure that after we split, each of the div is still smaller than 7
	div := found / n
	extra := 0
	count := 0
	if remainer > 6 {
		for {
			extra = remainer / 6
			fmt.Printf("Remainer: %v, Extra: %v\n", remainer, extra)
			count += 6
			if extra < 6 {
				count += extra
				break
			}
		}
	}
	maximumPlus := count / n
	if div+maximumPlus > 6 {
		return []int{}
	}
	for range n {
		if count == 0 {
			out = append(out, div)
			continue
		}
		if maximumPlus == 0 {
			out = append(out, div+1)
			count -= 1
			continue
		}
		out = append(out, div+maximumPlus)
		count -= maximumPlus
	}
	return out
}
