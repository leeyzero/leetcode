package algorithm

// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
// 解题思路：把B[i] = A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]
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