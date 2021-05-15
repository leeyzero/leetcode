package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/2-keys-keyboard/
// 题目：650. 只有两个键的键盘
// 难度：中等
// 描述：
// 最初在一个记事本上只有一个字符 'A'。你每次可以对这个记事本进行两种操作：
// Copy All (复制全部) : 你可以复制这个记事本中的所有字符(部分的复制是不允许的)。
// Paste (粘贴) : 你可以粘贴你上一次复制的字符。
// 给定一个数字 n 。你需要使用最少的操作次数，在记事本中打印出恰好 n 个 'A'。输出能够打印出 n 个 'A' 的最少操作次数。
// 思路：动态规划
// dp[i]表示延展到长度为i的最少操作次数, 对于每个位置j, 如果j可以被i整除，那么i就可以由长度j操作得到，操作的次数为i/j.
// base case: dp[1] = 0
// 状态转移方程：dp[i] = min(dp[i], dp[j] + i/j), (i % j = 0)
func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				dp[i] = base.Min(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}
