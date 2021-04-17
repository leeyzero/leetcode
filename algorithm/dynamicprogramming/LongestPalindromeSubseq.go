package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/longest-palindromic-subsequence/
// 题目：516. 最长回文子序列
// 难度：中等
// 描述：给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
// 思路：动态规划
// dp[i][j]：表示在子串s[i..j]中，最长回文子序列的长度
// 状态转移：
// dp[i][j] = 0, (i > j)
// dp[i][j] = 1, (i = j)
// dp[i][j] = dp[i+1][j-1] + 2, (s[i] == s[j])
// dp[i][j] = max(dp[i][j-1], dp[i+1][j])
func longestPalindromeSubseq(s string) int {
	if len(s) <= 0 {
		return 0
	}

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// 为了保证每次计算dp[i][j], 左下右方向的位置已经被计算出来, 需要斜着或反向遍历
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = base.Max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}
