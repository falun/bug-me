package itemstore

type CombinationType int

const (
	And CombinationType = iota
	Or
)

type ItemTestFn func(Item) bool

type ItemFilter struct {
	Next            *ItemFilter
	CombinationType CombinationType
	test            ItemTestFn
}

func NewFilter(f ItemTestFn) ItemFilter {
	return ItemFilter{test: f}
}

func (f ItemFilter) And(i ItemFilter) ItemFilter {
	f.Next = &i
	f.CombinationType = And
	return f
}

func (f ItemFilter) Or(i ItemFilter) ItemFilter {
	f.Next = &i
	f.CombinationType = Or
	return f
}

// does this work for groups?
func (f ItemFilter) Match(i Item) bool {
	cur := f.test(i)

	if f.Next == nil {
		return cur
	}

	if f.CombinationType == And {
		return cur && f.Next.Match(i)
	}

	if f.CombinationType == Or {
		return cur || f.Next.Match(i)
	}

	panic("Unhandled filter combination type")
}
