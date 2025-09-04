# 📦 Go Data Structures Library

[![Go Reference](https://pkg.go.dev/badge/github.com/Zubayear/sonic.svg)](https://pkg.go.dev/github.com/Zubayear/sonic)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zubayear/sonic)](https://goreportcard.com/report/github.com/Zubayear/sonic)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build Status](https://github.com/Zubayear/sonic/actions/workflows/go.yml/badge.svg)](https://github.com/Zubayear/sonic/actions)
[![codecov](https://codecov.io/gh/Zubayear/sonic/branch/main/graph/badge.svg)](https://codecov.io/gh/Zubayear/sonic)

*Ryushin* is a **robust**, **high-performance**, and **concurrency-safe** data structures library written in pure Go.

It is designed to help Go developers stop reinventing the wheel by providing production-ready collections with clean APIs and built-in thread-safety.

This library is designed with:
- **Performance** in mind – low allocations, cache-friendly.
- **Concurrency safety** – safe variants for multi-goroutine workloads.
- **Ease of use** – clean API, well-documented, and consistent.
- **Production-readiness** – tested and benchmarked.

---

## ✨ Features

- Common collections:
  - `Stack`
  - `Queue`
  - `Deque`
  - `LinkedList`
  - `Set`
- Tree structures:
  - `Binary Search Tree`
  - `AVL Tree`
  - `Red-Black Tree`
  - `Trie`
- Priority structures:
  - `Heap` (min & max)
  - `PriorityQueue`
- Thread-safe variants with `sync.RWMutex`.
- Custom iterators for all collections.

---
## 🚀 Why Sonic?
- **Performance-oriented**: Optimized for low allocations and cache-friendly operations
- **Concurrency-safe**: Multi-goroutine safe implementations
- **Ease of use**: Idiomatic APIs, consistent design
- **Production-ready**: Tested, benchmarked, and reliable

---
## 💡 Quick Example
```
package main

import (
    "fmt"
    "github.com/Zubayear/ryushin"
)

func main() {
    tr := NewTrie()

	words := []string{"hello", "helium", "he", "hero"}
	for _, w := range words {
		tr.Insert(w)
	}
	
	ok := tr.Search("hello")
}
```

---
## 💪 Benchmark
```
cpu: 12th Gen Intel(R) Core(TM) i7-1255U
pkg: github.com/Zubayear/ryushin/linkedlist
BenchmarkBinaryHeapAdd-12              	      46	  22188485 ns/op	 8923515 B/op	      28 allocs/op
BenchmarkBinaryHeapPeek-12             	42151258	        27.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapPoll-12             	      16	  67635712 ns/op	 8923558 B/op	      29 allocs/op
BenchmarkBinaryHeapClear-12            	31316138	        40.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapAddParallel-12      	      76	  30528826 ns/op	 8631298 B/op	      25 allocs/op
BenchmarkBinaryHeapPeekParallel-12     	 7108928	       161.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapPollParallel-12     	      46	  24318991 ns/op	 8923681 B/op	      29 allocs/op
BenchmarkBinaryHeapClearParallel-12    	 5849760	       184.9 ns/op	       0 B/op	       0 allocs/op

pkg: github.com/Zubayear/ryushin/priorityqueue
BenchmarkBinaryHeapAdd-12              	      72	  15216388 ns/op	 8923576 B/op	      28 allocs/op
BenchmarkBinaryHeapPeek-12             	52717822	        23.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapPoll-12             	      19	  59121358 ns/op	 8923823 B/op	      29 allocs/op
BenchmarkBinaryHeapClear-12            	32707885	        39.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapAddParallel-12      	      55	  22358065 ns/op	 8690696 B/op	      27 allocs/op
BenchmarkBinaryHeapPeekParallel-12     	 6294998	       181.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapPollParallel-12     	      50	  28703796 ns/op	 8923574 B/op	      29 allocs/op
BenchmarkBinaryHeapClearParallel-12    	 8554953	       134.8 ns/op	       0 B/op	       0 allocs/op

pkg: github.com/Zubayear/ryushin/queue
BenchmarkEnqueue-12            	    2134	    524674 ns/op	  261888 B/op	      10 allocs/op
BenchmarkDequeue-12            	    1317	    868599 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeek-12               	50060071	        23.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrint-12              	  143790	      8433 ns/op	    1200 B/op	      97 allocs/op
BenchmarkEnqueueParallel-12    	 9277118	       132.3 ns/op	      28 B/op	       0 allocs/op
BenchmarkDequeueParallel-12    	 6286309	       172.0 ns/op	      15 B/op	       0 allocs/op
BenchmarkPeekParallel-12       	 9191943	       132.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkEnqueueLarge-12       	     235	   5153847 ns/op	 2097104 B/op	      15 allocs/op

pkg: github.com/Zubayear/ryushin/stack
BenchmarkStackPush-12            	    2152	    540537 ns/op	  249475 B/op	       0 allocs/op
BenchmarkStackPop-12             	    1896	    668498 ns/op	  159918 B/op	    9994 allocs/op
BenchmarkStackPeek-12            	46704237	        25.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkStackPushParallel-12    	 5148144	       215.8 ns/op	      26 B/op	       0 allocs/op
BenchmarkStackPopParallel-12     	 6172747	       173.1 ns/op	      15 B/op	       0 allocs/op
BenchmarkStackPeekParallel-12    	 7440540	       161.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkStackPushLarge-12       	     228	   5278262 ns/op	 2354695 B/op	       0 allocs/op

pkg: github.com/Zubayear/ryushin/trie
BenchmarkTrieInsert-12                        	   86865	     14524 ns/op
BenchmarkTrieSearch-12                        	 8947917	       139.2 ns/op
BenchmarkTrieStartsWith-12                    	27064534	        44.18 ns/op
BenchmarkTrieGetWordsWithPrefix-12            	  654068	      1687 ns/op
BenchmarkTrieInsertParallel-12                	 2720186	       421.7 ns/op
BenchmarkTrieSearchParallel-12                	 6387027	       193.4 ns/op
BenchmarkTrieMapPrefixSearch-12               	 4287447	       273.1 ns/op
BenchmarkTrieInsertLarge-12                   	      26	  44391819 ns/op
BenchmarkTrieStartsWithParallel-12            	 6933982	       178.6 ns/op
BenchmarkTrieGetWordsWithPrefixParallel-12    	 1980440	       628.3 ns/op
BenchmarkTrieMapPrefixSearchParallel-12       	22304251	        49.68 ns/op
```
---

## 📦 Installation
```
go get github.com/Zubayear/ryushin
```
---
## 🤝 Contributing
We welcome contributions! Whether you want to:
- Fix bugs
- Add new data structures
- Improve documentation
- Write benchmarks or tests