/*
Package deque provides a generic, thread-safe double-ended queue (Deque) implementation in Go.

A Deque (double-ended queue) allows insertion, removal, and retrieval of elements
from both ends with O(1) complexity for operations at the front or rear.

This implementation is backed by a DoublyLinkedList from the linkedlist package,
which provides efficient O(1) head/tail operations and O(n) element search or removal.

Key Features:
  - OfferFirst / OfferLast: Add elements to the front or rear of the deque.
  - PollFirst / PollLast: Remove elements from the front or rear.
  - PeekFirst / PeekLast: Access elements at the front or rear without removal.
  - Remove: Delete the first occurrence of an element (O(n) operation).
  - Size / IsEmpty: Retrieve deque size or check for emptiness.

Concurrency:
  - All public methods are safe for concurrent use by multiple goroutines.
*/
package deque

import "github.com/Zubayear/ryushin/linkedlist"

// Deque is a generic double-ended queue backed by a doubly linked structure.
// It supports adding, removing, and peeking elements from both ends in O(1) time.
type Deque[T comparable] struct {
	data *linkedlist.DoublyLinkedList[T]
}

// NewDeque returns a new, empty Deque[T] backed by a doubly linked list.
// The returned deque is ready to use immediately.
//
// Time Complexity: O(1)
func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{
		data: linkedlist.NewLinkedList[T](),
	}
}

// OfferFirst inserts an element at the front of the deque.
// Algorithm: Add element to the head of the underlying doubly linked list.
//
// Time Complexity: O(1)
func (d *Deque[T]) OfferFirst(elem T) (bool, error) {
	return d.data.AddFirst(elem)
}

// PollFirst removes and returns the first element of the deque.
// Returns zero values and an error if the deque is empty.
// Algorithm: Remove element from the head of the linked list.
//
// Time Complexity: O(1)
func (d *Deque[T]) PollFirst() (T, error) {
	return d.data.RemoveFirst()
}

// PeekFirst retrieves the first element without removing it.
// Returns zero values and an error if the deque is empty.
// Algorithm: Access head element of the linked list.
//
// Time Complexity: O(1)
func (d *Deque[T]) PeekFirst() (T, error) {
	return d.data.PeekFirst()
}

// OfferLast inserts an element at the end of the deque.
// Algorithm: Add element to the tail of the underlying doubly linked list.
//
// Time Complexity: O(1)
func (d *Deque[T]) OfferLast(elem T) (bool, error) {
	return d.data.AddLast(elem)
}

// PollLast removes and returns the last element of the deque.
// Returns zero values and an error if the deque is empty.
// Algorithm: Remove element from the tail of the linked list.
//
// Time Complexity: O(1)
func (d *Deque[T]) PollLast() (T, error) {
	return d.data.RemoveLast()
}

// PeekLast retrieves the last element without removing it.
// Returns zero values and an error if the deque is empty.
// Algorithm: Access tail element of the linked list.
//
// Time Complexity: O(1)
func (d *Deque[T]) PeekLast() (T, error) {
	return d.data.PeekLast()
}

// Remove deletes the first occurrence of the specified element from the deque.
// Returns true if an element was removed, false otherwise.
// Algorithm: Traverse the linked list to find and remove the node.
//
// Time Complexity: O(n)
func (d *Deque[T]) Remove(elem T) bool {
	ok, err := d.data.Remove(elem)
	if err != nil {
		return false
	}
	return ok == elem
}

// Size returns the number of elements in the deque.
//
// Time Complexity: O(1)
func (d *Deque[T]) Size() int {
	return d.data.Size()
}

// IsEmpty reports whether the deque has no elements.
//
// Time Complexity: O(1)
func (d *Deque[T]) IsEmpty() bool {
	return d.data.IsEmpty()
}
