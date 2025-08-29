package linkedlist

import (
	"testing"
)

func BenchmarkLinkedListAddLast(b *testing.B) {
	dl := NewLinkedList[int]()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dl.AddLast(i)
	}
}

func BenchmarkLinkedListAddFirst(b *testing.B) {
	dl := NewLinkedList[int]()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = dl.AddFirst(i)
	}
}

func BenchmarkLinkedListRemoveFirst(b *testing.B) {
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

func BenchmarkLinkedListRemoveLast(b *testing.B) {
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

func BenchmarkLinkedListAddLastParallel(b *testing.B) {
	dl := NewLinkedList[int]()
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = dl.AddLast(1)
		}
	})
}

func BenchmarkLinkedListRemoveFirstParallel(b *testing.B) {
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
