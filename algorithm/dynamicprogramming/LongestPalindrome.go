package dynamicprogramming

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 题目：5. 最长回文子串
// 难度：中等
// 描述：给你一个字符串 s，找到 s 中最长的回文子串。
// 思路：动态规划
// 状态：dp[i][j]表示s[i..j]子串的是否为回文
// dp[i][j] = true, (i >= j)
// dp[i][j] = dp[i+1][j-1] (s[i] = s[j])
// dp[i][j] = false (s[i] != s[j])
// 由于dp[i][j]的状态依赖于dp[i+1][j-1]，而左下右的状态已经计算出来，所以可以斜向遍历或反向遍历
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if i > 0 {
			dp[i][i-1] = true
		}
	}

	start, maxLen := 0, 1
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}

			if dp[i][j] && j-i+1 > maxLen {
				start, maxLen = i, j-i+1
			}
		}
	}
	return s[start : start+maxLen]
}
