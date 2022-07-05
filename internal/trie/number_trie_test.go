package trie

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type foo struct {
	bar string
}

func TestNumberTrie_Insert(t *testing.T) {
	trie := NewNumberTrie[foo]()
	insertKey := "012"
	wantItem := foo{bar: "lorem"}

	require.NoError(t, trie.Insert(insertKey, wantItem))

	node := trie.root.children[0]
	require.Nil(t, node.value)

	node = node.children[1]
	require.Nil(t, node.value)

	node = node.children[2]
	require.Equal(t, wantItem, *node.value)
}

func TestNumberTrie_Insert_duplicateError(t *testing.T) {
	trie := NewNumberTrie[foo]()
	insertKey := "0123456789"
	wantItem := foo{bar: "lorem"}
	require.NoError(t, trie.Insert(insertKey, wantItem))
	require.Error(t, trie.Insert(insertKey, wantItem))
}

func TestNumberTrie_Get(t *testing.T) {
	insertKey := "0123456789"

	tests := []struct {
		name      string
		key       string
		wantFound bool
	}{
		{
			name:      "value found",
			key:       insertKey,
			wantFound: true,
		},
		{
			name:      "value not found",
			key:       "000",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewNumberTrie[foo]()
			wantItem := foo{bar: "lorem"}
			require.NoError(t, trie.Insert(insertKey, wantItem))

			gotItem, ok := trie.Get(tt.key)
			require.Equal(t, tt.wantFound, ok)
			if ok {
				require.Equal(t, wantItem, gotItem)
			} else {
				require.Empty(t, gotItem)
			}
		})
	}
}

func TestNumberTrie_FindByPrefix(t *testing.T) {
	trie := NewNumberTrie[foo]()
	wantItem := foo{bar: "lorem"}
	prefix := "48"
	wantNumItems := 50

	for i := 0; i < wantNumItems; i++ {
		key := prefix + strconv.Itoa(i)
		require.NoError(t, trie.Insert(key, wantItem))
	}

	tests := []struct {
		name      string
		prefix    string
		wantFound bool
	}{
		{
			name:      "items found for prefix",
			prefix:    prefix,
			wantFound: true,
		},
		{
			name:      "no items for prefix",
			prefix:    "00",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItems, ok := trie.FindByPrefix(tt.prefix)
			require.Equal(t, tt.wantFound, ok)
			if ok {
				require.Len(t, gotItems, wantNumItems)
			} else {
				require.Empty(t, gotItems)
			}
		})
	}
}

func TestNumberTrie_Delete(t *testing.T) {
	trie := NewNumberTrie[foo]()
	wantItem := foo{bar: "lorem"}
	require.NoError(t, trie.Insert("0", wantItem))
	require.NoError(t, trie.Insert("01", wantItem))
	require.NoError(t, trie.Insert("012", wantItem))

	trie.Delete("01")

	node := trie.root.children[0]
	require.Equal(t, wantItem, *node.value)

	node = node.children[1]
	require.Nil(t, node.value)

	node = node.children[2]
	require.Equal(t, wantItem, *node.value)
}
