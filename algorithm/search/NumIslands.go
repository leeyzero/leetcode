package search

// https://leetcode-cn.com/problems/number-of-islands/
// 题目：200. 岛屿数量
// 难度：中等
// 描述：
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// 思路：dfs
func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				numIslandsDFS(grid, i, j)
				ans++
			}
		}
	}
	return ans
}

func numIslandsDFS(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	numIslandsDFS(grid, i, j-1) // 左
	numIslandsDFS(grid, i-1, j) // 上
	numIslandsDFS(grid, i, j+1) // 右
	numIslandsDFS(grid, i+1, j) // 下
}
