package search

import (
	"bytes"
	"math"
)

// https://leetcode-cn.com/problems/n-queens/
// 题目：51. N 皇后
// 难度：困难
// 描述：
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 思路：回溯法
func solveNQueens(n int) [][]string {
	var ans [][]string
	var queens []int // queues[i] = j 表示(i, j)放置皇后
	backtracking(n, queens, &ans)
	return ans
}

func backtracking(n int, queens []int, ans *[][]string) {
	if len(queens) >= n {
		*ans = append(*ans, dumpBoard(queens))
		return
	}

	for col := 0; col < n; col++ {
		if !conflict(queens, col) {
			queens = append(queens, col)
			backtracking(n, queens, ans)
			queens = queens[:len(queens)-1]
		}
	}
}

// 检查第row行是否跟已经放置的皇后冲突
func conflict(queens []int, col int) bool {
	row := len(queens)
	for i := 0; i < row; i++ {
		diffX := int(math.Abs(float64(queens[i] - col)))
		if diffX == 0 || diffX == (row-i) {
			// 1. diffX == 0 表示第col列表已经放置了皇后,
			// 2. diffX == (row - i) 表示在同一个对角上，如[0, 3]和[3, 0], 可以理解为两点的斜率为1
			return true
		}
	}
	return false
}

func dumpBoard(queens []int) []string {
	var board []string
	for i, n := 0, len(queens); i < n; i++ {
		curr := bytes.Repeat([]byte{'.'}, n)
		curr[queens[i]] = 'Q'
		board = append(board, string(curr))
	}
	return board
}

// 另一种更直观的解法
func solveNQueens2(n int) [][]string {
	// 初始化棋盘
	var board [][]byte
	for i := 0; i < n; i++ {
		board = append(board, bytes.Repeat([]byte{'.'}, n))
	}

	var ans [][]string
	backtrackingSloveNQueens(board, 0, &ans)
	return ans
}

func backtrackingSloveNQueens(board [][]byte, row int, ans *[][]string) {
	if row >= len(board) {
		var curr []string
		for _, row := range board {
			curr = append(curr, string(row))
		}
		*ans = append(*ans, curr)
		return
	}

	// 对当前行row，有col=n种选择
	for col := 0; col < len(board); col++ {
		// 有冲突，跳过
		if isConflict(board, row, col) {
			continue
		}

		// 放置皇后
		board[row][col] = 'Q'
		// 递归处理下一行
		backtrackingSloveNQueens(board, row+1, ans)
		// 恢复
		board[row][col] = '.'
	}
}

func isConflict(board [][]byte, row int, col int) bool {
	// 检查row之前的
	for i := 0; i < row; i++ {
		// col列是否已经放置了皇后
		if board[i][col] == 'Q' {
			return true
		}

		// col的左对角是否放置了皇后
		j := col - (row - i)
		if (j >= 0) && (board[i][j] == 'Q') {
			return true
		}

		// col的右对角是否放置了皇后
		j = col + (row - i)
		if (j < len(board)) && (board[i][j] == 'Q') {
			return true
		}
	}
	return false
}
