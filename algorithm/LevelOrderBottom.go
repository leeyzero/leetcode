package algorithm

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
    ans := [][]int{}
    if root == nil {
        return ans
    }

    queue := []*TreeNode{root}
    for i := 0; len(queue) > 0; i++ {
        ans = append(ans, []int{})
        count := len(queue)
        for k := 0; k < count; k++ {
            node := queue[0]
            queue = queue[1:]
            ans[i] = append(ans[i], node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    return ans
}