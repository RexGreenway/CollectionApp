package items

// ItemType establishes the subset of strings that an Item can be.
type ItemType string

// Item defines the required methods for collectable items.
type Item interface {
	GetID() string
	GetName() string
	GetType() ItemType
}
