package LevelOrder

 // https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
 type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }
 
func levelOrder(root *TreeNode) []int {
    ans := []int{}
    if root == nil {
        return ans
    }

    queue := []*TreeNode{root}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        ans = append(ans, curr.Val)
        if curr.Left != nil {
            queue = append(queue, curr.Left)
        }
        if curr.Right != nil {
            queue = append(queue, curr.Right)
        }
    }
    return ans
}