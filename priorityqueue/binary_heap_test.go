package priorityqueue_test

import (
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
}
