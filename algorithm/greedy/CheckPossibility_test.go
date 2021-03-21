package greedy

import (
	"testing"
)

func TestCheckPossibility(t *testing.T) {
	tests := [][]interface{}{
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{3, 4, 2, 3}, false},
		{[]int{3, 4, 2, 5}, true},
	}
	for _, test := range tests {
		nums := (test[0]).([]int)
		want := (test[1]).(bool)
		if got := checkPossibility(nums); got != want {
			t.Errorf("checkPossibility(%v).got:%v want:%v", nums, got, want)
		}
	}
}
