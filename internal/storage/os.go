package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/RexGreenway/CollectionApp/internal/entities"
)

const OS StorageType = "os"

// osStorage ???
type osStorage struct {
	filePath string
	store    map[string]entities.Collection
}

func NewOSStorage(filePath string) (osStorage, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic("failed to read os storage file path")
	}
	decoder := json.NewDecoder(f)

	var store map[string]entities.Collection
	decoder.Decode(&store)

	return osStorage{
		filePath: filePath,
		store:    store,
	}, nil
}

// CreateCollection ???
func (s osStorage) CreateCollection(collection entities.Collection) error {
	if _, ok := s.store[collection.ID]; ok {
		return errors.New("collection already exists")
	}
	s.store[collection.ID] = collection
	return nil
}

// GetCollection ???
func (s osStorage) GetCollection(collectionID string) (entities.Collection, error) {
	collection, ok := s.store[collectionID]
	if !ok {
		return entities.Collection{}, errors.New("collection not found")
	}
	return collection, nil
}
