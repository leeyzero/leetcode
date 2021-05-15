package stack

// https://leetcode-cn.com/problems/valid-parentheses/
// 题目：20. 有效的括号
// 难度：简单
// 描述：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 思路：栈
func isValid(s string) bool {
	hash := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if v, ok := hash[s[i]]; ok {
			stack = append(stack, v)
		} else if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
