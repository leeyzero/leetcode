package dynamicprogramming

import (
	"strconv"
)

// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
// 题目：剑指 Offer 46. 把数字翻译成字符串
// 难度：中等
// 思路：动态规划
// num表示为X1X2X3...Xn
// 记f(i)表示以Xi结尾的字符串
// 当XnXn-1为整体翻译时，翻译的方案数为f(i-2)
// 当Xn为整体翻译时，翻译的方案数为f(i-1)
// 可得递推式：
// f(i) = f(i-1) + f(i-2), XiXi-1可翻译
// f(i) = f(i-1), Xi可翻译
func translateNum(num int) int {
	if num < 0 {
		return 0
	}

	s := strconv.FormatInt(int64(num), 10)
	a, b := 1, 1
	for i := 2; i <= len(s); i++ {
		subNum, _ := strconv.ParseInt(s[i-2:i], 10, 32)
		c := b
		if subNum >= 10 && subNum <= 25 {
			c = a + b
		}
		a, b = b, c
	}
	return b
}
