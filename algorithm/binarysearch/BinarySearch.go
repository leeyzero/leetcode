package binarysearch

// 对排序数组进行二分查找，闭区间[0, n-1]
func binarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] < target {
			lo = mid + 1
		} else if arr[mid] > target {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 左闭右开区间[0, n)
func binarySearch2(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if target > arr[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(arr) && arr[lo] == target {
		return lo
	}
	return -1
}
