package queue

import (
	"reflect"
	"sync"
	"testing"
)

// TestIteratorFIFOAfterWrap guards the regression where Iterator() copied the
// raw backing array from index 0, ignoring front/rear, so it yielded stale or
// zeroed slots after dequeues.
func TestIteratorFIFOAfterWrap(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("dequeue failed: %v", err)
		}
	}
	want := []int{5, 6, 7, 8, 9}

	var got []int
	it := q.Iterator()
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		got = append(got, v)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Iterator FIFO mismatch:\n want %v\n got  %v", want, got)
	}
}

// TestIteratorMatchesToArray ensures Iterator and ToArray agree, including after
// a resize triggered while front != 0.
func TestIteratorMatchesToArray(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 8; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 6; i++ {
		q.Dequeue()
	}
	// Force a resize while front is advanced and wrapped.
	for i := 100; i < 140; i++ {
		q.Enqueue(i)
	}

	want := q.ToArray()
	var got []int
	it := q.Iterator()
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		got = append(got, v)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Iterator vs ToArray mismatch:\n want %v\n got  %v", want, got)
	}
}

// TestSizeConcurrent exercises Size alongside writers; run with -race to detect
// the prior unsynchronized read of count.
func TestSizeConcurrent(t *testing.T) {
	q := NewQueue[int]()
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 5000; i++ {
				q.Enqueue(i)
				_ = q.Size()
				q.Dequeue()
			}
		}()
	}
	wg.Wait()
}
