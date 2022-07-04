package index

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type foo struct {
	lorem string
	ipsum string
}

func TestIndexes(t *testing.T) {
	loremIndex, ipsumIndex := 1, 2

	// Create lorem and ipsum indexes
	indexes := NewIndexes[foo](
		NewMapIndex[foo](loremIndex, func(foo foo) (string, bool) { return foo.lorem, true }),
		NewMapIndex[foo](ipsumIndex, func(foo foo) (string, bool) { return foo.ipsum, true }),
	)

	// Add items to indexes
	want1 := foo{lorem: "lorem1", ipsum: "ipsum1"}
	want2 := foo{lorem: "lorem1", ipsum: "ipsum2"}
	want3 := foo{lorem: "lorem2", ipsum: "ipsum1"}
	indexes.Add(want1)
	indexes.Add(want2)
	indexes.Add(want3)

	// Get items from indexes
	items, ok := indexes.Get(loremIndex, "lorem1")
	require.True(t, ok)
	require.ElementsMatch(t, []foo{want1, want2}, items)
	items, ok = indexes.Get(ipsumIndex, "ipsum1")
	require.True(t, ok)
	require.ElementsMatch(t, []foo{want1, want3}, items)

	// Delete items from indexes
	indexes.Delete(want1)
	indexes.Delete(want2)
	indexes.Delete(want3)

	// Items no longer found
	items, ok = indexes.Get(loremIndex, "lorem1")
	require.False(t, ok)
	require.Empty(t, items)
	items, ok = indexes.Get(loremIndex, "ipsum1")
	require.False(t, ok)
	require.Empty(t, items)
}
