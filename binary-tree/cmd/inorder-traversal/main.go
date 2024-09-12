package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:  8,
				Left: &TreeNode{Val: 9},
			},
		},
	}
	fmt.Printf("%v\n", inorderTraversal2(test1))
}

// Left most and depth
func inorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)
	if root == nil {
		return out
	}
	out = append(out, inorderTraversal(root.Left)...)
	out = append(out, root.Val)
	out = append(out, inorderTraversal(root.Right)...)
	return out
}

func inorderTraversal2(root *TreeNode) []int {
	out := make([]int, 0)
	stack := list.New()
	cur := root
	for stack.Len() != 0 || cur != nil {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		cur = stack.Back().Value.(*TreeNode)
		out = append(out, cur.Val)
		stack.Remove(stack.Back())
		cur = cur.Right
	}
	return out
}
