package dynamicprogramming

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
// 题目：剑指 Offer 10- II. 青蛙跳台阶问题
// 难度：简单
// 描述：
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
// 思路：动态规划
func numWays(n int) int {
	if n <= 1 {
		return 1
	}

	a, b := 1, 1
	for i := 2; i <= n; i++ {
		next := (a + b) % 1000000007
		a, b = b, next
	}
	return b
}
