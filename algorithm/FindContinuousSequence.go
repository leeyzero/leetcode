package algorithm

// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
// 解题思路：
// 1. 最直观的做法时穷取法，从i = [1,target/2]，依次判断已i为起始的连续序列
// 2. 活动窗口
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