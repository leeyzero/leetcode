package greedy

import (
	"testing"
)

func TestCandy(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		want := (test[1]).(int)
		if got := candy(p1); got != want {
			t.Errorf("Candy(%v).got:%v want:%v", p1, got, want)
		}
	}
}
