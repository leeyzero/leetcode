package algorithm

import (
    "testing"
)

func TestIsSubStructure1(t *testing.T) {
    p1 := []int{1, 2, -1, -1, 3, -1, -1}
    p2 := []int{3, 1, -1, -1, -1}
    A := makeTreeNode(p1) 
    B := makeTreeNode(p2)
    if got, want := isSubStructure(A, B), false; got != want {
        t.Errorf("isSubStructure(%v, %v). got:%v want:%v", p1, p2, got, want)
    }
}

func TestIsSubStructure2(t *testing.T) {
    p1 := []int{3, 4, 1, -1, -1, 2, -1, -1, 5, -1, -1}
    p2 := []int{4, 1, -1, -1, -1}
    A := makeTreeNode(p1) 
    B := makeTreeNode(p2)
    if got, want := isSubStructure(A, B), true; got != want {
        t.Errorf("isSubStructure(%v, %v). got:%v want:%v", p1, p2, got, want)
    }
}

