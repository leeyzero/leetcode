package algorithm

// 剑指 Offer 68 - II. 二叉树的最近公共祖先
// https://leetcode-cn.com/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// 解题思路：
func lowestCommonAncestor2(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode{
	if root == nil || p == nil || q == nil {
		return nil
	}

	pPath, qPath := []*TreeNode{}, []*TreeNode{}
	pFound, qFound := false, false
	findPath(root, p, &pPath, &pFound)
	findPath(root, q, &qPath, &qFound)
	if !pFound || !qFound {
		return nil
	}

	var ans *TreeNode
	for i, j := 0, 0; i < len(pPath) && j < len(qPath); i, j = i+1, j+1 {
		if pPath[i] == qPath[j] {
			ans = pPath[i]
		} else {
			break
		}
	}
	return ans
}

func findPath(node *TreeNode, target *TreeNode, path *[]*TreeNode, found *bool) {
	if node == nil || *found {
		return 
	}

	*path = append(*path, node)
	if node == target {
		*found = true
		return
	}

	findPath(node.Left, target, path, found)
	findPath(node.Right, target, path, found)
	if !(*found) {
		*path = (*path)[:len(*path)-1]
	}
}