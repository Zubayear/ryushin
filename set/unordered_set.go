package set

import "sync"

// UnorderedSet is a generic unordered set implementation with thread safety.
type UnorderedSet struct {
	lockObj sync.RWMutex
	items   map[interface{}]struct{}
}

// NewUnorderedSet creates a new instance of the UnorderedSet.
func NewUnorderedSet() *UnorderedSet {
	return &UnorderedSet{items: make(map[interface{}]struct{})}
}

// Insert an item to the set
func (us *UnorderedSet) Insert(item interface{}) {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	us.items[item] = struct{}{}
}

// Remove removes an item from set
func (us *UnorderedSet) Remove(item interface{}) {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	delete(us.items, item)
}

// Contain checks if an item is present in set
func (us *UnorderedSet) Contain(item interface{}) bool {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	_, ok := us.items[item]
	return ok
}

// Size return the size of the set
func (us *UnorderedSet) Size() int {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	return len(us.items)
}

// Clear removes all elements from the set.
func (us *UnorderedSet) Clear() {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	us.items = make(map[interface{}]struct{})
}

// Items returns a slice containing all elements in the set.
func (us *UnorderedSet) Items() []interface{} {
	us.lockObj.Lock()
	defer us.lockObj.Unlock()
	elements := make([]interface{}, 0, len(us.items))
	for element := range us.items {
		elements = append(elements, element)
	}
	return elements
}
