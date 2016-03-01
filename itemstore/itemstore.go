package itemstore

type ItemStore interface {
	Get(id string) (Item, error)
	Add(item Item) error

	List(opts *ListOptions, filters ...ItemFilter) (Items, PagingData, error)

	Pager() Pager
}

type Items []Item
type Item interface {
	Id() string
	Labels() ([]Label, error)

	AddLabel(label Label)
	RemoveLabel(label Label)

	// special relationships
	Parent() (Item, error)
	Children() (Items, error)
}

type Pager interface {
	Next(PagingData) (Items, error)
	Prev(PagingData) (Items, error)
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
