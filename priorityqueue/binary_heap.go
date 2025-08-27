package priorityqueue

import (
	"errors"
	"golang.org/x/exp/constraints"
)

var i any = -1

// BinaryHeap represents a generic binary min-heap implementation.
// It stores elements in a slice and ensures the smallest element
// is always at the root (index 0).
// T must satisfy constraints.Ordered, meaning the type supports
// comparison operators (<, >, etc.).
type BinaryHeap[T constraints.Ordered] struct {
	data []T
}

// NewBinaryHeap creates and returns a new empty BinaryHeap.
func NewBinaryHeap[T constraints.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{data: make([]T, 0)}
}

// IsEmpty checks if the heap is empty.
func (bh *BinaryHeap[T]) IsEmpty() bool {
	return bh.Size() == 0
}

// Clear removes all elements from the heap by setting the internal slice to nil.
func (bh *BinaryHeap[T]) Clear() {
	bh.data = nil
}

// Size returns the number of elements currently in the heap.
func (bh *BinaryHeap[T]) Size() int {
	return len(bh.data)
}

// Peek returns the smallest element in the heap without removing it.
// If the heap is empty, an error is returned.
func (bh *BinaryHeap[T]) Peek() (T, error) {
	if bh.IsEmpty() {
		return i.(T), errors.New("heap empty")
	}
	return bh.data[0], nil
}

// Poll removes and returns the smallest element (root) from the heap.
// If the heap is empty, an error is returned.
func (bh *BinaryHeap[T]) Poll() (T, error) {
	if bh.IsEmpty() {
		return i.(T), errors.New("heap empty")
	}
	return bh.removeAt(0)
}

// removeAt removes and returns the element at index k, then re-heapifies the tree.
// Internal method used by Poll. Returns an error if the heap is empty.
func (bh *BinaryHeap[T]) removeAt(k int) (T, error) {
	first := bh.data[k]
	last := bh.data[bh.Size()-1]
	bh.data[0] = last
	if bh.Size() > 0 {
		bh.data = bh.data[:bh.Size()-1]
	}

	parent := 0
	child := 2 * (parent + 1)
	for child < bh.Size()-1 {
		if bh.data[child+1] < bh.data[child] {
			child = child + 1
		}
		if bh.data[child] < bh.data[parent] {
			bh.swap(child, parent)
			parent = child
			child = 2 * child
		} else {
			break
		}
	}

	return first, nil
}

// Add inserts a new element into the heap and maintains the heap property.
func (bh *BinaryHeap[T]) Add(val T) {
	bh.data = append(bh.data, val)
	idxOfLastElem := bh.Size() - 1
	bh.swim(idxOfLastElem)
}

// swap exchanges the elements at indexes i and j.
func (bh *BinaryHeap[T]) swap(i, j int) {
	temp := bh.data[i]
	bh.data[i] = bh.data[j]
	bh.data[j] = temp
}

// swim moves the element at index k up the heap until the heap property is restored.
// This happens when a newly inserted element is smaller than its parent.
func (bh *BinaryHeap[T]) swim(k int) {
	parent := (k - 1) / 2
	// compare with parent if it's less then swap
	for k > 0 && bh.data[parent] > bh.data[k] {
		bh.swap(parent, k)
		k = parent
		parent = (k - 1) / 2
	}
}
