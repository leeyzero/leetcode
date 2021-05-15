package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/edit-distance/
// 题目：72. 编辑距离
// 难度：困难
// 描述：
// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 思路：动态规划
// 解决两个字符串的动态规划问题，一般都是用两个指针i,j分别指向两个字符串的最后，然后一步步往前走，缩小问题的规模
// dp[i][j]表示s1[0..i]和s2[0..j]的最小编辑距离
// base case:
// - 如果i走完s1时，j还没有走完s2, 那就只能用插入操作把s2剩下的字符全部插入s1
// - 如果j走完s2，i还没有走完s1，那么只能用删除操作把s1缩短为s2
// 状态转换：
// dp[i][j] = dp[i-1][j-1], s1[i] == s2[j]
// dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1], dp[i-1][j-1]+1) + 1, s1[i] != s2[j]
func minDistance(s1 string, s2 string) int {
	return minDisanceByDP(s1, s2)
}

// 递归：至上而下
func minDistanceByRecursive(s1 string, s2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if s1[i] == s2[j] {
			return dp(i-1, j-1) // 什么都不用做
		} else {
			return base.Min(base.Min(dp(i, j-1)+1, dp(i-1, j)+1), dp(i-1, j-1)+1)
		}
	}
	return dp(len(s1)-1, len(s2)-1)
}

// 备忘录：递归+剪枝
func minDistanceByMemo(s1 string, s2 string) int {
	type Pair [2]int
	memo := map[Pair]int{}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// 先查备忘录，避免重复计算
		if v, ok := memo[Pair{i, j}]; ok {
			return v
		}

		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if s1[i] == s2[j] {
			memo[Pair{i, j}] = dp(i-1, j-1) // 什么都不用做
		} else {
			memo[Pair{i, j}] = base.Min(base.Min(dp(i, j-1)+1, dp(i-1, j)+1), dp(i-1, j-1)+1)
		}
		return memo[Pair{i, j}]
	}
	return dp(len(s1)-1, len(s2)-1)
}

// 动态规划：至下而上
func minDisanceByDP(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = base.Min(base.Min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}
