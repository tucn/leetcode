package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("%v\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Printf("%v\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("%v\n", longestCommonPrefix([]string{"cir", "car"}))

}

func longestCommonPrefix(strs []string) string {
	// O(n) since we need to go through all strings
	// In fact, we only need to check the first 2 to make sure we have common prefix
	// The problem here is the length of prefix (for example: [1,2] = 3 prefix, but [1,2,3] = 2 prefix )
	// We will try to reduce the common prefix length as we go through the strings element
	commonPrefix := ""
	for index, val := range strs {
		if index == 0 {
			commonPrefix = val
			continue
		}
		checkLen := min(len(val), len(commonPrefix))
		newCommon := ""
		for idx := range checkLen {
			if commonPrefix[idx] == val[idx] {
				newCommon += string(commonPrefix[idx])
				continue
			}
			break
		}
		if len(newCommon) == 0 {
			return ""
		}
		commonPrefix = newCommon
	}
	return commonPrefix

}
