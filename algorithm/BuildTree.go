package algorithm

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 本质上是二叉树的前序遍历，先通过preorder找到根节点，然后通过根节点在inorder中找到左右子树的大小
// 构建出根节点后，在使用相同策略构建左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	root := (*TreeNode)(nil)
	buildTreeAux(&root, preorder, inorder)
	return root
}

func buildTreeAux(node **TreeNode, preorder []int, inorder []int) {
	if len(preorder) <= 0 {
		return
	}

	*node = &TreeNode{preorder[0], nil, nil}
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
