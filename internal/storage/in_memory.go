package storage

import (
	"errors"

	"github.com/RexGreenway/CollectionApp/internal/entities"
	"github.com/RexGreenway/CollectionApp/internal/entities/items"
)

var INMEM StorageType = "in_memory"

// ???
type inMemoryStorage struct {
	store map[string]entities.Collection[items.Item]
}

func newInMemoryStorage() (inMemoryStorage, error) {
	return inMemoryStorage{
		store: map[string]entities.Collection[items.Item]{},
	}, nil
}

func (s inMemoryStorage) CreateCollection(
	collection entities.Collection[items.Item],
) error {
	s.store[collection.ID] = collection
	return nil
}

func (s inMemoryStorage) GetCollection(
	collectionID string,
) (entities.Collection[items.Item], error) {
	if collection, ok := s.store[collectionID]; !ok {
		return entities.Collection[items.Item]{}, errors.New("not found")
	} else {
		return collection, nil
	}
}

func (s inMemoryStorage) UpdateCollection(
	collection entities.Collection[items.Item],
) error {
	if _, ok := s.store[collection.ID]; !ok {
		return errors.New("collection not found")
	}
	s.store[collection.ID] = collection
	return nil
}

func (s inMemoryStorage) DeleteCollection(collectionID string) error {
	if _, ok := s.store[collectionID]; !ok {
		return errors.New("collection not found")
	}
	delete(s.store, collectionID)
	return nil
}
