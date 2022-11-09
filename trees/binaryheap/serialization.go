// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import "github.com/rahul1534/go18ds/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Heap[int])(nil)
	var _ containers.JSONDeserializer = (*Heap[int])(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap[T]) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap[T]) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
