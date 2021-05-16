package math

// https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/
// 题目：剑指 Offer 14- II. 剪绳子 II
// 难度：中等
// 思路：
// 切分规则：
// 最优： 3 。把绳子尽可能切为多个长度为 3 的片段，留下的最后一段绳子的长度可能为 0,1,2 三种情况。
// 次优： 2 。若最后一段绳子长度为 2 ；则保留，不再拆为 1+1。
// 最差： 1 。若最后一段绳子长度为 1 ；则应把一份 3 + 1 替换为 2 + 2，因为 2 x 2 > 3 x 1
func cuttingRopeII(n int) int {
	if n <= 3 {
		return n - 1
	}

	var rem, x, p int64 = 1, 3, 1000000007
	for a := n/3 - 1; a > 0; a /= 2 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = (x * x) % p
	}

	ans := []int{int(rem * 3 % p), int(rem * 4 % p), int(rem * 6 % p)}
	return ans[n%3]
}
