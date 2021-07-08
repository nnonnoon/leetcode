package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	n := 5
	connections := [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}

	_ = n
	_ = connections
	res := makeConnected(n, connections)

	_ = res
}

func makeConnected(n int, connections [][]int) int {
	root := make([]int, n)

	for i := range root {
		root[i] = i
	}
	dup := 0
	for i := range connections {
		a := connections[i][0]
		b := connections[i][1]
		if !connect(a, b, root) {
			dup++
		}
		fmt.Println("1 : ", root)
	}
	set := make(map[int]bool, 0)

	for i := range root {
		root[i] = find(i, root) // ** ทำอีกรอบเพื่อเช็ค root อีกที
		set[root[i]] = true
	}
	fmt.Println("2 : ", root)
	fmt.Println(set)

	cnt := len(set)
	res := cnt - 1

	fmt.Println("cnt : ", cnt)
	fmt.Println("res : ", res)

	if dup < res {
		return -1
	} else {
		return res
	}
}

func connect(a int, b int, root []int) bool {
	ra := find(a, root)
	rb := find(b, root)

	fmt.Println("ra : ", ra)
	fmt.Println("rb : ", rb)
	if ra == rb {
		return false
	}
	root[rb] = ra // ** บอกว่าโหนดนั้น root เป็นใคร
	return true
}

func find(a int, root []int) int {
	if root[a] == a {
		return a
	}
	root[a] = find(root[a], root) // ** หา root ของแต่ละ node find ไปเรื่อยๆจนเจอ และเป็นโหนดนั้นให้บอกว่า root คือใคร
	return root[a]
}
