## 滑动窗口

滑动窗口是双指针技术的一种高级应用，简单来说就是使用两个指针，两个指针相隔的区间即为窗口大小。随着两个指针向前移动，窗口也随之移动，不断滑动，故称为滑动窗口。

## 算法框架
```
func slidingWindow(s string, t string) {
    window := map[byte]int{} // 记录窗口内的字符出现的次数
    need := map[byte]int{}   // 记录目标字符各字符需要的次数

    for _, c := range t {
        need[c] += 1
    }

    left, right := 0, 0
    valid := 0
    for right < len(s) {
        // c 为将要移入窗口的字符
        c := s[right]

        // 右指针向前移动扩大窗口
        right++

        // 进行窗口内数据的更新
        // ...

        // debug
        fmt.Printf("window: [%d, %d)\n", left, right)

        // 判断窗口是否需要收缩
        for window_needs_shrink {
            // d 是将要移出窗口的字符
            d := s[left]

            // 左指针向前移动缩小窗口
            left++

            // 进行窗口内数据的更新
        }
    }
}
```

## 经典问题
[1] [76.最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)
[2] [剑指 Offer 59 - I. 滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)
[3] [题目：567. 字符串的排列](https://leetcode.cn/problems/permutation-in-string/)

## 参考资料
[1] [Window Sliding Technique](https://www.geeksforgeeks.org/window-sliding-technique/)
[2] [labuladong的算法小抄](https://item.jd.com/12759911.html)