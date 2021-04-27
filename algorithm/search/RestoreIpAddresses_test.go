package search

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := [][]interface{}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"0000", []string{"0.0.0.0"}},
		{"1111", []string{"1.1.1.1"}},
		{"010010", []string{"0.10.0.10", "0.100.1.0"}},
	}
	for _, test := range tests {
		s := (test[0]).(string)
		want := (test[1]).([]string)
		if got := restoreIpAddresses(s); !reflect.DeepEqual(got, want) {
			t.Errorf("restoreIpAddresses(%v).got:%v want:%v", s, got, want)
		}
	}
}
