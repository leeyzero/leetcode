package DeleteNode

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// Definition for singly-linked list.
type ListNode struct {
    Val int
	Next *ListNode
}
 
func deleteNode(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{0, head}
    pre := sentinel
    for curr := head; curr != nil; curr, pre = curr.Next, pre.Next {
        if curr.Val == val {
            pre.Next = curr.Next
            curr = nil
            break
        }
    }
    return sentinel.Next
}