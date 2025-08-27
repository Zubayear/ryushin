package linkedlist

import (
	"testing"
)

func TestAddAndSize(t *testing.T) {
	list := NewLinkedList[int]()

	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty initially")
	}

	ok, _ := list.Add(10)
	if !ok {
		t.Errorf("Expected Add to return true")
	}
	ok, _ = list.Add(20)
	ok, _ = list.Add(30)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
}

func TestAddFirstAndAddLast(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.AddFirst(10)
	_, _ = list.AddFirst(20)
	_, _ = list.AddLast(30)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	val, _ := list.PeekFirst()
	if val != 20 {
		t.Errorf("Expected first element 20, got %d", val)
	}

	val, _ = list.PeekLast()
	if val != 30 {
		t.Errorf("Expected last element 30, got %d", val)
	}
}

func TestAddAtInvalidIndex(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)

	if _, err := list.AddAt(-1, 5); err == nil {
		t.Errorf("Expected error for negative index")
	}
	if _, err := list.AddAt(2, 5); err == nil {
		t.Errorf("Expected error for index > size")
	}
}

func TestPeekFirstAndLastOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	if _, err := list.PeekFirst(); err == nil {
		t.Errorf("Expected error on empty list for PeekFirst")
	}
	if _, err := list.PeekLast(); err == nil {
		t.Errorf("Expected error on empty list for PeekLast")
	}
}

func TestRemoveFirstAndLast(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)

	val, _ := list.RemoveFirst()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	val, _ = list.RemoveLast()
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
}

func TestRemoveOnEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	if _, err := list.RemoveFirst(); err == nil {
		t.Errorf("Expected error on empty list for RemoveFirst")
	}
	if _, err := list.RemoveLast(); err == nil {
		t.Errorf("Expected error on empty list for RemoveLast")
	}

	if _, err := list.Remove(1); err == nil {
		t.Errorf("Expected error on empty list for Remove")
	}
}

func TestRemoveSpecificElement(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)

	val, err := list.Remove(20)
	if err != nil || val != 20 {
		t.Errorf("Expected 20, got %d, err: %v", val, err)
	}

	if _, err := list.Remove(100); err == nil {
		t.Errorf("Expected error for element not in list")
	}
}

func TestRemoveAt(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)

	val, _ := list.RemoveAt(1)
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	val, _ = list.RemoveAt(0)
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	val, _ = list.RemoveAt(0)
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	if _, err := list.RemoveAt(0); err == nil {
		t.Errorf("Expected error on removing from empty list")
	}
}

func TestClear(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	list.Clear()

	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty after Clear")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 after Clear")
	}
}

func TestContains(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)

	contains, _ := list.Contains(20)
	if !contains {
		t.Errorf("Expected list to contain 20")
	}

	contains, _ = list.Contains(100)
	if contains {
		t.Errorf("Expected list to not contain 100")
	}
}

func TestIterate(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(1)
	_, _ = list.Add(2)
	_, _ = list.Add(3)

	expected := []int{1, 2, 3}
	i := 0
	for val := range list.Iterate() {
		if val != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], val)
		}
		i++
	}
}

func TestRemoveLastCases(t *testing.T) {
	list := NewLinkedList[int]()

	if _, err := list.RemoveLast(); err == nil {
		t.Errorf("Expected error when removing from empty list")
	}

	_, _ = list.Add(10)
	val, err := list.RemoveLast()
	if err != nil || val != 10 {
		t.Errorf("Expected 10, got %d, err: %v", val, err)
	}
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty after removing last element")
	}

	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)
	val, err = list.RemoveLast()
	if err != nil || val != 30 {
		t.Errorf("Expected 30, got %d, err: %v", val, err)
	}

	if last, _ := list.PeekLast(); last != 20 {
		t.Errorf("Expected last element to be 20, got %d", last)
	}
}

func TestIndexOf(t *testing.T) {
	list := NewLinkedList[int]()

	if idx, err := list.indexOf(10); err == nil || idx != -1 {
		t.Errorf("Expected -1 and error on empty list, got idx=%d, err=%v", idx, err)
	}

	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)

	if idx, err := list.indexOf(20); err != nil || idx != 1 {
		t.Errorf("Expected index 1 for element 20, got idx=%d, err=%v", idx, err)
	}

	if idx, err := list.indexOf(100); err == nil || idx != -1 {
		t.Errorf("Expected -1 and error for missing element, got idx=%d, err=%v", idx, err)
	}
}

func TestRemoveAtFirstHalf(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)
	_, _ = list.Add(40)

	val, err := list.RemoveAt(1)
	if err != nil || val != 20 {
		t.Errorf("Expected 20, got %d, err: %v", val, err)
	}

	expected := []int{10, 30, 40}
	i := 0
	for v := range list.Iterate() {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
		i++
	}
}

func TestAddAt(t *testing.T) {
	list := NewLinkedList[int]()

	ok, err := list.AddAt(0, 10)
	if !ok || err != nil {
		t.Errorf("Expected AddAt(0) to succeed, got err: %v", err)
	}
	if val, _ := list.PeekFirst(); val != 10 {
		t.Errorf("Expected first element to be 10, got %d", val)
	}

	ok, err = list.AddAt(list.Size(), 30)
	if !ok || err != nil {
		t.Errorf("Expected AddAt(size) to succeed, got err: %v", err)
	}
	if val, _ := list.PeekLast(); val != 30 {
		t.Errorf("Expected last element to be 30, got %d", val)
	}

	ok, err = list.AddAt(1, 20)
	if !ok || err != nil {
		t.Errorf("Expected AddAt(1) to succeed, got err: %v", err)
	}

	expected := []int{10, 20, 30}
	i := 0
	for val := range list.Iterate() {
		if val != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], val)
		}
		i++
	}

	if _, err := list.AddAt(-1, 5); err == nil {
		t.Errorf("Expected error for negative index")
	}

	if _, err := list.AddAt(list.Size()+1, 40); err == nil {
		t.Errorf("Expected error for index greater than size")
	}
}

func TestAddAtLoopCovered(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)

	ok, err := list.AddAt(2, 25)
	if !ok || err != nil {
		t.Errorf("Expected AddAt(2) to succeed, got err: %v", err)
	}

	expected := []int{10, 20, 25, 30}
	i := 0
	for val := range list.Iterate() {
		if val != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], val)
		}
		i++
	}
}

func TestRemoveNode_LastNode(t *testing.T) {
	list := NewLinkedList[int]()
	_, _ = list.Add(10)
	_, _ = list.Add(20)
	_, _ = list.Add(30)

	val, err := list.Remove(30)
	if err != nil || val != 30 {
		t.Errorf("Expected 30 removed, got %d, err: %v", val, err)
	}

	if last, _ := list.PeekLast(); last != 20 {
		t.Errorf("Expected last element to be 20, got %d", last)
	}
}
