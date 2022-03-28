// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import "github.com/daichi-m/go18ds/containers"

func assertEnumerableImplementation() {
	var _ containers.IndexedStreamer[string, *Set[string]] = (*Set[string])(nil)
}

// Each calls the given function once for each element, passing that element's index and value.
func (set *Set[T]) Each(f func(index int, value T)) {
	iterator := set.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (set *Set[T]) Map(f func(index int, value T) T) *Set[T] {
	newSet := New[T]()
	iterator := set.Iterator()
	for iterator.Next() {
		newSet.Add(f(iterator.Index(), iterator.Value()))
	}
	return newSet
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (set *Set[T]) Select(f func(index int, value T) bool) *Set[T] {
	newSet := New[T]()
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newSet.Add(iterator.Value())
		}
	}
	return newSet
}

// Filter returns a new container containing all elements for which the given function
// returns true. If none of the elements return true, it will return an empty container.
func (set *Set[T]) Filter(f func(index int, value T) bool) *Set[T] {
	iterator := set.Iterator()
	resultList := New[T]()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			resultList.Add(iterator.Value())
		}
	}
	return resultList
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (set *Set[T]) Find(f func(index int, value T) bool) (int, T) {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, *new(T)
}
