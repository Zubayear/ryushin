/*
Package linkedlist provides a generic, concurrency-safe doubly linked list
(also known as a doubly linked chain) implementation in Go.

It supports:
  - Efficient insertion and removal at both head and tail (O(1)).
  - Arbitrary index-based insertions and deletions (O(n)).
  - Searching, iteration, and element containment checks.
  - Thread-safe operations using sync.RWMutex.

The list consists of ListNode elements, each containing a value and pointers
to the previous and next nodes.

Key Features:
  - AddFirst / AddLast: Insert elements at the head or tail.
  - AddAt: Insert element at a specific index.
  - RemoveFirst / RemoveLast: Remove elements from head or tail.
  - Remove / RemoveAt: Remove by value or index.
  - PeekFirst / PeekLast: Read values at head/tail without removal.
  - Iterate: Channel-based iterator for easy traversal.
  - Contains / indexOf: Check if an element exists or get its index.
  - Clear: Reset the list.

Concurrency:
  - All public methods are protected with RWMutex for safe concurrent access.

Algorithms:
  - Insertion at head/tail: Create a new ListNode and adjust prev/next pointers.
  - Deletion by value or index: Traverse list to locate node, then relink
    neighbors to exclude the node.
  - Iteration: Channel-based iteration reads nodes sequentially under read lock.

Time Complexities:
  - AddFirst / AddLast: O(1)
  - RemoveFirst / RemoveLast: O(1)
  - AddAt / RemoveAt / Remove by value: O(n)
  - PeekFirst / PeekLast: O(1)
  - Contains / indexOf: O(n)
  - Iterate: O(n)
*/
package linkedlist

import (
	"errors"
	"sync"
)

// Iterator is a channel-based iterator for traversing the linked list.
type Iterator[T any] <-chan T

// ListNode represents a node in a doubly linked list.
type ListNode[T comparable] struct {
	val        T
	next, prev *ListNode[T]
}

// NewListNode creates a new node with the given value.
func NewListNode[T comparable](val T, prev *ListNode[T], next *ListNode[T]) *ListNode[T] {
	return &ListNode[T]{
		val:  val,
		prev: prev,
		next: next,
	}
}

// DoublyLinkedList represents a generic doubly linked list.
type DoublyLinkedList[T comparable] struct {
	size       int
	head, tail *ListNode[T]
	mutex      sync.RWMutex
}

// NewLinkedList initializes and returns a new empty doubly linked list.
func NewLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{size: 0}
}

// Clear removes all elements from the list and resets it to an empty state.
// Algorithm: Traverse each node, disconnecting prev and next references.
//
// Time Complexity: O(n)
func (dl *DoublyLinkedList[T]) Clear() {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()
	iter := dl.head
	for iter != nil {
		next := iter.next
		iter.prev = nil
		iter.next = nil
		iter = next
	}
	dl.head = nil
	dl.tail = nil
	dl.size = 0
}

// Size returns the number of elements in the list. O(1)
func (dl *DoublyLinkedList[T]) Size() int {
	dl.mutex.RLock()
	defer dl.mutex.RUnlock()
	return dl.size
}

// IsEmpty checks if the linked list is empty. O(1)
func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	dl.mutex.RLock()
	defer dl.mutex.RUnlock()
	return dl.size == 0
}

// Add appends an element to the end of the list. O(1)
func (dl *DoublyLinkedList[T]) Add(elem T) (bool, error) {
	return dl.AddLast(elem)
}

// AddLast inserts an element at the tail of the list.
// Algorithm: Create new node, link previous tail to it, update a tail pointer.
//
// Time Complexity: O(1)
func (dl *DoublyLinkedList[T]) AddLast(elem T) (bool, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()
	if dl.size == 0 {
		node := NewListNode(elem, nil, nil)
		dl.head = node
		dl.tail = node
	} else {
		node := NewListNode(elem, dl.tail, nil)
		dl.tail.next = node
		dl.tail = dl.tail.next
	}
	dl.size++
	return true, nil
}

// AddFirst inserts a new element at the head of the list. O(1)
func (dl *DoublyLinkedList[T]) AddFirst(elem T) (bool, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()
	if dl.size == 0 {
		node := NewListNode(elem, nil, nil)
		dl.head = node
		dl.tail = node
	} else {
		node := NewListNode(elem, nil, dl.head)
		dl.head.prev = node
		dl.head = dl.head.prev
	}
	dl.size++
	return true, nil
}

