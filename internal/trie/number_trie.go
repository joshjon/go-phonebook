package trie

import (
	"fmt"
)

const numbers = 10 // 0-9

type numberTrieNode[T any] struct {
	children [numbers]*numberTrieNode[T]
	item     *T
}

type NumberTrie[T any] struct {
	root *numberTrieNode[T]
}

func NewNumberTrie[T any]() *NumberTrie[T] {
	return &NumberTrie[T]{
		root: &numberTrieNode[T]{},
	}
}

func (t *NumberTrie[T]) Insert(key string, item T) error {
	numLen := len(key)
	current := t.root

	for i := 0; i < numLen; i++ {
		index := key[i] - '0'

		if current.children[index] == nil {
			current.children[index] = &numberTrieNode[T]{}
		}
		current = current.children[index]
	}

	if current.item == nil {
		current.item = &item
		return nil
	}

	return fmt.Errorf("key already exists: %s", key)
}

func (t *NumberTrie[T]) Get(key string) (T, bool) {
	if node, ok := find(key, t.root); ok && node.item != nil {
		return *node.item, true
	}

	var none T
	return none, false
}

func (t *NumberTrie[T]) FindByPrefix(prefix string) ([]T, bool) {
	var items []T

	prefixNode, ok := find(prefix, t.root)
	if !ok {
		return nil, false
	}

	// Search branches of prefix node and retrieve all items
	stack := []*numberTrieNode[T]{prefixNode}

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n] // Pop

		if node != nil {
			if node.item != nil {
				items = append(items, *node.item)
			}

			for _, childNode := range node.children {
				stack = append(stack, childNode)
			}
		}
	}

	if len(items) == 0 {
		return nil, false
	}

	return items, true
}

func (t *NumberTrie[T]) Delete(key string) {
	remove(key, t.root, 0)
}

func find[T any](key string, root *numberTrieNode[T]) (*numberTrieNode[T], bool) {
	numLen := len(key)
	current := root

	for i := 0; i < numLen; i++ {
		index := key[i] - '0'
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

func remove[T any](key string, node *numberTrieNode[T], depth int) *numberTrieNode[T] {
	if node == nil {
		return nil
	}

	if depth == len(key) {
		node.item = nil
		return node
	}

	// TODO: can optimize further by pruning branches
	index := key[depth] - '0'
	node.children[index] = remove(key, node.children[index], depth+1)
	return node
}
