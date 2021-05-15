package slidingwindow

import (
	"reflect"
	"testing"
)

func TestFindContinuousSequence(t *testing.T) {
	tests := [][]interface{}{
		{9, [][]int{{2, 3, 4}, {4, 5}}},
		{15, [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}}},
	}
	for _, test := range tests {
		p1 := (test[0]).(int)
		want := (test[1]).([][]int)
		if got := findContinuousSequence(p1); !reflect.DeepEqual(got, want) {
			t.Errorf("findContinuousSequence(%v).got:%v want:%v", p1, got, want)
		}
	}
}
