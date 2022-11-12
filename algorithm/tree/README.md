## 树

二叉树是树的一种特化，由于二叉树使用的比较多，这里讨论的树主要是指二叉树，不失一般性，二叉树的处理方法可以应用于N叉树。

## 结构定义

二叉树的结构定义

```
type BinaryTreeNode struct {
    Val   int
    Left  *BinaryTreeNode
    Right *BinaryTreeNode
}
```

N叉树的结构定义
```
type TreeNode struct {
    Val      int
    Children []*TreeNode
}
```

## 遍历框架

二叉树的遍历框架

```
func traverse(node *BinaryTreeNode) {
    // 递归结束条件
    if node == nil {
        return
    }

    // 前序遍历
    traverse(node.Left)
    // 中序遍历
    traverse(node.Right)
    // 后续遍历
}
```

N叉树的遍历框架

```
func traverse(node *TreeNode) {
    // 递归结束条件
    if node == nil {
        return
    }

    // 前序遍历
    for _, child := range node.Children {
        traverse(child)
    }
    // 后序遍历
}

```

## 经典问题
[105. 从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

难度：中等

解题思路：本质上是二叉树的遍历问题。需要清楚前序遍历和中序遍历的特点。前序序列的第一个元素即为二叉树的根，通过前序遍历的根元素，可以在序序列中找到这个根元素，根元素左边子序列即为左子树，右边子序列即为右子树，递归构建左右子树。

## 参考资料
[1] [Tree Traversals](https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/)
[2] [labuladong fucking-algorithm](https://github.com/labuladong/fucking-algorithm/tree/plugin)