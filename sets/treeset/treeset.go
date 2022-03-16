// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treeset implements a tree backed by a red-black tree.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package treeset

import (
	"fmt"
	"strings"

	"github.com/daichi-m/go18ds/sets"
	rbt "github.com/daichi-m/go18ds/trees/redblacktree"
	"github.com/daichi-m/go18ds/utils"
)

func assertSetImplementation() {
	var _ sets.Set[string] = (*Set[string])(nil)
}

// Set holds elements in a red-black tree
type Set[T comparable] struct {
	tree *rbt.Tree[T]
}

var itemExists = struct{}{}

// NewWith instantiates a new empty set with the custom comparator.
func NewWith[T comparable](comparator utils.Comparator[T], values ...T) *Set[T] {
	set := &Set[T]{tree: rbt.NewWith(comparator)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWithIntComparator instantiates a new empty set with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator(values ...int) *Set[int] {
	set := &Set[int]{tree: rbt.NewWithIntComparator()}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWithStringComparator instantiates a new empty set with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator(values ...string) *Set[string] {
	set := &Set[string]{tree: rbt.NewWithStringComparator()}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.tree.Get(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) Empty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[T]) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set.
func (set *Set[T]) Values() []T {
	return set.tree.Keys()
}

// String returns a string representation of container
func (set *Set[T]) String() string {
	str := "TreeSet\n"
	items := []string{}
	for _, v := range set.tree.Keys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}
