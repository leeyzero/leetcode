package search

// https://leetcode-cn.com/problems/number-of-provinces/
// 547. 省份数量
// 中等
// 思路：dfs
// 注意：图的表示方法，isConnected[i][j]表示第i个节点和第j个节点相连，一共有len(isConnected)个节点
func findCircleNum(isConnected [][]int) int {
	ans := 0
	visited := make([]bool, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		if visited[i] {
			continue
		}

		dfsFindCircleNum(isConnected, i, visited)
		ans++
	}
	return ans
}

func dfsFindCircleNum(isConnected [][]int, i int, visited []bool) {
	visited[i] = true
	for j := 0; j < len(isConnected[i]); j++ {
		if isConnected[i][j] == 0 || visited[j] {
			continue
		}
		dfsFindCircleNum(isConnected, j, visited)
	}
}
