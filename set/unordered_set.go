package set

import "sync"

// UnorderedSet represents a generic unordered set data structure.
// It supports thread-safe operations and stores unique elements of any type.
type UnorderedSet struct {
	lockObj sync.RWMutex
	items   map[any]struct{}
}

// NewUnorderedSet creates and returns a new instance of UnorderedSet.
func NewUnorderedSet() *UnorderedSet {
	return &UnorderedSet{items: make(map[any]struct{})}
}

// Insert adds an item to the set.
// If the item already exists, the set remains unchanged.
func (us *UnorderedSet) Insert(item any) {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	us.items[item] = struct{}{}
}

// Remove deletes an item from the set.
// If the item does not exist, the set remains unchanged.
func (us *UnorderedSet) Remove(item any) {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	delete(us.items, item)
}

// Contain checks whether the specified item exists in the set.
// Returns true if the item is present, false otherwise.
func (us *UnorderedSet) Contain(item any) bool {
	us.lockObj.RLock()
	defer us.lockObj.RUnlock()
	_, ok := us.items[item]
	return ok
}

// Size returns the number of elements currently in the set.
func (us *UnorderedSet) Size() int {
	us.lockObj.RLock()
	defer us.lockObj.RUnlock()
	return len(us.items)
}

// Clear removes all elements from the set, resetting it to empty.
func (us *UnorderedSet) Clear() {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	us.items = make(map[any]struct{})
}

// Items returns a slice containing all elements currently in the set.
// The order of elements is not guaranteed.
func (us *UnorderedSet) Items() []any {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	elements := make([]any, 0, len(us.items))
	for element := range us.items {
		elements = append(elements, element)
	}
	return elements
}
