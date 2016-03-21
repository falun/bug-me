package filter

type Tester interface {
	Test(Item) bool
}

type Matcher interface {
	Match(Item) bool
	Or(Matcher) Matcher
	And(Matcher) Matcher
}

type terminal struct{ Test ItemTest }

type Joined struct {
	Left     Tester
	JoinType CombinationType
	Right    Tester
}

var True struct{}
var False struct{}
var TestNot struct{ Nested ItemTester }

func (_ TestTrue) Test(_ Item) bool  { return true }
func (_ TestFalse) Test(_ Item) bool { return false }
func (tn TestNot) Test(i Item) bool  { return !tn.Nested.Test(i) }

type HasLabelTest struct{ Label string }

func HasLabel(label string) HasLabelTest {
	return HasLabelTest{label}
}

func HasAnyLabel(labels ...string) Matcher {
	if len(labels) == 0 {
		return True
	}

	m := False
	for _, l := range labels {
		m = m.Or(HasLabel(l))
	}

	return m
}

func HasAllLabels(labels ...string) Matcher {
	if len(labels) == 0 {
		return True
	}

	m := True
	for _, l := range labels {
		m.And(HasLabel(l))
	}

	return m
}
