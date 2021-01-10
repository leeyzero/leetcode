package algorithm

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
    if root == nil {
        return ans
    }

    queue := []*TreeNode{root}
    for i := 0; len(queue) > 0; i++ {
        curr := []int{}
        count := len(queue)
        for k := 0; k < count; k++ {
            node := queue[0]
            queue = queue[1:]
            curr = append(curr, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if i % 2 != 0 {
            for l, r := 0, len(curr)-1; l < r; l, r = l+1, r-1 {
                curr[l], curr[r] = curr[r], curr[l]
            }
        }
        ans = append(ans, curr)
    }
    return ans
}