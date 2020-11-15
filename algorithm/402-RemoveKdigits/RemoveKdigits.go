package RemoveKdigits

import (
    "strings"
)

// https://leetcode-cn.com/problems/remove-k-digits/
// 解题思路：对于有n个数字的序列[D0D1D2...Dn-1]，从左往右找到第一个位置i(0 <= i < n)，使得Di < Di-1
// 并删去Di-1。如果不存在这样的i，则说明整个序列单调递增，删除最后一个数字即可。基于此，我们可以对整个数字
// 字符序列执行一次这个策略，删去一个字符后，剩下的长度为n-1的字符形成新的子问题，可以继续使用形同策略，直到
// 删除k次。
// 对于上述思路，时间复杂度最差为O(kn)(数字时单调递减的情况下)，可以考虑从左往右增量构造最终答案
// 可以使用一个栈维护当前的答案，栈中的元素表示截至到当前为止，删除不超过k次数字后所得到的最小整数，根据前面
// 的讨论，栈中从栈底到栈顶，元素是单调递增的

// 常规实现，时间复杂度O(kn)
func removeKdigits(num string, k int) string {
   if len(num) <= k {
        return "0"
    }

    minNum := num
    for i := 0; i < k; i++ {
        minNum = removeOneDigits(minNum)
    }
    
    if minNum == "" {
        minNum = "0"
    }
    return minNum
}

func removeOneDigits(num string) string {
    if num == "" {
        return num
    }

    delIndex := len(num) - 1
    for i := 1; i < len(num); i++ {
        if num[i-1] > num[i] {
            delIndex = i-1
            break
        }
    }

    left := []byte(num)[0:delIndex]
    right := []byte(num)[delIndex+1:]
    result := string(left) + string(right)
    return strings.TrimLeft(result, "0")
}

// 使用栈实现，增量记录结果，时间复杂度O(n)
func removeKdigits2(num string, k int) string {
    stack := []byte{}
    for i := 0; i < len(num); i++ {
        for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, num[i])
    }

    if k > len(stack) {
        stack = stack[:0]
        k = len(stack)
    }
    stack = stack[:len(stack)-k]
    ans := strings.TrimLeft(string(stack), "0")
    if ans == "" {
        ans = "0"
    }
    return ans
}