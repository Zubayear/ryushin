/*
Package priorityqueue provides a generic, thread-safe binary max-heap(default) implementation in Go.

A BinaryHeap is a priority queue where the smallest element is always at the root.
It supports insertion, retrieval of the minimum element, and removal while maintaining
the heap property.

The type parameter T must satisfy constraints.Ordered (supports <, > operators).

Concurrency:
  - All operations are protected by a read-write mutex and safe for concurrent access.

Key Features:
  - Add: Insert a new element while maintaining the heap property (O(log n)).
  - Peek: Retrieve the smallest element without removing it (O(1)).
  - Poll: Remove and return the smallest element, re-heapifying the structure (O(log n)).
  - IsEmpty: Check if the heap is empty (O(1)).
  - Size: Return the number of elements in the heap (O(1)).
  - Clear: Remove all elements from the heap (O(1)).

Algorithm Notes:
  - Binary Heap is stored in a slice.
  - Parent and child relationships:
    parent index = (k-1)/2
    left child = 2*k + 1, right child = 2*k + 2
  - Swim operation: Moves a newly added element up until the heap property is restored.
  - RemoveAt operation: Replaces the removed element with the last element, then sinks it down.
*/
package priorityqueue

import (
	"errors"
	"sync"

	"golang.org/x/exp/constraints"
)

// BinaryHeap is a generic, thread-safe binary heap implementation.
//
// It supports both min-heap and max-heap behavior depending on the comparator
// function provided during construction.
//
// Internally, the heap is stored as a slice representing a complete binary tree.
//
// Array-based heap indexing rules:
//   - Root element: index 0
//   - For a node at index i:
//     Left child: 2*i + 1
//     Right child: 2*i + 2
//     Parent: (i - 1) / 2
//
// Thread-safety:
//
//	All operations on the heap are protected with a read-write mutex (RWMutex),
//	making it safe for concurrent access.
//
// Fields:
//   - data: slice of elements stored in heap order
//   - cmp: comparator function used to maintain heap property
//     (should return true if the first element has higher priority than the second)-
//   - mutex: RWMutex to ensure safe concurrent access
type BinaryHeap[T any] struct {
	data  []T               // slice storing heap elements
	cmp   func(a, b T) bool // comparator defining heap ordering
	mutex sync.RWMutex      // protects heap for concurrent access
}

// NewBinaryHeap creates a new BinaryHeap instance using the natural ordering of T.
//
// By default, this creates a `max-heap`, where the element with the largest value
// is at the root. It uses the built-in comparison operators of T (constraints.Ordered).
//
// Notes:
//   - For numeric types (int, float, etc.), the largest value will have the highest priority.
//   - For strings, lexicographically larger strings will have higher priority.
//   - For min-heap behavior, you can either:
//   - Provide negative values for numeric types, or
//   - Use NewBinaryHeapWithComparator with a custom comparator.
//
// Example usage:
//
//	// Max-heap of integers
//	h := NewBinaryHeap[int]()
//	h.Add(5)
//	h.Add(10)
//	h.Add(3)
//
//	// Polling repeatedly will give: 10, 5, 3
//
//	// Max-heap of strings (lexicographically the largest first)
//	sh := NewBinaryHeap[string]()
//	sh.Add("apple")
//	sh.Add("banana")
//	sh.Add("cherry")
//
//	// Polling repeatedly will give: "cherry", "banana", "apple"
func NewBinaryHeap[T constraints.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{
		data: make([]T, 0),
		cmp: func(a, b T) bool {
			return a > b
		},
	}
}

// NewBinaryHeapWithComparator creates and returns a new empty BinaryHeap
// with a custom comparator function.
//
// Parameters:
//
//	cmp: A function of type `func(a, b T) bool` that defines the heap ordering.
//	     - Should return `true` if element `a` has higher priority than `b`.
//	     - This allows you to define min-heaps, max-heaps, or custom ordering
//	       based on any field or combination of fields in T.
//
// Returns:
//
//	A pointer to an empty BinaryHeap[T] that uses the provided comparator.
//
// Example usage:
//
//	type Person struct {
//	    Name string
//	    Age  uint
//	}
//
//	// Max-heap: higher Age first, tie-breaker: longer Name
//	bh: = NewBinaryHeapWithComparator[Person](func(p1, p2 Person) bool {
//	    if p1.Age != p2.Age {
//	        return p1.Age > p2.Age
//	    }
//	    return len(p1.Name) > len(p2.Name)
//	})
//
//	bh.Add(Person{"Alice", 30})
//	bh.Add(Person{"Bob", 25})
func NewBinaryHeapWithComparator[T any](cmp func(a, b T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		data: make([]T, 0),
		cmp:  cmp,
	}
}

// IsEmpty checks whether the heap contains any elements.
//
// Returns:
//   - true if the heap is empty, false otherwise.
//
// Complexity: O(1)
func (bh *BinaryHeap[T]) IsEmpty() bool {
	bh.mutex.RLock()
	defer bh.mutex.RUnlock()
	return bh.Size() == 0
}

// Clear removes all elements from the heap.
//
// After calling Clear, the heap will be empty.
//
// Complexity: O(1)
func (bh *BinaryHeap[T]) Clear() {
	bh.mutex.Lock()
	defer bh.mutex.Unlock()
	bh.data = nil
}

