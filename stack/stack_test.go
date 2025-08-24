package stack_test

import (
	"testing"

	"github.com/Zubayear/sonic/stack"
)

func TestAllOperation(t *testing.T) {
	s := stack.NewStack[int]()
	_, _ = s.Push(1)
	_, _ = s.Push(2)
	_, _ = s.Push(3)

	value, err := s.Peek()
	if err != nil || value != 3 {
		t.Errorf("Expected top element to be 3, got %v\n", value)
	}

	size := s.Size()
	if size != 3 {
		t.Errorf("Expected size to be 3, got %v\n", size)
	}

	value, err = s.Pop()
	if err != nil || value != 3 {
		t.Errorf("Expected size to be 3, got %v\n", size)
	}

	size = s.Size()
	if size != 2 {
		t.Errorf("Expected size to be 3, got %v\n", size)
	}

}
