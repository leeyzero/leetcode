package dynamicprogramming

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 题目：剑指 Offer 10- I. 斐波那契数列
// 难度：中等
// 思路：动态规划+状态空间压缩
func fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for k := 2; k <= n; k++ {
		next := (a + b) % 1000000007
		a, b = b, next
	}
	return b
}
