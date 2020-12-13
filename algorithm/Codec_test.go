package algorithm

import (
	"testing"
)

func TestCodec(t *testing.T) {
	tests := [][]string{
		{"[1,2,3,null,null,4,5,null,null,null,null]", "[1,2,3,null,null,4,5,null,null,null,null]"},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		codec := &Codec{}
		if got := codec.serialize(codec.deserialize(p1)); got != want {
			t.Errorf("codec.serialize(codec.deserialize(%v)).got:%v want:%v", p1, got, want)
		}
	}
}