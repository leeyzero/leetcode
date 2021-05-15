package dynamicprogramming

import "testing"

func TestMinDistance(t *testing.T) {
	tests := [][]interface{}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, test := range tests {
		s1 := (test[0]).(string)
		s2 := (test[1]).(string)
		want := (test[2]).(int)
		if got := minDistance(s1, s2); got != want {
			t.Errorf("minDistance(%v, %v).got:%v want:%v", s1, s2, got, want)
		}
	}
}
