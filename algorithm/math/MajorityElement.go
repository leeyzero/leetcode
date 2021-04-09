package math

import (
	"sort"
)

// https://leetcode-cn.com/problems/majority-element/
// 169. 多数元素
// 难度：简单
//
// 解题思路：
// 方法一：哈希表统计法，遍历数组，用hash表统计各数字的次数，即可找出众数。时间复杂度O(N), 空间复杂度O(N)
// 方法二：数组排序，数组中点的元素一定时众数。时间复杂度O(NlogN), 空间复杂度O(1)
// 方法三：摩尔投票法，核心理念为票数正负抵消。时间复杂度O(N), 空间复杂度O(1)
// 方法四：基于partition思想。时间复杂度O(N), 空间复杂度O(1)

// 摩尔投票法：
// 参考：https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
// 设输入数组 nums 的众数为 x ，数组长度为 n 。
// 推论一： 若记 众数 的票数为 +1 ，非众数 的票数为 -1 ，则一定有所有数字的 票数和 >0 。
// 推论二： 若数组的前 a 个数字的 票数和 =0 ，则 数组剩余 (n-a)个数字的 票数和一定仍 >0 ，即后 (n-a)个数字的 众数仍为 x 。
// 根据以上推论，假设数组首个元素 n1 为众数，遍历并统计票数。当发生 票数和 =0 时，剩余数组的众数一定不变 ，这是由于：
// 当 n1 = x: 抵消的所有数字，有一半是众数x 。
// 当 n1 != x: 抵消的所有数字，少于或等于一半是众数x
// 利用此特性，每轮假设发生票数和=0 都可以缩小剩余数组区间。当遍历完成时，最后一轮假设的数字即为众数。
func majorityElement(nums []int) int {
	votes := 0
	ans := 0
	for _, num := range nums {
		if votes == 0 {
			ans = num
			votes = 1
		} else if num == ans {
			votes++
		} else {
			votes--
		}
	}
	return ans
}

func majorityElementByHash(nums []int) int {
	ans := 0
	hash := map[int]int{}
	for _, num := range nums {
		if cnt, ok := hash[num]; ok {
			hash[num] = cnt + 1
		} else {
			hash[num] = 1
		}

		if cnt, _ := hash[num]; cnt > len(nums)/2 {
			ans = num
			break
		}
	}
	return ans
}

func majorityElementBySort(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})
	return nums[len(nums)/2]
}
