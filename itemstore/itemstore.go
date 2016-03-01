package itemstore

type ItemID string
type ItemIDs []ItemID

type Items []Item
type Item interface {
	Id() string
	Labels() ([]Label, error)

	AddLabel(label Label)
	RemoveLabel(label Label)

	// special relationships and attributes

	Parent() (ItemID, error)
	Children() (ItemIDs, error)
	Priority() int
}

type ItemStore interface {
	Add(item Item) error
	Get(id ItemID) (Item, error)
	Update(newItem ItemID) error
	Remove(id ItemID) error

	List(opts *ListOptions, filter ItemFilter) (Items, PagingData, error)

	Pager() Pager
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
