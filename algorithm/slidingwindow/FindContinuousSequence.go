package slidingwindow

// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// 题目：剑指 Offer 57 - II. 和为s的连续正数序列
// 难度：简单
// 描述：输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
// 思路：
// 1. 最直观的做法时穷取法，从i = [1,target/2]，依次判断已i为起始的连续序列
// 2. 滑动窗口
func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	if target < 3 {
		return ans
	}

	small, big := 1, 2
	mid := target >> 1
	for small <= mid {
		sum := (small + big) * (big - small + 1) / 2
		if sum == target {
			ans = append(ans, makeSequence(small, big))
		}

		if sum <= target {
			// 扩大窗口
			big++
		} else {
			// 缩小窗口
			small++
		}
	}
	return ans
}

func makeSequence(small, big int) []int {
	ans := []int{}
	for i := small; i <= big; i++ {
		ans = append(ans, i)
	}
	return ans
}
