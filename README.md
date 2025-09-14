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

## 🚀 Why Ryushin?
- **Performance-oriented**: Optimized for low allocations and cache-friendly operations
- **Concurrency-safe**: Multi-goroutine safe implementations
- **Ease of use**: Idiomatic APIs, consistent design
- **Production-ready**: Tested, benchmarked, and reliable

## 💡 Quick Example
```go
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

## 💪 Benchmark
```bash
Running benchmarks for package: ./... with BENCH='.'
go test -run . -bench '.' -benchmem -benchtime=5s  \
        -count=1 -cpu 12 -tags '' -timeout 10m  ./...
goos: darwin
goarch: arm64
cpu: Apple M4 Pro

pkg: github.com/Zubayear/ryushin/deque
BenchmarkOfferFirst-12                  281625303               22.48 ns/op           24 B/op       1 allocs/op
BenchmarkOfferLast-12                   299147086               21.37 ns/op           24 B/op       1 allocs/op
BenchmarkPollFirst-12                   633121848                9.411 ns/op           0 B/op          0 allocs/op
BenchmarkPollLast-12                    612422688                9.729 ns/op           0 B/op          0 allocs/op
BenchmarkPeekFirst-12                   1000000000               3.783 ns/op           0 B/op          0 allocs/op
BenchmarkPeekLast-12                    1000000000               3.837 ns/op           0 B/op          0 allocs/op
BenchmarkMixed-12                       161494538               36.43 ns/op           24 B/op          1 allocs/op
BenchmarkOfferParallel-12               67327788                93.42 ns/op           24 B/op          1 allocs/op
BenchmarkParallelMixed-12               39336824               147.0 ns/op            24 B/op          1 allocs/op
BenchmarkRemove-12                        397366             15990 ns/op              24 B/op          1 allocs/op
BenchmarkSizeIsEmpty-12                 820596001                7.308 ns/op           0 B/op          0 allocs/op
BenchmarkCoordinatedParallel-12         34488332               166.9 ns/op            29 B/op          1 allocs/op

pkg: github.com/Zubayear/ryushin/linkedlist
BenchmarkLinkedListAddLast-12                   302317197               21.45 ns/op           24 B/op          1 allocs/op
BenchmarkLinkedListAddFirst-12                  288448120               22.68 ns/op           24 B/op          1 allocs/op
BenchmarkLinkedListRemoveFirst-12               426533461               14.16 ns/op           15 B/op          0 allocs/op
BenchmarkLinkedListRemoveLast-12                420148182               14.09 ns/op           15 B/op          0 allocs/op
BenchmarkLinkedListAddLastParallel-12           65639409                88.12 ns/op           24 B/op          1 allocs/op
BenchmarkLinkedListRemoveFirstParallel-12       74205880                85.45 ns/op           15 B/op          0 allocs/op

pkg: github.com/Zubayear/ryushin/priorityqueue
BenchmarkBinaryHeapAdd-12                                   1773           3332028 ns/op         8923504 B/op         28 allocs/op
BenchmarkBinaryHeapPeek-12                              1000000000               3.777 ns/op           0 B/op          0 allocs/op
BenchmarkBinaryHeapPoll-12                                   334          17841241 ns/op         8923584 B/op         30 allocs/op
BenchmarkBinaryHeapClear-12                             921234924                6.524 ns/op           0 B/op          0 allocs/op
BenchmarkBinaryHeapAddParallel-12                            600           9806767 ns/op         8531390 B/op         26 allocs/op
BenchmarkBinaryHeapPeekParallel-12                      46330063               129.2 ns/op             0 B/op          0 allocs/op
BenchmarkBinaryHeapPollParallel-12                          1664           3386634 ns/op         8923587 B/op         30 allocs/op
BenchmarkBinaryHeapClearParallel-12                     89793810                67.98 ns/op            0 B/op          0 allocs/op
BenchmarkBinaryHeapAddWithCustomComparator-12             272872             20765 ns/op           59432 B/op         12 allocs/op
BenchmarkBinaryHeapPollWithCustomComparator-12             98647             60544 ns/op           59432 B/op         12 allocs/op
BenchmarkBinaryHeapSort-12                                 38180            157167 ns/op          491584 B/op          3 allocs/op

