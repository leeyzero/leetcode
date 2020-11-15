package FindLadders

// https://leetcode-cn.com/problems/word-transformer-lcci/
// 解题思路：用一个栈记录结果，从beginWord出发，查找wordList中满足条件的单词，找到标记已访问，然后以该单词继续往下找，
// 如果未找到，退回到上个单词，重复上述策略，直到找到或栈为空
func findLadders(beginWord string, endWord string, wordList []string) []string {
	visited := make([]bool, len(wordList))
	for i, _ := range visited {
		visited[i] = false
	}

	ans := []string{beginWord}
	for len(ans) > 0 {
		curr := ans[len(ans)-1]
		if curr == endWord {
			break
		}

		if next := visitNextWord(curr, wordList, visited); next != "" {
			ans = append(ans, next)
		} else {
			// 当前单词无法到达结尾，回退到上一个单词
			ans = ans[:len(ans)-1]
		}
	}
	return ans
}

func visitNextWord(curr string, wordList []string, visited []bool) string {
	var next string
	for i, word := range wordList {
		if visited[i] {
			continue
		}
		if canTranslate(curr, word) {
			next = word
			visited[i] = true
			break
		}
	}
	return next
}

func canTranslate(curr string, next string) bool {
	if len(curr) != len(next) {
		return false
	}

	var diffcnt int
	for i := 0; i < len(curr) && diffcnt <= 1; i++ {
		if curr[i] != next[i] {
			diffcnt++
		}
	}
	return diffcnt == 1
}