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
	map[byte]int{ // stateBegin
		's': stateSignBeforeE,
		'd': stateDigitBeforeDot,
		'.': stateDigitAfterDotWithEmptyInteger,
	},
	map[byte]int{ // stateSignBeforeE
		'd': stateDigitBeforeDot,
		'.': stateDigitAfterDotWithEmptyInteger,
	},
	map[byte]int{ // stateDigitBeforeDot
		'd': stateDigitBeforeDot,
		'.': stateDigitAfterDot,
		'e': stateE,
	},
	map[byte]int{ // stateDigitAfterDot
		'd': stateDigitAfterDot,
		'e': stateE,
	},
	map[byte]int{ // stateDigitAfterDotWithEmptyInteger
		'd': stateDigitAfterDot,
	},
	map[byte]int{ // stateE
		's': stateSignAfterE,
		'd': stateDigitExponent,
	},
	map[byte]int{ // stateSignAfterE
		'd': stateDigitExponent,
	},
	map[byte]int{ // stateDigitExponent
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
