package array

// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
// 剑指 Offer 66. 构建乘积数组
// 难度：中等
// 描述：给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
// 解题思路：
// 把B[i] = A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]
// 看成A[0]*A[1]*...*A[i-1]和A[i+1]*...*A[n-2]*A[n-1]两部分的乘积，令
// C[i] = A[0]*A[1]*...*A[i-1]; D[i] = A[i+1]*...*A[n-2]*A[n-1]
// C[i]可以自上而下顺序计算出，即：C[i] = C[i-1] * A[i-1] (i>0)
// D[i]可以自下而上顺序计算出，即：D[i] = D[i+1] * A[i+1] (i<n-1)
func constructArr(a []int) []int {
	if len(a) <= 0 {
		return []int{}
	}

	ans := make([]int, len(a))
	// 构建C[i]
	ans[0] = 1
	for i := 1; i < len(a); i++ {
		ans[i] = ans[i-1] * a[i-1]
	}
	// 构建D[i]
	tmp := 1
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		ans[i] *= tmp
	}
	return ans
}
