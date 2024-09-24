package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	test1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	test2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	fmt.Printf("%v\n", addTwoNumbers(
		test1, test2,
	))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// This is used to store the extra value, for example: 9 + 9 = 18 stored 8 and extra is 1
	// store := num % 10
	// extra := num / 10
	extra := 0
	// Go through all l1 and l2
	output := &ListNode{}
	current := output
	for l1 != nil || l2 != nil || extra != 0 {
		current.Next = &ListNode{}
		current = current.Next
		if l1 != nil && l2 != nil {
			num := l1.Val + l2.Val + extra
			extra = 0
			if num >= 10 {
				extra = num / 10
			}
			current.Val = num % 10
			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		if l2 != nil {
			num := l2.Val + extra
			extra = 0
			if num >= 10 {
				extra = num / 10
			}
			extra = num / 10
			current.Val = num % 10
			l2 = l2.Next
			continue
		}
		if l1 != nil {
			num := l1.Val + extra
			extra = 0
			if num >= 10 {
				extra = num / 10
			}
			extra = num / 10
			current.Val = num % 10
			l1 = l1.Next
			continue
		}
		if extra != 0 {
			current.Val = extra
			break
		}
	}
	return output.Next
}
