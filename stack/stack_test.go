package stack

import (
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int](5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(10)
	stack.Push(10)
	got := stack.IsEmpty()
	if got {
		t.Errorf("got = %v, want %v", got, false)
	}
}

func TestStack_IsFull(t *testing.T) {
	stack := NewStack[int](5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(10)
	stack.Push(10)
	got := stack.IsFull()
	if !got {
		t.Errorf("got = %v, want %v", got, false)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int](5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(100)
	got, err := stack.Peek()
	if err != nil {
		return
	}
	if got != 100 {
		t.Errorf("got = %v, want = %v", got, 100)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int](5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(100)
	got, err := stack.Pop()
	if err != nil {
		return
	}
	if !got {
		t.Errorf("got = %v, want = %v", got, false)
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int](5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(100)
	got, err := stack.Peek()
	if err != nil {
		return
	}
	if got != 100 {
		t.Errorf("got = %v, want = %v", got, 100)
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int](5)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(100)
	got := stack.Size()
	if got != 4 {
		t.Errorf("got = %v, want = %v", got, 4)
	}
}
