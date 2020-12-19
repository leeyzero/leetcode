package algorithm

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	return reversePairsCore(nums, 0, len(nums)-1)
}

func reversePairsCore(nums []int, left int, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right - left) / 2
	cnt := reversePairsCore(nums, left, mid) + reversePairsCore(nums, mid + 1, right)
	tmp := []int{}
	i, j := left, mid + 1
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
			cnt += j - (mid + 1)
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += right - (mid + 1) + 1
	}
	for ; j <= right; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := left; i <= right; i++ {
		nums[i] = tmp[i - left]
	}
	return cnt
}