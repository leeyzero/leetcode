package search

import (
	"testing"
)

func TestLadderLength(t *testing.T) {
	tests := [][]interface{}{
		{"hot", "dog", []string{"hot", "dog", "cog", "pot", "dot"}, 3},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		// {"qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}, 5},
	}
	for _, test := range tests {
		beginWord := (test[0]).(string)
		endWord := (test[1]).(string)
		wordList := (test[2]).([]string)
		want := (test[3]).(int)
		if got := ladderLength(beginWord, endWord, wordList); got != want {
			t.Errorf("ladderLength(%v, %v, %v).got:%v want:%v", beginWord, endWord, wordList, got, want)
		}
	}
}
