package sort

// 归并排序
// Divide: 将原问题分解成一系列子问题
// Conquer: 递归地解各子问题。若子问题足够小，则直接求解
// Commbine: 将子问题的结果合并成原问题的解
func MergeSort(arr []int) {
	aux := make([]int, len(arr))
	mergeSort(arr, 0, len(arr)-1, aux)
}

func mergeSort(arr []int, lo int, hi int, aux []int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(arr, lo, mid, aux)
	mergeSort(arr, mid+1, hi, aux)
	mergeCore(arr, lo, mid, hi, aux)
}

func mergeCore(arr []int, lo int, mid int, hi int, aux []int) {
	for k := lo; k <= hi; k++ {
		aux[k] = arr[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > hi {
			arr[k] = aux[i]
			i++
		} else if aux[i] < arr[j] {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}
}
