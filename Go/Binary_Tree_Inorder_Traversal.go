package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	res := inorderTraversal(&root)
	println(res)

	// ***** Pointer in golang~
	i := 42
	p := &i
	print(*p)
}

func inorderTraversal(root *TreeNode) []int {
	xs := []int{}
	if root != nil {
		xs = append(xs, inorderTraversal(root.Left)...)
		xs = append(xs, root.Val)
		xs = append(xs, inorderTraversal(root.Right)...)
	}
	return xs
}
