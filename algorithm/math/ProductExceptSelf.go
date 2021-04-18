package math

// https://leetcode-cn.com/problems/product-of-array-except-self/
// 题目：238. 除自身以外数组的乘积
// 难度：中等
// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 思路：把结束想象成一个nxn的矩阵，对角线的值为1，分别计算矩阵的下三角和上三角，并在计算过程中存储过程值
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	if len(nums) <= 0 {
		return ans
	}
	ans[0] = 1

	// 下三角
	p := 1
	for i := 0; i < len(nums)-1; i++ {
		p *= nums[i]
		ans[i+1] = p
	}

	// 上三角
	q := 1
	for i := len(nums) - 1; i > 0; i-- {
		q *= nums[i]
		ans[i-1] *= q
	}

	return ans
}
