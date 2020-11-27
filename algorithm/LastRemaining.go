package algorithm

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// 解题思路：约琴夫环问题
// f(n, m) = (f(n-1, m) + m) % n
// f(1, m) = 0
func lastRemaining(n int, m int) int {
	if n <= 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}
