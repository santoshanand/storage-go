package main

import "fmt"

// Result -
type Result any // we can change any to interface{}
// Input -
type Input Result

// Storage -
type Storage map[string]Result

// NewStorage -
func NewStorage() Storage {
	return Storage{}
}

// Get - used to get value
func (s Storage) Get(key string) (Result, bool) {
	val, ok := s[key]
	return val, ok
}

// Set - used to set value
func (s Storage) Set(key string, value Input) {
	s[key] = value
}

// Delete - delete value by key
func (s Storage) Delete(key string) {
	delete(s, key)
}

// Keys - return all the keys
func (s Storage) Keys() []string {
	var keys []string
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}

// Contains - used to check storage has value or not by key
func (s Storage) Contains(key string) bool {
	_, ok := s[key]
	return ok
}
func main() {
	storage := NewStorage()
	storage.Set("1", "Orange")
	storage.Set("3", "Red")
	storage.Delete("1")
	fmt.Println(storage.Keys())
}
