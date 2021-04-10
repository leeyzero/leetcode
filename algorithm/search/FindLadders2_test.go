package search

import (
	"reflect"
	"testing"
)

func TestFindLaddersII(t *testing.T) {
	tests := [][]interface{}{
		{"red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}, [][]string{{"red", "ted", "tex", "tax"}, {"red", "ted", "tad", "tax"}, {"red", "rex", "tex", "tax"}}},
	}
	for _, test := range tests {
		beginWord := (test[0]).(string)
		endWord := (test[1]).(string)
		wordList := (test[2]).([]string)
		want := (test[3]).([][]string)
		if got := findLaddersII(beginWord, endWord, wordList); !reflect.DeepEqual(got, want) {
			t.Errorf("findLaddersII(%v, %v, %v).got:%v want:%v", beginWord, endWord, wordList, got, want)
		}
	}
}
