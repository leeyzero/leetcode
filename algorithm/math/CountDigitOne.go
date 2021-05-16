package math

// https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
// 题目：剑指 Offer 43. 1～n 整数中 1 出现的次数
// 难度：困难
// 描述：
// 输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
// 例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
func countDigitOne(n int) int {
	digit, ans := 1, 0
	curr, high, low := n%10, n/10, 0
	for high != 0 || curr != 0 {
		if curr == 0 {
			ans += high * digit
		} else if curr == 1 {
			ans += high*digit + low + 1
		} else {
			ans += (high + 1) * digit
		}

		low += curr * digit
		curr = high % 10
		high /= 10
		digit *= 10
	}
	return ans
}
