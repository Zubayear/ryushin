# 📦 Go Data Structures Library

[![Go Reference](https://pkg.go.dev/badge/github.com/<your-username>/<repo-name>.svg)](https://pkg.go.dev/github.com/<your-username>/<repo-name>)
[![Go Report Card](https://goreportcard.com/badge/github.com/<your-username>/<repo-name>)](https://goreportcard.com/report/github.com/<your-username>/<repo-name>)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A **robust, efficient, and extensible** data structures library written in **pure Go**.

This library is designed with:
- **Performance** in mind – low allocations, cache-friendly.
- **Concurrency safety** (optional) – safe variants for multi-goroutine workloads.
- **Ease of use** – clean API, well-documented, and consistent.
- **Production-readiness** – tested and benchmarked.

---

## ✨ Features

- Common collections:
  - `Stack`
  - `Queue`
  - `Deque`
  - `LinkedList` (singly & doubly linked)
  - `HashMap` (chained & open addressing)
  - `Set` (hash-based & tree-based)
- Tree structures:
  - `Binary Search Tree`
  - `AVL Tree`
  - `Red-Black Tree`
  - `Trie`
- Graph representations:
  - Adjacency list
  - Adjacency matrix
- Priority structures:
  - `Heap` (min & max)
  - `PriorityQueue`
- Thread-safe variants with `sync.Mutex` or `sync.RWMutex`.
- Custom iterators for all collections.

---

## 📦 Installation

```bash
go get github.com/<your-username>/<repo-name>
