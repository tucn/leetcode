package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	out := modifiedList([]int{1, 2, 3}, head)
	for {
		if out == nil {
			break
		}
		fmt.Println(out.Val)
		out = out.Next
	}
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	// The best way is to update the head it self
	// since max of n is 10^5, we can use numMap := [100001]bool{}
	// (reference from one of the solutions)
	numMap := [100001]bool{}
	for _, i := range nums {
		numMap[i] = true
	}
	rootpointer := &ListNode{
		Val:  -1,
		Next: head,
	}
	current := rootpointer
	for current.Next != nil {
		if numMap[current.Next.Val] {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return rootpointer.Next
}
