/*
Package set provides a generic, thread-safe implementation of an unordered set data structure in Go.

An UnorderedSet stores unique elements of any type and supports concurrent access using read-write locks.
It provides basic set operations such as insertion, removal, containment check, and retrieval of all elements.

Key Features:
  - Insert: Add elements to the set. Duplicate insertions are ignored.
  - Remove: Delete elements from the set.
  - Contain: Check if an element exists in the set.
  - Size: Get the number of elements in the set.
  - Clear: Remove all elements from the set.
  - Items: Retrieve all elements in the set as a slice (order not guaranteed).

Concurrency:
  - All operations are safe for concurrent use by multiple goroutines.
*/
package set

import "sync"

// UnorderedSet represents a generic unordered set data structure.
// It stores unique elements and ensures thread-safe operations.
type UnorderedSet[T comparable] struct {
	lockObj sync.RWMutex
	items   map[T]struct{}
}

// NewUnorderedSet creates and returns a new, empty UnorderedSet.
//
// Time Complexity: O(1)
func NewUnorderedSet[T comparable]() *UnorderedSet[T] {
	return &UnorderedSet[T]{items: make(map[T]struct{})}
}

// Insert adds an element to the set. Duplicate insertions are ignored.
// Algorithm: Map insertion ensures uniqueness. Lock acquired for thread-safety.
//
// Time Complexity: O(1) amortized
func (us *UnorderedSet[T]) Insert(item T) {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	us.items[item] = struct{}{}
}

// Remove deletes an element from the set.
// Algorithm: Map deletion removes the key if present. Lock acquired for thread-safety.
//
// Time Complexity: O(1)
func (us *UnorderedSet[T]) Remove(item T) {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	delete(us.items, item)
}

// Contain checks if an element exists in the set.
// Returns true if present, false otherwise.
// Algorithm: Map lookup. Lock acquired for reading.
//
// Time Complexity: O(1)
func (us *UnorderedSet[T]) Contain(item T) bool {
	us.lockObj.RLock()
	defer us.lockObj.RUnlock()
	_, ok := us.items[item]
	return ok
}

// Size returns the number of elements currently in the set.
// Algorithm: Map length retrieval. Lock acquired for reading.
//
// Time Complexity: O(1)
func (us *UnorderedSet[T]) Size() int {
	us.lockObj.RLock()
	defer us.lockObj.RUnlock()
	return len(us.items)
}

// Clear removes all elements from the set, resetting it to empty.
// Algorithm: Reinitialize the internal map. Lock acquired for writing.
//
// Time Complexity: O(1)
func (us *UnorderedSet[T]) Clear() {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	us.items = make(map[T]struct{})
}

// Items return a slice containing all elements in the set.
// The order of elements is not guaranteed.
// Algorithm: Iterate over the map keys and append to a slice. Lock acquired for writing.
//
// Time Complexity: O(n), where n = number of elements in the set
func (us *UnorderedSet[T]) Items() []T {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	elements := make([]T, 0, len(us.items))
	for element := range us.items {
		elements = append(elements, element)
	}
	return elements
}

// Iter returns a channel that streams elements of the set.
// It captures a snapshot at the time of the call, so later modifications
// to the set will not affect the iteration.
func (us *UnorderedSet[T]) Iter() <-chan T {
	ch := make(chan T)

	go func() {
		us.lockObj.RLock()
		items := make([]T, 0, len(us.items))
		for item := range us.items {
			items = append(items, item)
		}
		us.lockObj.RUnlock()

		for _, item := range items {
			ch <- item
		}
		close(ch)
	}()

	return ch
}
