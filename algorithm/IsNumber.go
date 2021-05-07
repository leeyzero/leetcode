package algorithm

import (
	"strings"
)

// https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
// 解题思路：状态机 A[.[B]][e|EC]
const (
	stateBegin = iota
	stateSignBeforeE
	stateDigitBeforeDot
	stateDigitAfterDot
	stateDigitAfterDotWithEmptyInteger
	stateE
	stateSignAfterE
	stateDigitExponent
)

var states = []map[byte]int{
	{ // stateBegin
		's': stateSignBeforeE,
		'd': stateDigitBeforeDot,
		'.': stateDigitAfterDotWithEmptyInteger,
	},
	{ // stateSignBeforeE
		'd': stateDigitBeforeDot,
		'.': stateDigitAfterDotWithEmptyInteger,
	},
	{ // stateDigitBeforeDot
		'd': stateDigitBeforeDot,
		'.': stateDigitAfterDot,
		'e': stateE,
	},
	{ // stateDigitAfterDot
		'd': stateDigitAfterDot,
		'e': stateE,
	},
	{ // stateDigitAfterDotWithEmptyInteger
		'd': stateDigitAfterDot,
	},
	{ // stateE
		's': stateSignAfterE,
		'd': stateDigitExponent,
	},
	{ // stateSignAfterE
		'd': stateDigitExponent,
	},
	{ // stateDigitExponent
		'd': stateDigitExponent,
	},
}

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	p := stateBegin
	for i := 0; i < len(s); i++ {
		c := s[i]
		if s[i] == '-' || s[i] == '+' {
			c = 's'
		} else if s[i] == '.' {
			c = s[i]
		} else if s[i] == 'e' || s[i] == 'E' {
			c = 'e'
		} else if s[i] >= '0' && s[i] <= '9' {
			c = 'd'
		}
		if _, ok := states[p][c]; !ok {
			return false
		}
		p = states[p][c]
	}
	return p == stateDigitBeforeDot || p == stateDigitAfterDot || p == stateDigitExponent
}
