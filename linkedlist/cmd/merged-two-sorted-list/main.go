package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 4},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Printf("%v\n", mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return &ListNode{}
	}
	// create return ListNode
	root := &ListNode{}
	// create current node
	cur := root
	for list1 != nil || list2 != nil {
		if list1.Val < list2.Val || list2 == nil {
			cur.Next = list1
			list1 = list1.Next
			continue
		}
		if list1 == nil {
			cur.Next = list2
			list2 = list2.Next
			continue
		}
		// list2.Val > list1.Val
		cur = cur.Next
	}
	return root.Next
}
