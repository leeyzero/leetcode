package search

// https://leetcode-cn.com/problems/word-ladder-ii/
// 题目：126. 单词接龙 II
// 难度：困难
// 思路：BFS+DFS
func findLaddersII(beginWord string, endWord string, wordList []string) [][]string {
	ans := [][]string{}
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}
	if _, ok := dict[endWord]; !ok {
		return ans
	}

	delete(dict, beginWord)
	delete(dict, endWord)
	q1 := map[string]bool{beginWord: true}
	q2 := map[string]bool{endWord: true}
	next := map[string][]string{}
	reversed, find := false, false
	for len(q1) > 0 {
		q := map[string]bool{}
		for w, _ := range q1 {
			s := []byte(w)
			for i := 0; i < len(s); i++ {
				ch := s[i]
				for j := 0; j < 26; j++ {
					s[i] = 'a' + byte(j)
					if v, ok := q2[string(s)]; v && ok {
						if reversed {
							next[string(s)] = append(next[string(s)], w)
						} else {
							next[w] = append(next[w], string(s))
						}
						find = true
					}
					if v, ok := dict[string(s)]; v && ok {
						if reversed {
							next[string(s)] = append(next[string(s)], w)
						} else {
							next[w] = append(next[w], string(s))
						}
						q[string(s)] = true
					}
				}
				s[i] = ch
			}
		}
		if find {
			break
		}

		for w := range q {
			delete(dict, w)
		}
		if len(q) <= len(q2) {
			q1 = q
		} else {
			reversed = !reversed
			q1 = q2
			q2 = q
		}
	}

	// fmt.Println(next)
	if find {
		backtrackingfindLaddersII(beginWord, endWord, next, []string{beginWord}, &ans)
	}
	return ans
}

func backtrackingfindLaddersII(src string, dst string, next map[string][]string, path []string, ans *[][]string) {
	if src == dst {
		curr := make([]string, len(path))
		copy(curr, path)
		*ans = append(*ans, curr)
		return
	}

	neighbors := next[src]
	for _, nei := range neighbors {
		path = append(path, nei)
		backtrackingfindLaddersII(nei, dst, next, path, ans)
		path = path[:len(path)-1]
	}
}
