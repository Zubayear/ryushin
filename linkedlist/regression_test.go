package linkedlist

import (
	"sync"
	"testing"
	"time"
)

// TestConcurrentMutationRace guards the regression where AddAt/Remove/RemoveAt
// and the middle-node path of removeNode mutated the list without holding the
// lock. Run with -race.
func TestConcurrentMutationRace(t *testing.T) {
	dl := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		dl.AddLast(i)
	}

	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(base int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				dl.AddAt(1, base+i)
				dl.Remove(base + i)
				dl.RemoveAt(0)
				dl.AddFirst(base + i)
			}
		}(g * 1000000)
	}
	wg.Wait()
}

// TestIterateEarlyBreakDoesNotDeadlock guards the regression where Iterate held
// the read lock for the lifetime of the streaming goroutine. A consumer that
// stops reading early must not block subsequent writers.
func TestIterateEarlyBreakDoesNotDeadlock(t *testing.T) {
	dl := NewLinkedList[int]()
	for i := 0; i < 100; i++ {
		dl.AddLast(i)
	}

	// Read a single value then abandon the iterator.
	it := dl.Iterate()
	<-it

	done := make(chan struct{})
	go func() {
		// These take the write lock; they must not block on a held read lock.
		dl.AddLast(999)
		_, _ = dl.RemoveFirst()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("DEADLOCK: writer blocked after iterator consumer stopped early")
	}
}
