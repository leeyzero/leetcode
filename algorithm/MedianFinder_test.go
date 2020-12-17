package algorithm

import (
	"testing"
	"math"
)

func TestMedianFinder(t *testing.T) {
	obj := NewMedianFinder()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(3)
	if got, want := obj.FindMedian(), float64(2); math.Abs(got - want) >= 0.000001 {
		t.Errorf("FindMedian.got:%v want:%v", got, want)
	} 
}