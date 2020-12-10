package algorithm

import (
	"strconv"
	"sort"
	"strings"
)

// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
// 解题思路：对于字符串x, y
// x + y > y + x, 说明y应该在前面
// x + y < y + x, 说明x应该在前面
func minNumber(nums []int) string {
	strNums := []string{}
	for _, num := range nums {
		strNums = append(strNums, strconv.FormatInt(int64(num), 10))
	}

	sort.Slice(strNums, func(x, y int) bool {
		if strNums[x] + strNums[y] < strNums[y] + strNums[x] {
			return true
		}
		return false
	})
	return strings.Join(strNums, "")
}