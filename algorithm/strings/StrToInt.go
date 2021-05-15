package strings

import (
	"math"
	"strings"
)

// https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
// 题目：剑指 Offer 67. 把字符串转换成整数
// 难度：中等
// 描述：
// 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
// 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
// 当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
// 该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
// 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
// 在任何情况下，若函数不能进行有效的转换时，请返回 0。
// 思路：关键时对细节的处理，如正负号、溢出、非法字符
func strToInt(str string) int {
	str = strings.Trim(str, " ")
	if str == "" {
		return 0
	}

	i := 0
	minus := false
	if str[i] == '-' {
		minus = true
	}
	if str[i] == '-' || str[i] == '+' {
		i++
	}

	var ans int64
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}

		ans = ans*int64(10) + int64(str[i]-'0')
		if minus && -ans < math.MinInt32 {
			return math.MinInt32
		} else if !minus && ans > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if minus {
		ans = -ans
	}
	return int(ans)
}
