package sort

import (
	"testing"
)

func TestReversePairs(t *testing.T) {
	tests := [][]interface{}{
		{[]int{7, 5, 6, 4}, 5},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).(int)
		if got := reversePairs(p1); got != want {
			t.Errorf("reversePairs(%v).got:%v want:%v", p1, got, want)
		}
	}
}
