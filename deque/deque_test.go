package deque

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestZeroValueDeque ensures the zero value is usable and returns errors on empty ops.
func TestZeroValueDeque(t *testing.T) {
	d := NewDeque[int]()

	if !d.IsEmpty() {
		t.Fatalf("expected zero-value deque to be empty")
	}
	if d.Size() != 0 {
		t.Fatalf("expected size 0, got %d", d.Size())
	}

	if _, err := d.PeekFirst(); err == nil {
		t.Fatalf("expected error on PeekFirst for empty deque")
	}
	if _, err := d.PeekLast(); err == nil {
		t.Fatalf("expected error on PeekLast for empty deque")
	}
	if _, err := d.PollFirst(); err == nil {
		t.Fatalf("expected error on PollFirst for empty deque")
	}
	if _, err := d.PollLast(); err == nil {
		t.Fatalf("expected error on PollLast for empty deque")
	}
}

// TestOfferAndPollFirst verifies front insertions and removals.
func TestOfferAndPollFirst(t *testing.T) {
	d := NewDeque[int]()

	if ok, err := d.OfferFirst(1); !ok || err != nil {
		t.Fatalf("OfferFirst failed: ok=%v err=%v", ok, err)
	}
	if ok, err := d.OfferFirst(2); !ok || err != nil {
		t.Fatalf("OfferFirst failed: ok=%v err=%v", ok, err)
	}
	if d.Size() != 2 {
		t.Fatalf("expected size 2, got %d", d.Size())
	}

	// LIFO from the front
	v, err := d.PollFirst()
	if err != nil || v != 2 {
		t.Fatalf("PollFirst expected 2, got %v err=%v", v, err)
	}
	v, err = d.PollFirst()
	if err != nil || v != 1 {
		t.Fatalf("PollFirst expected 1, got %v err=%v", v, err)
	}

	if !d.IsEmpty() || d.Size() != 0 {
		t.Fatalf("expected empty deque after removals")
	}
}

// TestOfferAndPollLast verifies back insertions and removals.
func TestOfferAndPollLast(t *testing.T) {
	d := NewDeque[int]()

	if ok, err := d.OfferLast(1); !ok || err != nil {
		t.Fatalf("OfferLast failed: ok=%v err=%v", ok, err)
	}
	if ok, err := d.OfferLast(2); !ok || err != nil {
		t.Fatalf("OfferLast failed: ok=%v err=%v", ok, err)
	}
	if d.Size() != 2 {
		t.Fatalf("expected size 2, got %d", d.Size())
	}

	// LIFO from the back
	v, err := d.PollLast()
	if err != nil || v != 2 {
		t.Fatalf("PollLast expected 2, got %v err=%v", v, err)
	}
	v, err = d.PollLast()
	if err != nil || v != 1 {
		t.Fatalf("PollLast expected 1, got %v err=%v", v, err)
	}
}

// TestMixedOperations tests a sequence of mixed operations and peek behavior.
func TestMixedOperations(t *testing.T) {
	d := NewDeque[string]()

	must := func(ok bool, err error) {
		if !ok || err != nil {
			t.Fatalf("operation failed: ok=%v err=%v", ok, err)
		}
	}

	must(d.OfferFirst("b"))
	must(d.OfferLast("c"))
	must(d.OfferFirst("a")) // deque: a, b, c

	if s := d.Size(); s != 3 {
		t.Fatalf("expected size 3, got %d", s)
	}

	first, err := d.PeekFirst()
	if err != nil || first != "a" {
		t.Fatalf("PeekFirst expected 'a', got %q err=%v", first, err)
	}
	last, err := d.PeekLast()
	if err != nil || last != "c" {
		t.Fatalf("PeekLast expected 'c', got %q err=%v", last, err)
	}

	// Peeks do not change size
	if s := d.Size(); s != 3 {
		t.Fatalf("expected size 3 after peeks, got %d", s)
	}

	// Remove from both ends
	v, err := d.PollFirst()
	if err != nil || v != "a" {
		t.Fatalf("PollFirst expected 'a', got %q err=%v", v, err)
	}
	v, err = d.PollLast()
	if err != nil || v != "c" {
		t.Fatalf("PollLast expected 'c', got %q err=%v", v, err)
	}

	// Only "b" remains
	v, err = d.PeekFirst()
	if err != nil || v != "b" {
		t.Fatalf("PeekFirst expected 'b', got %q err=%v", v, err)
	}
	v, err = d.PeekLast()
	if err != nil || v != "b" {
		t.Fatalf("PeekLast expected 'b', got %q err=%v", v, err)
	}
	v, err = d.PollFirst()
	if err != nil || v != "b" {
		t.Fatalf("PollFirst expected 'b', got %q err=%v", v, err)
	}

	if !d.IsEmpty() || d.Size() != 0 {
		t.Fatalf("expected empty deque at end")
	}
}

