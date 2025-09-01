package treemap

import "golang.org/x/exp/constraints"

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node[K constraints.Ordered, V any] struct {
	key    K
	value  V
	color  Color
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
}

type TreeMap[K constraints.Ordered, V any] struct {
	root *Node[K, V]
	size int
}

func NewTreeMap[K constraints.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{}
}

// Utility
func (t *TreeMap[K, V]) isRed(n *Node[K, V]) bool {
	if n == nil {
		return false
	}
	return n.color == Red
}

func (t *TreeMap[K, V]) getGrandParent(n *Node[K, V]) *Node[K, V] {
	if n == nil || n.parent == nil {
		return nil
	}
	return n.parent.parent
}

// Uncle is sibling of parent: check if parent is left child of grandparent.
func (t *TreeMap[K, V]) getUncle(n *Node[K, V]) *Node[K, V] {
	g := t.getGrandParent(n)
	if g == nil || n.parent == nil {
		return nil
	}
	if n.parent == g.left {
		return g.right
	}
	return g.left
}

func (t *TreeMap[K, V]) Put(key K, value V) {
	newNode := &Node[K, V]{key: key, value: value, color: Red}
	t.root = t.insertBST(t.root, newNode)
	t.fixInsert(newNode)
	t.size++
}

func (t *TreeMap[K, V]) rotateLeft(x *Node[K, V]) *Node[K, V] {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
	return y
}

func (t *TreeMap[K, V]) rotateRight(x *Node[K, V]) *Node[K, V] {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y
	return y
}

// Fix red-black properties after insertion
func (t *TreeMap[K, V]) fixInsert(n *Node[K, V]) {
	for n != t.root && t.isRed(n.parent) {
		g := t.getGrandParent(n)
		if g == nil {
			break
		}
		if n.parent == g.left {
			u := g.right
			if t.isRed(u) { // Case 1: uncle is red
				n.parent.color = Black
				u.color = Black
				g.color = Red
				n = g
			} else { // uncle is black
				if n == n.parent.right { // Case 2: left-right
					n = n.parent
					t.rotateLeft(n)
				}
				// Case 3: left-left
				n.parent.color = Black
				g.color = Red
				t.rotateRight(g)
			}
		} else {
			u := g.left
			if t.isRed(u) { // Case 1: uncle is red
				n.parent.color = Black
				u.color = Black
				g.color = Red
				n = g
			} else { // uncle is black
				if n == n.parent.left { // Case 2: right-left
					n = n.parent
					t.rotateRight(n)
				}
				// Case 3: right-right
				n.parent.color = Black
				g.color = Red
				t.rotateLeft(g)
			}
		}
	}
	if t.root != nil {
		t.root.color = Black
	}
}

func (t *TreeMap[K, V]) insertBST(root, node *Node[K, V]) *Node[K, V] {
	if root == nil {
		return node
	}
	if node.key < root.key {
		root.left = t.insertBST(root.left, node)
		root.left.parent = root
	} else if node.key > root.key {
		root.right = t.insertBST(root.right, node)
		root.right.parent = root
	} else {
		// key exists -> update
		root.value = node.value
		t.size-- // caller increments size; adjust for update
	}
	return root
}

