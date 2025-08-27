package linkedlist

import "errors"

var i any = -1

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
}

// NewLinkedList initializes and returns a new empty doubly linked list.
func NewLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{size: 0}
}

// Clear removes all elements from the list, resetting it to empty.
func (dl *DoublyLinkedList[T]) Clear() {
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
	return dl.size
}

// IsEmpty checks if the linked list is empty. O(1)
func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.size == 0
}

// Add appends an element to the end of the list. O(1)
func (dl *DoublyLinkedList[T]) Add(elem T) (bool, error) {
	return dl.AddLast(elem)
}

// AddLast appends a new element at the tail of the list. O(1)
func (dl *DoublyLinkedList[T]) AddLast(elem T) (bool, error) {
	if dl.IsEmpty() {
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
	if dl.IsEmpty() {
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

// AddAt inserts an element at a specific index in the list. O(n)
func (dl *DoublyLinkedList[T]) AddAt(idx int, elem T) (bool, error) {
	if idx < 0 || idx > dl.size {
		return false, errors.New("invalid index")
	}
	if idx == 0 {
		return dl.AddFirst(elem)
	}
	if idx == dl.Size() {
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
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	return dl.head.val, nil
}

// PeekLast returns the value of the last element. O(1)
func (dl *DoublyLinkedList[T]) PeekLast() (T, error) {
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	return dl.tail.val, nil
}

// RemoveFirst removes and returns the first element. O(1)
func (dl *DoublyLinkedList[T]) RemoveFirst() (T, error) {
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	value := dl.head.val
	dl.head = dl.head.next
	dl.size--
	if dl.IsEmpty() {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}
	return value, nil
}

// RemoveLast removes and returns the last element. O(1)
func (dl *DoublyLinkedList[T]) RemoveLast() (T, error) {
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	value := dl.tail.val
	dl.tail = dl.tail.prev
	dl.size--
	if dl.IsEmpty() {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}
	return value, nil
}

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
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}

	for traveler := dl.head; traveler != nil; traveler = traveler.next {
		if traveler.val == elem {
			return dl.removeNode(traveler)
		}
	}
	return i.(T), errors.New("value not found")
}

// RemoveAt removes and returns the element at a specific index. O(n)
func (dl *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	if idx < 0 || idx >= dl.size {
		return i.(T), errors.New("invalid index")
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
	if dl.IsEmpty() {
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
		defer close(iterChan)
		iterNode := dl.head
		for iterNode != nil {
			iterChan <- iterNode.val
			iterNode = iterNode.next
		}
	}()
	return iterChan
}
