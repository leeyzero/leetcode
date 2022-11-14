package slidingwindow

// https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
// 题目：剑指 Offer 59 - I. 滑动窗口的最大值
// 难度：困难
// 描述：给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
// 思路：
// 方法一： 最直观的解法是，用两个指针维护窗口大小，每次计算窗口中的最大值，时间复杂度为：O(kn)
// 方法二： 能不能从获取窗口中最大值的时间复杂度从O(k)降低到O(1)呢？
// 一个滑动窗口可以看成是一个队列，当窗口滑动时，处于窗口的第一个数字被删除，同时在窗口的末尾添加一个新的
// 数字。这看起来像是队列的『先进先出』特性，如果能从这个『队列』中找出最大值，就可以解决这个问题了。
func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	if k <= 0 || k > len(nums) {
		return ans
	}

	var queue []int
	for left, right := 0, 0; right < len(nums); {
		// 维护单调递减队列，让队头将始终保持最大值
		// 注意：队列中存储的是数组的下标
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[right] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)

		// 窗口还未完全打开，继续扩大窗口，窗口大小为：[left, right + 1）
		if (right - left + 1) < k {
			right++
			continue
		}

		// 队头指向的元素即为窗口的最大值
		ans = append(ans, nums[queue[0]])

		// 同时增大窗口指针，滑动窗口
		left, right = left+1, right+1

		// 需要移除队头元素已经不在窗口中的下标
		if len(queue) > 0 && queue[0] < left {
			queue = queue[1:]
		}
	}
	return ans
}
