package search

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/max-area-of-island/
// 695. 岛屿的最大面积
// 中等
// 思路：dfs
// 方法一：递归实现
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}

	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			currArea := 0
			dfs(grid, i, j, visited, &currArea)
			maxArea = base.Max(maxArea, currArea)
		}
	}
	return maxArea
}

func dfs(grid [][]int, i, j int, visited [][]bool, area *int) {
	if len(grid) <= 0 || i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	if visited[i][j] || grid[i][j] == 0 {
		return
	}

	visited[i][j] = true
	(*area)++

	// 左
	dfs(grid, i, j-1, visited, area)
	// 上
	dfs(grid, i-1, j, visited, area)
	// 右
	dfs(grid, i, j+1, visited, area)
	// 下
	dfs(grid, i+1, j, visited, area)
}

// 方法二：使用栈
func maxAreaOfIsland2(grid [][]int) int {
	type Pair [2]int

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			currArea := 0
			stack := []Pair{{i, j}}
			for len(stack) > 0 {
				x, y := stack[len(stack)-1][0], stack[len(stack)-1][1]
				stack = stack[:len(stack)-1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] == 0 {
					continue
				}

				currArea++
				grid[x][y] = 0

				// 左
				stack = append(stack, Pair{x, y - 1})
				// 上
				stack = append(stack, Pair{x - 1, y})
				// 右
				stack = append(stack, Pair{x, y + 1})
				// 下
				stack = append(stack, Pair{x + 1, y})
			}
			maxArea = base.Max(maxArea, currArea)
		}
	}
	return maxArea
}
