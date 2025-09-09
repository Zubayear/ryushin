package queue

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestQueueOperations(t *testing.T) {
	q := NewQueue[int]()
	isEmpty := q.IsEmpty()
	if !isEmpty {
		t.Errorf("Expected %v, got %v\n", false, isEmpty)
	}

	q.Enqueue(1)
	q.Enqueue(4)
	q.Enqueue(79)

	size := q.Size()
	if size != 3 {
		t.Errorf("Extected %v, got %v\n", 3, size)
	}
	value, err := q.Dequeue()
	if err != nil || value != 1 {
		t.Errorf("Expected %v, got %v\n", 1, value)
	}

	value, err = q.Peek()
	if err != nil || value != 4 {
		t.Errorf("Expected %v, got %v\n", 4, value)
	}
	isFull := q.IsFull()
	if isFull {
		t.Errorf("Expected %v, got %v\n", false, isFull)
	}
	result := q.ToString()
	if !reflect.DeepEqual(result, "[4, 79]") {
		t.Errorf("ToString() = %v; want %v", result, "[4, 79]")
	}

	q.Clear()
	if q.Size() != 0 {
		t.Errorf("Expected %v, got %v\n", 0, q.Size())
	}

	v, _ := q.Peek()
	if v != 0 {
		t.Errorf("Expected %v, got %v\n", 0, v)
	}

	v, _ = q.Dequeue()
	if v != 0 {
		t.Errorf("Expected %v, got %v\n", 0, v)
	}

	for i := 0; i < 50; i++ {
		q.Enqueue(i)
	}
}

func TestQueueToArray(t *testing.T) {
	q := NewQueue[string]()
	str := "To be or not to be, that is the question"
	arr := strings.Split(str, " ")
	for _, s := range arr {
		q.Enqueue(s)
	}
	fmt.Println(q.ToArray())

	fmt.Println("Iterator: ")
	it := q.Iterator()

	for v, ok := it.Next(); ok; v, ok = it.Next() {
		fmt.Print(" ", v)
	}
	fmt.Println()

	fmt.Println("Dequeuing Iterator")
	for v, ok := it.Next(); ok; v, ok = it.Next() {
		fmt.Print(" ", v)
		it.Dequeue()
	}

	fmt.Println("Enqueuing Data")
	str1 := "There are many people in our country are illiterate"
	arr1 := strings.Split(str1, " ")
	for _, s := range arr1 {
		q.Enqueue(s)
	}

	fmt.Println("Iterator: ")
	it1 := q.Iterator()

	for v, ok := it1.Next(); ok; v, ok = it1.Next() {
		fmt.Print(" ", v)
	}

	fmt.Println()
}
