package queue_test

import (
	"github.com/Zubayear/sonic/queue"
	"testing"
)

func TestQueueOperations(t *testing.T) {
	q := queue.NewQueue[int]()
	isEmpty := q.IsEmpty()
	if !isEmpty {
		t.Errorf("Expected %v, got %v\n", false, isEmpty)
	}

	_ = q.Enqueue(1)
	_ = q.Enqueue(4)
	_ = q.Enqueue(79)

	size := q.Size()
	if size != 3 {
		t.Errorf("Extected %v, got %v\n", 3, size)
	}
	value, err := q.Dequeue()
	if err != nil || value != 1 {
		t.Errorf("Expected %v, got %v\n", 1, value)
	}

	value, err = q.Peek()
	if err != nil || value != 4 {
		t.Errorf("Expected %v, got %v\n", 4, value)
	}
	isFull := q.IsFull()
	if isFull {
		t.Errorf("Expected %v, got %v\n", false, isFull)
	}
	q.Print()
}
