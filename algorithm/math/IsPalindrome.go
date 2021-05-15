package math

// https://leetcode-cn.com/problems/palindrome-number/
// 题目：9. 回文数
// 难度：简单
// 思路：对于负责，一定不是回文；对于正数，对数字进行逆置后判断跟原数是否相等
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversed, n := 0, x
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return x == reversed
}
