package stack

import "fmt"

var i interface{} = nil

// Stack representation of stack
type Stack[T any] struct {
	cap, top int
	data     []T
}

// NewStack initialize a new stack
func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{cap: cap, top: -1, data: make([]T, cap)}
}

// Push adds val to stack
func (s *Stack[T]) Push(val T) (bool, error) {
	if s.IsFull() {
		return false, fmt.Errorf("stack full")
	}
	s.top++
	s.data[s.top] = val
	return true, nil
}

// Pop removes the value at top
func (s *Stack[T]) Pop() (bool, error) {
	if s.IsEmpty() {
		return false, fmt.Errorf("stack empty")
	}
	s.top--
	return true, nil
}

// Peek returns value at the top of the stack
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return i.(T), fmt.Errorf("stack empty")
	}
	return s.data[s.top], nil
}

// Size returns the size of the stack
func (s *Stack[T]) Size() int {
	return s.top + 1
}

// IsEmpty check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

// IsFull check if stack is full
func (s *Stack[T]) IsFull() bool {
	return s.top > len(s.data)
}
