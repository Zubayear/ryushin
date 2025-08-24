package stack

import (
	"errors"
	"sync"
)

var i any = -1

// Stack represents a generic stack data structure with dynamic resizing.
// It uses a slice as underlying storage and a RWMutex for optional concurrency safety.
type Stack[T comparable] struct {
	size, top int
	data      []T
	lock      sync.RWMutex
}

// NewStack initializes a new stack with default capacity of 16.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		size: 16,
		top:  -1,
		data: make([]T, 16),
		lock: sync.RWMutex{},
	}
}

// increaseSize doubles the capacity of the underlying slice.
// It copies all existing elements to the new slice.
// This method is protected by a lock for thread-safety.
func (s *Stack[T]) increaseSize() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.size = s.size * 2
	newData := make([]T, s.size)
	copy(newData, s.data)
	s.data = newData
}

// Push adds a new element to the top of the stack.
// Automatically increases size if the stack is full.
// Returns true on success and an error if any.
func (s *Stack[T]) Push(val T) (bool, error) {
	if s.IsFull() {
		s.increaseSize()
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top++
	s.data[s.top] = val
	return true, nil
}

// Pop removes and returns the top element from the stack.
// Returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return i.(T), errors.New("stack empty")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	value := s.data[s.top]
	s.top--
	return value, nil
}

// Peek returns the value at the top of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return i.(T), errors.New("stack empty")
	}
	return s.data[s.top], nil
}

// Size returns the number of elements currently in the stack.
func (s *Stack[T]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top + 1
}

// IsEmpty checks if the stack has no elements.
// Returns true if empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

// IsFull checks if the stack has reached its current capacity.
// Returns true if full, false otherwise.
func (s *Stack[T]) IsFull() bool {
	return s.top == s.size-1
}

// ValueAt returns the value at a specific position from the top of the stack (0-based).
// pos = 0 returns the top element, pos = 1 returns one below top, etc.
// Returns an error if the stack is empty or the position is out of bounds.
func (s *Stack[T]) ValueAt(pos int) (T, error) {
	if s.IsEmpty() {
		return i.(T), errors.New("stack empty")
	}
	if s.top-pos+1 < 0 {
		return i.(T), errors.New("stack empty")
	}
	return s.data[s.top-pos+1], nil
}
