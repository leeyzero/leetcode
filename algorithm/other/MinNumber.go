package other

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
// 题目：剑指 Offer 45. 把数组排成最小的数
// 难度：中等
// 思路：对于字符串x, y
// x + y > y + x, 说明y应该在前面
// x + y < y + x, 说明x应该在前面
func minNumber(nums []int) string {
	strNums := []string{}
	for _, num := range nums {
		strNums = append(strNums, strconv.FormatInt(int64(num), 10))
	}

	sort.Slice(strNums, func(x, y int) bool {
		if strNums[x]+strNums[y] < strNums[y]+strNums[x] {
			return true
		}
		return false
	})
	return strings.Join(strNums, "")
}
