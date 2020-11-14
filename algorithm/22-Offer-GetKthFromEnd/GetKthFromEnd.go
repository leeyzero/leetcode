package GetKthFromEnd

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
type ListNode struct {
	Val int
	Next *ListNode
}

// 思路：快指针向前走k步，然后慢指针和快指针一起走，直到快指针到达结尾
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
    step := 0
    for (step < k) && (fast != nil) {
		fast = fast.Next
        step++
	}
    if step != k {
        return nil
    }

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
