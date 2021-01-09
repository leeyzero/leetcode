package algorithm

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
// 解题思路：递归算法比较简单，这里使用迭代实现
type SNode struct {
	node *Node
	i int
}

func postorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	stack := []*SNode{}
	curr, lastVisitNode := root, (*Node)(nil)
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, &SNode{curr, 0})
			if len(curr.Children) > 0 {
				curr = curr.Children[0]
			} else {
				curr = nil
			}
		}

		top := stack[len(stack)-1]
		count := len(top.node.Children)
		if count <= 0 || top.node.Children[count-1] == lastVisitNode {
			stack = stack[:len(stack)-1]
			ans = append(ans, top.node.Val)
			lastVisitNode = top.node
		} else {
			// 切换到下一个child
			top.i++
			curr = top.node.Children[top.i]
		}
	}
    return ans
}