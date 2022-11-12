package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 题目：105. 从前序与中序遍历序列构造二叉树
// 难度：中等
// 描述：根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 示例1
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
// 解题思路：
// 本质上是二叉树的前序遍历，先通过preorder找到根节点，然后通过根节点在inorder中找到左右子树的大小
// 构建出根节点后，在使用相同策略构建左右子树
func buildTree(preorder []int, inorder []int) *base.TreeNode {
	// 结点已处理完
	if len(preorder) == 0 {
		return nil
	}

	// 创建根结点
	root := &base.TreeNode{preorder[0], nil, nil}

	// 前序序列中的第一个元素为根，在中间序列中找到根，根左边的子序列为左子树，根右边的子序列为右子树
	var pivot int
	for i, val := range inorder {
		if val == preorder[0] {
			pivot = i
			break
		}
	}

	// 递归处理左右子树
	root.Left = buildTree(preorder[1:1+pivot], inorder[0:pivot])
	root.Right = buildTree(preorder[1+pivot:], inorder[pivot+1:])
	return root
}
