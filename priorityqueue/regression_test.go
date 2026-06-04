package priorityqueue

import (
	"sync"
	"testing"
	"time"
)

// TestIsEmptyConcurrentNoDeadlock guards the regression where IsEmpty() took a
// read lock and then called Size(), which took the read lock again. sync.RWMutex
// read locks are not reentrant, so with a contending writer this could deadlock.
func TestIsEmptyConcurrentNoDeadlock(t *testing.T) {
	bh := NewBinaryHeap[int]()
	done := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for w := 0; w < 4; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 200000; i++ {
					bh.Add(i)
					bh.Poll()
				}
			}()
		}
		for r := 0; r < 8; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 200000; i++ {
					_ = bh.IsEmpty()
				}
			}()
		}
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(20 * time.Second):
		t.Fatal("DEADLOCK: IsEmpty()+Size() recursive RLock with contending writer")
	}
}
