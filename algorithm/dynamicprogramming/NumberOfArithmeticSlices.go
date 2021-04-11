package dynamicprogramming

/*
https://leetcode-cn.com/problems/arithmetic-slices/
题目：413. 等差数列划分
难度：中等
描述：如果一个数列至少有三个元素，并且任意两个相邻元素之差相同，则称该数列为等差数列。

例如，以下数列为等差数列:

1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9
以下数列不是等差数列。

1, 1, 2, 5, 7

数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。

如果满足以下条件，则称子数组(P, Q)为等差数组：

元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。

函数要返回数组 A 中所有为等差数组的子数组个数。

思路：
这道题略微特殊，因为要求是等差数列，可以很自然的想到子数组必定满足num[i] - num[i-1]
= num[i-1] - num[i-2]。然而由于我们对于dp 数组的定义通常为以i 结尾的，满足某些条件的子数
组数量，而等差子数组可以在任意一个位置终结，因此此题在最后需要对dp 数组求和。
*/
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1] + 1
	}

	var ans int
	for i := 0; i < len(dp); i++ {
		ans += dp[i]
	}
	return ans
}
