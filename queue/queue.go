/*
Package queue provides a generic, concurrency-safe implementation of a queue
(First-In-First-Out) data structure in Go. It uses a circular array for
efficient memory usage and supports dynamic resizing when the queue is full.

Features:
  - Generic Type Support: Works with any comparable type.
  - Thread-Safety: All operations are protected using sync.RWMutex.
  - Dynamic Resizing: Doubles capacity automatically when full.
  - Utility Methods: Peek, IsEmpty, IsFull, Size, Clear, Print.

Use Cases:
  - Task scheduling and job queues.
  - Breadth-first search (BFS) in graphs.
  - Message buffering in concurrent systems.
  - Order processing systems.

Example:

	q := queue.NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	val, _ := q.Dequeue()
	fmt.Println(val) // 10

Implementation Details:
  - Uses a circular array to store elements.
  - `front` and `rear` track positions for dequeue and enqueue operations.
  - Automatically resizes when the number of elements equals capacity.
  - Protected by RWMutex for concurrent access.

Complexity:
  - Enqueue: O(1) amortized
  - Dequeue: O(1)
  - Peek: O(1)
  - Size: O(1)
  - Print: O(n)
*/
package queue

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// Queue represents a generic circular queue with dynamic resizing.
// It is concurrency-safe using sync.RWMutex for read/write operations.
//
// Internally, it uses a slice as a circular buffer to optimize memory usage.
//
// Type parameter:
//
//	T - The element type, which must be comparable.
//
// Example usage:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(10)
//	q.Enqueue(20)
//	val, _ := q.Dequeue()
//	fmt.Println(val) // Output: 10
type Queue[T comparable] struct {
	front, rear, cap, count int
	data                    []T
	mutex                   sync.RWMutex
}

// NewQueue creates and returns a new queue with an initial capacity of 16.
//
// Complexity: O(1)
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{cap: 16, front: 0, rear: 0, count: 0, data: make([]T, 16)}
}

// increaseSize doubles the capacity of the queue when it's full
// and rearranges existing elements to maintain the correct order.
//
// Algorithm Steps:
//  1. Double the current capacity.
//  2. Create a new slice with the updated capacity.
//  3. Copy all elements from the old slice to the new one in order.
//  4. Reset the front and rear pointers.
//
// Complexity: O(n), where n = current number of elements.
func (q *Queue[T]) increaseSize() {
	newCap := q.cap * 2
	newData := make([]T, newCap)

	// Copy elements in the correct order
	for i := 0; i < q.count; i++ {
		newData[i] = q.data[(q.front+i)%q.cap]
	}

	q.data = newData
	q.front = 0
	q.rear = q.count
	q.cap = newCap
}

// Enqueue adds an element to the rear of the queue.
// If the queue is full, it doubles its capacity before adding.
//
// Algorithm Steps:
//  1. If full, increase capacity using increaseSize().
//  2. Insert element at rear index (mod cap).
//  3. Increment rear and count.
//
// Complexity: O(1) amortized, O(n) when resizing.
func (q *Queue[T]) Enqueue(val T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.count == q.cap {
		q.increaseSize()
	}
	q.data[q.rear%q.cap] = val
	q.rear++
	q.count++
}

// Dequeue removes and returns the element from the front of the queue.
// Returns an error if the queue is empty.
//
// Algorithm Steps:
//  1. If empty, return error.
//  2. Retrieve element at the front index (mod cap).
//  3. Clear the element (optional).
//  4. Increment front and decrement count.
//
// Complexity: O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.count == 0 {
		return zero, errors.New("queue empty")
	}
	value := q.data[q.front%q.cap]
	q.data[q.front%q.cap] = zero
	q.front++
	q.count--
	return value, nil
}

// Peek returns the element at the front of the queue without removing it.
// Returns an error if the queue is empty.
//
// Complexity: O(1)
func (q *Queue[T]) Peek() (T, error) {
	var zero T
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.count == 0 {
		return zero, errors.New("queue empty")
	}
	return q.data[q.front%q.cap], nil
}

// IsFull checks if the queue has reached its current capacity.
//
// Complexity: O(1)
func (q *Queue[T]) IsFull() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.count == q.cap
}

// IsEmpty checks if the queue contains no elements.
//
// Complexity: O(1)
func (q *Queue[T]) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.count == 0
}

// Size returns the current number of elements in the queue.
//
// Complexity: O(1)
func (q *Queue[T]) Size() int {
	return q.count
}

// Deprecated: Use ToArray instead. ToArray returns a []T which can be
// easily converted to string using fmt.Sprint if needed.
// ToString returns a string representation of the queue elements in FIFO order.
//
// Example output:
//
//	[10, 20, 30]
//
// Complexity: O(n)
func (q *Queue[T]) ToString() string {
	q.mutex.Lock()
	defer q.mutex.Unlock()
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

// Clear removes all elements from the queue and resets it to the initial state.
// The capacity remains unchanged.
//
// Complexity: O(1)
func (q *Queue[T]) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.front = 0
	q.rear = 0
	q.count = 0
	q.cap = 16
}

// ToArray returns a array representation of the queue elements in FIFO order.
//
// Example output:
//
//	[10, 20, 30]
//
// Complexity: O(n)
func (q *Queue[T]) ToArray() []T {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	var result []T
	for r := q.front; r <= q.rear-1; r++ {
		value := q.data[r%q.cap]
		result = append(result, value)
	}
	return result
}

// Iterator represents a type to iterate queue.
// It is concurrency-safe using sync.RWMutex for read/write operations.
//
// Type parameter:
//
//	T - The element type, which must be comparable.
type Iterator[T comparable] struct {
	idx  int
	data []T
}

// Iterator returns a snapshot of the queue elements in FIFO order.
//
// # Use Next() to iterate values
//
// Complexity: O(n)
func (q *Queue[T]) Iterator() *Iterator[T] {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	// copy snapshot
	snapshot := make([]T, q.Size())
	copy(snapshot, q.data)

	return &Iterator[T]{data: snapshot, idx: 0}
}

// Next return queue elements in FIFO order
//
// Returns value from queue in FIFO order,
// If Queue has element returns T type value and
// bool true if next value exit, otherwise return false
//
// Example Return: (T, bool)
func (q *Iterator[T]) Next() (T, bool) {
	if q.idx >= len(q.data) {
		var zero T
		return zero, false
	}
	v := q.data[q.idx]
	q.idx++
	return v, true
}
