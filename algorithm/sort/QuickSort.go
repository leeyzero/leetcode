package sort

// 快速排序，核心是partion方法
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

// 快慢指针
// 选择一个pivot，使得坐标小于pivot的元素小于等于arr[pivot], 坐标大于pivot的元素大于arr[pivot]
func partition(arr []int, left, right int) int {
	pivot := left
	for i := left + 1; i <= right; i++ {
		if arr[i] <= arr[left] {
			pivot++
			arr[pivot], arr[i] = arr[i], arr[pivot]
		}
	}

	arr[left], arr[pivot] = arr[pivot], arr[left]
	return pivot
}
