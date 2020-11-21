package algorithm

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
// 解题思路：将num跟1进行&操作，如果为1，标示改为时1，然后将num右移1为，继续处理，直到num为0
func hammingWeight(num uint32) int {
    ans := 0
    for num > 0 {
        if (num&1) == 1 {
            ans += 1
        }
        num >>= 1
    }
    return ans
}