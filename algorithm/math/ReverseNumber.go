package math

import (
	"math"
)

// https://leetcode-cn.com/problems/reverse-integer/
// 题目：7. 整数反转
// 难度：简单
// 描述：给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。
// 思路：
// 通过循环将数字x的每一位拆开，在计算新值时每一步都判断是否溢出。
// 方法一：
// 溢出条件有两个，一个是大于整数最大值MAX_VALUE，另一个是小于整数最小值MIN_VALUE，设当前计算结果为ans，下一位为pop。
// 从ans * 10 + pop > MAX_VALUE这个溢出条件来看
// 当出现 ans > MAX_VALUE / 10 且 还有pop需要添加 时，则一定溢出
// 当出现 ans == MAX_VALUE / 10 且 pop > 7 时，则一定溢出，7是2^31 - 1的个位数
// 从ans * 10 + pop < MIN_VALUE这个溢出条件来看
// 当出现 ans < MIN_VALUE / 10 且 还有pop需要添加 时，则一定溢出
// 当出现 ans == MIN_VALUE / 10 且 pop < -8 时，则一定溢出，8是-2^31的个位数
// 方法二：
// 使用int64存储结果
func reverseNumber(x int) int {
	var ans int64
	for x != 0 {
		ans = ans*10 + int64(x%10)
		x /= 10
	}
	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}
	return int(ans)
}
