package index

import "github.com/deckarep/golang-set/v2"

// MapIndex uses a map to index records on one or more fields, which makes
// searching dramatically faster for the specified field(s) at O(1) time complexity.
type MapIndex[T comparable] struct {
	id    int
	index map[string]mapset.Set[T]
	// A function to specify which fields should be returned from the item for
	// the index. For multiple fields simply concatenating them will suffice.
	// For fields that are optional, return a false boolean flag to indicate that
	// the value is not present, in which it will be skipped.
	keyFn func(T) (string, bool)
}

// NewMapIndex returns a new MapIndex.
func NewMapIndex[T comparable](id int, keyFn func(T) (string, bool)) *MapIndex[T] {
	return &MapIndex[T]{
		id:    id,
		index: map[string]mapset.Set[T]{},
		keyFn: keyFn,
	}
}

// Add adds a new item to the map index.
func (i *MapIndex[T]) Add(item T) {
	if key, ok := i.keyFn(item); ok {
		if items, ok := i.index[key]; ok {
			items.Add(item)
		} else {
			i.index[key] = mapset.NewSet[T](item)
		}
	}
}

// Get returns the item for the specified key.
func (i *MapIndex[T]) Get(key string) ([]T, bool) {
	items, ok := i.index[key]
	if !ok {
		return nil, false
	}

	return items.ToSlice(), true
}

// Delete removes the specified item from the index.
func (i *MapIndex[T]) Delete(item T) {
	if key, ok := i.keyFn(item); ok {
		if items, ok := i.index[key]; ok {
			items.Remove(item)
			if len(items.ToSlice()) == 0 {
				delete(i.index, key)
			}
		}
	}
}

// Indexes holds multiple indexes to easily perform operations across.
type Indexes[T comparable] struct {
	indexes []*MapIndex[T]
}

// NewIndexes returns a new Indexes that contains the provided indexes.
func NewIndexes[T comparable](indexes ...*MapIndex[T]) *Indexes[T] {
	return &Indexes[T]{
		indexes: indexes,
	}
}

// Add adds the item to each index.
func (i *Indexes[T]) Add(item T) {
	for _, index := range i.indexes {
		index.Add(item)
	}
}

// Get returns all items from the specified index for the provided key.
func (i Indexes[T]) Get(id int, key string) ([]T, bool) {
	for _, index := range i.indexes {
		if index.id == id {
			return index.Get(key)
		}
	}
	return nil, false
}

// Delete removes the specified item from each index.
func (i Indexes[T]) Delete(item T) {
	for _, index := range i.indexes {
		index.Delete(item)
	}
}
