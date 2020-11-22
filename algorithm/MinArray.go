package algorithm

import (
	"math"
)

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 该题在leetcode中为简单，但要考虑的边界情况比较多，其实时比较难的
// 解题思路：二分法
// 注意：边界条件处理
func minArray(numbers []int) int {
	if len(numbers) <= 0 {
		return 0
	}

	left, right := 0, len(numbers)-1
	mid := left
	for numbers[left] >= numbers[right] {
		// 退出条件
		if right - left <= 1 {
			mid = right
			break
		}

		mid = (left + right) >> 1

		// 如果left,right,mid下边对应的值均相等，只能顺序查找最小值了
		if numbers[left] == numbers[right] && numbers[left] == numbers[mid] {
			return findMinInOrder(numbers[left:right+1])
		}

		if numbers[mid] >= numbers[left] {
			// mid在左递增区间
			left = mid
		} else if numbers[mid] <= numbers[right] {
			// mid在右递增区间
			right = mid
		}
	}
	return numbers[mid]
}

func findMinInOrder(numbers []int) int {
	ans := math.MaxInt32
	for _, num := range numbers {
		if num < ans {
			ans = num
		}
	}
	return ans
}