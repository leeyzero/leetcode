package math

// https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
// 题目：剑指 Offer 16. 数值的整数次方
// 难度：中等
// 描述：实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
// 思路：递归
func myPow(x float64, n int) float64 {
	if x == float64(0) && n < 0 {
		return float64(0)
	}

	exponent := n
	if n < 0 {
		exponent = -n
	}
	result := powerWithUnsingedExponent(x, exponent)
	if n < 0 {
		result = float64(1) / result
	}
	return result
}

func powerWithUnsingedExponent(base float64, exponent int) float64 {
	if exponent == 0 {
		return float64(1)
	}
	if exponent == 1 {
		return base
	}

	result := powerWithUnsingedExponent(base, exponent>>1)
	result *= result
	if exponent&0x1 == 1 {
		result *= base
	}
	return result
}
