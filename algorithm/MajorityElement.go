package algorithm

// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// 解题思路：
// 方法一：哈希表统计法，遍历数组，用hash表统计各数字的次数，即可找出众数。时间复杂度O(N), 空间复杂度O(N)
// 方法二：数组排序，数组中点的元素一定时众数。时间复杂度O(NlogN), 空间复杂度O(1)
// 方法三：摩尔投票法，核心理念为票数正负抵消。时间复杂度O(N), 空间复杂度O(1)
// 方法四：基于partition思想。时间复杂度O(N), 空间复杂度O(1)
func majorityElement(nums []int) int {
	votes := 0
	ans := 0
	for _, num := range nums {
		if votes == 0 {
			ans = num
		} 

		if num == ans {
			votes++
		} else {
			votes--
		}
	}
	return ans
}