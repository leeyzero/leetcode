package sort

// https://leetcode-cn.com/problems/sort-colors/
// 题目：75. 颜色分类（经典的「荷兰国旗问题」，由计算机科学家 Edsger W. Dijkstra 首先提出）
// 难度：中等
// 要求：原地进行排序
// 思路：双指针
func sortColors(nums []int) {
	head := -1
	// 第一次遍历将0放到头部
	for i := head + 1; i < len(nums); i++ {
		if nums[i] == 0 {
			head++
			nums[head], nums[i] = nums[i], nums[head]
		}
	}

	// 第二次遍历将1放到头部
	for i := head + 1; i < len(nums); i++ {
		if nums[i] == 1 {
			head++
			nums[head], nums[i] = nums[i], nums[head]
		}
	}
}

// [2, 0, 2, 1, 1, 0]
// 同上述方法，但只需要一次遍历，具体做法是分别用p0和p1指针0和1的头部，遍历时，只要nums[i] <= 1时，
// 先将p1指针前移，交换nums[i]和nums[p1];如果nums[i]为0，再将p0前移，交换nums[p0]和nums[p1]
func sortColors2(nums []int) {
	p0, p1 := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 1 {
			continue
		}

		p1++
		nums[p1], nums[i] = nums[i], nums[p1]
		if p0 < p1 && nums[p1] == 0 {
			p0++
			nums[p0], nums[p1] = nums[p1], nums[p0]
		}
	}
}
