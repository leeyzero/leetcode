package slidingwindow

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := [][]interface{}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	}
	for _, test := range tests {
		p1 := (test[0]).([]int)
		p2 := (test[1]).(int)
		want := (test[2]).([]int)
		if got := maxSlidingWindow(p1, p2); !reflect.DeepEqual(got, want) {
			t.Errorf("maxSlidingWindow(%v, %v).got:%v want:%v", p1, p2, got, want)
		}
	}
}
