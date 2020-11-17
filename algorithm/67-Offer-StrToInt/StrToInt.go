package StrToInt

import (
	"math"
	"strings"
)

// https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
// 解题思路：关键时对细节的处理，如正负号、溢出、非法字符
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

        ans = ans * int64(10) + int64(str[i] - '0')
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
