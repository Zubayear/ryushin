package linkedlist

import "errors"

var i any = -1

type Iterator[T any] <-chan T

type ListNode[T comparable] struct {
	val        T
	next, prev *ListNode[T]
}

func NewListNode[T comparable](val T) *ListNode[T] {
	return &ListNode[T]{val: val}
}

type doublyLinkedList[T comparable] struct {
	size       int
	head, tail *ListNode[T]
}

func NewLinkedList[T comparable]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{size: 0}
}

// Clear - empty this linked list
func (dl *doublyLinkedList[T]) Clear() {
	dl.head = nil
	dl.tail = nil
	dl.size = 0
}

// Size - return the size of the linked list, O(1)
func (dl *doublyLinkedList[T]) Size() int {
	return dl.size
}

// IsEmpty - return true if linked list is empty false otherwise, O(1)
func (dl *doublyLinkedList[T]) IsEmpty() bool {
	return dl.size == 0
}

// Add - add element to the tail of the linked list, O(1)
func (dl *doublyLinkedList[T]) Add(elem T) (bool, error) {
	return dl.AddLast(elem)
}

// AddLast - add a node to the tail of the linked list, O(1)
func (dl *doublyLinkedList[T]) AddLast(elem T) (bool, error) {
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

// AddFirst - add an element to the beginning of the linked list, O(1)
func (dl *doublyLinkedList[T]) AddFirst(elem T) (bool, error) {
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

// AddAt - add at specific index, O(n)
func (dl *doublyLinkedList[T]) AddAt(idx int, elem T) (bool, error) {
	if idx < 0 || idx > dl.size {
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

// PeekFirst - check the value of the first node if it exists, O(1)
func (dl *doublyLinkedList[T]) PeekFirst() (T, error) {
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	return dl.head.val, nil
}

// PeekLast - check the value of the last node if it exists, O(1)
func (dl *doublyLinkedList[T]) PeekLast() (T, error) {
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	return dl.tail.val, nil
}

// RemoveFirst - remove the first value at the head of the linked list, O(1)
func (dl *doublyLinkedList[T]) RemoveFirst() (T, error) {
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

// RemoveLast - remove the last value at the tail of the linked list, O(1)
func (dl *doublyLinkedList[T]) RemoveLast() (T, error) {
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

// Remove - remove a node with provided value, O(N)
func (dl *doublyLinkedList[T]) Remove(elem T) (T, error) {
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

// RemoveAt - remove node at idx, O(N)
func (dl *doublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	if idx < 0 || idx > dl.size {
		return i.(T), errors.New("invalid index")
	}
	if dl.IsEmpty() {
		return i.(T), errors.New("linked list empty")
	}
	iterNode := dl.head
	k := 0
	var result T
	if idx == 0 {
		return dl.RemoveFirst()
	} else if idx == dl.size-1 {
		return dl.RemoveLast()
	}
	for iterNode.next != nil {
		if k == idx-1 {
			//1<->40->78
			result = iterNode.next.val
			nnNode := iterNode.next.next
			iterNode.next = nnNode
			nnNode.prev = iterNode
			break
		}
		iterNode = iterNode.next
		k++
	}
	dl.size--
	return result, nil
}

// find the index of the element in linked list, O(N)
func (dl *doublyLinkedList[T]) indexOf(elem T) (int, error) {
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

// Contains - check if an element exists in the list, O(N)
func (dl *doublyLinkedList[T]) Contains(elem T) (bool, error) {
	result, err := dl.indexOf(elem)
	if err != nil {
		return false, err
	}
	return result >= 0, nil
}

func (dl *doublyLinkedList[T]) Iterate() Iterator[T] {
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
