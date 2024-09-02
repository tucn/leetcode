package main

import "fmt"

func main() {
	fmt.Printf("%v\n", chalkReplacer([]int{5, 1, 5}, 22))
	fmt.Printf("%v\n", chalkReplacer([]int{3, 4, 1, 2}, 25))
}

func _chalkReplacer(chalk []int, k int) int {
	i := 0
	for {

		if k < chalk[i] {
			return i
		}
		k -= chalk[i]
		// reset
		if i == len(chalk)-1 {
			i = 0
			continue
		}
		i++
	}
}

// In case we have a big number, we can reduce the loop by mod
func chalkReplacer(chalk []int, k int) int {
	total := 0
	for _, v := range chalk {
		total += v
	}
	if total < k {
		k %= total
	}
	return _chalkReplacer(chalk, k)
}
