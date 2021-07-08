package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// root := TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	// root.Left.Left = &TreeNode{Val: 3}
	// root.Left.Right = &TreeNode{Val: 3}
	// root.Left.Left.Left = &TreeNode{Val: 4}
	// root.Left.Left.Left.Left = &TreeNode{Val: 5}
	// root.Right = &TreeNode{Val: 2}

	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(root)

	res := isBalanced(&root)
	fmt.Println(res)
}

func isBalanced(root *TreeNode) bool {
	return maxDepth(root) != -1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	fmt.Println("Root", root)
	fmt.Println("Left", left)
	fmt.Println("Right", right)

	// ***** ตอนรวม 2 โหนดเข้าด้วยกัน ถ้าเป็น -1 จะเป็นตลอดไปไม่ต้องไปนั่งเช็คแล้ว
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
