package algorithm

import (
	"fmt"
)

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
// 解题思路：依次计算每个位数的数字占的位数
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
	num := start + (n-1) / digit // start为第0个数字

	// 确定所求的位数在num的哪一位
	s := fmt.Sprintf("%d", num)
	return int(s[(n-1) % digit] - '0')
}

