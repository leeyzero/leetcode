package algorithm

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
// 解题思路：沿用异或运算的思想，将数组中的每个元素对应的位相加，对应的位和如果是3的倍数，则说这些位
// 均时由相同的三个元素产生，如果不为0，则说明是唯一那个不同元素产生的
func singleNumber(nums []int) int {
	if len(nums) < 4 {
		return 0
	}

	bits := make([]int, 32)
	for _, num := range nums {
		mask := 1
		for idx := 0; idx < 32 && num > 0; idx++ {
			if (num & mask) == mask {
				bits[idx] += 1
			}
			mask <<= 1
		}
	}

	ans, mask := 0, 1
	for i := 0; i < len(bits); i++ {
		if bits[i] % 3 == 1 {
			ans |= (mask << i)
		}
	}
	return ans
}