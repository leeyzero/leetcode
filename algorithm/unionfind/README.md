# 并查集

并查集（Union-Find）算法主要是解决图论中“动态连通性”问题的。其中连通具有如下性质：
- 自反性：节点p和p是连通的。
- 对称性：如果节点p和q连通，那么q和p也连通。
- 传递性：如果节点p和q连通，q和r连通，则p和r也连通。

Union-Find算法的API描述：

```
type UionFind interface {
    // 将p和q连接
    Union(p int, q int)

    // 判断p和q是否连通
    Connected(p int, q int) bool

    // 图中有多少个连通分量
    Count() int
}
```

