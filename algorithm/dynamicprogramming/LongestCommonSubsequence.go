package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/longest-common-subsequence/
// 题目：1143. 最长公共子序列
// 难度：中等
// 描述：
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
// 思路：动态规划
// 定义：dp[i][j]表示到第一个字符串前i个字符结尾，第二个字符串前j个字符串结尾的最长公共子序列
// 状态转移方程：
// dp[i][j] = dp[i-1][j-1] (第一个字符串的第i个字符和与第二个字符串的第j个字符相同)
// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = base.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
