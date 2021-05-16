package math

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// 题目：剑指 Offer 62. 圆圈中最后剩下的数字
// 难度：简单
// 描述：
// 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
// 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
// 思路：约琴夫环问题
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
