package unionfind

import "testing"

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(6)
	if got, want := uf.Count(), 6; got != want {
		t.Errorf("uf.Count().got:%v want:%v", got, want)
	}

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)
	if got, want := uf.Count(), 3; got != want {
		t.Errorf("uf.Count().got:%v want:%v", got, want)
	}
	if got, want := uf.Connected(0, 3), true; got != want {
		t.Errorf("uf.Connect(0, 3).got:%v want:%v", got, want)
	}
}
