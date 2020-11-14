// https://leetcode-cn.com/problems/valid-mountain-array/
package ValidMountainArray

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	var i, j int
	for i = 0; i < len(A)-1; i++ {
		if A[i+1] <= A[i] {
			break
		}
	}
	for j = len(A) - 1; j > 0; j-- {
		if A[j-1] <= A[j] {
			break
		}
	}
	return (i == j) && (i > 0) && (i < len(A)-1)
}
