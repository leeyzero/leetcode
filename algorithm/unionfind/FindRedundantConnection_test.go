package unionfind

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	want := []int{2, 3}
	if got := findRedundantConnection(edges); !reflect.DeepEqual(got, want) {
		t.Errorf("findRedundantConnection(%v).got:%v want:%v", edges, got, want)
	}
}
