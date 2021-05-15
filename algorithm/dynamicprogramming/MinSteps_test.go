package dynamicprogramming

import "testing"

func TestMinStep(t *testing.T) {
	tests := [][]int{
		{3, 3},
	}
	for _, test := range tests {
		n := test[0]
		want := test[1]
		if got := minSteps(n); got != want {
			t.Errorf("minSteps(%v).got:%v want:%v", n, got, want)
		}
	}
}
