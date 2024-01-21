package storage

import (
	"github.com/RexGreenway/CollectionApp/internal/entities"
	"github.com/RexGreenway/CollectionApp/internal/entities/items"
)

// StorageType defines
type StorageType string

// Storage defines the methodologies for fetching and editing collection data.
type Storage interface {
	// COLLECTION LEVEL
	// CreateCollection adds the provided collection into storage.
	CreateCollection(entities.Collection[items.Item]) error
	// GetCollection looks up and returns a collection from storage given an ID.
	GetCollection(collectionID string) (entities.Collection[items.Item], error)
	// UpdateCollection [ID cannot be changed!] ???
	UpdateCollection(entities.Collection[items.Item]) error
	// DeleteCollection ???
	DeleteCollection(collectionID string) error
}

// StorageManager returns the Storage implementation
func StorageManager(val StorageType) (Storage, error) {
	switch val {
	case "in_memory":
		return newInMemoryStorage()

	// Defaults to in memory
	default:
		return newInMemoryStorage()
	}
}
