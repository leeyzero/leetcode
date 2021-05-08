package search

// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
// 题目：剑指 Offer 12. 矩阵中的路径
// 难度：中等
// 思路：回溯、DFS
func exist(board [][]byte, word string) bool {
	if len(board) <= 0 || len(board[0]) <= 0 || len(word) <= 0 {
		return false
	}

	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if hasPath(board, i, j, word, 0, visited) {
				return true
			}
		}
	}
	return false
}

func hasPath(board [][]byte, i, j int, word string, pos int, visited [][]bool) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		pos >= len(word) || visited[i][j] || board[i][j] != word[pos] {
		return false
	}
	if pos == len(word)-1 {
		return true
	}

	visited[i][j] = true
	ans := hasPath(board, i-1, j, word, pos+1, visited) ||
		hasPath(board, i, j+1, word, pos+1, visited) ||
		hasPath(board, i+1, j, word, pos+1, visited) ||
		hasPath(board, i, j-1, word, pos+1, visited)
	visited[i][j] = false
	return ans
}
