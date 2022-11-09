// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"github.com/rahul1534/go18ds/containers"
	rbt "github.com/rahul1534/go18ds/trees/redblacktree"
)

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithKey[string, string] = (*Iterator[string, string])(nil)
}

// Iterator holding the iterator's state
type Iterator[K comparable, V comparable] struct {
	iterator rbt.Iterator[K, V]
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (m *Map[K, V]) Iterator() Iterator[K, V] {
	return Iterator[K, V]{iterator: m.tree.Iterator()}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Next() bool {
	return iterator.iterator.Next()
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Prev() bool {
	return iterator.iterator.Prev()
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Value() V {
	return iterator.iterator.Value()
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, V]) Key() K {
	return iterator.iterator.Key()
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[K, V]) Begin() {
	iterator.iterator.Begin()
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[K, V]) End() {
	iterator.iterator.End()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iterator *Iterator[K, V]) First() bool {
	return iterator.iterator.First()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, V]) Last() bool {
	return iterator.iterator.Last()
}
