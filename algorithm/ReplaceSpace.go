package algorithm

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
// 解题思路：先计算出字符中的空格树，然后开辟len(s) + 2*空格数的空间，最后遍历s进行替换
func replaceSpace(s string) string {
	cnt := 0
	for _, c := range s {
		if c == ' ' {
			cnt++
		}
	}

	// 前提一个字符可以用byte表示
	data := make([]byte, len(s) + 2*cnt)
	i := 0
	for _, c := range s {
		if c == ' ' {
			data[i] = '%'
			data[i+1] = '2'
			data[i+2] = '0'
			i += 3
		} else {
			data[i] = byte(c)
			i += 1
		}
	}
	return string(data)
}