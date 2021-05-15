package search

// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
// 题目：剑指 Offer 13. 机器人的运动范围
// 难度：中等
// 描述：
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
// 思路：DFS
func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	ans := 0
	movingCountCore(0, 0, k, visited, &ans)
	return ans
}

func movingCountCore(i int, j int, k int, visited [][]bool, ans *int) {
	if i < 0 || i >= len(visited) || j < 0 || j >= len(visited[i]) || visited[i][j] || !isCoordValid(i, j, k) {
		return
	}

	visited[i][j] = true
	*ans++
	// 上
	movingCountCore(i-1, j, k, visited, ans)
	// 右
	movingCountCore(i, j+1, k, visited, ans)
	// 下
	movingCountCore(i+1, j, k, visited, ans)
	// 左
	movingCountCore(i, j-1, k, visited, ans)
}

func isCoordValid(i, j int, k int) bool {
	if sumNumber(i)+sumNumber(j) <= k {
		return true
	}
	return false
}

func sumNumber(i int) int {
	ans := 0
	for i > 0 {
		ans += i % 10
		i /= 10
	}
	return ans
}
