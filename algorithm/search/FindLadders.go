package search

// https://leetcode-cn.com/problems/word-transformer-lcci/
// 面试题 17.22. 单词转换
// 难度：困难
// 描述：
// 给定字典中的两个词，长度相等。写一个方法，把一个词转换成另一个词， 但是一次只能改变一个字符。每一步得到的新词都必须能在字典中找到。
// 编写一个程序，返回一个可能的转换序列。如有多个可能的转换序列，你可以返回任何一个。
// 思路：可以将每个单词当做一个节点，如果节点之间只有一个元素不同，则节点相连，可以使用DFS找到任意一条beginWord到endWord的路径
// 需要注意的是，对于已经访问过的结点需要进行剪枝，防止重复访问
// 方法一：使用栈进行迭代
func findLadders(beginWord string, endWord string, wordList []string) []string {
	visited := make([]bool, len(wordList))
	ans := []string{beginWord}
	for len(ans) > 0 {
		curr := ans[len(ans)-1]
		if curr == endWord {
			break
		}

		if next := visitNextWord(curr, wordList, visited); next >= 0 {
			ans = append(ans, wordList[next])
			visited[next] = true
		} else {
			// 当前单词无法到达结尾，回退到上一个单词
			ans = ans[:len(ans)-1]
		}
	}
	return ans
}

func visitNextWord(curr string, wordList []string, visited []bool) int {
	ans := -1
	for i, word := range wordList {
		if visited[i] {
			continue
		}
		if isNeighbor(curr, word) {
			return i
		}
	}
	return ans
}

// 方法二：回溯
func findLadders2(beginWord string, endWord string, wordList []string) []string {
	visited := map[string]bool{}
	var ans []string
	backtrackingfindLadders(beginWord, endWord, wordList, visited, &ans)
	return ans
}

func backtrackingfindLadders(curr string, endWord string, wordList []string, visited map[string]bool, ans *[]string) bool {
	visited[curr] = true
	*ans = append(*ans, curr)
	if curr == endWord {
		return true
	}

	isFind := false
	neighbors := findUnvisitedNeighbors(curr, wordList, visited)
	for _, nei := range neighbors {
		if backtrackingfindLadders(nei, endWord, wordList, visited, ans) {
			isFind = true
			break
		}
	}
	if !isFind {
		// 恢复状态
		*ans = (*ans)[:len(*ans)-1]
	}
	return isFind
}

func isNeighbor(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var diffcnt int
	for i := 0; i < len(word1) && diffcnt <= 1; i++ {
		if word1[i] != word2[i] {
			diffcnt++
		}
	}
	return diffcnt == 1
}

func findUnvisitedNeighbors(curr string, wordList []string, visited map[string]bool) []string {
	var neighbors []string
	for _, word := range wordList {
		if _, ok := visited[word]; ok {
			continue
		}

		if isNeighbor(curr, word) {
			neighbors = append(neighbors, word)
		}
	}
	return neighbors
}
