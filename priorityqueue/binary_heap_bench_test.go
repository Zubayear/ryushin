package priorityqueue

import (
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateData(n int) []string {
	data := make([]string, n)
	for i := 0; i < n; i++ {
		data[i] = randSeq(10)
	}
	return data
}

// ---------------------------
// Sequential Benchmarks
// ---------------------------

func BenchmarkBinaryHeapAdd(b *testing.B) {
	data := generateData(100000)
	bh := NewBinaryHeap[string]()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			bh.Add(v)
		}
		bh.Clear()
	}
}

func BenchmarkBinaryHeapPeek(b *testing.B) {
	data := generateData(100000)
	bh := NewBinaryHeap[string]()
	for _, v := range data {
		bh.Add(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = bh.Peek()
	}
}

func BenchmarkBinaryHeapPoll(b *testing.B) {
	data := generateData(100000)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bh := NewBinaryHeap[string]()
		for _, v := range data {
			bh.Add(v)
		}
		for !bh.IsEmpty() {
			_, _ = bh.Poll()
		}
	}
}

func BenchmarkBinaryHeapClear(b *testing.B) {
	data := generateData(100000)
	bh := NewBinaryHeap[string]()
	for _, v := range data {
		bh.Add(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bh.Clear()
	}
}

// ---------------------------
// Parallel Benchmarks
// ---------------------------

func BenchmarkBinaryHeapAddParallel(b *testing.B) {
	data := generateData(100000)
	bh := NewBinaryHeap[string]()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, v := range data {
				bh.Add(v)
			}
			bh.Clear()
		}
	})
}

func BenchmarkBinaryHeapPeekParallel(b *testing.B) {
	data := generateData(100000)
	bh := NewBinaryHeap[string]()
	for _, v := range data {
		bh.Add(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = bh.Peek()
		}
	})
}

func BenchmarkBinaryHeapPollParallel(b *testing.B) {
	data := generateData(100000)
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bh := NewBinaryHeap[string]()
			for _, v := range data {
				bh.Add(v)
			}
			for !bh.IsEmpty() {
				_, _ = bh.Poll()
			}
		}
	})
}

func BenchmarkBinaryHeapClearParallel(b *testing.B) {
	data := generateData(100000)
	bh := NewBinaryHeap[string]()
	for _, v := range data {
		bh.Add(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bh.Clear()
		}
	})
}
