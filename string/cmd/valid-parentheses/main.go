package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", isValid("()"))
	fmt.Printf("%v\n", isValid("()[]{}"))
	fmt.Printf("%v\n", isValid("(]"))
	fmt.Printf("%v\n", isValid("([])"))
	fmt.Printf("%v\n", isValid("){"))
}

func isValid(s string) bool {
	// Open brackets must be closed by the same type of brackets.
	// Open brackets must be closed in the correct order.
	// Every close bracket has a corresponding open bracket of the same type.
	//
	// Base analysis: two pointer ? Queue
	if len(s) < 1 {
		return false
	}
	// s length should be even
	if len(s)%2 == 1 {
		return false
	}
	match := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	// We could see that if we found 1st parentheses at position i
	// the second match should be in the next: (i + 1), (i + 2), (i + 4), (i + 6) since
	// there should be another couple between those
	stack := list.New()
	for i := range s {
		iVal := rune(s[i])
		if _, ok := match[iVal]; ok {
			stack.PushBack(iVal)
			continue
		}
		// if not open, then should be a close parentheses
		// get latest pushed element
		if stack.Len() == 0 {
			return false
		}
		latest := stack.Back().Value.(rune)
		if match[latest] != iVal {
			return false
		}
		// matched, remove latest
		stack.Remove(stack.Back())
	}
	return stack.Len() == 0
}
