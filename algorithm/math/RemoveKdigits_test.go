package math

import (
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	num := "1432219"
	r := removeKdigits(num, 3)
	e := "1219"
	if r != e {
		t.Errorf("case fail.input:%s got:%s expect:%s", num, r, e)
	}

	num = "1132219"
	r = removeKdigits(num, 2)
	e = "11219"
	if r != e {
		t.Errorf("case fail.input:%s got:%s expect:%s", num, r, e)
	}

	num = "17837"
	r = removeKdigits(num, 1)
	e = "1737"
	if r != e {
		t.Errorf("case fail.input:%s got:%s expect:%s", num, r, e)
	}
}

func TestRemoveKdigits2(t *testing.T) {
	num := "1432219"
	r := removeKdigits2(num, 3)
	e := "1219"
	if r != e {
		t.Errorf("case fail.input:%s got:%s expect:%s", num, r, e)
	}

	num = "1132219"
	r = removeKdigits2(num, 2)
	e = "11219"
	if r != e {
		t.Errorf("case fail.input:%s got:%s expect:%s", num, r, e)
	}

	num = "17837"
	r = removeKdigits2(num, 1)
	e = "1737"
	if r != e {
		t.Errorf("case fail.input:%s got:%s expect:%s", num, r, e)
	}
}
