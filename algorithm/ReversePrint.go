package algorithm

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	ans := []int{}
	reversePrintAux(head, &ans)
	return ans
}

func reversePrintAux(node *ListNode, ans *[]int) {
	if node == nil {
		return
	}

	reversePrintAux(node.Next, ans)
	*ans = append(*ans, node.Val)
}
