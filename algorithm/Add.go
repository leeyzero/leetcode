package algorithm

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
// 解题思路：
// 非进位求和：n = a ^ b
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