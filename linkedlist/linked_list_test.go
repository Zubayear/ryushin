package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/Zubayear/sonic/linkedlist"
)

func TestLinkedListOperations(t *testing.T) {
	ll := linkedlist.NewLinkedList[int]()
	_, _ = ll.AddFirst(1)
	_, _ = ll.AddLast(40)
	_, err := ll.AddAt(2, 78)

	if err.Error() != "invalid index" {
		t.Errorf("Expected %v, got %v\n", "invalid index", err)
	}

	value, err := ll.PeekLast()
	if err != nil || value != 40 {
		t.Fatalf("Expected %v, got %v\n", 40, value)
	}

	value, err = ll.RemoveAt(1)
	if err != nil || value != 40 {
		t.Fatalf("Expected %v, got %v\n", 40, value)
	}
	size := ll.Size()
	if size != 1 {
		t.Fatalf("Expected %v, got %v\n", 1, size)
	}

	value, err = ll.RemoveAt(1)
	if err.Error() != "invalid index" {
		t.Fatalf("Expected %v, got %v\n", "invalid index", err)
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
	if err != nil {
		t.Errorf("Expected %v, got %v\n", true, 1)
	}

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

func TestRemoveAt(t *testing.T) {
	ll := linkedlist.NewLinkedList[int]()
	_, _ = ll.AddLast(10)
	_, _ = ll.AddLast(20)
	_, _ = ll.AddLast(30)
	_, _ = ll.AddLast(40)
	_, _ = ll.AddLast(50)

	val, err := ll.RemoveAt(0)
	if err != nil || val != 10 {
		t.Errorf("expected 10, got %v, err: %v", val, err)
	}

	val, err = ll.RemoveAt(ll.Size() - 1)
	if err != nil || val != 50 {
		t.Errorf("expected 50, got %v, err: %v", val, err)
	}

	val, err = ll.RemoveAt(1)
	if err != nil || val != 30 {
		t.Errorf("expected 30, got %v, err: %v", val, err)
	}

	_, err = ll.RemoveAt(10)
	if err == nil {
		t.Errorf("expected error for invalid index")
	}

	val, err = ll.RemoveAt(0)
	if err != nil || val != 20 {
		t.Errorf("expected 20, got %v, err: %v", val, err)
	}

	_, err = ll.RemoveAt(1)
	if err.Error() != "invalid index" {
		t.Errorf("Expected %v, got %v", "invalid index", err)
	}

	val, err = ll.RemoveAt(0)
	if err != nil || val != 40 {
		t.Errorf("expected 20, got %v, err: %v", val, err)
	}

	_, err = ll.RemoveAt(0)
	if err == nil {
		t.Errorf("expected error when removing from empty list")
	}
}
