package linklist

import "github.com/leeyzero/leetcode/algorithm/base"

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 题目：剑指 Offer 06. 从尾到头打印链表
// 难度：简单
// 描述：输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
func reversePrint(head *base.ListNode) []int {
	ans := []int{}
	reversePrintAux(head, &ans)
	return ans
}

func reversePrintAux(node *base.ListNode, ans *[]int) {
	if node == nil {
		return
	}

	reversePrintAux(node.Next, ans)
	*ans = append(*ans, node.Val)
}
