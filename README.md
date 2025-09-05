# 📦 Go Data Structures Library

[![Go Reference](https://pkg.go.dev/badge/github.com/Zubayear/ryushin.svg)](https://pkg.go.dev/github.com/Zubayear/ryushin)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zubayear/ryushin)](https://goreportcard.com/report/github.com/Zubayear/ryushin)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build Status](https://github.com/Zubayear/ryushin/actions/workflows/go.yml/badge.svg)](https://github.com/Zubayear/ryushin/actions)
[![codecov](https://codecov.io/gh/Zubayear/ryushin/branch/main/graph/badge.svg)](https://codecov.io/gh/Zubayear/ryushin)

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
## 🚀 Why Ryushin?
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
make bench
"Running benchmarks for package: ./... with BENCH='.'"
go test -run . -bench '.' -benchmem -benchtime=5s  \
-count=1 -cpu 12 -tags '' -timeout 10m  ./...

cpu: 12th Gen Intel(R) Core(TM) i7-1255U
pkg: github.com/Zubayear/ryushin/deque
BenchmarkOfferFirst-12                  65113488                82.23 ns/op           24 B/op          1 allocs/op
BenchmarkOfferLast-12                   63029910                83.28 ns/op           24 B/op          1 allocs/op
BenchmarkPollFirst-12                   140796646               42.88 ns/op            0 B/op          0 allocs/op
BenchmarkPollLast-12                    138808090               41.84 ns/op            0 B/op          0 allocs/op
BenchmarkPeekFirst-12                   280668638               21.35 ns/op            0 B/op          0 allocs/op
BenchmarkPeekLast-12                    278238441               21.47 ns/op            0 B/op          0 allocs/op
BenchmarkMixed-12                       46330650               116.7 ns/op            24 B/op          1 allocs/op
BenchmarkOfferParallel-12               30619899               175.5 ns/op            24 B/op          1 allocs/op
BenchmarkParallelMixed-12               19881282               272.8 ns/op            24 B/op          1 allocs/op
BenchmarkRemove-12                        206761             28530 ns/op              24 B/op          1 allocs/op
BenchmarkSizeIsEmpty-12                 139107025               42.90 ns/op            0 B/op          0 allocs/op
BenchmarkCoordinatedParallel-12         23883600               266.9 ns/op            24 B/op          1 allocs/op

pkg: github.com/Zubayear/ryushin/linkedlist
BenchmarkLinkedListAddLast-12                   63643450                84.31 ns/op           24 B/op          1 allocs/op
BenchmarkLinkedListAddFirst-12                  66705429                81.71 ns/op           24 B/op          1 allocs/op
BenchmarkLinkedListRemoveFirst-12               100000000               59.75 ns/op           15 B/op          0 allocs/op
BenchmarkLinkedListRemoveLast-12                93877467                67.03 ns/op           15 B/op          0 allocs/op
BenchmarkLinkedListAddLastParallel-12           33084061               171.6 ns/op            24 B/op          1 allocs/op
BenchmarkLinkedListRemoveFirstParallel-12       43253436               161.9 ns/op            15 B/op          0 allocs/op

pkg: github.com/Zubayear/ryushin/priorityqueue
BenchmarkBinaryHeapAdd-12                    465          12803935 ns/op         8923519 B/op         28 allocs/op
BenchmarkBinaryHeapPeek-12              257329814               22.05 ns/op            0 B/op          0 allocs/op
BenchmarkBinaryHeapPoll-12                    93          57015260 ns/op         8923552 B/op         29 allocs/op
BenchmarkBinaryHeapClear-12             159918670               38.79 ns/op            0 B/op          0 allocs/op
BenchmarkBinaryHeapAddParallel-12            333          18108846 ns/op         8616459 B/op         25 allocs/op
BenchmarkBinaryHeapPeekParallel-12      100000000              151.2 ns/op             0 B/op          0 allocs/op
BenchmarkBinaryHeapPollParallel-12           295          19365797 ns/op         8923597 B/op         29 allocs/op
BenchmarkBinaryHeapClearParallel-12     57560858               115.2 ns/op             0 B/op          0 allocs/op

pkg: github.com/Zubayear/ryushin/queue
BenchmarkEnqueue-12                12412            448842 ns/op          261888 B/op         10 allocs/op
BenchmarkDequeue-12                 6207            807221 ns/op               0 B/op          0 allocs/op
BenchmarkPeek-12                264576570               21.79 ns/op            0 B/op          0 allocs/op
BenchmarkPrint-12                 663604              7726 ns/op            1200 B/op         97 allocs/op
BenchmarkEnqueueParallel-12     50850387               119.8 ns/op            21 B/op          0 allocs/op
BenchmarkDequeueParallel-12     38319889               156.3 ns/op            15 B/op          0 allocs/op
BenchmarkPeekParallel-12        77795079                90.87 ns/op            0 B/op          0 allocs/op
BenchmarkEnqueueLarge-12            1203           4826167 ns/op         2097105 B/op         15 allocs/op

pkg: github.com/Zubayear/ryushin/stack
BenchmarkStackPush-12                      10000            511898 ns/op          214749 B/op          0 allocs/op
BenchmarkStackPop-12                       10120            603000 ns/op          159984 B/op       9999 allocs/op
BenchmarkStackPeek-12                   274523893               22.38 ns/op            0 B/op          0 allocs/op
BenchmarkStackPushParallel-12           51585878               115.8 ns/op            20 B/op          0 allocs/op
BenchmarkStackPopParallel-12            35868969               165.2 ns/op            15 B/op          0 allocs/op
BenchmarkStackPeekParallel-12           29385316               202.9 ns/op             0 B/op          0 allocs/op
BenchmarkStackPushLarge-12                  1176           5192140 ns/op         1826091 B/op          0 allocs/op

pkg: github.com/Zubayear/ryushin/trie
BenchmarkTrieInsert-12                            683140              7381 ns/op            7696 B/op        121 allocs/op
BenchmarkTrieSearch-12                          57652837                93.44 ns/op            0 B/op          0 allocs/op
BenchmarkTrieStartsWith-12                      164599136               35.48 ns/op            0 B/op          0 allocs/op
BenchmarkTrieGetWordsWithPrefix-12               4593763              1286 ns/op             205 B/op         13 allocs/op
BenchmarkTrieInsertParallel-12                  21762350               249.8 ns/op             0 B/op          0 allocs/op
BenchmarkTrieSearchParallel-12                  30905372               191.9 ns/op             0 B/op          0 allocs/op
BenchmarkTrieMapPrefixSearch-12                 24322260               244.7 ns/op             0 B/op          0 allocs/op
BenchmarkTrieInsertLarge-12                          180          35236579 ns/op        11120945 B/op     240015 allocs/op
BenchmarkTrieStartsWithParallel-12              35996709               188.2 ns/op             0 B/op          0 allocs/op
BenchmarkTrieGetWordsWithPrefixParallel-12      10666076               546.3 ns/op           205 B/op         13 allocs/op
BenchmarkTrieMapPrefixSearchParallel-12         135850008               45.26 ns/op            0 B/op          0 allocs/op
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
