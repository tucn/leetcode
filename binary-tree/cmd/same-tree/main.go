package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	p := &TreeNode{
		1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	q := &TreeNode{
		1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	fmt.Printf("%v\n", isSameTree(p, q))
	p = &TreeNode{
		1,
		&TreeNode{2, nil, nil},
		nil,
	}
	q = &TreeNode{
		1,
		nil,
		&TreeNode{2, nil, nil},
	}
	fmt.Printf("%v\n", isSameTree(p, q))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
