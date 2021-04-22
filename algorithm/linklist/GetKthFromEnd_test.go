package linklist

import (
	"reflect"
	"testing"

	"github.com/leeyzero/leetcode/algorithm/base"
)

func TestGetKthFromEnd(t *testing.T) {
	head := base.MakeLinkList([]int{1, 2, 3, 4, 5})
	h1 := getKthFromEnd(head, 2)
	if !reflect.DeepEqual([]int{4, 5}, base.LinkListToSlice(h1)) {
		t.Errorf("test getKthFromEnd fail.got:%v", base.LinkListToSlice(h1))
	}
}
