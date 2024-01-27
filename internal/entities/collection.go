package entities

import (
	"errors"
	"fmt"

	"github.com/RexGreenway/CollectionApp/internal/entities/items"
)

// Collection defines a container for "collectable" Items with a unique ID and
// a name. Collections only contain Items of a single type stored as a map of
// IDs to Items.
//
// Note: Multi-'item type' collections could be possible in the future.
type Collection struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// Contained Item related fields
	Items    map[string]items.Item `json:"items"`
	ItemType items.ItemType        `json:"item_type"`
}

// String implements the stringer interface returning the Collection name.
func (c Collection) String() string {
	return fmt.Sprintf("%s Collection", c.Name)
}

// GetItem returns a "collectable" from a collection if it exists.
func (c Collection) GetItem(itemID string) (items.Item, error) {
	item, ok := c.Items[itemID]
	if !ok {
		// TODO: define custom errors.
		return nil, errors.New("item not found")
	}
	return item, nil
}

// AddItem adds a non-preexisting "collectible" Item to a collection.
func (c *Collection) AddItem(item items.Item) error {
	itemID := item.GetID()
	if _, ok := c.Items[itemID]; ok {
		return errors.New("item already exists")
	}
	c.Items[itemID] = item
	return nil
}
