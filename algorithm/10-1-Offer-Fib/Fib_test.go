package Fib

import (
    "testing"
)

func TestFib(t *testing.T) {
    tests := [][]int{
        []int{0, 0},
        []int{1, 1},
        []int{2, 1},
        []int{3, 2},
        []int{4, 3},
        []int{5, 5},
        []int{45, 134903163},
    }
    for _, test := range tests {
        if r := fib(test[0]); r != test[1] {
            t.Errorf("test fib fail.in:%d got:%d expect:%d", test[0], r, test[1])
        }
    }
}
