package dfs

// https://leetcode.cn/problems/surrounded-regions/
// 题目：130. 被围绕的区域
// 难度：中等
// 描述：给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
// 示例1：
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//
func solve(board [][]byte) {
	if len(board) <= 0 || len(board[0]) <= 0 {
		return
	}

	row, col := len(board), len(board[0])
	// 上边和下边
	for j := 0; j < col; j++ {
		dfs(board, 0, j)
		dfs(board, row-1, j)
	}
	// 左边和右边
	for i := 0; i < row; i++ {
		dfs(board, i, 0)
		dfs(board, i, col-1)
	}

	// 棋盘中剩下的所以'O'变成'X', 并恢复'#'为'O'
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// 使用dfs将边角的'O'替换成一个特殊字符，如'#'
func dfs(board [][]byte, i int, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] != 'O' {
		return
	}

	board[i][j] = '#'
	// 上
	dfs(board, i-1, j)
	// 右
	dfs(board, i, j+1)
	// 下
	dfs(board, i+1, j)
	// 左
	dfs(board, i, j-1)
}