func (t *TreeMap[K, V]) Get(key K) (V, bool) {
	current := t.root
	for current != nil {
		if key == current.key {
			return current.value, true
		} else if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	var zero V
	return zero, false
}

// ------------------ Deletion helpers & implementation ------------------

// transplant replaces subtree u with subtree v (u's parent now points to v)
// similar to CLRS transplant
func (t *TreeMap[K, V]) transplant(u, v *Node[K, V]) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (t *TreeMap[K, V]) minimum(n *Node[K, V]) *Node[K, V] {
	if n == nil {
		return nil
	}
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (t *TreeMap[K, V]) findNode(key K) *Node[K, V] {
	cur := t.root
	for cur != nil {
		if key == cur.key {
			return cur
		} else if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

// Remove returns the removed value and true if the key existed.
func (t *TreeMap[K, V]) Remove(key K) (V, bool) {
	z := t.findNode(key)
	var zero V
	if z == nil {
		return zero, false
	}

	removedValue := z.value

	y := z
	originalColor := y.color
	var x *Node[K, V]       // node that moves into y's original position or nil
	var xParent *Node[K, V] // used when x is nil to know its parent for fixup

	if z.left == nil {
		x = z.right
		xParent = z.parent
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		xParent = z.parent
		t.transplant(z, z.left)
	} else {
		// z has two children: replace z with its in-order successor y
		y = t.minimum(z.right)
		originalColor = y.color
		x = y.right
		if y.parent == z {
			// x's parent should be y (even if x is nil)
			if x != nil {
				x.parent = y
			}
			xParent = y
		} else {
			// transplant y with its right child
			xParent = y.parent
			t.transplant(y, y.right)
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}
		}
		t.transplant(z, y)
		y.left = z.left
		if y.left != nil {
			y.left.parent = y
		}
		y.color = z.color
	}

	// If the removed node (or moved node) was black, fix the tree
	if originalColor == Black {
		// we need to call fixDelete with x (may be nil) and its parent
		t.fixDelete(x, xParent)
	}

	t.size--
	return removedValue, true
}

// fixDelete handles the "double-black" situations after deletion.
// x may be nil; parent is the parent of x (or where x would be).
func (t *TreeMap[K, V]) fixDelete(x *Node[K, V], parent *Node[K, V]) {
	// Loop until x is root or x is red (we can color it black and finish)
	for (x != t.root) && (x == nil || !t.isRed(x)) {
		var sib *Node[K, V]
		if parent == nil {
			// This can happen when tree becomes empty; break to avoid nil deref
			break
		}
		if x == parent.left {
			sib = parent.right
			// Case 1: sibling is red
			if t.isRed(sib) {
				sib.color = Black
				parent.color = Red
				t.rotateLeft(parent)
				// update sibling after rotation
				sib = parent.right
			}
			// Case 2: sibling is black and both sibling's children are black
			if sib == nil || (!t.isRed(sib.left) && !t.isRed(sib.right)) {
				if sib != nil {
					sib.color = Red
				}
				x = parent
				parent = x.parent
			} else {
				// Case 3: sibling is black and sibling's right child is black -> rotate right at sibling
				if !t.isRed(sib.right) {
					// sibling.left must be red
					if sib.left != nil {
						sib.left.color = Black
					}
					sib.color = Red
					t.rotateRight(sib)
					sib = parent.right
				}
				// Case 4: sibling is black and sibling's right child is red
				if sib != nil {
					sib.color = parent.color
					if sib.right != nil {
						sib.right.color = Black
					}
				}
				parent.color = Black
				t.rotateLeft(parent)
				x = t.root
				parent = nil
			}
		} else {
			// mirror cases: x is right child
			sib = parent.left
			// Case 1
			if t.isRed(sib) {
				sib.color = Black
				parent.color = Red
				t.rotateRight(parent)
				sib = parent.left
			}
			// Case 2
			if sib == nil || (!t.isRed(sib.left) && !t.isRed(sib.right)) {
				if sib != nil {
					sib.color = Red
				}
				x = parent
				parent = x.parent
			} else {
				// Case 3: sibling.left is black, sibling.right is red -> rotate left at sibling
				if !t.isRed(sib.left) {
					if sib.right != nil {
						sib.right.color = Black
					}
					sib.color = Red
					t.rotateLeft(sib)
					sib = parent.left
				}
				// Case 4
				if sib != nil {
					sib.color = parent.color
					if sib.left != nil {
						sib.left.color = Black
					}
				}
				parent.color = Black
				t.rotateRight(parent)
				x = t.root
				parent = nil
			}
		}
	}
	if x != nil {
		x.color = Black
	}
}

// ------------------ Navigation helpers ------------------

// FirstKey returns the smallest key and true if tree non-empty.
func (t *TreeMap[K, V]) FirstKey() (K, bool) {
	var zero K
	if t.root == nil {
		return zero, false
	}
	n := t.root
	for n.left != nil {
		n = n.left
	}
	return n.key, true
}

// LastKey returns the largest key and true if tree non-empty.
func (t *TreeMap[K, V]) LastKey() (K, bool) {
	var zero K
	if t.root == nil {
		return zero, false
	}
	n := t.root
	for n.right != nil {
		n = n.right
	}
	return n.key, true
}

// CeilingKey returns smallest key >= given key
func (t *TreeMap[K, V]) CeilingKey(key K) (K, bool) {
	var zero K
	cur := t.root
	var candidate *Node[K, V]
	for cur != nil {
		if key == cur.key {
			return cur.key, true
		} else if key < cur.key {
			candidate = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if candidate == nil {
		return zero, false
	}
	return candidate.key, true
}

// FloorKey returns largest key <= given key
func (t *TreeMap[K, V]) FloorKey(key K) (K, bool) {
	var zero K
	cur := t.root
	var candidate *Node[K, V]
	for cur != nil {
		if key == cur.key {
			return cur.key, true
		} else if key < cur.key {
			cur = cur.left
		} else {
			candidate = cur
			cur = cur.right
		}
	}
	if candidate == nil {
		return zero, false
	}
	return candidate.key, true
}

func (t *TreeMap[K, V]) ContainsKey(key K) bool {
	_, ok := t.Get(key)
	return ok
}

func (t *TreeMap[K, V]) Size() int {
	return t.size
}

func (t *TreeMap[K, V]) Min() (K, V) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV
	}
	// leftmost node
	current := t.root
	for current.left != nil {
		current = current.left
	}
	return current.key, current.value
}

func (t *TreeMap[K, V]) Max() (K, V) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV
	}
	// rightmost node
	current := t.root
	for current.right != nil {
		current = current.right
	}
	return current.key, current.value
}

func (t *TreeMap[K, V]) Keys() []K {
	var result []K
	t.inorder(t.root, &result)
	return result
}

func (t *TreeMap[K, V]) inorder(n *Node[K, V], result *[]K) {
	if n != nil {
		t.inorder(n.left, result)
		*result = append(*result, n.key)
		t.inorder(n.right, result)
	}
}
