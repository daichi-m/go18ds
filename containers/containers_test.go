// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// All data structures must implement the container structure

package containers

import (
	"testing"

	"github.com/daichi-m/go18ds/utils"
)

// For testing purposes
type ContainerTest[T ~int | ~string] struct {
	values []T
}

func (container ContainerTest[T]) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest[T]) Size() int {
	return len(container.values)
}

func (container ContainerTest[T]) Clear() {
	container.values = []T{}
}

func (container ContainerTest[T]) Values() []T {
	return container.values
}

func TestGetSortedValuesInts(t *testing.T) {
	var container Container[int] = ContainerTest[int]{
		values: []int{5, 1, 3, 2, 4},
	}
	values := GetSortedValues(container, utils.NumberComparator[int])
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	var container Container[string] = ContainerTest[string]{
		values: []string{"g", "a", "d", "e", "f", "c", "b"},
	}
	values := GetSortedValues(container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}
