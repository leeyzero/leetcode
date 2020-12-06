package algorithm

// https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
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

    result := powerWithUnsingedExponent(base, exponent >> 1)
    result *= result
    if exponent & 0x1 == 1 {
        result *= base
    }
    return result
}