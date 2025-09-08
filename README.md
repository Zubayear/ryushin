# 📦 Go Data Structures Library
[![Gitpod](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod&style=flat-square)](https://gitpod.io/#https://github.com/Zubayear/ryushin)
[![Go Reference](https://pkg.go.dev/badge/github.com/Zubayear/ryushin.svg)](https://pkg.go.dev/github.com/Zubayear/ryushin)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zubayear/ryushin)](https://goreportcard.com/report/github.com/Zubayear/ryushin)
[![Continuous Integration](https://github.com/Zubayear/ryushin/actions/workflows/ci.yml/badge.svg)](https://github.com/Zubayear/ryushin/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![codecov](https://codecov.io/gh/Zubayear/ryushin/branch/main/graph/badge.svg)](https://codecov.io/gh/Zubayear/ryushin)
![](https://img.shields.io/github/repo-size/Zubayear/ryushin.svg?label=Repo%20size&style=flat-square)&nbsp;
![GitHub Release](https://img.shields.io/github/v/release/Zubayear/ryushin)
[![Security Policy](https://img.shields.io/badge/security-policy-blue.svg)](./SECURITY.md)
[![Contributing](https://img.shields.io/badge/contributing-guide-brightgreen.svg)](./CONTRIBUTING.md)

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
  - `TreeMap(Red-Black Tree/AVL Tree)`
  - `Trie`
- Priority structures:
  - `PriorityQueue(Binary Heap)` (min & max)
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

pkg: github.com/Zubayear/ryushin/deque
cpu: 12th Gen Intel(R) Core(TM) i7-1255U
BenchmarkOfferFirst-12             	47256696	       190.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkOfferLast-12              	64287664	        92.05 ns/op	      24 B/op	       1 allocs/op
BenchmarkPollFirst-12              	127510657	        45.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkPollLast-12               	185325477	        49.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekFirst-12              	455091012	        13.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeekLast-12               	458355780	        12.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkMixed-12                  	77434741	        74.48 ns/op	      24 B/op	       1 allocs/op
BenchmarkOfferParallel-12          	48596216	       129.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkParallelMixed-12          	28872110	       211.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkRemove-12                 	  364267	     18197 ns/op	      24 B/op	       1 allocs/op
BenchmarkSizeIsEmpty-12            	243685191	        25.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkCoordinatedParallel-12    	32917041	       177.6 ns/op	      24 B/op	       1 allocs/op

pkg: github.com/Zubayear/ryushin/linkedlist
BenchmarkLinkedListAddLast-12                	100000000	        57.73 ns/op	      24 B/op	       1 allocs/op
BenchmarkLinkedListAddFirst-12               	126720588	        53.49 ns/op	      24 B/op	       1 allocs/op
BenchmarkLinkedListRemoveFirst-12            	143900203	        37.39 ns/op	      15 B/op	       0 allocs/op
BenchmarkLinkedListRemoveLast-12             	156739274	        39.89 ns/op	      15 B/op	       0 allocs/op
BenchmarkLinkedListAddLastParallel-12        	50725714	       115.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkLinkedListRemoveFirstParallel-12    	66724568	        97.01 ns/op	      15 B/op	       0 allocs/op

pkg: github.com/Zubayear/ryushin/priorityqueue
BenchmarkBinaryHeapAdd-12                         	     637	   8708605 ns/op	 8923504 B/op	      28 allocs/op
BenchmarkBinaryHeapPeek-12                        	458584401	        13.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapPoll-12                        	     174	  34229119 ns/op	 8923584 B/op	      30 allocs/op
BenchmarkBinaryHeapClear-12                       	255704010	        22.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapAddParallel-12                 	     434	  12803666 ns/op	 8594236 B/op	      25 allocs/op
BenchmarkBinaryHeapPeekParallel-12                	87129517	        82.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapPollParallel-12                	     490	  11563992 ns/op	 8923586 B/op	      30 allocs/op
BenchmarkBinaryHeapClearParallel-12               	72046628	        85.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryHeapAddWithCustomComparator-12     	   77326	     84117 ns/op	   59432 B/op	      12 allocs/op
BenchmarkBinaryHeapPollWithCustomComparator-12    	   26112	    228994 ns/op	   59433 B/op	      12 allocs/op
BenchmarkBinaryHeapSort-12                        	    7824	    689732 ns/op	  491584 B/op	       3 allocs/op

pkg: github.com/Zubayear/ryushin/queue
BenchmarkEnqueue-12            	   14925	    384116 ns/op	  261888 B/op	      10 allocs/op
BenchmarkDequeue-12            	   10000	    551869 ns/op	       0 B/op	       0 allocs/op
BenchmarkPeek-12               	383929602	        15.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrint-12              	  830155	      6211 ns/op	    1200 B/op	      97 allocs/op
BenchmarkEnqueueParallel-12    	62687997	        91.91 ns/op	      17 B/op	       0 allocs/op
BenchmarkDequeueParallel-12    	52404309	       126.9 ns/op	      15 B/op	       0 allocs/op
BenchmarkPeekParallel-12       	100000000	        56.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkEnqueueLarge-12       	    1614	   4521053 ns/op	 2097105 B/op	      15 allocs/op

pkg: github.com/Zubayear/ryushin/set
BenchmarkUnorderedSet_Insert-12             	31984255	       201.7 ns/op	      75 B/op	       0 allocs/op
BenchmarkUnorderedSet_Contain-12            	27392319	       211.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedSet_Remove-12             	45859218	       224.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedSet_Items-12              	    5298	   1148956 ns/op	  802817 B/op	       1 allocs/op
BenchmarkUnorderedSet_StringKeys-12         	21114708	       393.4 ns/op	      92 B/op	       1 allocs/op
BenchmarkUnorderedSet_ParallelInsert-12     	62360078	       199.0 ns/op	       4 B/op	       0 allocs/op
BenchmarkUnorderedSet_ParallelContain-12    	87302314	        71.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnorderedSet_ParallelRemove-12     	50900384	       284.2 ns/op	       0 B/op	       0 allocs/op

pkg: github.com/Zubayear/ryushin/stack
BenchmarkStackPush-12            	   18462	    322790 ns/op	  232638 B/op	       0 allocs/op
BenchmarkStackPop-12             	   14610	    407967 ns/op	  159989 B/op	    9999 allocs/op
BenchmarkStackPeek-12            	470524170	        14.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkStackPushParallel-12    	75873013	        73.38 ns/op	      28 B/op	       0 allocs/op
BenchmarkStackPopParallel-12     	63473793	       102.8 ns/op	      15 B/op	       0 allocs/op
BenchmarkStackPeekParallel-12    	64488361	        94.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkStackPushLarge-12       	    1867	   3288300 ns/op	 2300464 B/op	       0 allocs/op

pkg: github.com/Zubayear/ryushin/trie
BenchmarkTrieInsert-12                        	 1150882	      5197 ns/op	    7696 B/op	     121 allocs/op
BenchmarkTrieSearch-12                        	100000000	        57.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieStartsWith-12                    	302768622	        20.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieGetWordsWithPrefix-12            	 8320234	       847.3 ns/op	     205 B/op	      13 allocs/op
BenchmarkTrieInsertParallel-12                	33411962	       178.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieSearchParallel-12                	82843863	        75.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieMapPrefixSearch-12               	35491470	       146.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieInsertLarge-12                   	     295	  19556853 ns/op	11120963 B/op	  240015 allocs/op
BenchmarkTrieStartsWithParallel-12            	88285148	        68.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkTrieGetWordsWithPrefixParallel-12    	24377301	       281.0 ns/op	     205 B/op	      13 allocs/op
BenchmarkTrieMapPrefixSearchParallel-12       	211659241	        26.12 ns/op	       0 B/op	       0 allocs/op
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
