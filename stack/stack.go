package stack

import (
	"errors"
	"sync"
)

// Stack represents a generic stack data structure with dynamic resizing.
// It uses a slice as underlying storage and a RWMutex for optional concurrency safety.
type Stack[T comparable] struct {
	cap, top int
	data     []T
	lock     sync.RWMutex
}

// NewStack initializes a new stack with default capacity of 16.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		cap:  16,
		top:  -1,
		data: make([]T, 16),
		lock: sync.RWMutex{},
	}
}

// increaseSize doubles the capacity of the underlying slice.
// It copies all existing elements to the new slice.
// This method is protected by a lock for thread-safety.
func (s *Stack[T]) increaseSize() {
	s.cap = s.cap * 2
	newData := make([]T, s.cap)
	copy(newData, s.data)
	s.data = newData
}

// Push adds a new element to the top of the stack.
// Automatically increases size if the stack is full.
// Returns true on success and an error if any.
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
// Returns an error if the stack is empty.
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

// Peek returns the value at the top of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack empty")
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
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.top == -1
}

// IsFull checks if the stack has reached its current capacity.
// Returns true if full, false otherwise.
func (s *Stack[T]) IsFull() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.top == s.cap-1
}

// ValueAt returns the value at a specific position from the top of the stack (0-based).
// pos = 0 returns the top element, pos = 1 returns one below top, etc.
// Returns an error if the stack is empty or the position is out of bounds.
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

func (s *Stack[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.top = -1
	s.data = nil
}
