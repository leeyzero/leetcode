package search

// https://leetcode-cn.com/problems/shortest-bridge/
// 题目：934. 最短的桥
// 难度：中等
// 描述：在给定的二维二进制数组 A 中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。）
// 现在，我们可以将 0 变为 1，以使两座岛连接起来，变成一座岛。
// 返回必须翻转的 0 的最小数目。（可以保证答案至少是 1 。）
// 思路：DFS+BFS
// 在寻找这两座岛时，我们使用深度优先搜索。在向外延伸时，我们使用广度优先搜索。
type Pair [2]int

func shortestBridge(A [][]int) int {
	if len(A) <= 0 {
		return 0
	}

	r, c := len(A), len(A[0])
	colors := colorize(A)
	q := []Pair{}
	target := map[Pair]bool{}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if colors[i][j] == 1 {
				q = append(q, Pair{i, j})
			} else if colors[i][j] == 2 {
				target[Pair{i, j}] = true
			}
		}
	}

	directs := [4]Pair{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	distance := 0
	for len(q) > 0 {
		var nextQ []Pair
		for i := 0; i < len(q); i++ {
			if _, ok := target[q[i]]; ok {
				return distance - 1
			}
			for _, d := range directs {
				x, y := q[i][0]+d[0], q[i][1]+d[1]
				if x < 0 || x >= r || y < 0 || y >= c {
					continue
				}
				if colors[x][y] != 1 {
					nextQ = append(nextQ, Pair{x, y})
					colors[x][y] = 1
				}
			}
		}
		q = nextQ
		distance++
	}
	return -1
}

// 对值为1的岛进行着色
func colorize(A [][]int) [][]int {
	if len(A) <= 0 {
		return [][]int{}
	}

	row, col := len(A), len(A[0])
	colors := make([][]int, row)
	for i := 0; i < row; i++ {
		colors[i] = make([]int, col)
	}

	color := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if colors[i][j] == 0 && A[i][j] == 1 {
				// 开始dfs着色
				color++
				dfsColorize(A, i, j, colors, color)
			}
		}
	}
	return colors
}

func dfsColorize(A [][]int, i, j int, colors [][]int, color int) {
	if i < 0 || i >= len(A) || j < 0 || j >= len(A[0]) {
		return
	}
	if colors[i][j] != 0 || A[i][j] != 1 {
		return
	}

	colors[i][j] = color
	dfsColorize(A, i, j-1, colors, color)
	dfsColorize(A, i-1, j, colors, color)
	dfsColorize(A, i, j+1, colors, color)
	dfsColorize(A, i+1, j, colors, color)
}
