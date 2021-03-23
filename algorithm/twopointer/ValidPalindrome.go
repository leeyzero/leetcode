package twopointer

// https://leetcode-cn.com/problems/valid-palindrome-ii/
// 题目：680. 验证回文字符串 Ⅱ
// 难度：简单
// 思路：双指针
func ValidPalindrome(s string) bool {
	return validPalindromeRecursive(s, 1)
}

func validPalindromeRecursive(s string, delcnt int) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}

		if delcnt > 1 {
			return false
		}
		return validPalindromeRecursive(s[left+1:right+1], delcnt+1) ||
			validPalindromeRecursive(s[left:right], delcnt+1)
	}
	return true
}
