package TwoSum

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum(nums []int, target int) []int {
	ans := []int{}
    if len(nums) <= 1 {
        return ans
    }

    for i, j := 0, len(nums)-1; i < j; {
        sum := nums[i] + nums[j] 
        if target == sum {
            ans = []int{nums[i], nums[j]}
            break
        } else if sum > target {
            j--
        } else {
            i++
        }
    }
    return ans
}