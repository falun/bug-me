package itemstore

import (
	"fmt"
)

// CombinationType describes ways in which Matchers can be logically joined.
type CombinationType int

const (
	And CombinationType = iota
	Or
)

// Matcher is the building block of a complex check.
type Matcher interface {
	Match(Item) bool
	Or(Matcher) Matcher
	And(Matcher) Matcher
}

// Joined is an implementation af a Matcher.
type Joined struct {
	Left     Matcher
	JoinType CombinationType
	Right    Matcher
}

func (j Joined) And(m Matcher) Matcher { return Joined{j, And, m} }
func (j Joined) Or(m Matcher) Matcher  { return Joined{j, Or, m} }
func (j Joined) Match(i Item) bool {
	switch j.JoinType {
	case And:
		return j.Left.Match(i) && j.Right.Match(i)
	case Or:
		return j.Left.Match(i) || j.Right.Match(i)
	default:
		panic(fmt.Sprintf("bad combination type: %d", j.JoinType))
	}
}

type Tester interface {
	Test(Item) bool
}

type Leaf struct {
	Test Tester
}

func (l Leaf) Match(i Item) bool     { return l.Test.Test(i) }
func (l Leaf) And(m Matcher) Matcher { return Joined{l, And, m} }
func (l Leaf) Or(m Matcher) Matcher  { return Joined{l, Or, m} }

func MatcherFromTester(t Tester) Matcher { return Leaf{t} }

//

/*
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
*/
