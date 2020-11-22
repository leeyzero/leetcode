package algorithm

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 该题在leetcode中为简单，但要考虑的边界情况比较多，其实时比较难的
// 解题思路：二分法
// 注意：边界条件处理
func minArray(numbers []int) int {
	if len(numbers) <= 0 {
		return 0
	}

	low, high := 0, len(numbers)-1
	for low < high {
		pivot := low + (high - low) >> 1
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

