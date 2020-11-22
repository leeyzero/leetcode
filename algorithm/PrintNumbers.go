package algorithm

// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
// 解题思路：不考虑大数的情况
func printNumbers(n int) []int {
    limit := 1
    for i := 1; i <= n; i++ {
        limit *= 10
    }

    ans := []int{}
    for i := 1; i < limit; i++ {
        ans = append(ans, i)
    }
    return ans
}