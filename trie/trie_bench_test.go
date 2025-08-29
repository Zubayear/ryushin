package trie

import (
	"fmt"
	"strings"
	"testing"
)

var words = []string{
	"apple", "app", "application", "apply", "banana", "band", "bandana",
	"cat", "cater", "catering", "dog", "dodge", "zebra",
}

func generateWords(n int) []string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = fmt.Sprintf("word%d", i)
	}
	return words
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := NewTrie()
		for _, word := range words {
			t.Insert(word)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	t := NewTrie()
	for _, word := range words {
		t.Insert(word)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.Search("application")
	}
}

func BenchmarkStartsWith(b *testing.B) {
	t := NewTrie()
	for _, word := range words {
		t.Insert(word)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		t.StartsWith("app")
	}
}

func BenchmarkGetWordsWithPrefix(b *testing.B) {
	t := NewTrie()
	for _, word := range words {
		t.Insert(word)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.GetWordsWithPrefix("app")
	}
}

func BenchmarkInsertParallel(b *testing.B) {
	largeWords := generateWords(10000)
	t := NewTrie()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			word := largeWords[i%len(largeWords)]
			t.Insert(word)
			i++
		}
	})
}

func BenchmarkSearchParallel(b *testing.B) {
	t := NewTrie()
	largeWords := generateWords(10000)
	for _, w := range largeWords {
		t.Insert(w)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			t.Search(largeWords[i%len(largeWords)])
			i++
		}
	})
}

func BenchmarkMapPrefixSearch(b *testing.B) {
	wordMap := make(map[string]bool)
	for _, w := range words {
		wordMap[w] = true
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		prefix := "app"
		for key := range wordMap {
			if strings.HasPrefix(key, prefix) {
				_ = key
			}
		}
	}
}

func BenchmarkInsertLarge(b *testing.B) {
	largeWords := generateWords(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := NewTrie()
		for _, w := range largeWords {
			t.Insert(w)
		}
	}
}

func BenchmarkStartsWithParallel(b *testing.B) {
	t := NewTrie()
	for _, word := range words {
		t.Insert(word)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = t.StartsWith("app")
		}
	})
}

func BenchmarkGetWordsWithPrefixParallel(b *testing.B) {
	t := NewTrie()
	for _, word := range words {
		t.Insert(word)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = t.GetWordsWithPrefix("app")
		}
	})
}

func BenchmarkMapPrefixSearchParallel(b *testing.B) {
	wordMap := make(map[string]bool)
	for _, w := range words {
		wordMap[w] = true
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			prefix := "app"
			for key := range wordMap {
				if strings.HasPrefix(key, prefix) {
					_ = key
				}
			}
		}
	})
}
