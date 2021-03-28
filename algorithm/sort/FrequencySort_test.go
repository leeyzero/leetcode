package sort

import (
	"testing"
)

func inStringArray(needle string, hasty []string) bool {
	for _, v := range hasty {
		if v == needle {
			return true
		}
	}
	return false
}

func TestFrequencySort(t *testing.T) {
	tests := [][]interface{}{
		{"tree", []string{"eert", "eetr"}},
		{"cccaaa", []string{"cccaaa"}},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		want := (test[1]).([]string)
		if got := frequencySort(s); !inStringArray(got, want) {
			t.Errorf("frequencySort(%v).got:%v want:%v", s, got, want)
		}
	}
}
