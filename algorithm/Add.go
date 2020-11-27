package algorithm

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
// 解题思路：
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