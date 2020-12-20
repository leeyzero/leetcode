package algorithm

// https://leetcode-cn.com/problems/valid-parentheses/
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