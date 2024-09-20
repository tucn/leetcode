package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%d\n", uniqueFreq([]int{1, 2, 2, 1}))
	fmt.Printf("%d\n", uniqueFreq([]int{1, 1, 2, 3, 3, 3, 4, 4, 4}))
}

func uniqueFreq(N []int) int {
	fmap := make(map[int]int, len(N))
	// map frequencies so that we will have fmap := {1:2, 2:2}
	// incase of [1,1,2,2]
	for _, i := range N {
		fmap[i]++
	}
	// make a count map
	fcount := make([]int, len(fmap))
	for _, f := range fmap {
		fcount = append(fcount, f)
	}
	// Sort in the descending order
	sort.Sort(sort.Reverse(sort.IntSlice(fcount)))
	used := make(map[int]bool)
	deletions := 0
	for _, f := range fcount {
		for f > 0 && used[f] {
			f--
			deletions++
		}
		if f > 0 {
			used[f] = true
		}
	}
	return deletions
}
