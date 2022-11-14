package slidingwindow

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	cases := [][]interface{}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}
	for _, c := range cases {
		p1 := c[0].(string)
		p2 := c[1].(string)
		want := c[2].([]int)
		if got := findAnagrams(p1, p2); !reflect.DeepEqual(got, want) {
			t.Errorf("findAnagrams(%v, %v) failed.got:%v want:%v", p1, p2, got, want)
		}
	}
}
