package dynamicprogramming

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := [][]interface{}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(int)
		if got := lengthOfLIS(nums); got != want {
			t.Errorf("lengthOfLIS(%v).got:%v want:%v", nums, got, want)
		}
	}
}
