/*
Package trie provides an implementation of a prefix tree (Trie) data structure.

A Trie is a tree-like data structure commonly used for storing and retrieving
strings efficiently, especially for prefix-based operations. It supports the
following features:

  - Insert: Add a string to the trie in O(n) time, where n is the length of the string.
  - Search: Check if a string exists in the trie in O(n) time.
  - StartsWith: Check if any string in the trie starts with a given prefix in O(n) time.
  - Delete: Remove a string from the trie, adjusting nodes as needed in O(n) time.
  - Thread Safety: All operations are concurrency-safe using sync.RWMutex.

Use Cases:
  - Autocomplete systems
  - Spell checking
  - IP routing tables
  - Dictionary or prefix matching

Example usage:

	t := trie.NewTrie()
	t.Insert("go")
	t.Insert("gopher")
	fmt.Println(t.Search("go"))        // true
	fmt.Println(t.StartsWith("gop"))   // true
	fmt.Println(t.Search("java"))      // false

Implementation Details:
  - Each node contains a map of rune to *Node for children.
  - An `isEnd` flag marks the end of a valid word.
  - The trie dynamically grows as new words are added.
  - A stack from github.com/Zubayear/ryushin/stack may be used internally for traversal or deletion.

Time Complexity:
  - Insert: O(n)
  - Search: O(n)
  - StartsWith: O(n)
  - Delete: O(n)

Space Complexity:
  - O(m * n), where m is the number of words and n is the average length of each word.
*/
package trie

import (
	"sync"

	"github.com/Zubayear/ryushin/stack"
)

// Node represents a single node in the Trie data structure.
//
// Each node contains:
//   - children: a map of rune to Node pointers representing possible next characters.
//   - isEnd: a boolean flag that indicates whether this node marks the end of a complete word.
type Node struct {
	children map[rune]*Node // maps each character to its next node
	isEnd    bool           // true if this node marks the end of a valid word
}

// NewTrieNode creates and returns a new Trie node.
//
// The returned node has:
//   - an empty children map
//   - isEnd set to false
func NewTrieNode() *Node {
	return &Node{make(map[rune]*Node), false}
}

// Trie represents a thread-safe Trie (prefix tree) implementation.
//
// Fields:
//   - root: the root node of the Trie
//   - size: the number of complete words stored in the Trie
//   - mutex: a read-write mutex (RWMutex) to ensure concurrent safety
//
// Operations supported:
//   - Insert: Add a word to the Trie
//   - Search: Check if a word exists
//   - StartsWith: Check if a prefix exists
//   - GetWordsWithPrefix: Retrieve all words with a given prefix
//   - Remove: Delete a word from the Trie
//   - Size / IsEmpty: Utility functions
type Trie struct {
	root  *Node
	size  int
	mutex sync.RWMutex
}

// NewTrie creates and returns an empty Trie instance.
//
// Example:
//
//	t := NewTrie()
//	t.Insert("hello")
//	fmt.Println(t.Search("hello")) // true
func NewTrie() *Trie {
	return &Trie{NewTrieNode(), 0, sync.RWMutex{}}
}

// Size returns the total number of complete words stored in the Trie.
//
// Time Complexity: O(1)
func (t *Trie) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

// IsEmpty returns true if the Trie contains no words, false otherwise.
//
// Time Complexity: O(1)
func (t *Trie) IsEmpty() bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size == 0
}

// Insert adds a word into the Trie.
//
// Notes:
//   - If the word already exists, it does not increase the size again.
//   - The method is case-sensitive and does not trim spaces.
//
// Algorithm Steps:
//   - Start from the root node
//   - For each character in the word
//   - If the character's child does not exist, create a new node
//   - Move to the child node
//   - After constructing the branch mark the last node as terminal node if it's not already marked
//
// Time Complexity: O(N), where N = length of the word
//
// Space Complexity: O(N) for new nodes (if needed)
func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	current := t.root
	for _, ch := range word {
		if current.children[ch] == nil {
			current.children[ch] = NewTrieNode()
		}
		current = current.children[ch]
	}
	if !current.isEnd {
		current.isEnd = true
		t.size++
	}
}

