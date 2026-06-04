package stack

import (
	"sync"
	"testing"
)

// TestClearThenPushDoesNotPanic guards the regression where Clear() nil-ed the
// backing slice but left cap unchanged, causing the next Push to index into a
// nil slice and panic.
func TestClearThenPushDoesNotPanic(t *testing.T) {
	s := NewStack[int]()
	if _, err := s.Push(1); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	s.Clear()

	if _, err := s.Push(2); err != nil {
		t.Fatalf("Push after Clear failed: %v", err)
	}
	v, err := s.Peek()
	if err != nil {
		t.Fatalf("Peek after Clear+Push failed: %v", err)
	}
	if v != 2 {
		t.Fatalf("want 2 got %d", v)
	}
	if got := s.Size(); got != 1 {
		t.Fatalf("want size 1 got %d", got)
	}
}

// TestPeekConcurrent exercises Peek alongside writers; run with -race to detect
// the prior unsynchronized read of the backing slice in Peek.
func TestPeekConcurrent(t *testing.T) {
	s := NewStack[int]()
	s.Push(42)

	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 5000; i++ {
				if i%2 == 0 {
					s.Push(i)
				} else {
					s.Pop()
				}
				_, _ = s.Peek()
				_ = s.Size()
			}
		}()
	}
	wg.Wait()
}
