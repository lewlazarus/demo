package data

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

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

const FileLocation = "./current_value.txt"

// FileStorage is an implementation of StorageInterface which uses
// local file for data persistence. For simplicity, file location is
// hardcoded as FileLocation constant
type FileStorage struct {
	value float64
	mu    sync.Mutex
}

func NewFileStorage() *FileStorage {
	return &FileStorage{
		value: 1,
		mu:    sync.Mutex{},
	}
}

func (r *FileStorage) Init() error {
	if _, err := os.Stat(FileLocation); err != nil {
		r.value = 0
		return nil
	}

	data, err := os.ReadFile(FileLocation)
	if err != nil {
		return err
	}

	res, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}

	r.value = res
	return nil
}

func (r *FileStorage) BeginTx() error {
	r.mu.Lock()
	return nil
}

func (r *FileStorage) EndTx() error {
	r.mu.Unlock()
	return nil
}

func (r *FileStorage) Get() (float64, error) {
	return r.value, nil
}

func (r *FileStorage) Set(value float64) error {
	r.value = value
	return nil
}

func (r *FileStorage) Persist() error {
	str := fmt.Sprintf("%f", r.value)
	if err := os.WriteFile(FileLocation, []byte(str), 0777); err != nil {
		return err
	}

	return nil
}
