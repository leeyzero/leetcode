package algorithm

// https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
    ans := [][]int{}
    permuteCore(nums, 0, &ans)
    return ans
}

func permuteCore(nums []int, pos int, ans *[][]int) {
    if pos >= len(nums) - 1 {
        curr := make([]int, len(nums))
        copy(curr, nums)
        *ans = append(*ans, curr)
        return
    }

    for i := pos; i < len(nums); i++ {
        nums[pos], nums[i] = nums[i], nums[pos]
        permuteCore(nums, pos+1, ans)
        nums[pos], nums[i] = nums[i], nums[pos]
    }
}