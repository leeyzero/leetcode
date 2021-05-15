package math

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
// 题目：剑指 Offer 15. 二进制中1的个数
// 描述：请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
// 思路：将num跟1进行&操作，如果为1，标示改为时1，然后将num右移1为，继续处理，直到num为0
func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		if (num & 1) == 1 {
			ans += 1
		}
		num >>= 1
	}
	return ans
}
