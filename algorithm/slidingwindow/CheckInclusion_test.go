package slidingwindow

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	cases := [][]interface{}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}
	for _, c := range cases {
		p1 := c[0].(string)
		p2 := c[1].(string)
		want := c[2].(bool)
		if got := checkInclusion(p1, p2); got != want {
			t.Errorf("checkInclusion(%v, %v) failed.got:%v, want:%v", p1, p2, got, want)
		}
	}
}
