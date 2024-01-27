package storage

import (
	"github.com/RexGreenway/CollectionApp/internal/entities"
)

// Storage defines the methodologies for fetching and editing collection data.
type Storage interface {
	// COLLECTION LEVEL
	// CreateCollection adds the provided collection into storage.
	CreateCollection(entities.Collection) error
	// GetCollection looks up and returns a collection from storage given an ID.
	GetCollection(collectionID string) (entities.Collection, error)
	// UpdateCollection [ID cannot be changed!] ???
	UpdateCollection(entities.Collection) error
	// DeleteCollection ???
	DeleteCollection(collectionID string) error
}
