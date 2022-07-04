package index

import "github.com/deckarep/golang-set/v2"

type MapIndex[T comparable] struct {
	id    int
	index map[string]mapset.Set[T]
	keyFn func(T) (string, bool)
}

func NewMapIndex[T comparable](id int, keyFn func(T) (string, bool)) *MapIndex[T] {
	return &MapIndex[T]{
		id:    id,
		index: map[string]mapset.Set[T]{},
		keyFn: keyFn,
	}
}

func (i *MapIndex[T]) Add(item T) {
	if key, ok := i.keyFn(item); ok {
		if items, ok := i.index[key]; ok {
			items.Add(item)
		} else {
			i.index[key] = mapset.NewSet[T](item)
		}
	}
}

func (i *MapIndex[T]) Get(key string) ([]T, bool) {
	items, ok := i.index[key]
	if !ok {
		return nil, false
	}

	return items.ToSlice(), true
}

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

type Indexes[T comparable] struct {
	indexes []*MapIndex[T]
}

func NewIndexes[T comparable](indexes ...*MapIndex[T]) *Indexes[T] {
	return &Indexes[T]{
		indexes: indexes,
	}
}

func (i *Indexes[T]) Add(item T) {
	for _, index := range i.indexes {
		index.Add(item)
	}
}

func (i Indexes[T]) Get(id int, key string) ([]T, bool) {
	for _, index := range i.indexes {
		if index.id == id {
			return index.Get(key)
		}
	}
	return nil, false
}

func (i Indexes[T]) Delete(item T) {
	for _, index := range i.indexes {
		index.Delete(item)
	}
}
