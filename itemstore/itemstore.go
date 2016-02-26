package itemstore

type ItemStore interface {
	List(filters ...ItemFilter) ([]Item, error)
	Get(id string) (Item, error)
}

type Item interface {
	Id() string
	Labels() ([]Label, error)

	AddLabel(label Label)
	RemoveLabel(label Label)
}

type Label string

type ItemFilter interface {
}
