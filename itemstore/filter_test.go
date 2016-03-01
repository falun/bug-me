package itemstore

import (
	"testing"
)

func testTrue(i Item) bool  { return true }
func testFalse(i Item) bool { return false }

func TestItemFilterSimpleAnd(t *testing.T) {
	f := NewFilter(testTrue)
	f = f.And(NewFilter(testTrue))
	if !f.Match(nil) {
		t.Errorf("true && true must be true")
	}

	f = NewFilter(testFalse)
	f = f.And(NewFilter(testTrue))
	if f.Match(nil) {
		t.Errorf("false && true must be false")
	}

	f = NewFilter(testTrue)
	f = f.And(NewFilter(testFalse))
	if f.Match(nil) {
		t.Errorf("true && false must be false")
	}

	f = NewFilter(testFalse)
	f = f.And(NewFilter(testFalse))
	if f.Match(nil) {
		t.Errorf("false && false must be false")
	}
}

func TestItemFilterSimpleOr(t *testing.T) {
	f := NewFilter(testTrue)
	f = f.Or(NewFilter(testTrue))
	if !f.Match(nil) {
		t.Errorf("true || true must be true")
	}

	f = NewFilter(testFalse)
	f = f.Or(NewFilter(testTrue))
	if !f.Match(nil) {
		t.Errorf("false || true must be true")
	}

	f = NewFilter(testTrue)
	f = f.Or(NewFilter(testFalse))
	if !f.Match(nil) {
		t.Errorf("true || false must be true")
	}

	f = NewFilter(testFalse)
	f = f.Or(NewFilter(testFalse))
	if f.Match(nil) {
		t.Errorf("false || false must be false")
	}
}

func TestItemFilterGroupRight(t *testing.T) {
	// true && (false || true)
	f := NewFilter(testTrue)
	f2 := NewFilter(testFalse)
	f2 = f2.Or(NewFilter(testTrue))
	f = f.And(f2)
	if !f.Match(nil) {
		t.Errorf("true && (false || true) must be true")
	}
}

func TestItemFilterGroupLeft(t *testing.T) {
	// (false || true) && true
	f := NewFilter(testFalse)
	f = f.Or(NewFilter(testTrue))
	f2 := NewFilter(testTrue)
	f = f.And(f2)
	if !f.Match(nil) {
		t.Errorf("(false || true) && true must be true")
	}
}

/*
func TestItemFilter(t *testing.T) {
}

func TestItemFilter(t *testing.T) {
}
*/
