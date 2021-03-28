package binarysearch

// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// 题目：剑指 Offer 53 - I. 在排序数组中查找数字 I
// 描述：统计一个数字在排序数组中出现的次数。
// 难度：简单
// 思路：排序数组中的搜索问题，首先想到二分法
// 本题要求统计数字target中的次数，可以转化为使用二分法搜索target的左右边界left,right
// 则target的数量为right-left+1
func search(nums []int, target int) int {
	rbound := rightBound(nums, target)
	if rbound == -1 {
		return 0
	}
	lbound := leftBound(nums, target)
	return rbound - lbound + 1
}

// target的右边界的索引，不存在返回-1
func rightBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target >= arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && arr[right] == target {
		return right
	}
	return -1
}

// target的左边界的索引，不存在返回-1
func leftBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target <= arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(arr) && arr[left] == target {
		return left
	}
	return -1
}
