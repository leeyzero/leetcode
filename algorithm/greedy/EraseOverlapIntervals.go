package greedy

import (
	"sort"
)

// 题目：435. 无重叠区间
// https://leetcode-cn.com/problems/non-overlapping-intervals/submissions/
// 难度：中等
// 思路：贪心
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var ans int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prev[1] {
			ans++
		} else {
			prev = intervals[i]
		}
	}
	return ans
}
