package trie

import (
	"fmt"
)

const numbers = 10 // 0-9

type numberTrieNode[T any] struct {
	children [numbers]*numberTrieNode[T]
	value    *T
}

// NumberTrie is a trie of paths with int keys ranging from 0-9. It is used to
// store numbers as keys with an associated value. All children of a node have a
// common number prefix, which allows fast retrieval of a number. Worst case is
// O(m), where m is the length of the number searched. Given that each node has a
// value associated with it, a non nil value indicates the end of a valid number.
type NumberTrie[T any] struct {
	root *numberTrieNode[T]
}

// NewNumberTrie returns a new NumberTrie.
func NewNumberTrie[T any]() *NumberTrie[T] {
	return &NumberTrie[T]{
		root: &numberTrieNode[T]{},
	}
}

// Insert adds the specified number to the trie and stores the associated value
// within the final node of the number. Inserted numbers must be unique.
func (t *NumberTrie[T]) Insert(number string, value T) error {
	numLen := len(number)
	current := t.root

	for i := 0; i < numLen; i++ {
		index := number[i] - '0'

		if current.children[index] == nil {
			current.children[index] = &numberTrieNode[T]{}
		}
		current = current.children[index]
	}

	if current.value == nil {
		current.value = &value
		return nil
	}

	return fmt.Errorf("number already exists: %s", number)
}

// Get returns the value associated with the specified number.
func (t *NumberTrie[T]) Get(number string) (T, bool) {
	if node, ok := find(number, t.root); ok && node.value != nil {
		return *node.value, true
	}

	var none T
	return none, false
}

// FindByPrefix returns all values that are associated with numbers beginning
// with the specified number prefix.
func (t *NumberTrie[T]) FindByPrefix(numberPrefix string) ([]T, bool) {
	var values []T

	prefixNode, ok := find(numberPrefix, t.root)
	if !ok {
		return nil, false
	}

	// Search branches of prefix node and retrieve all values
	stack := []*numberTrieNode[T]{prefixNode}

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n] // Pop

		if node != nil {
			if node.value != nil {
				values = append(values, *node.value)
			}

			for _, childNode := range node.children {
				stack = append(stack, childNode)
			}
		}
	}

	if len(values) == 0 {
		return nil, false
	}

	return values, true
}

// Delete removes the specified number from the trie.
func (t *NumberTrie[T]) Delete(number string) {
	remove(number, t.root, 0)
}

func find[T any](number string, root *numberTrieNode[T]) (*numberTrieNode[T], bool) {
	numLen := len(number)
	current := root

	for i := 0; i < numLen; i++ {
		index := number[i] - '0'
		if current.children[index] == nil {
			return nil, false
		}
		current = current.children[index]
	}

	if current != nil {
		return current, true
	}

	return nil, false
}

func remove[T any](number string, node *numberTrieNode[T], depth int) *numberTrieNode[T] {
	if node == nil {
		return nil
	}

	if depth == len(number) {
		node.value = nil
		return node
	}

	// TODO: can optimize further by pruning branches
	index := number[depth] - '0'
	node.children[index] = remove(number, node.children[index], depth+1)
	return node
}
