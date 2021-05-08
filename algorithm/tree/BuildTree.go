package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 题目：105. 从前序与中序遍历序列构造二叉树
// 难度：中等
// 描述：根据一棵树的前序遍历与中序遍历构造二叉树。
// 本质上是二叉树的前序遍历，先通过preorder找到根节点，然后通过根节点在inorder中找到左右子树的大小
// 构建出根节点后，在使用相同策略构建左右子树
func buildTree(preorder []int, inorder []int) *base.TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	root := (*base.TreeNode)(nil)
	buildTreeAux(&root, preorder, inorder)
	return root
}

func buildTreeAux(node **base.TreeNode, preorder []int, inorder []int) {
	if len(preorder) <= 0 {
		return
	}

	*node = &base.TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	var pivot int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			pivot = i
			break
		}
	}

	buildTreeAux(&(*node).Left, preorder[1:1+pivot], inorder[:pivot])
	buildTreeAux(&(*node).Right, preorder[1+pivot:], inorder[pivot+1:])
}
