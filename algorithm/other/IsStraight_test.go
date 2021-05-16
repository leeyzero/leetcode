package other

import (
	"testing"
)

func TestIsStraight(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{0, 0, 1, 2, 5}, true},
		{[]int{3, 4, 2, 1, 5}, true},
		{[]int{0, 0, 2, 2, 5}, false},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).(bool)
		if got := isStraight(p1); got != want {
			t.Errorf("isStraight(%v).got:%v want:%v", p1, got, want)
		}
	}
}
