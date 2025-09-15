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

// BenchmarkUnorderedSet_Insert-12                 21979269               315.3 ns/op            55 B/op          0 allocs/op
//BenchmarkUnorderedSet_Contain-12                15175297               380.9 ns/op             0 B/op          0 allocs/op
//BenchmarkUnorderedSet_Remove-12                 26081410               455.4 ns/op             0 B/op          0 allocs/op
//BenchmarkUnorderedSet_Items-12                      1653           3092915 ns/op          802818 B/op          1 allocs/op
//BenchmarkUnorderedSet_StringKeys-12              6781633               746.9 ns/op            74 B/op          1 allocs/op
