package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// 题目：剑指 Offer 68 - II. 二叉树的最近公共祖先
// 难度：简单
// 思路：先遍历树分别找到p，q的路径，然后从p，q路径开始找第一个分叉点，则分叉点的前一个即为公共祖先
func lowestCommonAncestor2(root *base.TreeNode, p *base.TreeNode, q *base.TreeNode) *base.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	pPath, qPath := []*base.TreeNode{}, []*base.TreeNode{}
	pFound, qFound := false, false
	findNodePath(root, p, &pPath, &pFound)
	findNodePath(root, q, &qPath, &qFound)
	if !pFound || !qFound {
		return nil
	}

	var ans *base.TreeNode
	for i, j := 0, 0; i < len(pPath) && j < len(qPath); i, j = i+1, j+1 {
		if pPath[i] == qPath[j] {
			ans = pPath[i]
		} else {
			break
		}
	}
	return ans
}

func findNodePath(node *base.TreeNode, target *base.TreeNode, path *[]*base.TreeNode, found *bool) {
	if node == nil || *found {
		return
	}

	*path = append(*path, node)
	if node == target {
		*found = true
		return
	}

	findNodePath(node.Left, target, path, found)
	findNodePath(node.Right, target, path, found)
	if !(*found) {
		*path = (*path)[:len(*path)-1]
	}
}
