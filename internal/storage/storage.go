package storage

// StorageType defines
type StorageType string

// StorageManager returns the Storage implementation
func StorageManager(val StorageType) (Storage, error) {
	switch val {
	case InMemory:
		return NewInMemoryStorage()

	// case OS:
	// 	return NewOSStorage("test-path")

	// Defaults to in memory
	default:
		return NewInMemoryStorage()
	}
}
