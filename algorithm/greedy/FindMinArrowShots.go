package greedy

import (
	"sort"
)

// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
// 452. 用最少数量的箭引爆气球
// 难度：中等
// 思路：贪心
func findMinArrowShots(points [][]int) int {
	if len(points) <= 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] <= points[j][1]
	})

	ans := 1
	for i, curr := 1, points[0]; i < len(points); i++ {
		if points[i][0] > curr[1] {
			curr = points[i]
			ans++
		}
	}
	return ans
}
