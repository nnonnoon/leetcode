package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	// fmt.Print(root)
	res := maxDepth(&root)
	fmt.Println(res)

	// TODO space array
	a := []int{1, 2, 3}
	b := []int{5, 6, 7}
	a = append(a, b...)
	fmt.Println(a)
}

func maxDepth(root *TreeNode) int {
	fmt.Print(root)
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
