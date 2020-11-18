package algorithm

func makeLinkList(vals []int) *ListNode {
	sentinel := &ListNode{}
	curr := sentinel
	for _, val := range vals {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return sentinel.Next
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func convertToSlice(head *ListNode) []int {
	result := []int{}
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}
