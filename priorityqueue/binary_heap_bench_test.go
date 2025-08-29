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
