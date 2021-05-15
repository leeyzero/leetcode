package dynamicprogramming

import (
	"math"
)

// https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
// 题目：剑指 Offer 60. n个骰子的点数
// 难度：中等
// 描述：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
// 思路：每个骰子有6个面，分别时[1,6]，n个骰子的排列组合一个有6^n种组合，最小值为n（每个面都是1），最大6n（每个面都是6）
// 先将骰子分成两堆，第一堆有1个骰子，第二堆有n-1个骰子，需要计算出第一堆中1~6中每一个点数和n-1堆中的点数和。对于n-1堆，可以
// 采用同样的策略，一直到骰子分完（剩下0个骰子）
// 方法一：递归，递归会出现很多重复计算，时间效率比较低
// 方法二：动态规划，
// 表示状态：dp[i][j]表示投出i个骰子后点数和为j的次数
// 状态转移：dp[n][j] = dp[n][j] + dp[n-1][j-i] (i = [1, 6])
// 初始状态：dp[1][j] = 1 (j = [1,6])
const (
	MAX_VALUE = 6
)

func dicesProbability(n int) []float64 {
	if n <= 0 {
		return []float64{}
	}

	// ans := dicesProbabilityByRecursion(n)
	// ans := dicesProbabilityByDp(n)
	ans := dicesProbabilityByDp2(n)
	return ans
}

// 方法一：递归
func dicesProbabilityByRecursion(n int) []float64 {
	a := make([]float64, n*MAX_VALUE-n+1)
	s := 0
	total := math.Pow(float64(MAX_VALUE), float64(n))
	dicesProbabilityByRecursionCore(n, n, s, a)
	for i := 0; i < len(a); i++ {
		a[i] /= total
	}
	return a
}

func dicesProbabilityByRecursionCore(n int, curr int, s int, a []float64) {
	if curr <= 0 {
		a[s-n]++
		return
	}

	for i := 1; i <= MAX_VALUE; i++ {
		dicesProbabilityByRecursionCore(n, curr-1, s+i, a)
	}
}

// 方法二：动态规划
func dicesProbabilityByDp(n int) []float64 {
	// n个骰子
	dp := make([][]float64, n+1)
	for i := 0; i < len(dp); i++ {
		// i个骰子可能出现的点数
		dp[i] = make([]float64, MAX_VALUE*n+1)
	}

	// 初始值, 1个骰子出现点数为1-6的次数均是一次
	for j := 1; j <= MAX_VALUE; j++ {
		dp[1][j] = float64(1)
	}

	// 从第2个骰子开始
	for i := 2; i <= n; i++ {
		// 计算第i个骰子，出现点数和为j的次数
		for j := i; j <= MAX_VALUE*i; j++ {
			// 计算第i个骰子可能出现的值
			for k := 1; k <= MAX_VALUE && j > k; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}

	// 计算所有点数[n, 6n]的概率
	total := math.Pow(float64(MAX_VALUE), float64(n))
	ans := []float64{}
	for j := n; j <= MAX_VALUE*n; j++ {
		ans = append(ans, dp[n][j]/total)
	}
	return ans
}

// 方法三：优化版本的动态规划
// 从方法二中可以看出，计算第i个骰子时，只依赖i-1个骰子的取值，所以不用保留所有骰子的取值
// 用一维数组来保存第i个骰子的状态，然后对i+1个骰子可能出现的点数j从大到小遍历，实现从i到i+1的转换。
func dicesProbabilityByDp2(n int) []float64 {
	dp := make([]float64, MAX_VALUE*n+1)

	// 初始1个骰子的情况
	for j := 1; j <= MAX_VALUE; j++ {
		dp[j] = 1
	}

	// 从第2个骰子开始计算
	for i := 2; i <= n; i++ {
		// 计算i个骰子点数和为j=[6i..i]的次数
		for j := MAX_VALUE * i; j >= i; j-- {
			// 由于复用了数组，每次计算之前需要重置
			dp[j] = float64(0)

			// 第i个骰子可能的取值
			for k := 1; k <= MAX_VALUE; k++ {
				// 第i个骰子只依赖于i-1个骰子的状态，对于i-1个骰子，最小和为i-1
				if j-k < i-1 {
					break
				}
				dp[j] += dp[j-k]
			}
		}
	}

	// 计算所有点数[n, 6n]的概率
	total := math.Pow(float64(MAX_VALUE), float64(n))
	ans := []float64{}
	for j := n; j <= MAX_VALUE*n; j++ {
		ans = append(ans, dp[j]/total)
	}
	return ans
}
