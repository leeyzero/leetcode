package math

// https://leetcode-cn.com/problems/gray-code/
// 题目：89. 格雷编码
// 难度：中等
// 描述：
// 格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
// 给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。即使有多个不同答案，你也只需要返回其中一种。
// 格雷编码序列必须以 0 开头。
// 思路：镜像反映法
// 设G(n)为n阶格雷码，给G(n)每个元素的二进制前加0得G1(n)
// 设R(n)G(n)的倒序(镜像)为，给R(n)每个元素的二进制前加1得R1(n)
// G(n+1) = G1(n) + R1(n)，其中G1(n) = G(n)，由于最高位前默认为 00
func grayCode(n int) []int {
	ans := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, head+ans[j])
		}
		head = head << 1
	}
	return ans
}
