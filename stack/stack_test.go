package stack_test

import (
	"errors"
	"testing"

	"github.com/Zubayear/sonic/stack"
)

func TestAllOperation(t *testing.T) {
	s := stack.NewStack[int]()
	_, _ = s.Push(1)
	_, _ = s.Push(2)
	_, _ = s.Push(3)

	size := s.Size()
	if size != 3 {
		t.Errorf("Expected %v, got %v\n", size, 3)
	}

	v, _ := s.ValueAt(0)
	if v != 3 {
		t.Errorf("Expected %v, got %v\n", 3, v)
	}

	v, err := s.ValueAt(100)
	if errors.Is(err, errors.New("invalid position")) {
		t.Errorf("Expected %v, got %v\n", errors.New("invalid position"), err)
	}
	value, err := s.Peek()
	if err != nil || value != 3 {
		t.Errorf("Expected top element to be 3, got %v\n", value)
	}

	size = s.Size()
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

	_, _ = s.Pop()
	_, _ = s.Pop()
	_, err = s.Peek()
	if errors.Is(err, errors.New("stack empty")) {
		t.Errorf("Expected %v, got %v\n", errors.New("stack empty"), err)
	}

	_, err = s.Pop()
	if errors.Is(err, errors.New("stack empty")) {
		t.Errorf("Expected %v, got %v\n", errors.New("stack empty"), err)
	}

	for i := 0; i < 100; i++ {
		_, _ = s.Push(i)
	}

}
