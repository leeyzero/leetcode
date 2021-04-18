package tree

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 题目：236. 二叉树的最近公共祖先
// 难度：中等
// 描述：给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 思路：
func lowestCommonAncestor(root, p, q *base.TreeNode) *base.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	var path1, path2 []*base.TreeNode
	var find1, find2 bool
	findPath(root, p, &path1, &find1)
	findPath(root, q, &path2, &find2)
	if !find1 || !find2 {
		return nil
	}

	commonAncessor := (*base.TreeNode)(nil)
	commonLen := len(path1)
	if len(path2) < commonLen {
		commonLen = len(path2)
	}
	for i := 0; i < commonLen && path1[i] == path2[i]; i++ {
		commonAncessor = path1[i]
	}
	return commonAncessor
}

func findPath(node *base.TreeNode, target *base.TreeNode, path *[]*base.TreeNode, find *bool) {
	if node == nil || *find {
		return
	}

	*path = append(*path, node)
	if node == target {
		*find = true
		return
	}

	findPath(node.Left, target, path, find)
	findPath(node.Right, target, path, find)
	if !(*find) {
		*path = (*path)[:len(*path)-1]
	}
}
