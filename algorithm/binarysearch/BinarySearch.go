package binarysearch

// 对排序数组进行二分查找，找到target所在的下标，不存在返回-1
// 左闭右闭区间[0, n-1]
func binarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] < target {
			lo = mid + 1
		} else if arr[mid] > target {
			hi = mid - 1
		} else if arr[mid] == target {
			return mid
		}
	}
	// 未找到
	return -1
}

// 左闭右开区间[0, n)
func binarySearch2(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if target > arr[mid] {
			lo = mid + 1
		} else if target < arr[mid] {
			hi = mid // 注意右侧是开区间，所以使用mid
		} else if target == arr[mid] {
			return mid
		}
	}

	// 由于结束条件是lo == hi，这个元素未检测到，打个补丁
	if lo < len(arr) && arr[lo] == target {
		return lo
	}
	return -1
}
