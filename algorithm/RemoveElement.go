package algorithm

// https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
    j := -1
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            j++
            nums[j] = nums[i]
        }
    }
    return j + 1
}