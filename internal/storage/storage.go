package storage

import "fmt"

// StorageType defines
type StorageType string

// StorageManager returns the Storage implementation
func StorageManager(impl StorageType) (Storage, error) {
	switch impl {
	// case InMemory:
	// 	return newInMemoryStorage()

	case OS:
		return newOSStorage()

	// Defaults to in memory
	default:
		return nil, fmt.Errorf("no such storage implementation %q", impl)
	}
}
