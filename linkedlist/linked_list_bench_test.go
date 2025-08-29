package linkedlist

import (
	"testing"
)

func BenchmarkAddLast(b *testing.B) {
	dl := NewLinkedList[int]()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dl.AddLast(i)
	}
}

func BenchmarkAddFirst(b *testing.B) {
	dl := NewLinkedList[int]()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dl.AddFirst(i)
	}
}

func BenchmarkRemoveFirst(b *testing.B) {
	dl := NewLinkedList[int]()
	for i := 0; i < 100000; i++ {
		_, _ = dl.AddLast(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dl.RemoveFirst()
	}
}

func BenchmarkRemoveLast(b *testing.B) {
	dl := NewLinkedList[int]()
	for i := 0; i < 100000; i++ {
		_, _ = dl.AddLast(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dl.RemoveLast()
	}
}

func BenchmarkAddLastParallel(b *testing.B) {
	dl := NewLinkedList[int]()
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = dl.AddLast(1)
		}
	})
}

// Optional: Parallel remove benchmark
func BenchmarkRemoveFirstParallel(b *testing.B) {
	dl := NewLinkedList[int]()
	for i := 0; i < 100000; i++ {
		_, _ = dl.AddLast(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = dl.RemoveFirst()
		}
	})
}
