package priorityqueue_test

import (
	"errors"
	"sync"
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

// Test adding and polling strings in lexicographical order
func TestBinaryHeapStringBasic(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[string]()
	values := []string{"apple", "banana", "cat", "aardvark", "dog"}

	for _, v := range values {
		bh.Add(v)
	}

	expectedOrder := []string{"aardvark", "apple", "banana", "cat", "dog"}
	for _, expected := range expectedOrder {
		val, err := bh.Poll()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != expected {
			t.Errorf("expected %s, got %s", expected, val)
		}
	}

	if !bh.IsEmpty() {
		t.Errorf("heap should be empty after polling all elements")
	}
}

// Test Peek on empty and non-empty string heap
func TestBinaryHeapStringPeek(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[string]()

	// Peek on empty heap
	if _, err := bh.Peek(); err == nil {
		t.Error("expected error on empty heap Peek()")
	}

	bh.Add("zebra")
	val, err := bh.Peek()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "zebra" {
		t.Errorf("expected 'zebra', got %s", val)
	}
}

// Test Poll on empty string heap
func TestBinaryHeapStringPollEmpty(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[string]()
	if _, err := bh.Poll(); err == nil {
		t.Error("expected error on empty heap Poll()")
	}
}

// Test Clear with strings
func TestBinaryHeapStringClear(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[string]()
	bh.Add("apple")
	bh.Add("banana")
	bh.Clear()

	if !bh.IsEmpty() {
		t.Error("heap should be empty after Clear()")
	}

	if _, err := bh.Poll(); err == nil {
		t.Error("expected error on empty heap after Clear()")
	}
}

// Test duplicates in string heap
func TestBinaryHeapStringDuplicates(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[string]()
	bh.Add("apple")
	bh.Add("apple")
	bh.Add("apple")

	for i := 0; i < 3; i++ {
		val, err := bh.Poll()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != "apple" {
			t.Errorf("expected 'apple', got %s", val)
		}
	}

	if !bh.IsEmpty() {
		t.Error("heap should be empty after polling all duplicates")
	}
}

// Test concurrent Add and Poll with strings
func TestBinaryHeapStringConcurrent(t *testing.T) {
	bh := priorityqueue.NewBinaryHeap[string]()
	var wg sync.WaitGroup

	stringsToAdd := []string{"apple", "banana", "cat", "dog", "aardvark"}

	// Concurrent adds
	for i := 0; i < len(stringsToAdd); i++ {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			bh.Add(val)
		}(stringsToAdd[i])
	}

	wg.Wait()

	if bh.Size() != len(stringsToAdd) {
		t.Errorf("expected size %d after concurrent adds, got %d", len(stringsToAdd), bh.Size())
	}

	// Concurrent polls
	wg = sync.WaitGroup{}
	results := make(chan string, len(stringsToAdd))
	for i := 0; i < len(stringsToAdd); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, err := bh.Poll()
			if err == nil {
				results <- val
			}
		}()
	}

	wg.Wait()
	close(results)

	if len(results) != len(stringsToAdd) {
		t.Errorf("expected %d results after concurrent polls, got %d", len(stringsToAdd), len(results))
	}

	if !bh.IsEmpty() {
		t.Error("heap should be empty after all concurrent polls")
	}
}
