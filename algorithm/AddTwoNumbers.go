package algorithm

// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sentinel := ListNode{0, nil}
    curr, curr1, curr2 := &sentinel, l1, l2
    carry := 0
    for curr1 != nil || curr2 != nil || carry > 0 {
        val := carry
        if curr1 != nil {
            val += curr1.Val
            curr1 = curr1.Next
        }
        if curr2 != nil {
            val += curr2.Val
            curr2 = curr2.Next
        }
        carry = val / 10
        curr.Next = &ListNode{val%10, nil}
        curr = curr.Next
    }
    return sentinel.Next
}