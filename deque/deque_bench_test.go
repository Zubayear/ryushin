package deque

import (
	"strconv"
	"sync"
	"testing"
)

// Benchmark OfferFirst on a growing deque.
func BenchmarkOfferFirst(b *testing.B) {
	var d Deque[int]
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := d.OfferFirst(i); err != nil {
			b.Fatalf("OfferFirst error: %v", err)
		}
	}
}

// Benchmark OfferLast on a growing deque.
func BenchmarkOfferLast(b *testing.B) {
	var d Deque[int]
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := d.OfferLast(i); err != nil {
			b.Fatalf("OfferLast error: %v", err)
		}
	}
}

// Benchmark PollFirst by preloading then draining exactly b.N elements.
func BenchmarkPollFirst(b *testing.B) {
	var d Deque[int]
	for i := 0; i < b.N; i++ {
		if _, err := d.OfferLast(i); err != nil {
			b.Fatalf("OfferLast preload error: %v", err)
		}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := d.PollFirst(); err != nil {
			b.Fatalf("PollFirst error at i=%d: %v", i, err)
		}
	}
}

// Benchmark PollLast by preloading then draining exactly b.N elements.
func BenchmarkPollLast(b *testing.B) {
	var d Deque[int]
	for i := 0; i < b.N; i++ {
		if _, err := d.OfferLast(i); err != nil {
			b.Fatalf("OfferLast preload error: %v", err)
		}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := d.PollLast(); err != nil {
			b.Fatalf("PollLast error at i=%d: %v", i, err)
		}
	}
}

// Benchmark PeekFirst; maintains at least one element to avoid errors.
func BenchmarkPeekFirst(b *testing.B) {
	var d Deque[int]
	if _, err := d.OfferLast(1); err != nil {
		b.Fatalf("OfferLast preload error: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := d.PeekFirst(); err != nil {
			b.Fatalf("PeekFirst error: %v", err)
		}
	}
}

// Benchmark PeekLast; maintains at least one element to avoid errors.
func BenchmarkPeekLast(b *testing.B) {
	var d Deque[int]
	if _, err := d.OfferLast(1); err != nil {
		b.Fatalf("OfferLast preload error: %v", err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := d.PeekLast(); err != nil {
			b.Fatalf("PeekLast error: %v", err)
		}
	}
}

// Benchmark a mixed workload: alternating front/back push and pop.
func BenchmarkMixed(b *testing.B) {
	var d Deque[int]
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			if _, err := d.OfferFirst(i); err != nil {
				b.Fatalf("OfferFirst error: %v", err)
			}
		} else {
			if _, err := d.OfferLast(i); err != nil {
				b.Fatalf("OfferLast error: %v", err)
			}
		}
		// Keep size bounded to avoid unbounded growth.
		if d.Size() > 0 && i%3 == 0 {
			if i%2 == 0 {
				if _, err := d.PollLast(); err != nil {
					b.Fatalf("PollLast error: %v", err)
				}
			} else {
				if _, err := d.PollFirst(); err != nil {
					b.Fatalf("PollFirst error: %v", err)
				}
			}
		}
	}
}

// Parallel benchmark for OfferFirst/OfferLast on a shared deque.
func BenchmarkOfferParallel(b *testing.B) {
	var d Deque[int]
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		x := 0
		for pb.Next() {
			// Alternate ends to exercise both code paths under contention.
			if x%2 == 0 {
				if _, err := d.OfferFirst(x); err != nil {
					b.Fatalf("OfferFirst error: %v", err)
				}
			} else {
				if _, err := d.OfferLast(x); err != nil {
					b.Fatalf("OfferLast error: %v", err)
				}
			}
			x++
		}
	})
}

// Parallel mixed producer/consumer: each iteration does one push and one pop to avoid emptiness.
func BenchmarkParallelMixed(b *testing.B) {
	var d Deque[int]
	// Preload a small buffer to reduce initial empty errors.
	for i := 0; i < 1024; i++ {
		_, _ = d.OfferLast(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			// One offer and one poll per iteration to keep the deque balanced.
			if i%2 == 0 {
				_, _ = d.OfferFirst(i)
				if _, err := d.PollLast(); err != nil {
					// If empty due to races, compensate with an extra offer.
					_, _ = d.OfferLast(i)
				}
			} else {
				_, _ = d.OfferLast(i)
				if _, err := d.PollFirst(); err != nil {
					_, _ = d.OfferFirst(i)
				}
			}
			i++
		}
	})
}

// Benchmark Remove for present and absent keys using strings to avoid integer equality shortcuts.
func BenchmarkRemove(b *testing.B) {
	var d Deque[string]
	// Preload with duplicates and a target key.
	for i := 0; i < 10000; i++ {
		_, _ = d.OfferLast("k" + strconv.Itoa(i%100))
	}
	// Ensure target exists many times.
	target := "k42"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			_ = d.Remove(target) // likely true
			_, _ = d.OfferLast(target)
		} else {
			_ = d.Remove("absent-key") // false path
		}
	}
}

// Benchmark Size and IsEmpty for overhead.
func BenchmarkSizeIsEmpty(b *testing.B) {
	var d Deque[int]
	var sink int
	var sinkBool bool
	_, _ = d.OfferLast(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink += d.Size()
		sinkBool = d.IsEmpty()
		if sinkBool {
			_, _ = d.OfferLast(i)
		}
	}
	_ = sink
	_ = sinkBool
}

// Optional: benchmark under mild contention with coordinated producers and consumers.
func BenchmarkCoordinatedParallel(b *testing.B) {
	var d Deque[int]
	var wg sync.WaitGroup
	iters := b.N

	producers := 4
	consumers := 4
	itemsPerProducer := iters / producers

	b.ReportAllocs()
	b.ResetTimer()

	wg.Add(producers)
	for p := 0; p < producers; p++ {
		go func(offset int) {
			defer wg.Done()
			for i := 0; i < itemsPerProducer; i++ {
				_, _ = d.OfferLast(offset + i)
			}
		}(p * itemsPerProducer)
	}

	wg.Add(consumers)
	for c := 0; c < consumers; c++ {
		go func() {
			defer wg.Done()
			drained := 0
			for drained < itemsPerProducer {
				if _, err := d.PollFirst(); err == nil {
					drained++
				}
			}
		}()
	}

	wg.Wait()
}
