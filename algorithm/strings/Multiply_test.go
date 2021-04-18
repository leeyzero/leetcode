package strings

import "testing"

func TestMultiply(t *testing.T) {
	tests := [][]interface{}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"9", "9", "81"},
		{"9133", "0", "0"},
		{"0", "52", "0"},
	}
	for _, test := range tests {
		num1 := (test[0]).(string)
		num2 := (test[1]).(string)
		want := (test[2]).(string)
		if got := multiply(num1, num2); got != want {
			t.Errorf("multiply(%v, %v).got:%v want:%v", num1, num2, got, want)
		}
	}
}
