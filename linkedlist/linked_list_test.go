package linkedlist_test

import (
	"fmt"
	"github.com/Zubayear/sonic/linkedlist"
	"testing"
)

func TestLinkedListOperations(t *testing.T) {
	ll := linkedlist.NewLinkedList[int]()
	_, _ = ll.AddFirst(1)
	_, _ = ll.AddLast(40)
	_, _ = ll.AddAt(2, 78)

	value, err := ll.PeekLast()
	if err != nil || value != 78 {
		t.Fatalf("Expected %v, got %v\n", 78, value)
	}

	value, err = ll.RemoveAt(1)
	if err != nil || value != 40 {
		t.Fatalf("Expected %v, got %v\n", 40, value)
	}
	size := ll.Size()
	if size != 2 {
		t.Fatalf("Expected %v, got %v\n", 2, size)
	}

	value, err = ll.RemoveAt(1)
	if err != nil || value != 78 {
		t.Fatalf("Expected %v, got %v\n", 78, value)
	}

	size = ll.Size()
	if size != 1 {
		t.Fatalf("Expected %v, got %v\n", 1, size)
	}

	contains, err := ll.Contains(90)
	if err != nil && contains != false {
		t.Fatalf("Expected %v, got %v\n", false, contains)
	}

	last, err := ll.PeekLast()
	first, err := ll.PeekFirst()
	if err != nil || last != first {
		t.Fatalf("Expected %v, got %v\n", true, contains)
	}

	remove, err := ll.Remove(90)
	if err != nil && remove != -1 {
		t.Fatalf("Expected %v, got %v\n", -1, remove)
	}

	remove, err = ll.Remove(1)
	if err != nil || remove != 1 {
		t.Fatalf("Expected %v, got %v\n", 1, remove)
	}

	size = ll.Size()
	if size != 0 {
		t.Fatalf("Expected %v, got %v\n", 0, size)
	}

	_, _ = ll.Add(10)
	_, _ = ll.Add(20)
	ll.Clear()

	size = ll.Size()
	if size != 0 {
		t.Fatalf("Expected %v, got %v\n", 0, size)
	}

	iter := ll.Iterate()
	for elem := range iter {
		fmt.Println(elem)
	}
}
