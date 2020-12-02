package algorithm

import (
	"fmt"
)
// https://leetcode-cn.com/problems/reverse-integer/
func reverseNumber(x int) int {
	ans := 0
	for x != 0 {
		fmt.Println(x % 10)
		ans = ans * 10 + x % 10
		x /= 10
	}
	return ans
}