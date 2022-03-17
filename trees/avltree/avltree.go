// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package avltree implements an AVL balanced binary tree.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/AVL_tree
package avltree

import (
	"fmt"

	"github.com/daichi-m/go18ds/trees"
	"github.com/daichi-m/go18ds/utils"
)

func assertTreeImplementation() {
	var _ trees.Tree[string] = new(Tree[int, string])
}

// Tree holds elements of the AVL tree.
type Tree[K comparable, V comparable] struct {
	Root       *Node[K, V]         // Root node
	Comparator utils.Comparator[K] // Key comparator
	size       int                 // Total number of keys in the tree
}

// Node is a single element within the tree
type Node[K comparable, V comparable] struct {
	Key      K
	Value    V
	Parent   *Node[K, V]    // Parent node
	Children [2]*Node[K, V] // Children nodes
	b        int8
}

// NewWith instantiates an AVL tree with the custom comparator.
func NewWith[K comparable, V comparable](comparator utils.Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{Comparator: comparator}
}

// NewWithIntComparator instantiates an AVL tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator[V comparable]() *Tree[int, V] {
	return &Tree[int, V]{Comparator: utils.NumberComparator[int]}
}

// NewWithStringComparator instantiates an AVL tree with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[V comparable]() *Tree[string, V] {
	return &Tree[string, V]{Comparator: utils.StringComparator}
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Put(key K, value V) {
	t.put(key, value, nil, &t.Root)
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Get(key K) (value V, found bool) {
	n := t.Root
	for n != nil {
		cmp := t.Comparator(key, n.Key)
		switch {
		case cmp == 0:
			return n.Value, true
		case cmp < 0:
			n = n.Children[0]
		case cmp > 0:
			n = n.Children[1]
		}
	}
	return *new(V), false
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Remove(key K) {
	t.remove(key, &t.Root)
}

// Empty returns true if tree does not contain any nodes.
func (t *Tree[K, V]) Empty() bool {
	return t.size == 0
}

// Size returns the number of elements stored in the tree.
func (t *Tree[K, V]) Size() int {
	return t.size
}

// Keys returns all keys in-order
func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (t *Tree[K, V]) Values() []V {
	values := make([]V, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

// Left returns the minimum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree[K, V]) Left() *Node[K, V] {
	return t.bottom(0)
}

// Right returns the maximum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree[K, V]) Right() *Node[K, V] {
	return t.bottom(1)
}

// Floor Finds floor node of the input key, return the floor node or nil if no ceiling is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.Children[0]
		case c > 0:
			floor, found = n, true
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree is smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Ceiling(key K) (floor *Node[K, V], found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			floor, found = n, true
			n = n.Children[0]
		case c > 0:
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (t *Tree[K, V]) Clear() {
	t.Root = nil
	t.size = 0
}

// String returns a string representation of container
func (t *Tree[K, V]) String() string {
	str := "AVLTree\n"
	if !t.Empty() {
		output(t.Root, "", true, &str)
	}
	return str
}

func (n *Node[K, V]) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func (t *Tree[K, V]) put(key K, value V, p *Node[K, V], qp **Node[K, V]) bool {
	q := *qp
	if q == nil {
		t.size++
		*qp = &Node[K, V]{Key: key, Value: value, Parent: p}
		return true
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		q.Key = key
		q.Value = value
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	var fix bool
	fix = t.put(key, value, q, &q.Children[a])
	if fix {
		return putFix(int8(c), qp)
	}
	return false
}

func (t *Tree[K, V]) remove(key K, qp **Node[K, V]) bool {
	q := *qp
	if q == nil {
		return false
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		t.size--
		if q.Children[1] == nil {
			if q.Children[0] != nil {
				q.Children[0].Parent = q.Parent
			}
			*qp = q.Children[0]
			return true
		}
		fix := removeMin(&q.Children[1], &q.Key, &q.Value)
		if fix {
			return removeFix(-1, qp)
		}
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	fix := t.remove(key, &q.Children[a])
	if fix {
		return removeFix(int8(-c), qp)
	}
	return false
}

func removeMin[K comparable, V comparable](qp **Node[K, V], minKey *K, minVal *V) bool {
	q := *qp
	if q.Children[0] == nil {
		*minKey = q.Key
		*minVal = q.Value
		if q.Children[1] != nil {
			q.Children[1].Parent = q.Parent
		}
		*qp = q.Children[1]
		return true
	}
	fix := removeMin(&q.Children[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

func putFix[K comparable, V comparable](c int8, t **Node[K, V]) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.Children[(c+1)/2].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return false
}

func removeFix[K comparable, V comparable](c int8, t **Node[K, V]) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return false
	}

	if s.b == -c {
		s.b = 0
		return true
	}

	a := (c + 1) / 2
	if s.Children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}

	if s.Children[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return true
}

func singlerot[K comparable, V comparable](c int8, s *Node[K, V]) *Node[K, V] {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doublerot[K comparable, V comparable](c int8, s *Node[K, V]) *Node[K, V] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0
	return p
}

func rotate[K comparable, V comparable](c int8, s *Node[K, V]) *Node[K, V] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r
	return r
}

func (t *Tree[K, V]) bottom(d int) *Node[K, V] {
	n := t.Root
	if n == nil {
		return nil
	}

	for c := n.Children[d]; c != nil; c = n.Children[d] {
		n = c
	}
	return n
}

// Prev returns the previous element in an inorder
// walk of the AVL tree.
func (n *Node[K, V]) Prev() *Node[K, V] {
	return n.walk1(0)
}

// Next returns the next element in an inorder
// walk of the AVL tree.
func (n *Node[K, V]) Next() *Node[K, V] {
	return n.walk1(1)
}

func (n *Node[K, V]) walk1(a int) *Node[K, V] {
	if n == nil {
		return nil
	}

	if n.Children[a] != nil {
		n = n.Children[a]
		for n.Children[a^1] != nil {
			n = n.Children[a^1]
		}
		return n
	}

	p := n.Parent
	for p != nil && p.Children[a] == n {
		n = p
		p = p.Parent
	}
	return p
}

func output[K comparable, V comparable](node *Node[K, V], prefix string,
	isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}
}
