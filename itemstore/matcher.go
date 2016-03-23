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