package math

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := [][]interface{}{
		{float64(2.00000), 10, float64(1024.00000)},
		{float64(2.10000), 3, float64(9.26100)},
		{float64(2.00000), -2, float64(0.25000)},
	}
	for _, test := range tests {
		p1 := (test[0]).(float64)
		p2 := (test[1]).(int)
		want := (test[2]).(float64)
		if got := myPow(p1, p2); math.Abs(got-want) > 0.000000001 {
			t.Errorf("myPow(%v, %v).got:%v want:%v abs(got-want):%v", p1, p2, got, want, math.Abs(got-want))
		}
	}
}
