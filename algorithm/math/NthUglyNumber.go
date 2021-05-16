package math

// https://leetcode-cn.com/problems/chou-shu-lcof/
// 题目：剑指 Offer 49. 丑数
// 难度：中等
// 描述：我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		min := minOfTree(dp[p2]*2, dp[p3]*3, dp[p5]*5)
		dp[i] = min
		for dp[p2]*2 <= min {
			p2++
		}
		for dp[p3]*3 <= min {
			p3++
		}
		for dp[p5]*5 <= min {
			p5++
		}
	}
	return dp[n]
}

func minOfTree(x, y, z int) int {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}
