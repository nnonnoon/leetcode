package main

import (
	"fmt"
	_ "fmt"
)

var visited [][]bool
var count int
var m int
var n int

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	res := numIslands(grid)

	fmt.Print(res)

}

func numIslands(grid [][]byte) int {
	visited = make([][]bool, len(grid))
	m = len(grid)
	n = len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	count = 0

	for i := 0; i < m; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == '0' {
		return
	}

	visited[i][j] = true

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)

}
