package other

import (
	"testing"
)

func TestSumNums(t *testing.T) {
	tests := [][]int{
		{3, 6},
		{9, 45},
	}
	for _, test := range tests {
		if got, want := sumNums(test[0]), test[1]; got != want {
			t.Errorf("sumNums(%v).got:%v want:%v", test[0], got, want)
		}
	}
}
