package dynamicprogramming

// https://leetcode-cn.com/problems/decode-ways/
// 题目：91. 解码方法
// 难度：中等
// 思路：动态规划
// dp[i]表示s的前i个字符的解码方法数
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
