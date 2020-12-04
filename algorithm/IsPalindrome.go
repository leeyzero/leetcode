package algorithm

// https://leetcode-cn.com/problems/palindrome-number/
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
