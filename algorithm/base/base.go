package base

import (
	"encoding/json"
)

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

func Forward(head *ListNode, length int) *ListNode {
	curr := head
	for i := 0; curr != nil && i < length; i++ {
		curr = curr.Next
	}
	return curr
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

func LevelOrderTraverseTreeNode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		ans = append(ans, curr.Val)
		if curr.Left != nil {
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}
	return ans
}

// 按层序遍历反序列化成一棵树，如：
// 输入：
// [3,9,20,-1,-1,15,7]
// 输出：
//    3
// 9    20
//    15  7
func UnmarshalTreeNode(content string) (*TreeNode, error) {
	var numbers []int
	if err := json.Unmarshal([]byte(content), &numbers); err != nil {
		return nil, err
	}
	if len(numbers) <= 0 || numbers[0] == -1 {
		return nil, nil
	}

	root := &TreeNode{numbers[0], nil, nil}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 && i < len(numbers) {
		curr := q[0]
		q = q[1:]
		children := []**TreeNode{&curr.Left, &curr.Right}
		for j := 0; j < len(children); j++ {
			if numbers[i+j] == -1 {
				continue
			}

			*children[j] = &TreeNode{numbers[i+j], nil, nil}
			q = append(q, *children[j])
		}
		i += len(children)
	}
	return root, nil
}
