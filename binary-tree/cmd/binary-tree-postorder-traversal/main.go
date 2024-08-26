package main

import "fmt"

func main() {
	test1 := &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			nil,
		},
	}
	// [3 2 1]
	fmt.Println(postOrderTraversal(test1))

	test2 := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				6,
				nil,
				nil,
			},
		},
	}
	// [4 5 2 6 3 1]
	fmt.Println(postOrderTraversal(test2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// With recursive
func postOrderTraversal(root *TreeNode) []int {
	// post-order: left, right, root
	if root == nil {
		return nil
	}
	out := make([]int, 0)
	if root.Left != nil {
		left := postOrderTraversal(root.Left)
		out = append(out, left...)
	}
	if root.Right != nil {
		right := postOrderTraversal(root.Right)
		out = append(out, right...)
	}
	out = append(out, root.Val)
	return out
}
