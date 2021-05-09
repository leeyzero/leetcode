package algorithm

import (
	"testing"
)

func TestCodec(t *testing.T) {
	tests := [][]string{
		{"1,2,3,#,#,4,5,#,#,#,#", "1,2,3,#,#,4,5,#,#,#,#"},
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

func TestCodec2(t *testing.T) {
	tests := [][]string{
		{"1,2,#,#,3,4,5,#,#,#,#", "1,2,#,#,3,4,5,#,#,#,#"},
	}
	for _, test := range tests {
		p1 := test[0]
		want := test[1]
		codec := &Codec{}
		if got := codec.serialize2(codec.deserialize2(p1)); got != want {
			t.Errorf("codec.serialize(codec.deserialize(%v)).got:%v want:%v", p1, got, want)
		}
	}
}
