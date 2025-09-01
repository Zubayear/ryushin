package treemap

import (
	"math/rand"
	"testing"

	_ "golang.org/x/exp/constraints"
)

// go test ./... -coverprofile=coverage.out
// go tool cover -html="coverage.out"

func TestPutAndGet(t *testing.T) {
	tree := NewTreeMap[int, string]()

	tree.Put(10, "ten")
	tree.Put(20, "twenty")
	tree.Put(5, "five")

	// Test retrieval
	if val, ok := tree.Get(10); !ok || val != "ten" {
		t.Errorf("Expected 'ten', got %v", val)
	}
	if val, ok := tree.Get(20); !ok || val != "twenty" {
		t.Errorf("Expected 'twenty', got %v", val)
	}
	if val, ok := tree.Get(5); !ok || val != "five" {
		t.Errorf("Expected 'five', got %v", val)
	}

	// Non-existing key
	if _, ok := tree.Get(100); ok {
		t.Errorf("Expected key 100 to not exist")
	}
}

func TestOverwriteValue(t *testing.T) {
	tree := NewTreeMap[int, string]()
	tree.Put(10, "ten")
	tree.Put(10, "TEN") // overwrite

	if val, ok := tree.Get(10); !ok || val != "TEN" {
		t.Errorf("Expected 'TEN', got %v", val)
	}
}

func TestSize(t *testing.T) {
	tree := NewTreeMap[int, string]()
	tree.Put(10, "ten")
	tree.Put(20, "twenty")
	tree.Put(30, "thirty")
	tree.Put(20, "TWENTY") // overwrite should not increase size

	if tree.Size() != 3 {
		t.Errorf("Expected size 3, got %d", tree.Size())
	}
}

func TestDelete(t *testing.T) {
	tree := NewTreeMap[int, string]()
	tree.Put(10, "ten")
	tree.Put(20, "twenty")
	tree.Put(5, "five")

	tree.Remove(20)

	if _, ok := tree.Get(20); ok {
		t.Errorf("Expected key 20 to be deleted")
	}
	if tree.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", tree.Size())
	}

	tree.Remove(100) // deleting non-existing key
	if tree.Size() != 2 {
		t.Errorf("Size changed after deleting non-existing key")
	}
}

func TestMinAndMax(t *testing.T) {
	tree := NewTreeMap[int, string]()
	tree.Put(10, "ten")
	tree.Put(1, "one")
	tree.Put(50, "fifty")

	if key, val := tree.Min(); key != 1 || val != "one" {
		t.Errorf("Expected Min (1, one), got (%d, %s)", key, val)
	}
	if key, val := tree.Max(); key != 50 || val != "fifty" {
		t.Errorf("Expected Max (50, fifty), got (%d, %s)", key, val)
	}
}

func TestEmptyMinMax(t *testing.T) {
	tree := NewTreeMap[int, string]()

	k, v := tree.Min()
	if k != 0 {
		t.Errorf("Expected 0, Got %v\n", k)
	}
	if v != "" {
		t.Errorf("Expected \"\", Got %v\n", v)
	}
	k, v = tree.Max()
	if k != 0 {
		t.Errorf("Expected 0, Got %v\n", k)
	}
	if v != "" {
		t.Errorf("Expected \"\", Got %v\n", v)
	}
}

func TestEdgeCases(t *testing.T) {
	tree := NewTreeMap[int, string]()

	// Insert duplicate keys
	tree.Put(10, "ten")
	tree.Put(10, "TEN")
	if val, _ := tree.Get(10); val != "TEN" {
		t.Errorf("Expected TEN after overwrite, got %s", val)
	}

	// Delete root
	tree.Remove(10)
	if _, ok := tree.Get(10); ok {
		t.Errorf("Expected root to be deleted")
	}

	// Insert in sorted order (worst case)
	for i := 1; i <= 1000; i++ {
		tree.Put(i, "val")
	}
	if tree.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", tree.Size())
	}
}

func TestRandomInsertDelete(t *testing.T) {
	tree := NewTreeMap[int, int]()
	n := 1000

	// Insert random values
	for i := 0; i < n; i++ {
		val := rand.Intn(10000)
		tree.Put(val, val)
	}

	// Check size (should be <= n due to duplicates)
	if tree.Size() > n {
		t.Errorf("Size too large: %d", tree.Size())
	}

	// Delete some random keys
	for i := 0; i < n/2; i++ {
		key := rand.Intn(10000)
		tree.Remove(key)
	}
}

func TestWithStrings(t *testing.T) {
	tree := NewTreeMap[string, string]()
	tree.Put("apple", "fruit")
	tree.Put("banana", "fruit")
	tree.Put("carrot", "vegetable")

	if val, ok := tree.Get("apple"); !ok || val != "fruit" {
		t.Errorf("Expected apple -> fruit, got %v", val)
	}
	if key, val := tree.Min(); key != "apple" || val != "fruit" {
		t.Errorf("Expected Min (apple, fruit), got (%s, %s)", key, val)
	}
	if key, val := tree.Max(); key != "carrot" || val != "vegetable" {
		t.Errorf("Expected Max (carrot, vegetable), got (%s, %s)", key, val)
	}
}

