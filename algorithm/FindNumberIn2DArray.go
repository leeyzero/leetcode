package algorithm

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
// 解题思路：选取数组右上角的数字，如果该数字等于要查找的数字，则查找结束；
// 如果该数字大于要查找的数字，则剔除这个数字所在列；
// 如果该数字小于要查找的数字，则剔除这个数字所在行；
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return false
	}

	x, y := 0, len(matrix[0])-1
	for x < len(matrix) && y >= 0 {
		// fmt.Println(x, y)
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

// 另一种方法：每次一左上角数字，在该数字行和列中通过二分法查找target，如果找到则查找结束；
// 否则找到该数字所在行和列中小于该数字的坐标，左上角数字向前和下移动移位，继续使用上述策略查找；
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return false
	}

	x0, y0, x1, y1 := 0, 0, len(matrix)-1, len(matrix[0])-1
	return findNumberIn2DArrayCore(matrix, x0, y0, x1, y1, target)
}

func findNumberIn2DArrayCore(matrix [][]int, x0, y0 int, x1, y1 int, target int) bool {
	for x0 <= x1 && y0 <= y1 {
		y1 = findNubmerIn2DArrayRow(matrix, x0, y0, y1, target)
		if matrix[x0][y1] == target {
			return true
		}

		x1 = findNumberIn2DArrayCol(matrix, x0, x1, y0, target)
		if matrix[x1][y0] == target {
			return true
		}
		x0, y0 = x0+1, y0+1
	}
	return false
}

// 找到x0行中<=target的位置
func findNubmerIn2DArrayRow(matrix [][]int, x0 int, y0 int, y1 int, target int) int {
	left, right := y0, y1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[x0][mid] == target {
			right = mid
			break
		} else if matrix[x0][mid] < target {
			left = mid + 1
		} else if matrix[x0][mid] > target {
			right = mid - 1
		}
	}
	if right < y0 {
		return y0
	}
	return right
}

// 找到y0列中<=target的位置
func findNumberIn2DArrayCol(matrix [][]int, x0 int, x1 int, y0 int, target int) int {
	top, down := x0, x1
	for top <= down {
		mid := top + (down-top)/2
		if matrix[mid][y0] == target {
			down = mid
			break
		} else if matrix[mid][y0] < target {
			top = mid + 1
		} else if matrix[mid][y0] > target {
			down = mid - 1
		}
	}
	if down < x0 {
		return x0
	}
	return down
}
