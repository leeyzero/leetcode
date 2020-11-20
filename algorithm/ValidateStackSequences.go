package algorithm

// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
// 解题思路：要充分了解栈的特性，使用一个栈来保存pushed的元素，每次push后，将top值和poped的顶部元素比较
// 如果相等，出栈，继续比较下一个栈顶元素
// 如果不相等，继续入栈，知道pushed中的所有元素均已入栈
// 如果最后栈为空，则说明顺序时合法的
func validateStackSequences(pushed []int, popped []int) bool {
	stack, j := []int{}, 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}