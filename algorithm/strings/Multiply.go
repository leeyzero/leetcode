package strings

import (
	"bytes"
)

// https://leetcode-cn.com/problems/multiply-strings/
// 题目：43. 字符串相乘
// 难度：中等
// 描述：
// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 思路：按普通加法思路做，注意进位细节问题
func multiply(num1 string, num2 string) string {
	ans := "0"
	for i, k := len(num2)-1, 0; i >= 0; i, k = i-1, k+1 {
		curr := multiplyTenTimes(multiplyN(num1, int(num2[i]-'0')), k)
		ans = addStringDigit(ans, curr)
	}
	return ans
}

func multiplyN(num string, n int) string {
	if n == 0 || num == "0" {
		return "0"
	}

	var carry int
	var ans []byte
	for i := len(num) - 1; i >= 0; i-- {
		result := n*int(num[i]-'0') + carry
		carry = result / 10
		ans = append(ans, byte(result%10+'0'))
	}
	if carry > 0 {
		ans = append(ans, byte(carry+'0'))
	}

	return string(reverseDigits(ans))
}

func multiplyTenTimes(num string, n int) string {
	if num == "0" {
		return num
	}
	return num + string(bytes.Repeat([]byte{'0'}, n))
}

func addStringDigit(num1 string, num2 string) string {
	var ans []byte
	var carry int

	p1, p2 := len(num1)-1, len(num2)-1
	for p1 >= 0 || p2 >= 0 || carry > 0 {
		result := carry
		if p1 >= 0 && p2 >= 0 {
			result = int(num1[p1]-'0') + int(num2[p2]-'0') + carry
			p1, p2 = p1-1, p2-1
		} else if p1 >= 0 {
			result = int(num1[p1]-'0') + carry
			p1--
		} else if p2 >= 0 {
			result = int(num2[p2]-'0') + carry
			p2--
		}

		carry = result / 10
		ans = append(ans, byte(result%10+'0'))
	}

	return string(reverseDigits(ans))
}

func reverseDigits(digits []byte) []byte {
	for lo, hi := 0, len(digits)-1; lo < hi; lo, hi = lo+1, hi-1 {
		digits[lo], digits[hi] = digits[hi], digits[lo]
	}
	return digits
}
