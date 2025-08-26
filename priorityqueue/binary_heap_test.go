package priorityqueue_test

import (
	"errors"
	"testing"

	"github.com/Zubayear/sonic/priorityqueue"
)

func TestBinaryHeapOperations(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[int]()
	isEmpty := bh.IsEmpty()
	if !isEmpty {
		t.Fatalf("Expected %v, got %v\n", false, isEmpty)
	}

	bh.Add(10)
	bh.Add(5)
	bh.Add(30)
	bh.Add(20)
	bh.Add(40)
	bh.Add(35)
	bh.Add(15)

	size := bh.Size()
	if size != 7 {
		t.Fatalf("Extected %v, got %v\n", 7, size)
	}

	top, _ := bh.Peek()
	if top != 5 {
		t.Errorf("Expected %v, got %v\n", 5, top)
	}

	top, _ = bh.Poll()
	if top != 5 {
		t.Errorf("Expected %v, got %v\n", 5, top)
	}

	bh.Clear()
	size = bh.Size()
	if size != 0 {
		t.Errorf("Expected %v, got %v\n", 0, size)
	}

	_, err := bh.Peek()
	if errors.Is(err, errors.New("heap empty")) {
		t.Errorf("Expected %v, got %v\n", errors.New("heap empty"), err)
	}

	_, err = bh.Poll()
	if errors.Is(err, errors.New("heap empty")) {
		t.Errorf("Expected %v, got %v\n", errors.New("heap empty"), err)
	}
}
