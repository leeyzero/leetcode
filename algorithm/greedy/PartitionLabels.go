package greedy

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/partition-labels/
// 题目：763. 划分字母区间
// 难度：中等
// 思路：贪心
// 由于每个字母只能出现在同一片段中，显示同一个字母第一次出现的位置和最后一次出现的位置必须在同一片段.
// 可以先计算出每个字母最后一次出现的位置，然后再次遍历字符串
func partitionLabels(S string) []int {
	last := map[uint8]int{}
	for i := 0; i < len(S); i++ {
		last[S[i]] = i
	}

	ans := []int{}
	start, end := 0, 0
	for i := 0; i < len(S); i++ {
		end = base.Max(end, last[S[i]])
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
