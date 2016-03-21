package itemstore

type ItemTester interface {
	Test(Item) bool
}

type Matcher interface {
	Match(Item) bool
	Or(Matcher) Matcher
	And(Matcher) Matcher
}

type terminal struct{ Test ItemTest }

type Joined struct {
	Left     ItemTester
	JoinType CombinationType
	Right    ItemTester
}

var True struct{}
var False struct{}
var TestNot struct{ Nested ItemTester }

func (_ TestTrue) Test(_ Item) bool  { return true }
func (_ TestFalse) Test(_ Item) bool { return false }
func (tn TestNot) Test(i Item) bool  { return !tn.Nested.Test(i) }