pkg: github.com/Zubayear/ryushin/queue
BenchmarkEnqueue-12                70687             84433 ns/op          261889 B/op         10 allocs/op
BenchmarkDequeue-12                45230            132657 ns/op               0 B/op          0 allocs/op
BenchmarkPeek-12                1000000000               3.786 ns/op           0 B/op          0 allocs/op
BenchmarkPrint-12                2065641              2849 ns/op            1200 B/op         97 allocs/op
BenchmarkEnqueueParallel-12     98679813                62.03 ns/op           21 B/op          0 allocs/op
BenchmarkDequeueParallel-12     70389729                83.71 ns/op           15 B/op          0 allocs/op
BenchmarkPeekParallel-12        45079846               133.2 ns/op             0 B/op          0 allocs/op
BenchmarkEnqueueLarge-12            7184            817269 ns/op         2097115 B/op         15 allocs/op

pkg: github.com/Zubayear/ryushin/set
BenchmarkUnorderedSet_Insert-12                 80975028               158.5 ns/op            59 B/op          0 allocs/op
BenchmarkUnorderedSet_Contain-12                139308354               43.13 ns/op            0 B/op          0 allocs/op
BenchmarkUnorderedSet_Remove-12                 100000000              153.3 ns/op             0 B/op          0 allocs/op
BenchmarkUnorderedSet_Items-12                     13315            449705 ns/op          802823 B/op          1 allocs/op
BenchmarkUnorderedSet_StringKeys-12             52859263               214.8 ns/op            75 B/op          1 allocs/op
BenchmarkUnorderedSet_ParallelInsert-12         44905024               168.2 ns/op             5 B/op          0 allocs/op
BenchmarkUnorderedSet_ParallelContain-12        48636406               122.7 ns/op             0 B/op          0 allocs/op
BenchmarkUnorderedSet_ParallelRemove-12         41448546               175.4 ns/op             0 B/op          0 allocs/op

pkg: github.com/Zubayear/ryushin/stack
BenchmarkStackPush-12                      76725             78738 ns/op          223915 B/op          0 allocs/op
BenchmarkStackPop-12                       41113            146522 ns/op          159997 B/op       9999 allocs/op
BenchmarkStackPeek-12                   1000000000               3.819 ns/op           0 B/op          0 allocs/op
BenchmarkStackPushParallel-12           91891528                65.34 ns/op           23 B/op          0 allocs/op
BenchmarkStackPopParallel-12            67882197                86.95 ns/op           15 B/op          0 allocs/op
BenchmarkStackPeekParallel-12           45767899               130.5 ns/op             0 B/op          0 allocs/op
BenchmarkStackPushLarge-12                  7726            879429 ns/op         2223643 B/op          0 allocs/op

pkg: github.com/Zubayear/ryushin/trie
BenchmarkTrieInsert-12                           2441874              2456 ns/op            7696 B/op        121 allocs/op
BenchmarkTrieSearch-12                          150718530               36.52 ns/op            0 B/op          0 allocs/op
BenchmarkTrieStartsWith-12                      562287691               10.69 ns/op            0 B/op          0 allocs/op
BenchmarkTrieGetWordsWithPrefix-12              12639706               468.2 ns/op           205 B/op         13 allocs/op
BenchmarkTrieInsertParallel-12                  29085496               202.9 ns/op             0 B/op          0 allocs/op
BenchmarkTrieSearchParallel-12                  44884323               134.0 ns/op             0 B/op          0 allocs/op
BenchmarkTrieMapPrefixSearch-12                 62977689                88.46 ns/op            0 B/op          0 allocs/op
BenchmarkTrieInsertLarge-12                          435          13728165 ns/op        11120996 B/op     240015 allocs/op
BenchmarkTrieStartsWithParallel-12              44893503               136.0 ns/op             0 B/op          0 allocs/op
BenchmarkTrieGetWordsWithPrefixParallel-12      36916459               162.4 ns/op           205 B/op         13 allocs/op
BenchmarkTrieMapPrefixSearchParallel-12         388200724               15.65 ns/op            0 B/op          0 allocs/op
```

## 📦 Installation
```go
go get github.com/Zubayear/ryushin
```

## 🤝 Contributing
We welcome contributions! Whether you want to:
- Fix bugs
- Add new data structures
- Improve documentation
- Write benchmarks or tests
