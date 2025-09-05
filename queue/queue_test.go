package queue_test

import (
	"reflect"
	"testing"

	"github.com/Zubayear/ryushin/queue"
)

func TestQueueOperations(t *testing.T) {
	q := queue.NewQueue[int]()
	isEmpty := q.IsEmpty()
	if !isEmpty {
		t.Errorf("Expected %v, got %v\n", false, isEmpty)
	}

	q.Enqueue(1)
	q.Enqueue(4)
	q.Enqueue(79)

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
	result := q.ToString()
	if !reflect.DeepEqual(result, "[4, 79]") {
		t.Errorf("ToString() = %v; want %v", result, "[4, 79]")
	}

	q.Clear()
	if q.Size() != 0 {
		t.Errorf("Expected %v, got %v\n", 0, q.Size())
	}

	v, _ := q.Peek()
	if v != 0 {
		t.Errorf("Expected %v, got %v\n", 0, v)
	}

	v, _ = q.Dequeue()
	if v != 0 {
		t.Errorf("Expected %v, got %v\n", 0, v)
	}

	for i := 0; i < 50; i++ {
		q.Enqueue(i)
	}
}
