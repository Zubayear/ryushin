package stack_test

import (
	"testing"

	"github.com/Zubayear/sonic/stack"
)

func TestStackBasicOperations(t *testing.T) {
	s := stack.NewStack[int]()

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	for i := 0; i < 16; i++ {
		ok, err := s.Push(i)
		if !ok || err != nil {
			t.Errorf("Push failed at i=%d, err=%v", i, err)
		}
	}
	if s.Size() != 16 {
		t.Errorf("Expected size 16, got %d", s.Size())
	}
	if !s.IsFull() {
		t.Errorf("Expected stack to be full")
	}

	ok, err := s.Push(16)
	if !ok || err != nil {
		t.Errorf("Push failed after resizing, err=%v", err)
	}
	if s.Size() != 17 {
		t.Errorf("Expected size 17 after resize, got %d", s.Size())
	}
	if s.IsFull() {
		t.Errorf("Stack should not be full after resize")
	}

	val, err := s.Peek()
	if err != nil || val != 16 {
		t.Errorf("Peek expected 16, got %d, err=%v", val, err)
	}

	for i := 16; i >= 0; i-- {
		val, err := s.Pop()
		if err != nil || val != i {
			t.Errorf("Pop expected %d, got %d, err=%v", i, val, err)
		}
	}
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty after popping all elements")
	}

	if _, err := s.Pop(); err == nil {
		t.Errorf("Expected error when popping from empty stack")
	}

	if _, err := s.Peek(); err == nil {
		t.Errorf("Expected error when peeking empty stack")
	}

	for i := 0; i < 5; i++ {
		_, _ = s.Push(i)
	}

	val, err = s.ValueAt(0)
	if err != nil || val != 4 {
		t.Errorf("ValueAt(0) expected 4, got %d, err=%v", val, err)
	}
	val, err = s.ValueAt(2)
	if err != nil || val != 2 {
		t.Errorf("ValueAt(2) expected 2, got %d, err=%v", val, err)
	}

	if _, err := s.ValueAt(-1); err == nil {
		t.Errorf("Expected error for ValueAt(-1)")
	}
	if _, err := s.ValueAt(5); err == nil {
		t.Errorf("Expected error for ValueAt(pos >= size)")
	}

	s.Clear()
	if _, err := s.ValueAt(0); err == nil {
		t.Errorf("Expected %v, got %v", "stack empty", err)
	}
}
