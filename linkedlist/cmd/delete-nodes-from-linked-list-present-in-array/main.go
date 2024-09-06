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
	numMap := make(map[int]int, 0)
	for _, i := range nums {
		numMap[i] = 0
	}
	// The logic is as follows:
	// - root is a pointer to head that contain very first value
	// - if the head.Val include in nums, then we need to point the root.Next to the next node
	// - For example: [1,2,3,2] with nums = [2]
	//       - rootnode = pointer to node1
	//       - repeat
	//       	- check rootnode.Next to see if they in nums, since node2 value is 2, change the rootnode.Next to node2.Next
	//          - if rootnode.Next is nil return root
	//          - else rootnode = pointer to rootnode.Next
	rootpointer := &ListNode{
		Val:  -1,
		Next: head,
	}
	nextpointer := head
	for {
		if nextpointer == nil {
			return head
		}
		if _, ok := numMap[nextpointer.Val]; ok {
			nextpointer = nextpointer.Next
			continue
		}

		rootpointer.Next = nextpointer
		nextpointer = nextpointer.Next
	}
	return head
}
