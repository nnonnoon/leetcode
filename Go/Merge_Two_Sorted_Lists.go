package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next = &ListNode{}

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{}

	res := mergeTwoLists(&l1, &l2)

	fmt.Println(res)

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println(l1)
	newList := &ListNode{} // ***** สร้างโหนดใหม่ Tail
	fmt.Println(newList)
	out := newList // ***** ประกาศ Header

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			newList.Next = l2      // ***** เชื่อมก่อน
			l2 = l2.Next           // ***** ขยับ l2
			newList = newList.Next // ***** ขยับ newList ไปที่ l2
		} else if l2.Val > l1.Val {
			newList.Next = l1
			l1 = l1.Next
			newList = newList.Next
		}
	}

	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return out.Next // ***** ส่งตัวที่ถัดจาก Header ออกไป
}
