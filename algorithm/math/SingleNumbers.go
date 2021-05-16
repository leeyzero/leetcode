package math

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
// 剑指 Offer 56 - I. 数组中数字出现的次数
// 难度：中等
// 描述：一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
// 思路：先考虑只有一个元素不同的情况。对于两个相同元素的异或运算，其值为0
// 如果数组中只有一个元素不同，则对数组逐一进行异或运算，得到的最终结果及时不同的元素
// 在考虑数组中有两个元素不同的情况，任然对数组的值进行异或运算，得到的值时不同元素异或的结果，肯定不为0
// 如果我们找出结果中一个为1的位，以这个位将这两个元素分到两个组中，分别对每个组进行异或运算就可以找到这
// 两个不同的元素了
func singleNumbers(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	xorResult := 0
	for _, num := range nums {
		xorResult ^= num
	}

	idxBit := findFirstBitIsOne(xorResult)
	num1, num2 := 0, 0
	for _, num := range nums {
		if isIndexBitIsOne(num, idxBit) {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}

func findFirstBitIsOne(num int) int {
	index := 0
	for num > 0 && num&1 == 0 {
		num >>= 1
		index++
	}
	return index
}

func isIndexBitIsOne(num int, idxBit int) bool {
	num >>= idxBit
	return (num&1 == 1)
}
