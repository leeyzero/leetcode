package search

// https://leetcode-cn.com/problems/word-search/
// 题目：79. 单词搜索
// 难度：中等
// 描述：给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用。
// 思路：回溯
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtrackingExist(board, i, j, word, 0, visited) {
				return true
			}
		}
	}
	return false
}

func backtrackingExist(board [][]byte, i, j int, word string, pos int, visited [][]bool) bool {
	if pos >= len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if visited[i][j] || board[i][j] != word[pos] {
		return false
	}

	visited[i][j] = true
	ans := backtrackingExist(board, i, j-1, word, pos+1, visited) ||
		backtrackingExist(board, i-1, j, word, pos+1, visited) ||
		backtrackingExist(board, i, j+1, word, pos+1, visited) ||
		backtrackingExist(board, i+1, j, word, pos+1, visited)
	visited[i][j] = false
	return ans
}
