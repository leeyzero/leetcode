package sort

// 插入排序
// 一共进行N-1轮，每轮将待插入的数跟前面已排序的数逐一比较，找到合适的位置进行插入
func insertSort(arr []int) []int {
	ans := make([]int, len(arr))
	copy(ans, arr)
	for i := 1; i < len(ans); i++ {
		j := i
		curr := ans[i]
		for ; j > 0 && curr < ans[j-1]; j-- {
			ans[j] = ans[j-1]
		}
		ans[j] = curr
	}
	return ans
}