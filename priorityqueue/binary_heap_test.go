package priorityqueue

import (
	"errors"
	"reflect"
	"sync"
	"testing"
)

func TestBinaryHeapOperations(t *testing.T) {
	bh := NewBinaryHeap[int]()
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
	if top != 40 {
		t.Errorf("Expected %v, got %v\n", 5, top)
	}

	top, _ = bh.Poll()
	if top != 40 {
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

func TestBinaryHeapStringBasic(t *testing.T) {
	bh := NewBinaryHeap[string]()
	values := []string{"apple", "banana", "cat", "aardvark", "dog"}

	for _, v := range values {
		bh.Add(v)
	}

	expectedOrder := []string{"dog", "cat", "banana", "apple", "aardvark"}
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

func TestBinaryHeapStringPeek(t *testing.T) {
	bh := NewBinaryHeap[string]()

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

func TestBinaryHeapStringPollEmpty(t *testing.T) {
	bh := NewBinaryHeap[string]()
	if _, err := bh.Poll(); err == nil {
		t.Error("expected error on empty heap Poll()")
	}
}

func TestBinaryHeapStringClear(t *testing.T) {
	bh := NewBinaryHeap[string]()
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

func TestBinaryHeapStringDuplicates(t *testing.T) {
	bh := NewBinaryHeap[string]()
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

func TestBinaryHeapStringConcurrent(t *testing.T) {
	bh := NewBinaryHeap[string]()
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

// Person represents a person with a Name and years they lived.
type Person struct {
	Name  string
	Lived uint
}

func TestBinaryHeapCustomComparator(t *testing.T) {
	// Custom comparator:
	// - Higher Lived first
	// - If Lived is equal, longer Name first
	bh := NewBinaryHeapWithComparator[Person](func(p1, p2 Person) bool {
		if p1.Lived != p2.Lived {
			return p1.Lived > p2.Lived
		}
		return len(p1.Name) > len(p2.Name)
	})

	people := []Person{
		{"Fyodor Dostoevsky", 46},
		{"George Orwell", 46},
		{"Ernest Hemingway", 61},
		{"Leo Tolstoy", 82},
		{"Friedrich Nietzsche", 55},
		{"Franz Kafka", 40},
	}

	// Add all people to heap
	for _, p := range people {
		bh.Add(p)
	}

	// Expected order according to comparator
	expectedOrder := []Person{
		{"Leo Tolstoy", 82},         // highest Lived
		{"Ernest Hemingway", 61},    // next highest Lived
		{"Friedrich Nietzsche", 55}, // next highest Lived
		{"Fyodor Dostoevsky", 46},   // tie Lived 46, longer name
		{"George Orwell", 46},       // tie Lived 46, shorter name
		{"Franz Kafka", 40},         // lowest Lived
	}

	// Poll elements and verify order
	for i, exp := range expectedOrder {
		p, err := bh.Poll()
		if err != nil {
			t.Fatalf("Poll failed at index %d: %v", i, err)
		}
		if p != exp {
			t.Errorf("Poll order incorrect at index %d: got %+v, want %+v", i, p, exp)
		}
	}

	// Heap should now be empty
	if !bh.IsEmpty() {
		t.Errorf("Heap should be empty after polling all elements")
	}

	// Poll on empty heap should return error
	_, err := bh.Poll()
	if err == nil {
		t.Errorf("Expected error when polling empty heap, got nil")
	}

	// Peek on empty heap should return error
	_, err = bh.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking empty heap, got nil")
	}

}

func TestBinaryHeapEdgeCases(t *testing.T) {

	// Edge case: Adding duplicates
	bh := NewBinaryHeapWithComparator[Person](func(p1, p2 Person) bool {
		return p1.Lived > p2.Lived
	})

	dup := Person{"John Doe", 40}
	for i := 0; i < 5; i++ {
		bh.Add(dup)
	}

	if bh.Size() != 5 {
		t.Errorf("Expected heap size 5 after adding duplicates, got %d", bh.Size())
	}

	// Poll all duplicates
	for i := 0; i < 5; i++ {
		p, err := bh.Poll()
		if err != nil {
			t.Fatalf("Poll failed at duplicate index %d: %v", i, err)
		}
		if p != dup {
			t.Errorf("Poll returned wrong element at index %d: got %+v, want %+v", i, p, dup)
		}
	}
}

func TestBinaryHeapSort(t *testing.T) {
	bh := NewBinaryHeap[int]()
	val := []int{10, 20, 30, 40, 50, 60}
	expected := []int{60, 50, 40, 30, 20, 10}
	for _, v := range val {
		bh.Add(v)
	}
	result := bh.Sort()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Got wrong sort order")
	}
}

func TestBinaryHeapConcurrency_Add(t *testing.T) {
	bh := NewBinaryHeap[int]()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				bh.Add(start*100 + j)
			}
		}(i)
	}
	wg.Wait()

	if bh.Size() != 50*100 {
		t.Errorf("Expected %d elements, got %d", 50*100, bh.Size())
	}
}

func TestBinaryHeapConcurrency_Peek(t *testing.T) {
	bh := NewBinaryHeap[int]()
	for i := 0; i < 1000; i++ {
		bh.Add(i)
	}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				_, _ = bh.Peek()
			}
		}()
	}
	wg.Wait()
}

