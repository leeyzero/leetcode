package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
// 题目：剑指 Offer 14- I. 剪绳子
// 描述：
// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
//
// 思路：f(n)表示长度为n的绳子切分后得到的最大值
// 每次将一段绳子剪成两端时，剩下的部分可以继续剪，也可以不剪，可以得到递推式：
// f(n) = max(i * (n-i), i * f(n-i)) i=1,2,3...,n-2
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			// max中的dp[i]是为了在内层循环中保持最大值
			dp[i] = base.Max(dp[i], base.Max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
