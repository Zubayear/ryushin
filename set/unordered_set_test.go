package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestUnorderedSet_Clear(t *testing.T) {
	set := NewUnorderedSet[string]()

	_ = set.Insert("apple")
	_ = set.Insert("banana")
	_ = set.Insert("cherry")

	set.Clear()

	if set.Size() != 0 {
		t.Errorf("Unexpected set size. Expected: %d, Got: %d", 0, set.Size())
	}

	elements := set.Items()
	if len(elements) != 0 {
		t.Error("Unexpected elements in the set after clearing")
	}
}

func TestUnorderedSet_Insert(t *testing.T) {
	set := NewUnorderedSet[string]()
	_ = set.Insert("How")
	_ = set.Insert("Are")
	_ = set.Insert("How")
	_ = set.Insert("You")

	notOk := set.Insert("You")
	if notOk {
		t.Errorf("Expected false got %v\n", notOk)
	}

	if set.Size() != 3 {
		t.Errorf("Unexpected set size. Expected: %d, Got: %d", 3, set.Size())
	}

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
	set := NewUnorderedSet[string]()

	_ = set.Insert("apple")
	_ = set.Insert("banana")
	_ = set.Insert("cherry")

	elements := set.Items()

	if len(elements) != 3 {
		t.Errorf("Unexpected number of elements. Expected: %d, Got: %d", 3, len(elements))
	}

	expectedElements := []string{"apple", "banana", "cherry"}
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
	set := NewUnorderedSet[string]()

	_ = set.Insert("apple")
	_ = set.Insert("banana")
	_ = set.Insert("cherry")

	ok := set.Remove("banana")
	if !ok {
		t.Errorf("Expected true, Got %v\n", ok)
	}

	notOk := set.Remove("king")
	if notOk {
		t.Errorf("Expected false, Got %v\n", notOk)
	}
	if set.Size() != 2 {
		t.Errorf("Unexpected set size. Expected: %d, Got: %d", 2, set.Size())
	}

	if set.Contain("banana") {
		t.Error("Element 'banana' still found in the set after removal")
	}
}

func TestUnorderedSet_Iter(t *testing.T) {
	set := NewUnorderedSet[string]()
	authors := []string{"Franz Kafka", "Fyodor Dostoevsky", "Leo Tolstoy", "Friedrich Nietzsche"}
	for _, author := range authors {
		set.Insert(author)
	}
	var actual []string
	for r := range set.Iter() {
		actual = append(actual, r)
	}
	sort.Strings(actual)
	sort.Strings(authors)
	if !reflect.DeepEqual(actual, authors) {
		t.Errorf("Expected %v, Got %v\n", authors, actual)
	}
}
