package deque

import "github.com/Zubayear/ryushin/linkedlist"

// Deque is a generic double-ended queue backed by a doubly linked structure.
// It supports adding, removing, and peeking elements from both ends in O(1) time.
// Concurrency: All methods on Deque are safe for concurrent use.
type Deque[T comparable] struct {
	data *linkedlist.DoublyLinkedList[T]
}

// NewDeque returns a new, empty Deque[T] backed by a doubly linked list.
// The returned deque is ready to use immediately.
// Time Complexity: O(1)
// Concurrency: Safe for concurrent use by multiple goroutines.
func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{
		data: linkedlist.NewLinkedList[T](),
	}
}

// OfferFirst inserts the specified element at the front of the deque.
// Returns true on success along with a nil error.
// Time Complexity: O(1)
func (d *Deque[T]) OfferFirst(elem T) (bool, error) {
	return d.data.AddFirst(elem)
}

// PollFirst retrieves and removes the first element of the deque.
// Returns the zero value of T and a non-nil error if the deque is empty.
// Time Complexity: O(1)
func (d *Deque[T]) PollFirst() (T, error) {
	return d.data.RemoveFirst()
}

// PeekFirst retrieves, but does not remove, the first element of the deque.
// Returns the zero value of T and a non-nil error if the deque is empty.
// Time Complexity: O(1)
func (d *Deque[T]) PeekFirst() (T, error) {
	return d.data.PeekFirst()
}

// OfferLast inserts the specified element at the end of the deque.
// Returns true on success along with a nil error.
// Time Complexity: O(1)
func (d *Deque[T]) OfferLast(elem T) (bool, error) {
	return d.data.AddLast(elem)
}

// PollLast retrieves and removes the last element of the deque.
// Returns the zero value of T and a non-nil error if the deque is empty.
// Time Complexity: O(1)
func (d *Deque[T]) PollLast() (T, error) {
	return d.data.RemoveLast()
}

// PeekLast retrieves, but does not remove, the last element of the deque.
// Returns the zero value of T and a non-nil error if the deque is empty.
// Time Complexity: O(1)
func (d *Deque[T]) PeekLast() (T, error) {
	return d.data.PeekLast()
}

// Remove deletes the first occurrence of the specified element from the deque.
// It returns true if an element was removed, or false otherwise.
// Time Complexity: O(n)
func (d *Deque[T]) Remove(elem T) bool {
	ok, err := d.data.Remove(elem)
	if err != nil {
		return false
	}
	return ok == elem
}

// Size returns the number of elements in the deque.
// Time Complexity: O(1)
func (d *Deque[T]) Size() int {
	return d.data.Size()
}

// IsEmpty reports whether the deque has no elements.
// Time Complexity: O(1)
func (d *Deque[T]) IsEmpty() bool {
	return d.data.IsEmpty()
}
