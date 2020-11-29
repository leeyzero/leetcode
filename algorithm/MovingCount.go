package algorithm

import (
	"fmt"
)
// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
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
	if i < 0 || i >= len(visited) || j < 0 || j >= len(visited[i]) || visited[i][j] {
		return
	}
	
	visited[i][j] = true
	if !isCoordValid(i, j, k) {
		return
	}

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
	if sumNumber(i) + sumNumber(j) <= k {
		fmt.Println(i, j, k)
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