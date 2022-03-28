// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package containers

// EnumerableWithIndex provides functions for ordered containers whose values can be fetched by an index.
type IndexedStreamer[T comparable, C IndexedContainer[T]] interface {
	// Each calls the given function once for each element, passing that element's index and value.
	Each(func(index int, value T))

	// Map invokes the given function once for each element and returns a
	// container containing the values returned by the given function.
	Map(f func(index int, value T) T) C

	// Select returns a new container containing all elements for which the given function
	// returns a true value.
	Select(func(index int, value T) bool) C

	// Filter returns a new container containing all elements for which the given function
	// returns true. If none of the elements return true, it will return an empty container.
	Filter(func(index int, value T) bool) C

	// Find passes each element of the container to the given function and returns
	// the first (index,value) for which the function is true or -1,nil otherwise
	// if no element matches the criteria.
	Find(func(index int, value T) bool) (int, T)
}

// EnumerableWithKey provides functions for ordered containers whose values whose elements are key/value pairs.
type KeyedStreamer[K comparable, V comparable, C KeyedContainer[K, V]] interface {
	// Each calls the given function once for each element, passing that element's key and value.
	Each(func(key K, value V))

	// Map invokes the given function once for each element and returns a container
	// containing the values returned by the given function as key/value pairs.
	Map(func(key K, value V) (K, V)) C

	// Select returns a new container containing all elements for which the given
	// function returns a true value.
	Select(func(key K, value V) bool) C

	// Filter returns a new container for which the given function returns true.
	// If none of the elements return true, it will return an empty container.
	Filter(func(key K, value V) bool) C

	// Find passes each element of the container to the given function and returns
	// the first (key,value) for which the function is true or nil,nil otherwise if no element
	// matches the criteria.
	Find(func(key K, value V) bool) (K, V)
}
