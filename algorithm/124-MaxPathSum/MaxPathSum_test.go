package MaxSumPath

import (
    "testing"
)

func TestMaxPathSum(t *testing.T) {
    root := &TreeNode{1, nil, nil}
    root.Left = &TreeNode{2, nil, nil}
    root.Right = &TreeNode{3, nil, nil}
    sum := maxPathSum(root)
    exp := 6
    if sum != exp {
        t.Errorf("maxPathSum fail.got:%v want:%v", sum, exp)
    }
}