func TestBinaryHeapConcurrency_Poll(t *testing.T) {
	bh := NewBinaryHeap[int]()
	for i := 0; i < 5000; i++ {
		bh.Add(i)
	}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				_, err := bh.Poll()
				if err != nil {
					break
				}
			}
		}()
	}
	wg.Wait()

	if !bh.IsEmpty() {
		t.Errorf("Heap should be empty after polling all elements")
	}
}

func TestBinaryHeapConcurrency_ClearAndIsEmpty(t *testing.T) {
	bh := NewBinaryHeap[int]()
	for i := 0; i < 1000; i++ {
		bh.Add(i)
	}

	var wg sync.WaitGroup
	// Concurrent Clears
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bh.Clear()
		}()
	}

	// Concurrent IsEmpty checks
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = bh.IsEmpty()
		}()
	}
	wg.Wait()
}

func TestBinaryHeapConcurrency_Size(t *testing.T) {
	bh := NewBinaryHeap[int]()
	for i := 0; i < 1000; i++ {
		bh.Add(i)
	}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				_ = bh.Size()
			}
		}()
	}
	wg.Wait()
}

func TestBinaryHeapConcurrency_Sort(t *testing.T) {
	bh := NewBinaryHeap[int]()
	for i := 0; i < 1000; i++ {
		bh.Add(i)
	}

	var wg sync.WaitGroup
	numGoroutines := 50

	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				sorted := bh.Sort()
				// quick sanity check: the first element should be largest (max-heap)
				if len(sorted) > 0 && sorted[0] < sorted[len(sorted)-1] {
					t.Errorf("Sort order incorrect")
				}
			}
		}()
	}

	wg.Wait()
}

func TestBinaryHeapConcurrencyIssue(t *testing.T) {
	bh := NewBinaryHeap[int]()
	wg := sync.WaitGroup{}
	numGoroutines := 50
	numOps := 1000

	// Writer goroutines: add and poll
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				bh.Add(j)
				_, _ = bh.Poll()
			}
		}(i)
	}

	// Reader goroutines: Size, Peek, Sort
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				_ = bh.Size() // unsafe if RLock removed
				_, _ = bh.Peek()
				_ = bh.Sort()
			}
		}(i)
	}

	wg.Wait()
}

func TestBinaryHeapRemoveInEmptyHeap(t *testing.T) {
	bh := NewBinaryHeap[int]()
	_, err := bh.removeAt(1)
	if errors.Is(err, errors.New("heap empty")) {
		t.Errorf("Expected heap empty error")
	}
}
