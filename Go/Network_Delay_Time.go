package main

import (
	"fmt"
	"math"
)

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	res := networkDelayTime(times, n, k)
	fmt.Print(res)
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([]map[int]int, n+1)

	for i := 0; i < len(graph); i++ {
		graph[i] = make(map[int]int)
	}

	for _, v := range times {
		graph[v[0]][v[1]] = v[2]
	}

	minpath := make([]int, n+1)
	visited := make([]bool, n+1)
	visited[k] = true

	for { // ** วนให้ครบทุกชั้นของ graph
		node, l := find(graph, minpath, visited)

		if node == 0 {
			break
		}
		visited[node] = true
		minpath[node] = l
	}

	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			return -1
		}
	}

	res := 0
	for i := 1; i < len(minpath); i++ {
		res = max(res, minpath[i])
	}
	return res
}

func find(graph []map[int]int, minpath []int, visited []bool) (int, int) {
	l := math.MaxInt64
	newnode := 0

	for i, v := range graph { // ** วนไปทีละชั้นของ graph
		if visited[i] {
			for node, weight := range v {
				if !visited[node] {
					if minpath[i]+weight < l { // ** เช็คภายในชั้นว่า node แต่ละตัวตัวไหนน้อยกว่าก็ return ตัวนั้น
						l = minpath[i] + weight
						newnode = node
					}
				}
			}
		}
	}

	return newnode, l
}

func max(res int, minpath int) int {
	if res > minpath {
		return res
	}
	return minpath
}