func TestContainsKey(t *testing.T) {
	tree := NewTreeMap[string, string]()
	tree.Put("apple", "fruit")
	tree.Put("banana", "fruit")
	tree.Put("carrot", "vegetable")
	if !tree.ContainsKey("apple") {
		t.Errorf("Expected %v, Got %v\n", true, tree.ContainsKey("apple"))
	}
	if tree.ContainsKey("king") {
		t.Errorf("Expected %v, Got %v\n", true, tree.ContainsKey("king"))
	}
}

func TestFirstKeyAndLastKey(t *testing.T) {
	tree := NewTreeMap[int, string]()

	// Empty tree
	if key, ok := tree.FirstKey(); ok {
		t.Errorf("Expected false for empty tree FirstKey(), got %v", key)
	}
	if key, ok := tree.LastKey(); ok {
		t.Errorf("Expected false for empty tree LastKey(), got %v", key)
	}

	// Single element
	tree.Put(10, "ten")
	if key, ok := tree.FirstKey(); !ok || key != 10 {
		t.Errorf("Expected FirstKey 10, got %v", key)
	}
	if key, ok := tree.LastKey(); !ok || key != 10 {
		t.Errorf("Expected LastKey 10, got %v", key)
	}

	// Multiple elements
	tree.Put(5, "five")
	tree.Put(20, "twenty")
	tree.Put(15, "fifteen")

	if key, ok := tree.FirstKey(); !ok || key != 5 {
		t.Errorf("Expected FirstKey 5, got %v", key)
	}
	if key, ok := tree.LastKey(); !ok || key != 20 {
		t.Errorf("Expected LastKey 20, got %v", key)
	}
}

func TestCeilingKey(t *testing.T) {
	tree := NewTreeMap[int, string]()
	tree.Put(10, "ten")
	tree.Put(20, "twenty")
	tree.Put(30, "thirty")

	// Exact match
	if key, ok := tree.CeilingKey(20); !ok || key != 20 {
		t.Errorf("Expected CeilingKey 20, got %v", key)
	}

	// Between keys
	if key, ok := tree.CeilingKey(25); !ok || key != 30 {
		t.Errorf("Expected CeilingKey 30, got %v", key)
	}

	// Smaller than min
	if key, ok := tree.CeilingKey(5); !ok || key != 10 {
		t.Errorf("Expected CeilingKey 10, got %v", key)
	}

	// Larger than max
	if key, ok := tree.CeilingKey(35); ok {
		t.Errorf("Expected false for CeilingKey 35, got %v", key)
	}

	// Empty tree
	empty := NewTreeMap[int, string]()
	if key, ok := empty.CeilingKey(10); ok {
		t.Errorf("Expected false for empty tree, got %v", key)
	}
}

func TestFloorKey(t *testing.T) {
	tree := NewTreeMap[int, string]()
	tree.Put(10, "ten")
	tree.Put(20, "twenty")
	tree.Put(30, "thirty")

	// Exact match
	if key, ok := tree.FloorKey(20); !ok || key != 20 {
		t.Errorf("Expected FloorKey 20, got %v", key)
	}

	// Between keys
	if key, ok := tree.FloorKey(25); !ok || key != 20 {
		t.Errorf("Expected FloorKey 20, got %v", key)
	}

	// Larger than max
	if key, ok := tree.FloorKey(35); !ok || key != 30 {
		t.Errorf("Expected FloorKey 30, got %v", key)
	}

	// Smaller than min
	if key, ok := tree.FloorKey(5); ok {
		t.Errorf("Expected false for FloorKey 5, got %v", key)
	}

	// Empty tree
	empty := NewTreeMap[int, string]()
	if key, ok := empty.FloorKey(10); ok {
		t.Errorf("Expected false for empty tree, got %v", key)
	}
}

func TestGetUncle(t *testing.T) {
	tree := NewTreeMap[int, string]()

	/*
	        10(B)
	       /   \
	     5(R)  20(R)
	    /
	   2(R)
	*/

	tree.Put(10, "ten")    // root
	tree.Put(5, "five")    // left child
	tree.Put(20, "twenty") // right child
	tree.Put(2, "two")     // left-left child

	// getUncle for node 2 should return 20
	node2 := tree.root.left.left
	uncle := tree.getUncle(node2)

	if uncle == nil {
		t.Errorf("Expected uncle to exist for node 2")
	} else if uncle.key != 20 {
		t.Errorf("Expected uncle key 20, got %d", uncle.key)
	}

	// getUncle for node 5 should return nil (root has no parent)
	node5 := tree.root.left
	uncle2 := tree.getUncle(node5)
	if uncle2 != nil {
		t.Errorf("Expected nil uncle for node 5, got %v", uncle2.key)
	}

	// getUncle for root should return nil
	uncleRoot := tree.getUncle(tree.root)
	if uncleRoot != nil {
		t.Errorf("Expected nil uncle for root, got %v", uncleRoot.key)
	}
}
