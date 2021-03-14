package sort

// 插入排序
// 一共进行N-1轮，每轮将待插入的数跟前面已排序的数逐一比较，找到合适的位置进行插入
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		curr := arr[i]
		for ; j > 0 && curr < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = curr
	}
}