// TestRemoveExistingAndNonExisting verifies Remove behavior for present/absent items.
func TestRemoveExistingAndNonExisting(t *testing.T) {
	d := NewDeque[int]()

	// Add elements including zero value to exercise edge cases.
	if ok, err := d.OfferLast(0); !ok || err != nil {
		t.Fatalf("OfferLast failed for 0: ok=%v err=%v", ok, err)
	}
	if ok, err := d.OfferLast(1); !ok || err != nil {
		t.Fatalf("OfferLast failed for 1: ok=%v err=%v", ok, err)
	}
	if ok, err := d.OfferLast(2); !ok || err != nil {
		t.Fatalf("OfferLast failed for 2: ok=%v err=%v", ok, err)
	}

	// Remove existing values should return true.
	if removed := d.Remove(1); !removed {
		t.Fatalf("Remove(1) expected true, got false")
	}
	if d.Size() != 2 {
		t.Fatalf("expected size 2 after removal, got %d", d.Size())
	}
	if removed := d.Remove(0); !removed {
		t.Fatalf("Remove(0) expected true (existing zero value), got false")
	}

	// Removing non-existing value should return false.
	if removed := d.Remove(42); removed {
		t.Fatalf("Remove(42) expected false, got true")
	}

	// Removing zero value again should return false.
	if removed := d.Remove(0); removed {
		t.Fatalf("Remove(0) expected false when not present, got true")
	}
}

// TestErrorsOnEmptyAfterDrains ensures error paths after draining the deque.
func TestErrorsOnEmptyAfterDrains(t *testing.T) {
	d := NewDeque[int]()
  
	_, _ = d.OfferFirst(10)
	_, _ = d.OfferLast(20)
	_, _ = d.PollFirst()
	_, _ = d.PollLast()

	if !d.IsEmpty() {
		t.Fatalf("expected empty after draining")
	}
	if _, err := d.PollFirst(); err == nil {
		t.Fatalf("expected error on PollFirst after draining")
	}
	if _, err := d.PollLast(); err == nil {
		t.Fatalf("expected error on PollLast after draining")
	}
	if _, err := d.PeekFirst(); err == nil {
		t.Fatalf("expected error on PeekFirst after draining")
	}
	if _, err := d.PeekLast(); err == nil {
		t.Fatalf("expected error on PeekLast after draining")
	}
}

func TestConcurrency(t *testing.T) {
	const (
		producers   = 8
		consumers   = 8
		perProducer = 1000
	)
	total := producers * perProducer

	d := NewDeque[int]()

	var consumed int64

	var wgProducers sync.WaitGroup
	wgProducers.Add(producers)

	// Start producers (half OfferFirst, half OfferLast).
	for p := 0; p < producers; p++ {
		p := p
		go func() {
			defer wgProducers.Done()
			start := p * perProducer
			for i := 0; i < perProducer; i++ {
				val := start + i
				if p%2 == 0 {
					if _, err := d.OfferFirst(val); err != nil {
						t.Errorf("OfferFirst error: %v", err)
					}
				} else {
					if _, err := d.OfferLast(val); err != nil {
						t.Errorf("OfferLast error: %v", err)
					}
				}
			}
		}()
	}

	var producersDone int32
	go func() {
		wgProducers.Wait()
		atomic.StoreInt32(&producersDone, 1)
	}()

	var wgConsumers sync.WaitGroup
	wgConsumers.Add(consumers)

	for c := 0; c < consumers; c++ {
		go func() {
			defer wgConsumers.Done()
			for {
				// First try from front.
				if _, err := d.PollFirst(); err == nil {
					if atomic.AddInt64(&consumed, 1) == int64(total) {
						return
					}
					continue
				}
				// Then try from back.
				if _, err := d.PollLast(); err == nil {
					if atomic.AddInt64(&consumed, 1) == int64(total) {
						return
					}
					continue
				}
				// If nothing to consume:
				// - If producers finished and deque is empty, we're done.
				if atomic.LoadInt32(&producersDone) == 1 && d.IsEmpty() {
					return
				}
				// Yield to reduce busy-wait pressure.
				runtime.Gosched()
			}
		}()
	}

	// Global timeout to avoid hanging the test.
	timeout := time.After(10 * time.Second)

	// Wait until all consumers complete or timeout fires.
	done := make(chan struct{})
	go func() {
		wgConsumers.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Completed
	case <-timeout:
		t.Fatalf("TestConcurrency timed out")
	}

	// Validate results.
	if got := int(atomic.LoadInt64(&consumed)); got != total {
		t.Fatalf("consumed %d items; expected %d", got, total)
	}
	if !d.IsEmpty() || d.Size() != 0 {
		t.Fatalf("expected deque to be empty at the end; size=%d", d.Size())
	}
}
