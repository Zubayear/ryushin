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
func NewListNode[T comparable](val T) *ListNode[T] {
	return &ListNode[T]{val: val}
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
	newNode := NewListNode(elem)
	if dl.IsEmpty() {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.prev = dl.tail
		dl.tail.next = newNode
		dl.tail = newNode
	}
	dl.size++
	return true, nil
}

// AddFirst inserts a new element at the head of the list. O(1)
func (dl *DoublyLinkedList[T]) AddFirst(elem T) (bool, error) {
	newNode := NewListNode(elem)
	if dl.IsEmpty() {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.prev = dl.head
		dl.head.next = newNode
		dl.head = newNode
	}
	dl.size++
	return true, nil
}

// AddAt inserts an element at a specific index in the list. O(n)
func (dl *DoublyLinkedList[T]) AddAt(idx int, elem T) (bool, error) {
	if idx < 0 || idx >= dl.size {
		return false, errors.New("invalid index")
	}
	newNode := NewListNode(elem)

	if dl.IsEmpty() {
		dl.head = newNode
		dl.tail = newNode
	} else if idx == 0 {
		return dl.AddFirst(elem)
	} else if idx == dl.size {
		return dl.AddLast(elem)
	} else {
		k := 0
		iterNode := dl.head
		for iterNode != nil {
			if k == idx-1 {
				newNode.prev = iterNode
				newNode.next = iterNode.next
				iterNode.next = newNode
				iterNode.next.prev = newNode
				break
			}
			iterNode = iterNode.next
			k++
		}
	}
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

// Remove deletes the first occurrence of a given element. O(n)
func (dl *DoublyLinkedList[T]) Remove(elem T) (T, error) {
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	if elem == dl.head.val {
		return dl.RemoveFirst()
	}
	if elem == dl.tail.val {
		return dl.RemoveLast()
	}
	iterNode := dl.head
	var returnValue T
	c := 0
	for iterNode.next != nil {
		if iterNode.next.val == elem {
			returnValue = iterNode.next.val
			iterNode.next = iterNode.next.next
			iterNode.next.next.prev = iterNode
			iterNode.next = nil
			c++
		}
		iterNode = iterNode.next
	}
	if c > 0 {
		dl.size--
		return returnValue, nil
	} else {
		return i.(T), errors.New("element not found in linked list")
	}
}

// RemoveAt removes and returns the element at a specific index. O(n)
func (dl *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	if idx < 0 || idx >= dl.size {
		return i.(T), errors.New("invalid index")
	}
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	if idx == 0 {
		return dl.RemoveFirst()
	} else if idx == dl.size-1 {
		return dl.RemoveLast()
	}
	iterNode := dl.head
	for k := 0; iterNode != nil; k++ {
		if k == idx-1 {
			toRemove := iterNode.next
			iterNode.next = toRemove.next
			if toRemove.next != nil {
				toRemove.next.prev = iterNode
			}
			dl.size--
			return toRemove.val, nil
		}
		iterNode = iterNode.next
	}
	return i.(T), errors.New("invalid index")
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
