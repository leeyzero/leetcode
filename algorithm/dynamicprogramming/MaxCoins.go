package dynamicprogramming

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode.cn/problems/burst-balloons/
// 312. 戳气球
// 难度：困难
/*
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
求所能获得硬币的最大数量。

示例 1：
输入：nums = [3,1,5,8]
输出：167
解释：
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167

示例 2：
输入：nums = [1,5]
输出：10

提示：
n == nums.length
1 <= n <= 300
0 <= nums[i] <= 100

思路：动态规划
dp[i][j] = x表示戳破气球i和气球j之间（不包括i和j）的所有气球获得的最大硬币数量
*/

func maxCoins(nums []int) int {
	points := make([]int, len(nums)+2)
	points[0], points[len(points)-1] = 1, 1
	for i := 1; i < len(points)-1; i++ {
		points[i] = nums[i-1]
	}

	// dp[i][j]表示戳破(i, j)之间的气球（不包括i和j）获得的硬币最大数量
	dp := make([][]int, len(points))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(points))
	}

	// 0 X X X
	// 0 0 X X
	// 0 0 0 X
	// 0 0 0 0
	n := len(points)
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = base.Max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[k]*points[j])
			}
		}
	}
	return dp[0][len(dp)-1]
}
