package main

import (
	"fmt"
)

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]

Example 3:
Input: root = []
Output: []

Example 4:
Input: root = [1]
Output: [1]

Constraints:
	The number of nodes in the tree is in the range [0, 100].
	-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traverse(root, &result)
	return result
}

func traverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, result)
	*result = append(*result, root.Val)
	traverse(root.Right, result)
}

func main() {
	root1 := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	root2 := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}, &TreeNode{3, nil, &TreeNode{8, &TreeNode{9, nil, nil}, nil}}}
	root4 := &TreeNode{1, nil, nil}

	result1 := inorderTraversal(root1)
	result2 := inorderTraversal(root2)
	result3 := inorderTraversal(nil)
	result4 := inorderTraversal(root4)

	fmt.Println("Case 1:", result1)
	fmt.Println("Case 2:", result2)
	fmt.Println("Case 3:", result3)
	fmt.Println("Case 4:", result4)
}
