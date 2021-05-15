package dynamicprogramming

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	testIsMatch(IsMatchByRecursive, t)
	testIsMatch(IsMatchByMemo, t)
	testIsMatch(IsMatchByDp, t)
}

func testIsMatch(fn func(s, p string) bool, t *testing.T) {
	s := "abc"
	p := "abc"
	ans := fn(s, p)
	if !ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = "abc"
	p = "abd"
	ans = fn(s, p)
	if ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = "abc"
	p = "a.c"
	ans = fn(s, p)
	if !ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = "abc"
	p = "a.d"
	ans = fn(s, p)
	if ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = "abc"
	p = "d*a*bc"
	ans = fn(s, p)
	if !ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = ""
	p = ""
	ans = fn(s, p)
	if !ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = "abc"
	p = ""
	ans = fn(s, p)
	if ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}

	s = ""
	p = ".*"
	ans = fn(s, p)
	if !ans {
		t.Errorf("test s=%s p=%s fail", s, p)
	}
}
