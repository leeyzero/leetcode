package math

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
// 题目：剑指 Offer 65. 不用加减乘除做加法
// 描述：写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
// 思路：
// a(i)   b(i)   无进位和n(i)   进位c(i+1)
// 0      0      0             0
// 0      1      1             0
// 1      0      1             0
// 1      1      0             1
// 非进位和跟异或运算相同，进位跟与运算相同且需左移移位
// 非进位和：n = a ^ b
// 进位：c = (a & b) << 1
// (和 s) = (非进位和 n) + (进位 c), 即s = a + b转化为 s = n + c
func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}
