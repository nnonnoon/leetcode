package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	p := TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 5}
	p.Right = &TreeNode{Val: 2}
	p.Right.Right = &TreeNode{Val: 6}
	p.Right.Left = &TreeNode{Val: 3}
	p.Right.Left.Left = &TreeNode{Val: 8}

	res := postorderTraversal(&p)
	fmt.Println("final = ", res)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	s := []*TreeNode{root}
	res := []int{}
	fmt.Println(s)
	fmt.Println(res)
	for len(s) > 0 {
		fmt.Println("before_s = ", s)
		n := s[len(s)-1]                   // ***** address ของ root (0xc00000c030) จะออกมาเป็น &{1 <nil> 0xc00000c048} คือ pop  address ค่าสุดท้ายมาคิด และเก็บไว้ที่ตัวแปร n
		s = s[:len(s)-1]                   // ***** เอาข้อมูลตั้งแต่ตัวที่ 0 ถึงก่อนสุดท้าย คือเอาตัวสุดท้ายออก
		res = append([]int{n.Val}, res...) // ***** res =  [1] -> [2 1] -> [3 2 1]
		// ***** res = append(res, n.Val)  // ***** res =  [1] -> [1 2] -> [1 2 3]
		fmt.Println("len(s) = ", len(s))
		fmt.Println("n = ", n)
		fmt.Println("after_s = ", s)
		fmt.Println("res = ", res)

		if n.Left != nil {
			s = append(s, n.Left)
		}
		if n.Right != nil {
			s = append(s, n.Right)
		}
		fmt.Println("s = ", s)
		fmt.Println("------------------------------")
	}
	return res
}
