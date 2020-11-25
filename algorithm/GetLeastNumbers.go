package algorithm

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
// 解题思路：利用partition思想，找到k大元素
func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	if len(arr) <= 0 || len(arr) <= k {
		return arr
	}

	left, right, mid := 0, len(arr)-1, k-1
	idx := partition(arr, left, right)
	for idx != mid {
		if idx < mid {
			left = idx + 1
		} else {
			right = idx - 1
		}
		idx = partition(arr, left, right)
	}
	return arr[:k]
}

// [1] [1,2] [2,1] [1,2,3] [3,2,1]
// 快慢指针, 把第一个元素做为pivot
func partition(arr []int, start, end int) int {
	j := start
	for i := start + 1; i <= end; i++ {
		if arr[i] < arr[start] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[start], arr[j] = arr[j], arr[start]
	return j
}

// 最直观的办法，使用插入排序，数组超过k的大小，删除最大元素
func getLeastNumbers2(arr []int, k int) []int {
	ans := []int{}
	for _, num := range arr {
		ans = append(ans, num)
		j := len(ans) - 1
		for j > 0 && num < ans[j-1] {
			ans[j] = ans[j-1]
			j--
		}
		ans[j] = num
		if len(ans) > k {
			ans = ans[:k]
		}
	}
	return ans
}
