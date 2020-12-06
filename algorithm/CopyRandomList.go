package algorithm

// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
// 解题思路：
// 方法一：hash表
// 方法二：拼接+拆分
type RNode struct {
    Val int
    Next *RNode
    Random *RNode
}

func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}

	hash := map[*RNode]*RNode{}
	for curr := head; curr != nil; curr = curr.Next {
		hash[curr] = &RNode{curr.Val, nil, nil}
	}
	for curr := head; curr != nil; curr = curr.Next {
		hash[curr].Next = hash[curr.Next]
		hash[curr].Random = hash[curr.Random]
	}
    return hash[head]
}