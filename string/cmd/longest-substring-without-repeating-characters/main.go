package main

import "fmt"

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
}

// sliding window
func lengthOfLongestSubstring(s string) int {
	out := 0
	check := make([]int, 256)
	for i, j := 0, 0; j < len(s); j++ {
		check[s[j]]++
		// if we found duplicated
		// we move the first of sliding windows to one
		for check[s[j]] > 1 {
			check[s[i]]--
			i++
		}
		if j-i+1 > out {
			out = j - i + 1
		}
	}
	return out
}
