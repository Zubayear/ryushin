package priorityqueue

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var i any = -1

type BinaryHeap[T constraints.Ordered] struct {
	data []T
}

func NewBinaryHeap[T constraints.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{data: make([]T, 0)}
}

func (bh *BinaryHeap[T]) IsEmpty() bool {
	return bh.Size() == 0
}

func (bh *BinaryHeap[T]) Clear() {
	bh.data = nil
}

func (bh *BinaryHeap[T]) Size() int {
	return len(bh.data)
}

func (bh *BinaryHeap[T]) Peek() (T, error) {
	if bh.IsEmpty() {
		return i.(T), errors.New("heap empty")
	}
	return bh.data[0], nil
}

func (bh *BinaryHeap[T]) Poll() (T, error) {
	if bh.IsEmpty() {
		return i.(T), errors.New("heap empty")
	}
	return bh.removeAt(0)
}

func (bh *BinaryHeap[T]) removeAt(k int) (T, error) {
	if bh.IsEmpty() {
		return i.(T), errors.New("heap empty")
	}

	first := bh.data[k]
	last := bh.data[bh.Size()-1]
	bh.data[0] = last
	if bh.Size() > 0 {
		bh.data = bh.data[:bh.Size()-1]
	}

	parent := 0
	child := 2 * (parent + 1)
	for child < bh.Size()-1 {
		if bh.data[child+1] < bh.data[child] {
			child = child + 1
		}
		if bh.data[child] < bh.data[parent] {
			bh.swap(child, parent)
			parent = child
			child = 2 * child
		} else {
			break
		}
	}

	return first, nil
}

func (bh *BinaryHeap[T]) Add(val T) {
	bh.data = append(bh.data, val)
	idxOfLastElem := bh.Size() - 1
	bh.swim(idxOfLastElem)
}

func (bh *BinaryHeap[T]) swap(i, j int) {
	temp := bh.data[i]
	bh.data[i] = bh.data[j]
	bh.data[j] = temp
}

func (bh *BinaryHeap[T]) swim(k int) {
	parent := (k - 1) / 2
	// compare with parent if it's less then swap
	for k > 0 && bh.data[parent] > bh.data[k] {
		bh.swap(parent, k)
		k = parent
		parent = (k - 1) / 2
	}
}

//func (bh *binaryHeap[T]) resizeHeap() {
//	bh.cap = bh.cap * 2
//	newData := make([]T, bh.cap)
//	copy(newData, bh.data)
//	bh.data = newData
//}

func (bh *BinaryHeap[T]) Print() {
	fmt.Println(bh.data)
}
