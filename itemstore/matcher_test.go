package itemstore

import (
	"testing"
)

type TestTrue struct{}
type TestFalse struct{}

func (_ TestTrue) Test(_ Item) bool  { return true }
func (_ TestFalse) Test(_ Item) bool { return false }

var True Matcher = Leaf{TestTrue{}}
var False Matcher = Leaf{TestFalse{}}

func TestItemFilterSimpleAnd(t *testing.T) {
	f := True
	f = f.And(True)
	if !f.Match(nil) {
		t.Errorf("true && true must be true")
	}

	f = False
	f = f.And(True)
	if f.Match(nil) {
		t.Errorf("false && true must be false")
	}

	f = True
	f = f.And(False)
	if f.Match(nil) {
		t.Errorf("true && false must be false")
	}

	f = False
	f = f.And(False)
	if f.Match(nil) {
		t.Errorf("false && false must be false")
	}
}

func TestItemFilterSimpleOr(t *testing.T) {
	f := True
	f = f.Or(True)
	if !f.Match(nil) {
		t.Errorf("true || true must be true")
	}

	f = False
	f = f.Or(True)
	if !f.Match(nil) {
		t.Errorf("false || true must be true")
	}

	f = True
	f = f.Or(False)
	if !f.Match(nil) {
		t.Errorf("true || false must be true")
	}

	f = False
	f = f.Or(False)
	if f.Match(nil) {
		t.Errorf("false || false must be false")
	}
}

func TestItemFilterGroupRight(t *testing.T) {
	// true && (false || true)
	f := True
	f2 := False
	f2 = f2.Or(True)
	f = f.And(f2)
	if !f.Match(nil) {
		t.Errorf("true && (false || true) must be true")
	}
}

func TestItemFilterGroupLeft(t *testing.T) {
	// (false || true) && true
	f := False
	f = f.Or(True)
	f2 := True
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
