package entities

import (
	"errors"
	"fmt"

	"github.com/RexGreenway/CollectionApp/internal/entities/items"
)

// NOTE: Does the concept of a collection need an interface?

// Collection defines a struct defining a collection of items of the same type.
type Collection[I items.Item] struct {
	// Collection related fields
	ID   string `json:"id"`
	Name string `json:"name"`

	// Contained Item related fields
	Items    map[string]I   `json:"items"`
	ItemType items.ItemType `json:"item_type"`
}

// String implements the stringer interface.
func (c Collection[I]) String() string {
	return fmt.Sprintf("Collection: %v", c.Name)
}

// ??? Getters should not be on pointers and return errors
// (Should define custom errors too)
func (c Collection[I]) GetItem(itemID string) (item I, err error) {
	item, ok := c.Items[itemID]
	if !ok {
		err = errors.New("item not found")
	}
	return
}

// ??? Setters should be on pointers
func (c *Collection[I]) Add(items ...I) error {
	return nil
}
