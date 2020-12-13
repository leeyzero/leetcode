package algorithm

import (
	"testing"
	"math"
)

func TestDicesProbability(t *testing.T) {
	tests := [][]interface{}{
		{1, []float64{0.16667,0.16667,0.16667,0.16667,0.16667,0.16667}},
		{2, []float64{0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778}},
	}
	for _, test := range tests {
		p1 := (test[0]).(int)
		want := (test[1]).([]float64)
		if got := dicesProbability(p1); !checkResult(got, want) {
			t.Errorf("dicesProbability(%v).got:%v want:%v", p1, got, want)
		}
	}
}

func checkResult(got, want []float64) bool {
	if len(got) != len(want) {
		return false
	}

	for i := 0; i < len(got); i++ {
		if math.Abs(got[i] - want[i]) > 0.00001 {
			return false
		}
	}
	return true
}