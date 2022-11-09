// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"fmt"
	"strings"

	"github.com/rahul1534/go18ds/maps"
	"github.com/rahul1534/go18ds/trees/redblacktree"
	"github.com/rahul1534/go18ds/utils"
)

func assertMapImplementation() {
	var _ maps.BidiMap[string, string] = (*Map[string, string])(nil)
}

// Map holds the elements in two red-black trees.
type Map[K comparable, V comparable] struct {
	forwardMap      redblacktree.Tree[K, *data[K, V]]
	inverseMap      redblacktree.Tree[V, *data[K, V]]
	keyComparator   utils.Comparator[K]
	valueComparator utils.Comparator[V]
}

type data[K comparable, V comparable] struct {
	key   K
	value V
}

// NewWith instantiates a bidirectional map.
func NewWith[K comparable, V comparable](keyComparator utils.Comparator[K],
	valueComparator utils.Comparator[V]) *Map[K, V] {
	return &Map[K, V]{
		forwardMap:      *redblacktree.NewWith[K, *data[K, V]](keyComparator),
		inverseMap:      *redblacktree.NewWith[V, *data[K, V]](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewWithIntComparators instantiates a bidirectional map with the IntComparator for key and value, i.e. keys and values are of type int.
func NewWithIntComparators() *Map[int, int] {
	return NewWith(utils.NumberComparator[int], utils.NumberComparator[int])
}

// NewWithStringComparators instantiates a bidirectional map with the StringComparator for key and value, i.e. keys and values are of type string.
func NewWithStringComparators() *Map[string, string] {
	return NewWith(utils.StringComparator, utils.StringComparator)
}

// Put inserts element into the map.
func (m *Map[K, V]) Put(key K, value V) {
	if d, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(d.value)
	}
	if d, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(d.key)
	}
	d := &data[K, V]{key: key, value: value}
	m.forwardMap.Put(key, d)
	m.inverseMap.Put(value, d)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d.value, true
	}
	return *new(V), false
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, V]) GetKey(value V) (key K, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d.key, true
	}
	return *new(K), false
}

// Remove removes the element from the map by key.
func (m *Map[K, V]) Remove(key K) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d.value)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (ordered).
func (m *Map[K, V]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
func (m *Map[K, V]) Values() []V {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "TreeBidiMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
