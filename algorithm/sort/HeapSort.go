package sort

// 堆排序
// 步骤1：先用buildMaxHeap将数据构建成一个最大堆，即最大元素为arr[0]，然后将它和arr[len-1]交换来达到正确的位置
// 步骤2：对arr[0,len-1)执行一个maxHeapify，使arr[0,len-1)仍然保持最大堆特性，继续执行步骤1
func HeapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, 0, i-1)
	}
}

// 将数组arr构建成一个最大堆，arr[len/2+1, len)都是叶子，可以看成是只有一个元素的堆
func buildMaxHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		maxHeapify(arr, i, len(arr)-1)
	}
}

// 下降算法(递归实现)，arr[lo...hi]为一个近似堆，即只有arr[lo]不满足堆次序条件，通过下降算法调整为堆
func maxHeapify(arr []int, lo int, hi int) {
	left, right := left(lo), right(lo)
	largest := lo
	if left <= hi && arr[left] > arr[lo] {
		largest = left
	}
	if right <= hi && arr[right] > arr[largest] {
		largest = right
	}
	if largest != lo {
		arr[lo], arr[largest] = arr[largest], arr[lo]
		maxHeapify(arr, largest, hi)
	}
}

// 下降算法(迭代实现)
func maxHeapify2(arr []int, lo int, hi int) {
	c := left(lo)
	for c <= hi {
		if c < hi && arr[c] < arr[c+1] {
			c = c + 1
		}
		if arr[c] > arr[lo] {
			arr[lo], arr[c] = arr[c], arr[lo]
			lo = c
			c = left(lo)
		} else {
			break
		}
	}
}

// 结点i的左孩子
func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
