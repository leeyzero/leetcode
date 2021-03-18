package base

func MakeLinkList(vals []int) *ListNode {
	sentinel := &ListNode{}
	curr := sentinel
	for _, val := range vals {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return sentinel.Next
}

func LinkListToSlice(head *ListNode) []int {
	result := []int{}
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
