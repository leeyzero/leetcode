package sort

import (
	"sort"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// 题目：56. 合并区间
// 链接：https://leetcode-cn.com/problems/merge-intervals/
// 难度：中等
// 思路：排序

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = base.Max(ans[len(ans)-1][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
