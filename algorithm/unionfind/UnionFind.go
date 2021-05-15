package unionfind

// 并查集数据结构
type UnionFind struct {
	// 记录连通分量
	count int

	// 节点i的父节点是parent[i]
	parent []int

	// 记录每棵树的重量
	size []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		count:  n, // 初始化互不连通
		parent: make([]int, n),
		size:   make([]int, n),
	}

	// 父节点指针初始指向自己
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

// 将p和q结点连接
func (t *UnionFind) Union(p, q int) {
	rootP := t.find(p)
	rootQ := t.find(q)
	if rootP == rootQ {
		return
	}

	// 将小树连接到大树下面，较平衡
	if t.size[rootP] > t.size[rootQ] {
		t.parent[rootQ] = rootP
		t.size[rootP] += t.size[rootQ]
	} else {
		t.parent[rootP] = rootQ
		t.size[rootQ] += t.size[rootP]
	}
	t.count--
}

// 判断p和q结点是否连通
func (t *UnionFind) Connected(p, q int) bool {
	rootP := t.find(p)
	rootQ := t.find(q)
	return rootP == rootQ
}

// 返回图中有多少个连通分量
func (t *UnionFind) Count() int {
	return t.count
}

// 返回节点x的根节点
func (t *UnionFind) find(x int) int {
	for t.parent[x] != x {
		// 进行路径压缩
		t.parent[x] = t.parent[t.parent[x]]
		x = t.parent[x]
	}
	return x
}
