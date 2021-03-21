package algorithm

import (
	"reflect"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	in := [][]int{
		{7, 0}, {4, 4},
		{7, 1}, {5, 0},
		{6, 1}, {5, 2},
	}
	want := [][]int{
		{5, 0}, {7, 0},
		{5, 2}, {6, 1},
		{4, 4}, {7, 1},
	}
	if got := reconstructQueue(in); !reflect.DeepEqual(got, want) {
		t.Errorf("reconstructQueue.in:%v got:%v want:%v", in, got, want)
	}
}
