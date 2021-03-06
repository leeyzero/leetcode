package tree

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
// 题目：剑指 Offer 26. 树的子结构
// 难度：中等
// 思路：对A采用先序遍历，逐一判断B是否是每个结点的子结构
// 判断子结构时，如果B结点为空时，说明完成匹配, 返回true
// 如果A结点为空时，说明已经越过A的叶结点，匹配失败，返回false
// 如果A、B结点的值不相同，匹配失败
func isSubStructure(A *base.TreeNode, B *base.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isSubStructureCore(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubStructureCore(A *base.TreeNode, B *base.TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return (A.Val == B.Val) && isSubStructureCore(A.Left, B.Left) && isSubStructureCore(A.Right,
		B.Right)
}
