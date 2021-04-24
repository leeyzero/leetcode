package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
// 题目：剑指 Offer 48. 最长不含重复字符的子字符串
// 难度：中等
// 思路：动态规划
// 状态定义：f(j)表示以s[j]为结尾的最长不重复子字符串的长度
// 转移方程：固定右边界j, 设s[j]左边距离最近的相同字符为s[i], 即s[i] = s[j]
// 1. 当i < 0, 即s[j]左边无相同字符，则f(j) = f(j-1) + 1
// 2. 当f(j-1) < j - i, 说明s[i]在f(j-1)子字符串之外，则f(j) = f(j-1) + 1
// 3. 当f(j-1) >= j - i, 说明s[i]在f(j-1)子字符串之内，则f(j) = j - i
// 返回值：max(f(i))
func lengthOfLongestSubstring(s string) int {
	hash := map[byte]int{}
	ans, tmp := 0, 0
	for j := 0; j < len(s); j++ {
		i := -1
		if v, ok := hash[s[j]]; ok {
			i = v
		}
		// 更新hash表
		hash[s[j]] = j
		if tmp < j-i {
			tmp += 1
		} else {
			tmp = j - i
		}
		ans = base.Max(ans, tmp)
	}
	return ans
}
