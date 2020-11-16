package NumWays

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func numWays(n int) int {
    if n <= 1 {
        return 1
    }

    a, b := 1, 1
    for i := 2; i <= n; i++ {
        next := (a+b) % 1000000007
        a, b = b, next
    }
    return b
}