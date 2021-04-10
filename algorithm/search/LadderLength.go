package search

// https://leetcode-cn.com/problems/word-ladder/
// 题目：127. 单词接龙
// 难度：困难
// 描述：
// 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
// 序列中第一个单词是 beginWord 。
// 序列中最后一个单词是 endWord 。
// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典 wordList 中的单词。
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。
// 思路：寻找最短路径问题一般使用BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := make([]bool, len(wordList))
	q := []string{beginWord}
	step := 1
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]
			for i, word := range wordList {
				if visited[i] {
					continue
				}
				if !isNeighbor(curr, word) {
					continue
				}
				if word == endWord {
					return step + 1
				}

				visited[i] = true
				q = append(q, word)
			}
		}
		step++
	}
	return 0
}
