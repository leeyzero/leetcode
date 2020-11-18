package algorithm

import (
	"reflect"
	"testing"
)

func TestFindLadders(t *testing.T) {
	ans := findLadders("hot", "dog", []string{"hot", "dog", "dot"})
	exp := []string{"hot", "dot", "dog"}
	if !reflect.DeepEqual(ans, exp) {
		t.Errorf("test case fail.got:%v expect:%v", ans, exp)
	}

	ans = findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	exp = []string{"hit", "hot", "dot", "dog", "log", "cog"}
	if !reflect.DeepEqual(ans, exp) {
		t.Errorf("test case fail.got:%v expect:%v", ans, exp)
	}
}
