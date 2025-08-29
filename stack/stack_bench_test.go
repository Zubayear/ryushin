package stack

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

func BenchmarkPush(b *testing.B) {
	data := generateData(10000)
	s := NewStack[int]()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			_, _ = s.Push(v)
		}
	}
}

func BenchmarkPop(b *testing.B) {
	data := generateData(10000)
	s := NewStack[int]()
	for _, v := range data {
		_, _ = s.Push(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(data); j++ {
			_, _ = s.Pop()
		}
	}
}

func BenchmarkPeek(b *testing.B) {
	data := generateData(10000)
	s := NewStack[int]()
	for _, v := range data {
		_, _ = s.Push(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = s.Peek()
	}
}

func BenchmarkPushParallel(b *testing.B) {
	data := generateData(10000)
	s := NewStack[int]()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_, _ = s.Push(data[i%len(data)])
			i++
		}
	})
}

func BenchmarkPopParallel(b *testing.B) {
	data := generateData(10000)
	s := NewStack[int]()
	for _, v := range data {
		_, _ = s.Push(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = s.Pop()
		}
	})
}

func BenchmarkPeekParallel(b *testing.B) {
	data := generateData(10000)
	s := NewStack[int]()
	for _, v := range data {
		_, _ = s.Push(v)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = s.Peek()
		}
	})
}

func BenchmarkPushLarge(b *testing.B) {
	data := generateData(100000)
	s := NewStack[int]()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			_, _ = s.Push(v)
		}
	}
}
