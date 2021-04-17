package algorithm

// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
// 题目：剑指 Offer 29. 顺时针打印矩阵 / 54. 螺旋矩阵
// 难度：简单
// 描述：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return []int{}
	}

	ans := []int{}
	row, col := len(matrix), len(matrix[0])
	il, ih := 0, row-1
	jl, jh := 0, col-1
	for il <= ih && jl <= jh {
		// single row
		if il == ih {
			for j := jl; j <= jh; j++ {
				ans = append(ans, matrix[il][j])
			}
			break
		}
		// single col
		if jl == jh {
			for i := il; i <= ih; i++ {
				ans = append(ans, matrix[i][jl])
			}
			break
		}

		// left->right
		for i, j := il, jl; j < jh; j++ {
			ans = append(ans, matrix[i][j])
		}
		// top->down
		for i, j := il, jh; i < ih; i++ {
			ans = append(ans, matrix[i][j])
		}
		// right->left
		for i, j := ih, jh; j > jl; j-- {
			ans = append(ans, matrix[i][j])
		}
		// down->top
		for i, j := ih, jl; i > il; i-- {
			ans = append(ans, matrix[i][j])
		}

		il, ih = il+1, ih-1
		jl, jh = jl+1, jh-1
	}

	return ans
}
