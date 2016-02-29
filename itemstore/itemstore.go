package itemstore

type Item interface {
	Id() string
	Labels() ([]Label, error)

	AddLabel(label Label)
	RemoveLabel(label Label)
}

type ItemStore interface {
	Get(id string) (Item, error)
	Add(item Item) error

	List(opts *ListOptions, filters ...ItemFilter) ([]Item, PagingData, error)

	Pager() Pager
}

type Pager interface {
	Next(PagingData) ([]Item, error)
	Prev(PagingData) ([]Item, error)
}

type PagingData struct {
	More bool
	Prev string
	Next string
}

type Label string

type ListOptions struct {
	PageSize int
}

type ItemFilter interface {
	And(ItemFilter) ItemFilter
	Or(ItemFilter) ItemFilter
}
