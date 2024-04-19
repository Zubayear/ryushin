package queue

import (
	"errors"
	"fmt"
	"strings"
)

var i interface{} = -1

type Queue[T comparable] struct {
	front, rear, cap, count int
	data                    []T
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{cap: 16, front: 0, rear: 0, count: 0, data: make([]T, 16)}
}

func (q *Queue[T]) increaseSize() {
	q.cap = q.cap * 2
	newData := make([]T, q.cap)
	copy(newData, q.data)
	q.data = newData
}

func (q *Queue[T]) Enqueue(val T) error {
	if q.IsFull() {
		q.increaseSize()
	}
	q.data[q.rear%q.cap] = val
	q.rear++
	q.count++
	return nil
}

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

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return i.(T), errors.New("queue empty")
	}
	return q.data[q.front%q.cap], nil
}

func (q *Queue[T]) IsFull() bool {
	return q.count == q.cap
}

func (q *Queue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue[T]) Size() int {
	return q.count
}

func (q *Queue[T]) Print() {
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
	fmt.Println(result.String())
}
