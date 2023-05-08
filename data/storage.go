package data

// StorageInterface is a general persistent storage interface
type StorageInterface interface {
	// Init initialisation of storage
	Init() error
	// Get returns last value
	Get() (float64, error)
	// Set sets new (next) value
	Set(value float64) error
	// Persist performs data persistence
	Persist() error
	// BeginTx begins 'transaction' (locks mutex)
	BeginTx() error
	// EndTx ends 'transaction' (unlocks mutex)
	EndTx() error
}
