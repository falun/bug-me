// The filter package collects a set of basic filtering options that can be
// used to make subselection over the set of items on the todo list.
//
// The intent is that each additional backend should handle these as they see
// fit but to support both pre and post query processing. Post query filtering
// is trivial and can be done by
package filter

import "github.com/falun/bug-me/itemstore"

var True = itemstore.MatcherFromTester(TestTrue{})
var False = itemstore.MatcherFromTester(TestFalse{})

func Not(m itemstore.Matcher) itemstore.Matcher {
	return itemstore.MatcherFromTester(TestNot{m})
}

type TestTrue struct{}
type TestFalse struct{}
type TestNot struct{ Nested itemstore.Matcher }

func (_ TestTrue) Test(_ itemstore.Item) bool  { return true }
func (_ TestFalse) Test(_ itemstore.Item) bool { return false }
func (tn TestNot) Test(i itemstore.Item) bool  { return !tn.Nested.Match(i) }

type HasLabelTest struct{ Label string }
type HasAnyLabelTest struct{ Labels []string }
type HasAllLabelTest struct{ Labels []string }

func HasLabel(label string) itemstore.Matcher {
	return itemstore.MatcherFromTester(HasLabelTest{label})
}

func (t HasLabelTest) Test(i itemstore.Item) bool {
	l := i.Labels()
	return l[t.Label]
}

func HasAnyLabel(labels ...string) itemstore.Matcher {
	return itemstore.MatcherFromTester(HasAnyLabelTest{labels})
}

func (t HasAnyLabelTest) Test(i itemstore.Item) bool {
	if len(t.Labels) == 0 {
		return true
	}

	labels := i.Labels()
	for _, l := range t.Labels {
		if labels[l] {
			return true
		}
	}

	return false
}

func HasAllLabels(labels ...string) itemstore.Matcher {
	return itemstore.MatcherFromTester(HasAllLabelTest{labels})
}

func (t HasAllLabelTest) Test(i itemstore.Item) bool {
	if len(t.Labels) == 0 {
		return true
	}

	labels := i.Labels()
	for _, l := range t.Labels {
		if !labels[l] {
			return false
		}
	}

	return true
}
