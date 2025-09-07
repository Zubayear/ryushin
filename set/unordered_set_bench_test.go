package set

import (
	"strconv"
	"testing"
)

func BenchmarkUnorderedSet_Insert(b *testing.B) {
	set := NewUnorderedSet[int]()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Insert(i)
	}
}

func BenchmarkUnorderedSet_Contain(b *testing.B) {
	set := NewUnorderedSet[float32]()
	for i := 0; i < 10000000; i++ {
		_ = set.Insert(float32(i))
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contain(float32(i % 10000000))
	}
}

func BenchmarkUnorderedSet_Remove(b *testing.B) {
	set := NewUnorderedSet[int]()
	for i := 0; i < b.N; i++ {
		_ = set.Insert(i)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = set.Remove(i)
	}
}

func BenchmarkUnorderedSet_Items(b *testing.B) {
	set := NewUnorderedSet[int]()
	for i := 0; i < 100000; i++ {
		_ = set.Insert(i)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = set.Items()
	}
}

func BenchmarkUnorderedSet_StringKeys(b *testing.B) {
	set := NewUnorderedSet[string]()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = set.Insert(strconv.Itoa(i))
	}
}

func BenchmarkUnorderedSet_ParallelInsert(b *testing.B) {
	set := NewUnorderedSet[int]()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = set.Insert(i)
			i++
		}
	})
}

func BenchmarkUnorderedSet_ParallelContain(b *testing.B) {
	set := NewUnorderedSet[int]()
	for i := 0; i < 100000; i++ {
		_ = set.Insert(i)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = set.Contain(i % 100000)
			i++
		}
	})
}

func BenchmarkUnorderedSet_ParallelRemove(b *testing.B) {
	set := NewUnorderedSet[int]()
	for i := 0; i < b.N; i++ {
		_ = set.Insert(i)
	}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			_ = set.Remove(i)
			i++
		}
	})
}