// Size returns the number of elements currently stored in the heap.
//
// Complexity: O(1)
func (bh *BinaryHeap[T]) Size() int {
	bh.mutex.RLock()
	defer bh.mutex.RUnlock()
	return len(bh.data)
}

// Peek returns the root element of the heap without removing it.
// The root is either the minimum or maximum element based on the comparator.
//
// Returns:
//   - the root element
//   - error if the heap is empty
//
// Complexity: O(1)
func (bh *BinaryHeap[T]) Peek() (T, error) {
	var zero T
	bh.mutex.RLock()
	defer bh.mutex.RUnlock()
	if len(bh.data) == 0 {
		return zero, errors.New("heap empty")
	}
	return bh.data[0], nil
}

// Poll removes and returns the root element of the heap.
// The root is either the minimum or maximum element based on the comparator.
//
// Returns:
//   - the root element
//   - error if the heap is empty
//
// Complexity: O(log n) due to re-heapification
func (bh *BinaryHeap[T]) Poll() (T, error) {
	var zero T
	bh.mutex.Lock()
	defer bh.mutex.Unlock()
	if len(bh.data) == 0 {
		return zero, errors.New("heap empty")
	}
	return bh.removeAt(0) // we can only remove the root
}

// removeAt removes the element at index k from the heap and returns it.
//
// Steps:
//  1. Replace the element at index k with the last element in the heap.
//  2. Remove the last element from the slice.
//  3. Re-heapify by comparing the new root with its children using the comparator:
//     - Select the smaller child (based on the comparator).
//     - Swap with the parent if the child violates the heap property.
//  4. Continue until the heap property is restored.
//
// Returns:
//   - the removed element
//   - error if the heap is empty
//
// Note: This is an internal helper method, primarily used by Poll.
//
// Complexity: O(log n)
func (bh *BinaryHeap[T]) removeAt(k int) (T, error) {
	size := len(bh.data)
	if size == 0 {
		var zero T
		return zero, errors.New("heap empty")
	}
	removed := bh.data[k]
	last := bh.data[size-1]
	bh.data[k] = last
	bh.data = bh.data[:size-1]

	parent := k
	child := 2*parent + 1
	for child < len(bh.data) {
		// pick the child with higher priority according to comparator
		// ex: min-heap -> compare left and right child
		// ex: for min-heap if the right < left, then use that
		if child+1 < len(bh.data) && bh.cmp(bh.data[child+1], bh.data[child]) {
			child = child + 1
		}
		// compare parent and child
		// if child has higher priority than parent, swap
		// ex: for min-heap if child < parent then interchange
		if bh.cmp(bh.data[child], bh.data[parent]) {
			bh.swap(child, parent)
			parent = child
			child = 2*parent + 1
		} else {
			break
		}
	}

	return removed, nil
}

// Add inserts a new element into the heap and restores the heap property.
//
// Parameters:
//   - val: the value to be added to the heap
//
// Complexity: O(log n)
func (bh *BinaryHeap[T]) Add(val T) {
	bh.mutex.Lock()
	defer bh.mutex.Unlock()
	bh.data = append(bh.data, val)
	idxOfLastElem := len(bh.data) - 1
	bh.swim(idxOfLastElem)
}

// Swap exchanges the elements at indexes i and j.
//
// Complexity: O(1)
func (bh *BinaryHeap[T]) swap(i, j int) {
	bh.data[i], bh.data[j] = bh.data[j], bh.data[i]
}

// Swim moves the element at index k up the heap until the heap property is satisfied.
//
// This is used after adding a new element to restore heap order.
//
// Complexity: O(log n)
func (bh *BinaryHeap[T]) swim(k int) {
	for k > 0 {
		parent := (k - 1) / 2
		// compare with parent
		// if it returns true i.e., for min-heap k < parent; then we move the k
		if bh.cmp(bh.data[k], bh.data[parent]) {
			bh.swap(k, parent)
			k = parent
		} else {
			break
		}
	}
}

// Sort returns a slice of all elements in the heap in order according to the heap's comparator.
// The original heap remains intact; this operation does not modify bh.
//
// Implementation details:
//  1. Creates a copy of the current heap's internal slice to avoid mutating the original heap.
//  2. Constructs a temporary BinaryHeap with the same comparator using the copied data.
//  3. Repeatedly polls the temporary heap to extract elements in sorted order.
//  4. Appends each polled element to the result slice.
//
// Complexity: O(n log n) because each Poll operation takes O(log n) and we perform n polls.
// Returns: a slice of elements sorted according to the heap's comparator.
func (bh *BinaryHeap[T]) Sort() []T {
	bh.mutex.RLock()
	defer bh.mutex.RUnlock()
	size := len(bh.data)
	copyHeap := make([]T, size)
	copy(copyHeap, bh.data)

	result := make([]T, 0, size)
	tempHeap := &BinaryHeap[T]{
		data: copyHeap,
		cmp:  bh.cmp,
	}

	for i := 0; i < size; i++ {
		v, _ := tempHeap.Poll()
		result = append(result, v)
	}
	return result
}
