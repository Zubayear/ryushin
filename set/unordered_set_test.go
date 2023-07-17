package set

import (
	"testing"
)

func TestUnorderedSet_Clear(t *testing.T) {
	set := NewUnorderedSet()

	// Add elements to the set
	set.Insert("apple")
	set.Insert("banana")
	set.Insert("cherry")

	// Clear the set
	set.Clear()

	// Check the size of the set
	if set.Size() != 0 {
		t.Errorf("Unexpected set size. Expected: %d, Got: %d", 0, set.Size())
	}

	// Check if all elements are removed from the set
	elements := set.Items()
	if len(elements) != 0 {
		t.Error("Unexpected elements in the set after clearing")
	}
}

func TestUnorderedSet_Insert(t *testing.T) {
	set := NewUnorderedSet()
	set.Insert("How")
	set.Insert("Are")
	set.Insert("How")
	set.Insert("You")

	if set.Size() != 3 {
		t.Errorf("Unexpected set size. Expected: %d, Got: %d", 3, set.Size())
	}

	// Check if elements are present in the set
	if !set.Contain("How") {
		t.Error("Element 'How' not found in the set")
	}
	if !set.Contain("Are") {
		t.Error("Element 'Are' not found in the set")
	}
	if !set.Contain("You") {
		t.Error("Element 'You' not found in the set")
	}
}

func TestUnorderedSet_Items(t *testing.T) {
	set := NewUnorderedSet()

	// Add elements to the set
	set.Insert("apple")
	set.Insert("banana")
	set.Insert("cherry")

	// Get the elements from the set
	elements := set.Items()

	// Check the number of elements
	if len(elements) != 3 {
		t.Errorf("Unexpected number of elements. Expected: %d, Got: %d", 3, len(elements))
	}

	// Check if all expected elements are present
	expectedElements := []interface{}{"apple", "banana", "cherry"}
	for _, element := range expectedElements {
		found := false
		for _, e := range elements {
			if e == element {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Element '%v' not found in the set", element)
		}
	}
}

func TestUnorderedSet_Remove(t *testing.T) {
	set := NewUnorderedSet()

	// Add elements to the set
	set.Insert("apple")
	set.Insert("banana")
	set.Insert("cherry")

	// Remove an element from the set
	set.Remove("banana")

	// Check the size of the set
	if set.Size() != 2 {
		t.Errorf("Unexpected set size. Expected: %d, Got: %d", 2, set.Size())
	}

	// Check if removed element is no longer present in the set
	if set.Contain("banana") {
		t.Error("Element 'banana' still found in the set after removal")
	}
}
