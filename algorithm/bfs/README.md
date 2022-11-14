# 深度优先遍历

BFS（Broad First Search，广度优先搜索），其核心思想是把问题抽象成图，从一个点开始，向四周扩散。BFS算法基本上都是使用队列，每次将一个结点头队头取出后，将该节点的所有邻节点加入到队尾。

BFS与DFS的主要区别是：BFS找到的路径一定是最短的，但代价是空间复杂度比DFS高。

## 算法框架
```
# 伪代码
func BFS(start *Node, target *Node) {
    // 核心数据结构：队列
    queue := []*Node{} 
    // 记录访问过的结点
    visited := map[*Node]bool{}

    // 起始点加入队列
    queue = append(queue, start)
    // 起始点加入访问记录
    visited[start] = true
    
    step := 0 // 记录扩散步数
    for len(queue) > 0 {
        for i, n := 0; len(queue); i < n; i++ {
            // 取出队头元素
            curr := queue[0]
            queue = queue[1:]

            // 判断是否到达终点
            if curr == target {
                return step
            }

            // 将当前结点的邻节点加入
            for _, adj := range curr.adajcent() {
                // 忽略访问过的结点
                if _, ok := visited[adj]; ok {
                    continue
                }

                // 邻节点加入队列
                queue = append(queue, adj)
                // 记录邻节点已经访问过
                visited[adj] = true
            }
        }

        // 更新步数 
        step++
    }

    // 未找到
    return -1
}
```

## 经典问题
[752. 打开转盘锁](https://leetcode.cn/problems/open-the-lock/)

## 参考资料
[1] [Breadth First Search or BFS for a Graph](https://www.geeksforgeeks.org/breadth-first-search-or-bfs-for-a-graph/)
[2] [labuladong的算法小抄](https://item.jd.com/12759911.html)
[3] [Breadth-first search](https://en.wikipedia.org/wiki/Breadth-first_search)