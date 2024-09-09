package main

import "fmt"

func main() {
	//4, 2, 8
	headListNodes := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 8,
			},
		},
	}
	// 1, 4, 4, nil, 2, 2, nil, 1, nil, 6, 8, nil, nil, nil, nil, 1, 3
	rootTreeNodes := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 4, Left: nil, Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
		}},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			},
			Right: nil},
	}
	fmt.Printf("%t\n", isSubPath(headListNodes, rootTreeNodes))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	return dfs(head, head, root)
}

// Thing that need to be notice:
//   - We need to reset count if we find the root.Val != current.Val (or event with the head not match)
//   - We need to return true in case the list reach its end
func dfs(head *ListNode, current *ListNode, root *TreeNode) bool {
	// if we reach to the end, then we are done
	if current == nil {
		return true
	}
	// if we reach to the end of tree without finish the current, then not found
	if root == nil {
		return false
	}
	// If the current value is the same as the root value, then we start the match
	// go to the next node
	if current.Val == root.Val {
		current = current.Next
	} else if head.Val == root.Val {
		head = head.Next
	} else {
		current = head
	}
	// Go to the left or right and get the result
	return dfs(head, current, root.Left) || dfs(head, current, root.Right)
}