// Search checks if a complete word exists in the Trie.
//
// Returns true if the word exists and is marked as a complete word.
// Does NOT return true for prefixes only.
//
// Algorithm steps:
//   - Start from root
//   - Iterate over the word
//   - For each character check if its children have it, if not return false
//   - Return the state of terminal node
//
// Time Complexity: O(N), where N = length of the word
func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	current := t.root
	for _, ch := range word {
		if current.children[ch] == nil {
			return false
		}
		current = current.children[ch]
	}
	return current.isEnd
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix.
//
// Returns true if the prefix exists in the Trie, even if it is not a complete word.
//
// Algorithm Steps:
//   - Traverse the Trie for each character in the prefix.
//   - If at any point a character is missing, return false.
//   - If at any point a character is missing, return false.
//   - If traversal succeeds, return true.
//
// Time Complexity: O(K), where K = length of the prefix
func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	current := t.root
	for _, ch := range prefix {
		if current.children[ch] == nil {
			return false
		}
		current = current.children[ch]
	}
	return true
}

// Dfs performs a depth-first search starting from the given node
// and collects all words that stem from the current prefix.
//
// Algorithm Steps:
//   - Initialize a result slice.
//   - If the current node marks the end of a word, append the prefix to result.
//   - For each child, recursively call DFS with an updated prefix.
//   - Start DFS from the given node and prefix.
//   - Return the result slice.
//
// Time Complexity: O(M * L), where M = number of words from a node, L = average word length
func (t *Trie) dfs(node *Node, prefix string) []string {
	var result []string
	var dfs func(node *Node, prefix string)
	dfs = func(node *Node, prefix string) {
		if node.isEnd {
			result = append(result, prefix)
		}
		for ch, child := range node.children {
			dfs(child, prefix+string(ch))
		}
	}
	dfs(node, prefix)
	return result
}

// findNodeForPrefix returns the node corresponding to the last character of the given prefix.
// If the prefix does not exist in the Trie, it returns nil.
//
// Time Complexity: O(K), where K = length of the prefix
func (t *Trie) findNodeForPrefix(prefix string) *Node {
	current := t.root
	for _, ch := range prefix {
		if current.children[ch] == nil {
			return nil
		}
		current = current.children[ch]
	}
	return current
}

// GetWordsWithPrefix retrieves all words in the Trie that start with the given prefix.
//
// Returns:
//   - A slice of words that start with the prefix
//   - An empty slice if the prefix does not exist
//
// Algorithm Steps:
//   - Traverse the Trie to find the node corresponding to the prefix.
//   - If the prefix is not found, return an empty slice.
//   - Perform DFS from that node to collect all words with a prefix.
//
// Time Complexity: O(K + M * L)
//   - K = length of prefix
//   - M = number of matching words
//   - L = average length of matching words
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	if len(prefix) == 0 {
		return nil
	}
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	var result []string
	current := t.findNodeForPrefix(prefix)
	if current == nil {
		return result
	}
	return t.dfs(current, prefix)
}

// Remove deletes a word from the Trie if it exists.
//
// Returns true if the word was successfully removed, false otherwise.
// It also removes unnecessary nodes to keep the Trie compact.
//
// Algorithm Steps:
//   - Traverse the word and push (node, char) pairs into a stack for backtracking.
//   - If the word does not exist or is not marked as the end, return false.
//   - Mark the last node as not the end.
//   - Backtrack and remove nodes that are no longer needed (no children and not end).
//   - Decrement size and return true.
//
// Time Complexity: O(N), where N = length of the word
//
// Space Complexity: O(N) for the stack used to track nodes
func (t *Trie) Remove(word string) bool {
	if len(word) == 0 {
		return false
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	current := t.root
	type Pair struct {
		node *Node
		ch   rune
	}

	s := stack.NewStack[Pair]()
	for _, ch := range word {
		next := current.children[ch]
		if next == nil {
			return false
		}
		_, _ = s.Push(Pair{current, ch})
		current = next
	}
	if !current.isEnd {
		return false
	}
	current.isEnd = false

	for !s.IsEmpty() {
		val, _ := s.Pop()
		parent := val.node
		ch := val.ch
		child := parent.children[ch]
		if len(child.children) == 0 && !child.isEnd {
			delete(parent.children, ch)
		} else {
			break
		}
	}
	t.size--
	return true
}
