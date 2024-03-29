package tree

import (
	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
// 114. 二叉树展开为链表
// 难度：中等
//
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
// 示例 1：
//    1
//  2  5
// 3 4  6
//
// 1 -> 2 -> 3 -> 4 -> 5 -> 6
//
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
//
// 方法一：最简单的办法是进行一次前序遍历，并将结果保存一个列表中，然后再对列表进行一次遍历，并修改各个节点的指针
// 这种方法时间复杂度为O(N)，空间复杂度也是O（N），对于前序遍历，采用递归比较简单，但在生产环境需要考虑栈溢出的情况。
// 也可以采用迭代，使用一个栈模拟递归调用，但实现相对复杂些。
//
// 方法二：寻找前驱节点。
// 如果一个节点的左子节点为空，则该节点不需要进行展开操作。
// 如果一个节点的左子节点不为空，则该节点的左子树中的最后一个节点被访问之后，该节点的右子节点被访问。
// 该节点的左子树中最后一个被访问的节点是左子树中的最右边的节点，也是该节点的前驱节点。因此，问题转化成寻找当前节点的前驱节点。
// 具体做法是：对于当前节点，如果其左子节点不为空，则在其左子树中找到最右边的节点，作为前驱节点，
// 将当前节点的右子节点赋给前驱节点的右子节点，然后将当前节点的左子节点赋给当前节点的右子节点，并将当前节点的左子节点设为空。
// 对当前节点处理结束后，继续处理链表中的下一个节点，直到所有节点都处理结束。
// 这种方法时间复杂度为O（N），空间复杂度为O（1）
func flatten(root *base.TreeNode) {
	curr := root
	for curr != nil {
		// 左节点不为nil，寻找当前节点的前驱节点，并将当前节点转为单向链表
		if curr.Left != nil {
			next := curr.Left
			// 当前左子树的最右节点为当前节点右节点的前驱节点
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			// 前驱节点的右指针指向当前节点的右节点
			predecessor.Right = curr.Right
			// 当前节点的右指针指向下一个节点
			curr.Right = next
			// 当前节点的左指针置为空
			curr.Left = nil
		}

		// 处理下一个节点
		curr = curr.Right
	}
}
