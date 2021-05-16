package unionfind

// https://leetcode-cn.com/problems/redundant-connection/
// 题目：684. 冗余连接
// 难度：中等
// 描述：
// 在本问题中, 树指的是一个连通且无环的无向图。
// 输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。
// 结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。
// 返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 [u, v] 应满足相同的格式 u < v。
// 思路：并查集
func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges) + 1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if uf.Connected(u, v) {
			return []int{u, v}
		}
		uf.Union(u, v)
	}
	return []int{-1, -1}
}
