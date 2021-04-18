package sort

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := [][]interface{}{
		{[]int{-1, 2, 1, -4}, 1, 2},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		target := (test[1]).(int)
		want := (test[2]).(int)
		if got := threeSumClosest(nums, target); got != want {
			t.Errorf("threeSumClosest(%v, %v).got:%v want:%v", nums, target, got, want)
		}
	}
}
