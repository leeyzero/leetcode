package greedy

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := [][]interface{}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}
	for _, test := range tests {
		str := (test[0]).(string)
		want := (test[1]).([]int)
		if got := partitionLabels(str); !reflect.DeepEqual(got, want) {
			t.Errorf("partitionLabels(%v).got:%v want:%v", str, got, want)
		}
	}
}
