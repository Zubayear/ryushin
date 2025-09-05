/*
Package stack provides a generic, concurrency-safe implementation of a stack
(Last-In-First-Out) data structure in Go. It supports dynamic resizing and
utility methods for stack manipulation.

Features:
  - Generic Type Support: Works with any comparable type.
  - Thread-Safety: All operations are protected using sync.RWMutex.
  - Dynamic Resizing: The underlying slice doubles in capacity when full.
  - Utility Methods: Peek, ValueAt, Clear, Size, IsEmpty, IsFull.

Use Cases:
  - Expression evaluation (e.g., postfix, infix).
  - Undo functionality in text editors.
  - Recursive function call simulation.
  - Depth-first search (DFS) in graphs.

Example:

	s := stack.NewStack[int]()
	s.Push(10)
	s.Push(20)
	val, _ := s.Pop()
	fmt.Println(val) // 20

Complexity:
  - Push: O(1) amortized
  - Pop: O(1)
  - Peek: O(1)
  - ValueAt: O(1)
  - Clear: O(1)

Implementation Details:
  - Internally uses a slice for storage.
  - `top` tracks the index of the last inserted element.
  - Capacity grows dynamically when the stack is full.
  - Protected by RWMutex for concurrent access.
*/
package stack

import (
	"errors"
	"sync"
)

// Stack represents a generic stack (LIFO) data structure with dynamic resizing.
// It is safe for concurrent use as sync.RWMutex guards all operations.
//
// Internally, it uses a slice as the underlying storage and automatically
// doubles the capacity when needed.
//
// Type parameter:
//
//	T - The element type, which must be comparable.
//
// Example usage:
//
//	s := stack.NewStack[int]()
//	s.Push(10)
//	val, _ := s.Pop()
//	fmt.Println(val) // Output: 10
type Stack[T comparable] struct {
	cap, top int
	data     []T
	lock     sync.RWMutex
}

// NewStack creates and returns a new Stack with a default initial capacity of 16.
//
// Complexity: O(1)
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		cap:  16,
		top:  -1,
		data: make([]T, 16),
		lock: sync.RWMutex{},
	}
}

// increaseSize doubles the capacity of the underlying slice
// while preserving existing elements.
//
// Algorithm:
//  1. Multiply current capacity by 2.
//  2. Create a new slice with the updated capacity.
//  3. Copy all elements from the old slice to the new one.
//  4. Replace the old slice with the new one.
//
// Complexity: O(N), where N is the current number of elements.
func (s *Stack[T]) increaseSize() {
	s.cap = s.cap * 2
	newData := make([]T, s.cap)
	copy(newData, s.data)
	s.data = newData
}

// Push adds an element to the top of the stack.
// If the stack is full, it automatically increases the capacity.
//
// Algorithm:
//  1. Check if the stack is full; if yes, double the capacity.
//  2. Increment top and insert the element.
//
// Returns true on success and nil error. Returns an error only in rare cases.
//
// Complexity: Amortized O(1)
func (s *Stack[T]) Push(val T) (bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top == s.cap-1 {
		s.increaseSize()
	}
	s.top++
	s.data[s.top] = val
	return true, nil
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns an error.
//
// Algorithm:
//  1. If the stack is empty, return an error.
//  2. Retrieve the top element and decrement top pointer.
//
// Complexity: O(1)
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == -1 {
		return zero, errors.New("stack empty")
	}

	value := s.data[s.top]
	s.top--
	return value, nil
}

// Peek returns the element at the top of the stack without removing it.
// Returns an error if the stack is empty.
//
// Complexity: O(1)
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack empty")
	}
	return s.data[s.top], nil
}

// Size returns the number of elements currently in the stack.
//
// Complexity: O(1)
func (s *Stack[T]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top + 1
}

// IsEmpty checks whether the stack has no elements.
// Returns true if empty, false otherwise.
//
// Complexity: O(1)
func (s *Stack[T]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.top == -1
}

// IsFull checks whether the stack has reached its current capacity.
// Returns true if full, false otherwise.
//
// Complexity: O(1)
func (s *Stack[T]) IsFull() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.top == s.cap-1
}

// ValueAt returns the element at a specific position from the top of the stack (0-based index).
//
// Note
//   - pos = 0 returns the top element,
//   - pos = 1 returns the element just below the top, and so on.
//
// Returns an error if:
//   - The stack is empty
//   - The position is out of range
//
// Complexity: O(1)
func (s *Stack[T]) ValueAt(pos int) (T, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var zero T
	if s.top == -1 {
		return zero, errors.New("stack empty")
	}
	if pos < 0 || pos >= s.top+1 {
		return zero, errors.New("invalid position")
	}
	return s.data[s.top-pos], nil
}

// Clear removes all elements from the stack and resets it to an empty state.
// After clearing, the underlying slice is set to nil to free memory.
//
// Complexity: O(1)
func (s *Stack[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top = -1
	s.data = nil
}