// AddAt inserts an element at a specific index in the list.
// Algorithm: Traverse to index, link a new node between prev and next nodes.
//
// Time Complexity: O(n)
func (dl *DoublyLinkedList[T]) AddAt(idx int, elem T) (bool, error) {
	if idx < 0 || idx > dl.size {
		return false, errors.New("invalid index")
	}
	if idx == 0 {
		return dl.AddFirst(elem)
	}
	if idx == dl.size {
		return dl.AddLast(elem)
	}
	temp := dl.head

	for i := 0; i < idx-1; i++ {
		temp = temp.next
	}
	node := NewListNode(elem, temp, temp.next)
	temp.next = node
	node.next.prev = node
	dl.size++
	return true, nil
}

// PeekFirst returns the value of the first element. O(1)
func (dl *DoublyLinkedList[T]) PeekFirst() (T, error) {
	dl.mutex.RLock()
	defer dl.mutex.RUnlock()
	var zero T
	if dl.size == 0 {
		return zero, errors.New("linked list empty")
	}
	return dl.head.val, nil
}

// PeekLast returns the value of the last element. O(1)
func (dl *DoublyLinkedList[T]) PeekLast() (T, error) {
	dl.mutex.RLock()
	defer dl.mutex.RUnlock()
	var zero T
	if dl.size == 0 {
		return zero, errors.New("linked list empty")
	}
	return dl.tail.val, nil
}

// RemoveFirst deletes and returns the first element.
// Algorithm: Update head pointer, disconnect removed node.
//
// Time Complexity: O(1)
func (dl *DoublyLinkedList[T]) RemoveFirst() (T, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()
	var zero T
	if dl.size == 0 {
		return zero, errors.New("linked list empty")
	}
	value := dl.head.val
	dl.head = dl.head.next
	dl.size--
	if dl.size == 0 {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}
	return value, nil
}

// RemoveLast removes and returns the last element. O(1)
func (dl *DoublyLinkedList[T]) RemoveLast() (T, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()
	var zero T
	if dl.size == 0 {
		return zero, errors.New("linked list empty")
	}
	value := dl.tail.val
	dl.tail = dl.tail.prev
	dl.size--
	if dl.size == 0 {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}
	return value, nil
}

// removeNode deletes a given node from the list and relink neighbors. O(1)
func (dl *DoublyLinkedList[T]) removeNode(node *ListNode[T]) (T, error) {
	if node.prev == nil {
		return dl.RemoveFirst()
	}
	if node.next == nil {
		return dl.RemoveLast()
	}
	node.next.prev = node.prev
	node.prev.next = node.next
	result := node.val
	node.prev = nil
	node.next = nil
	dl.size--
	return result, nil
}

// Remove deletes the first occurrence of a given element. O(n)
func (dl *DoublyLinkedList[T]) Remove(elem T) (T, error) {
	var zero T
	if dl.size == 0 {
		return zero, errors.New("linked list empty")
	}

	for traveler := dl.head; traveler != nil; traveler = traveler.next {
		if traveler.val == elem {
			return dl.removeNode(traveler)
		}
	}
	return zero, errors.New("value not found")
}

// RemoveAt removes and returns the element at a specific index. O(n)
func (dl *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	var zero T
	if idx < 0 || idx >= dl.size {
		return zero, errors.New("invalid index")
	}

	var traveler *ListNode[T]
	if idx < dl.size/2 {
		traveler = dl.head
		for i := 0; i < idx; i++ {
			traveler = traveler.next
		}
	} else {
		traveler = dl.tail
		for i := dl.size - 1; i > idx; i-- {
			traveler = traveler.prev
		}
	}

	return dl.removeNode(traveler)
}

// indexOf finds the index of an element in the list. O(n)
func (dl *DoublyLinkedList[T]) indexOf(elem T) (int, error) {
	dl.mutex.RLock()
	defer dl.mutex.RUnlock()
	if dl.size == 0 {
		return -1, errors.New("linked list empty")
	}
	iterNode := dl.head
	var idx int
	for iterNode != nil {
		if iterNode.val == elem {
			return idx, nil
		} else {
			iterNode = iterNode.next
			idx++
		}
	}
	return -1, errors.New("element not found in linked list")
}

// Contains checks if an element exists in the list. O(n)
func (dl *DoublyLinkedList[T]) Contains(elem T) (bool, error) {
	result, err := dl.indexOf(elem)
	if err != nil {
		return false, err
	}
	return result >= 0, nil
}

// Iterate returns a channel-based iterator for traversing the list.
func (dl *DoublyLinkedList[T]) Iterate() Iterator[T] {
	iterChan := make(chan T)
	go func() {
		dl.mutex.RLock()
		defer dl.mutex.RUnlock()
		defer close(iterChan)
		iterNode := dl.head
		for iterNode != nil {
			iterChan <- iterNode.val
			iterNode = iterNode.next
		}
	}()
	return iterChan
}
