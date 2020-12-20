package algorithm

// https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/
func countDigitOne(n int) int {
	digit, ans := 1, 0
	curr, high, low := n % 10, n / 10, 0
	for high != 0 || curr != 0 {
		if curr == 0 {
			ans += high * digit
		} else if curr == 1 {
			ans += high * digit + low + 1
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