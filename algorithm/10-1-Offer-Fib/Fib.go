package Fib

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
func fib(n int) int {
    if n <= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    a, b := 0, 1
    for k := 2; k <= n; k++ {
        next := (a + b)%1000000007
        a, b = b, next
    }
    return b
}
