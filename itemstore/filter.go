package itemstore

type CombinationType int

const (
	And CombinationType = iota
	Or
)

type ItemTestFn func(Item) bool

type Matcher interface {
	Match(Item) bool
	Or(Matcher) Matcher
	And(Matcher) Matcher
}

var LeafTrue = Leaf{func(_ Item) bool { return true }}
var LeafFalse = Leaf{func(_ Item) bool { return false }}

type Leaf struct {
	TestFn ItemTestFn
}

type Complex struct {
	Left     Matcher
	JoinType CombinationType
	Right    Matcher
}

func NewFilter(f ItemTestFn) Matcher {
	return Leaf{f}
}

func (l Leaf) Match(i Item) bool {
	return l.TestFn(i)
}

func (l Leaf) And(r Matcher) Matcher {
	return Complex{l, And, r}
}

func (l Leaf) Or(r Matcher) Matcher {
	return Complex{l, Or, r}
}

func (c Complex) Match(i Item) bool {
	if c.JoinType == And {
		return c.Left.Match(i) && c.Right.Match(i)
	} else if c.JoinType == Or {
		return c.Left.Match(i) || c.Right.Match(i)
	}

	panic("Unhandled filter combination type")
}

func (c Complex) And(r Matcher) Matcher {
	return Complex{c, And, r}
}

func (c Complex) Or(r Matcher) Matcher {
	return Complex{c, Or, r}
}
