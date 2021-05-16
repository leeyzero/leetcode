package other

// https://leetcode-cn.com/problems/qiu-12n-lcof/
// 题目：剑指 Offer 64. 求1+2+…+n
// 难度：中等
// 描述：
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
func sumNums(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNums(n-1)
}
