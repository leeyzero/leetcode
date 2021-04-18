package array

// https://leetcode-cn.com/problems/spiral-matrix-ii/
// 题目：59. 螺旋矩阵 II
// 难度：中等
// 描述：给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
// 思路：双指针
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	num := 1
	for lo, hi := 0, n-1; lo <= hi; lo, hi = lo+1, hi-1 {
		if lo == hi {
			matrix[lo][lo] = num
			break
		}

		// left -> right
		for i, j := lo, lo; j < hi; j++ {
			matrix[i][j] = num
			num++
		}
		// top -> bottom
		for i, j := lo, hi; i < hi; i++ {
			matrix[i][j] = num
			num++
		}
		// right -> left
		for i, j := hi, hi; j > lo; j-- {
			matrix[i][j] = num
			num++
		}
		// bottom -> top
		for i, j := hi, lo; i > lo; i-- {
			matrix[i][j] = num
			num++
		}
	}
	return matrix
}
