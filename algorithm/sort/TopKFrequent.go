package sort

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/top-k-frequent-elements/
// 题目：347. 前 K 个高频元素
// 难度：中等
// 要求：算法时间复杂度必须优于O(N*logN)
// 思路：桶排序
// 基于堆排序和快速排序的思路请参考：https://leetcode-cn.com/problems/top-k-frequent-elements/solution/qian-k-ge-gao-pin-yuan-su-by-leetcode-solution/
func topKFrequent(nums []int, k int) []int {
	// 先用一个hash统计每个数字出现的频率
	hash := map[int]int{}
	maxFrequent := 0
	for _, n := range nums {
		hash[n]++
		maxFrequent = base.Max(maxFrequent, hash[n])
	}

	// 使用频次桶存储数字
	buckets := make([][]int, maxFrequent+1)
	for key, cnt := range hash {
		buckets[cnt] = append(buckets[cnt], key)
	}

	// 对桶从后往前遍历，找到k个高频元素
	ans := []int{}
	for i := len(buckets) - 1; i >= 0 && len(ans) < k; i-- {
		for _, n := range buckets[i] {
			ans = append(ans, n)
			if len(ans) >= k {
				break
			}
		}
	}
	return ans
}
