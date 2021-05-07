package dynamicprogramming

// https://leetcode-cn.com/problems/word-break/
// 题目：139. 单词拆分
// 难度：中等
// 思路：动态规划
// dp[i]表示字符串s前i个字符组成的字符串s[0..i-1]是否能被空格拆分成若干个字典中出现的单词
// 状态转移方程：dp[i] = dp[j] && check(s[j..i−1])
func wordBreak(s string, wordDict []string) bool {
	hash := map[string]bool{}
	for _, word := range wordDict {
		hash[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && hash[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
