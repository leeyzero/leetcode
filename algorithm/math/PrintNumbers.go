package math

// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
// 题目：剑指 Offer 17. 打印从1到最大的n位数
// 难度：简单
// 描述：输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
// 思路：不考虑大数的情况
func printNumbers(n int) []int {
	limit := 1
	for i := 1; i <= n; i++ {
		limit *= 10
	}

	ans := []int{}
	for i := 1; i < limit; i++ {
		ans = append(ans, i)
	}
	return ans
}
