package storage

import (
	"github.com/RexGreenway/CollectionApp/internal/entities"
	"github.com/RexGreenway/CollectionApp/internal/entities/items"
)

// jsonStorage caches the most recent results in mem?
// probably not a good idea
type jsonStorage struct {
	filePath string
}

func NewJSONStorage(filePath string) jsonStorage {
	return jsonStorage{
		filePath: filePath,
	}
}

func (s jsonStorage) GetCollection(
	collectionID string,
) (entities.Collection[items.Item], error) {
	// read the json

	// How to unmarshal the items into the correct type?

	// Return collection object
	return entities.Collection[items.Item]{
		ID:   collectionID,
		Name: "test",

		Items: map[string]items.Item{},
	}, nil
}
