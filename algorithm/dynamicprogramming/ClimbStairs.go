package dynamicprogramming

// https://leetcode-cn.com/problems/climbing-stairs/
// 题目：70. 爬楼梯
// 难度：简单
// 描述：
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		next := a + b
		a, b = b, next
	}
	return b
}

func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
