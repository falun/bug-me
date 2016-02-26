package itemstore

type ItemStore interface {
	List(opts *ListOptions, filters ...ItemFilter) ([]Item, PagingData, error)
	Get(id string) (Item, error)
}

type Pager interface {
	HasNext(PagingData) (bool, error)
	Next(PagingData) ([]Item, error)
	Prev(PagingData) ([]Item, error)
}

type Item interface {
	Id() string
	Labels() ([]Label, error)

	AddLabel(label Label)
	RemoveLabel(label Label)
}

type Label string

type ListOptions struct{}

type ItemFilter interface{}
