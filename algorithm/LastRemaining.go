package algorithm

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// 解题思路：约琴夫环问题
// f(n, m) = (f(n-1, m) + m) % n
// f(1, m) = 0
// 简单的推理可以表述为：从0开始数m个数，让位置在m-1数出局后只剩下n-1个人，然后计算数n-1个人中选出的答案后
// 再加上一个相对位移k，最后对n取余即可得到答案
func lastRemaining(n int, m int) int {
	if n <= 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}

// 根据以上递推公式，可以得到迭代式子
func lastRemaining2(n int, m int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}
