package priorityqueue

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	bn, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[bn.Int64()]
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

// comparator: higher Lived first, then longer Name
func personComparator(p1, p2 Person) bool {
	if p1.Lived != p2.Lived {
		return p1.Lived > p2.Lived
	}
	return len(p1.Name) > len(p2.Name)
}

// generatePeople generates n dummy Person entries for testing
func generatePeople(n int) []Person {
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		people[i] = Person{
			Name:  "Person_" + strconv.Itoa(i),
			Lived: uint(i % 100), // ages between 0-99
		}
	}
	return people
}

// Benchmark adding elements to heap with custom comparator
func BenchmarkBinaryHeapAddWithCustomComparator(b *testing.B) {
	people := generatePeople(1000) // adjust size as needed

	for i := 0; i < b.N; i++ {
		h := NewBinaryHeapWithComparator(personComparator)
		for _, p := range people {
			h.Add(p)
		}
	}
}

// Benchmark polling all elements from the heap with custom comparator
func BenchmarkBinaryHeapPollWithCustomComparator(b *testing.B) {
	people := generatePeople(1000) // adjust size as needed

	for i := 0; i < b.N; i++ {
		h := NewBinaryHeapWithComparator(personComparator)
		for _, p := range people {
			h.Add(p)
		}
		for !h.IsEmpty() {
			_, _ = h.Poll()
		}
	}
}

// BenchmarkBinaryHeapSort benchmarks the Sort() method on a BinaryHeap with custom comparator.
func BenchmarkBinaryHeapSort(b *testing.B) {
	bn, _ := rand.Int(rand.Reader, big.NewInt(10000))

	type Person struct {
		Name  string
		Lived uint
	}

	// Custom comparator: higher Lived first, tie-breaker longer Name
	cmp := func(p1, p2 Person) bool {
		if p1.Lived != p2.Lived {
			return p1.Lived > p2.Lived
		}
		return len(p1.Name) > len(p2.Name)
	}

	// Generate N random elements
	N := 10000
	people := make([]Person, N)
	for i := 0; i < N; i++ {
		people[i] = Person{
			Name:  "Person_" + strconv.Itoa(int(bn.Int64())),
			Lived: uint(bn.Uint64()),
		}
	}

	bh := NewBinaryHeapWithComparator(cmp)
	for _, p := range people {
		bh.Add(p)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bh.Sort()
	}
}
