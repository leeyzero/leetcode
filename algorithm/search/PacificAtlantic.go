package search

// https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
// 题目：417. 太平洋大西洋水流问题
// 难度：中等
// 描述：
// 给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。
// 规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。
// 请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。
// 思路：如果对每个高度进行判断，即要考虑是否能流入太平洋，也要考虑流入大西洋，可能反过来考虑：
// 从太平洋和大西洋的两条边出发，分别看能流入哪些节点，最后遍历一下数组，就能知道即能注入太平洋又能流入大西洋的节点了
var directs = [4][2]int{
	{0, -1}, // 左
	{-1, 0}, // 上
	{0, 1},  // 右
	{1, 0},  // 下
}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) <= 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])
	visitedP := make([][]bool, m)
	visitedA := make([][]bool, m)
	for i := 0; i < m; i++ {
		visitedP[i] = make([]bool, n)
		visitedA[i] = make([]bool, n)
	}

	// left
	for i := 0; i < m; i++ {
		dfsPacificAtlantic(heights, i, 0, visitedP)
	}
	// top
	for j := 0; j < n; j++ {
		dfsPacificAtlantic(heights, 0, j, visitedP)
	}
	// right
	for i := 0; i < m; i++ {
		dfsPacificAtlantic(heights, i, n-1, visitedA)
	}
	// bottom
	for j := 0; j < n; j++ {
		dfsPacificAtlantic(heights, m-1, j, visitedA)
	}

	ans := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visitedP[i][j] && visitedA[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func dfsPacificAtlantic(heights [][]int, i, j int, visited [][]bool) {
	if visited[i][j] {
		return
	}

	visited[i][j] = true
	for _, d := range directs {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= len(heights) || y < 0 || y >= len(heights[0]) {
			continue
		}
		if heights[x][y] < heights[i][j] {
			continue
		}
		dfsPacificAtlantic(heights, x, y, visited)
	}
}
