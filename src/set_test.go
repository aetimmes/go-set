package set

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	NewSet[int]()
	NewSet[float64]()
	NewSet[string]()
	var x int

	i := NewSet(1, 2, -3)
	x = len(i.m)
	if x != 3 {
		t.Errorf("len i.m = %d, want 3", x)
	}

	f := NewSet(1.5, 2.1, -3.2)
	x = len(f.m)
	if x != 3 {
		t.Errorf("len f.m = %d, want 3", x)
	}
	g := NewSet("foo", "bar", "baz")
	x = len(g.m)
	if x != 3 {
		t.Errorf("len g.m = %d, want 3", x)
	}
}

func TestAdd(t *testing.T) {
	var x any
	i := NewSet[int]()
	Add(&i, 1)
	x = i.m[1]
	if x != struct{}{} {
		t.Errorf("i[1] returning %s, want struct{}{}", x)
	}

	Add(&i, 1)
	Add(&i, 1)
	Add(&i, 1)

	x = len(i.m)
	if x != 1 {
		t.Errorf("len(i.m) returning %d, want 1", x)
	}

	j := NewSet[float64]()
	Add(&j, -1.01)
	x = j.m[-1.01]
	if x != struct{}{} {
		t.Errorf("j[1] returning %s, want struct{}{}", x)
	}

	Add(&j, -1.01)
	Add(&j, -1.01)
	Add(&j, -1.01)

	x = len(j.m)
	if x != 1 {
		t.Errorf("len(j.m) returning %d, want 1", x)
	}

	k := NewSet[string]()
	Add(&k, "foo")
	x = k.m["foo"]
	if x != struct{}{} {
		t.Errorf("k[1] returning %s, want struct{}{}", x)
	}

	Add(&k, "foo")
	Add(&k, "foo")
	Add(&k, "foo")

	x = len(k.m)
	if x != 1 {
		t.Errorf("len(k.m) returning %d, want 1", x)
	}
}
