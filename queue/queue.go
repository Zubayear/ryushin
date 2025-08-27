package queue

import (
	"errors"
	"fmt"
	"strings"
)

var i any = -1

// Queue represents a generic queue using a circular array.
// Supports dynamic resizing when the queue is full.
type Queue[T comparable] struct {
	front, rear, cap, count int
	data                    []T
}

// NewQueue creates and returns a new queue with an initial capacity of 16.
// Time Complexity: O(1)
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{cap: 16, front: 0, rear: 0, count: 0, data: make([]T, 16)}
}

// increaseSize doubles the capacity of the queue when it's full.
// Time Complexity: O(n) where n is the current number of elements.
func (q *Queue[T]) increaseSize() {
	q.cap = q.cap * 2
	newData := make([]T, q.cap)
	copy(newData, q.data)
	q.data = newData
}

// Enqueue adds an element to the rear of the queue.
// If the queue is full, it dynamically increases its size.
// Time Complexity: O(1) amortized
// Space Complexity: O(1) (O(n) when resizing)
func (q *Queue[T]) Enqueue(val T) {
	if q.IsFull() {
		q.increaseSize()
	}
	q.data[q.rear%q.cap] = val
	q.rear++
	q.count++
}

// Dequeue removes and returns the element from the front of the queue.
// Returns an error if the queue is empty.
// Time Complexity: O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return i.(T), errors.New("queue empty")
	}
	value := q.data[q.front%q.cap]
	q.data[q.front%q.cap] = i.(T)
	q.front++
	q.count--
	return value, nil
}

// Peek returns the element at the front of the queue without removing it.
// Returns an error if the queue is empty.
// Time Complexity: O(1)
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return i.(T), errors.New("queue empty")
	}
	return q.data[q.front%q.cap], nil
}

// IsFull checks if the queue has reached its capacity.
// Time Complexity: O(1)
func (q *Queue[T]) IsFull() bool {
	return q.count == q.cap
}

// IsEmpty checks if the queue has no elements.
// Time Complexity: O(1)
func (q *Queue[T]) IsEmpty() bool {
	return q.count == 0
}

// Size returns the current number of elements in the queue.
// Time Complexity: O(1)
func (q *Queue[T]) Size() int {
	return q.count
}

// Print returns a string representation of the queue elements in order.
// Time Complexity: O(n)
func (q *Queue[T]) Print() string {
	var result strings.Builder
	result.WriteString("[")
	for r := q.front; r <= q.rear-1; r++ {
		value := q.data[r%q.cap]
		str := fmt.Sprint(value)
		result.WriteString(str)
		if r != q.rear-1 {
			result.WriteString(", ")
		}
	}
	result.WriteString("]")
	return result.String()
}

func (q *Queue[T]) Clear() {
	q.front = 0
	q.rear = 0
	q.count = 0
	q.cap = 16
}
