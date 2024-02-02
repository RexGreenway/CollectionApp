package entities

// ItemType establishes the subset of strings that an Item can be.
type ItemType string

// Pre-defined Item Types
const (
	ITEM_UNSPECIFIED ItemType = "Unspecified"
	ITEM_FILM        ItemType = "Film"
)

// Item defines a collectable.
type Item struct {
	ID   string            `json:"id"`
	Name string            `json:"name"`
	Type ItemType          `json:"item_type,omitempty"`
	Info map[string]string `json:"item_info,omitempty"`
}
