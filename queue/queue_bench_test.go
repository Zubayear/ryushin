package queue

import (
	"testing"
)

func generateData(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

func BenchmarkEnqueue(b *testing.B) {
	data := generateData(10000)
	q := NewQueue[int]()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			q.Enqueue(v)
		}
		q.Clear()
	}
}

func BenchmarkDequeue(b *testing.B) {
	data := generateData(10000)
	q := NewQueue[int]()
	for _, v := range data {
		q.Enqueue(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data); j++ {
			_, _ = q.Dequeue()
		}
		for _, v := range data {
			q.Enqueue(v)
		}
	}
}

func BenchmarkPeek(b *testing.B) {
	data := generateData(10000)
	q := NewQueue[int]()
	for _, v := range data {
		q.Enqueue(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = q.Peek()
	}
}

func BenchmarkPrint(b *testing.B) {
	data := generateData(100)
	q := NewQueue[int]()
	for _, v := range data {
		q.Enqueue(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = q.Print()
	}
}

func BenchmarkEnqueueParallel(b *testing.B) {
	data := generateData(10000)
	q := NewQueue[int]()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			q.Enqueue(data[i%len(data)])
			i++
		}
	})
}

func BenchmarkDequeueParallel(b *testing.B) {
	data := generateData(10000)
	q := NewQueue[int]()
	for _, v := range data {
		q.Enqueue(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = q.Dequeue()
		}
	})
}

func BenchmarkPeekParallel(b *testing.B) {
	data := generateData(10000)
	q := NewQueue[int]()
	for _, v := range data {
		q.Enqueue(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = q.Peek()
		}
	})
}

func BenchmarkEnqueueLarge(b *testing.B) {
	data := generateData(100000) // 100K elements
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q := NewQueue[int]()
		for _, v := range data {
			q.Enqueue(v)
		}
	}
}
