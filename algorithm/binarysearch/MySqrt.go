package binarysearch

// https://leetcode-cn.com/problems/sqrtx/
// 69. x 的平方根
// 简单
// 思路：由于x平方根的整数部分ans是满足 k^2 <= x的最大k值，因此我们可以对k进行二分查找。
func mySqrt(x int) int {
	lo, hi, ans := 0, x, -1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if mid*mid <= x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}
