package stack

import (
	"errors"
	"sync"
)

var i interface{} = -1

// Stack representation of stack
type Stack[T comparable] struct {
	size, top int
	data      []T
	lock      sync.RWMutex
}

// NewStack initialize a new stack
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		size: 16,
		top:  -1,
		data: make([]T, 16),
		lock: sync.RWMutex{},
	}
}

// increase size of the underlying slice
// and copy everything from old slice to new one
func (s *Stack[T]) increaseSize() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.size = s.size * 2
	newData := make([]T, s.size)
	copy(newData, s.data)
	s.data = newData
}

// Push adds val to stack
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

// Pop removes the value at top
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

// Peek returns value at the top of the stack
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return i.(T), errors.New("stack empty")
	}
	return s.data[s.top], nil
}

// Size returns the size of the stack
func (s *Stack[T]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top + 1
}

// IsEmpty check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

// IsFull check if stack is full
func (s *Stack[T]) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack[T]) ValueAt(pos int) (T, error) {
	if s.IsEmpty() {
		return i.(T), errors.New("stack empty")
	}
	if s.top-pos+1 < 0 {
		return i.(T), errors.New("stack empty")
	}
	return s.data[s.top-pos+1], nil
}
