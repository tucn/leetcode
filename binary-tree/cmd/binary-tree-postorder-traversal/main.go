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
	// fmt.Println("Without Recursive")
	// fmt.Println(postOrderTraversalWR(test1))
	// fmt.Println(postOrderTraversalWR(test2))

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

// Define a struct Node with integer data, pointer to left child and pointer to right child.
// Define a helper function called "postorder" which takes a pointer to the head of the tree.
// Create a pointer "temp" and an unordered set "visited".
// While "temp" is not NULL and "temp" is not visited before:
//
//	a. If "temp" has a left child and the left child is not visited before, then set "temp" to its left child and continue the loop.
//	b. If "temp" does not have a left child or the left child is already visited, check if "temp" has a right child and the right child is not visited before. If yes, set "temp" to its right child and continue the loop.
//	c. If "temp" does not have a left child or the left child is already visited, and "temp" does not have a right child or the right child is already visited, then print the data of "temp", insert "temp" into "visited" set, and set "temp" to the head of the tree.
//
// Define a function called "newNode" which takes an integer data as input and returns a new Node with the given data, NULL left pointer, and NULL right pointer.
// @Todo
func postOrderTraversalWR(root *TreeNode) []int {
	return nil
}
