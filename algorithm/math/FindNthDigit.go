package math

import (
	"fmt"
)

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
// 题目：剑指 Offer 44. 数字序列中某一位的数字
// 难度：中等
// 描述：数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
// 请写一个函数，求任意第n位对应的数字。
// 思路：依次计算每个位数的数字占的位数
// 1、确定n所在数字的位数，记为 digit；
// 2、确定n所在的数字，记为 num；
// 3、确定n是num中的哪一数位，并返回结果；
func findNthDigit(n int) int {
	if n <= 0 {
		return 0
	}

	// 确定n所在的数字位数
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = 9 * start * digit
	}

	// 确定n所在的数字
	num := start + (n-1)/digit // start为第0个数字

	// 确定所求的位数在num的哪一位
	s := fmt.Sprintf("%d", num)
	return int(s[(n-1)%digit] - '0')
}
