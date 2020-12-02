package algorithm

// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
// 解题思路：f(n)表示长度为n的绳子切分后得到的最大值
// 每次将一段绳子剪成两端时，剩下的部分可以继续剪，也可以不剪，可以得到递推式：
// f(n) = max(i * (n-i), i * f(n-i)) i=1,2,3...,n-2
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i - 1; j++ {
			// max中的dp[i]是为了在内层循环中保持最大值
			dp[i] = max(dp[i], max(j * (i-j), j * dp[i-j]))
		}
	}
	return dp[n]
}