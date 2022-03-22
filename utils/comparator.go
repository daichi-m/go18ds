// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"time"
)

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator[T comparable] func(a, b T) int

func assertComparatorImplementation() {
	var _ Comparator[string] = StringComparator

	var _ Comparator[int] = NumberComparator[int]
	var _ Comparator[int8] = NumberComparator[int8]
	var _ Comparator[int16] = NumberComparator[int16]
	var _ Comparator[int32] = NumberComparator[int32]
	var _ Comparator[int64] = NumberComparator[int64]
	var _ Comparator[uint8] = NumberComparator[uint8]
	var _ Comparator[uint16] = NumberComparator[uint16]
	var _ Comparator[uint32] = NumberComparator[uint32]
	var _ Comparator[uint64] = NumberComparator[uint64]

	var _ Comparator[float32] = NumberComparator[float32]
	var _ Comparator[float64] = NumberComparator[float64]
	var _ Comparator[byte] = NumberComparator[byte]
	var _ Comparator[rune] = NumberComparator[rune]

	var _ Comparator[time.Time] = TimeComparator
}

// StringComparator provides a fast comparison on strings
func StringComparator(s1, s2 string) int {
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// IntComparator provides a basic comparison on int
func NumberComparator[T Number](a, b T) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

// TimeComparator provides a basic comparison on time.Time
func TimeComparator(a, b time.Time) int {
	switch {
	case a.After(b):
		return 1
	case a.Before(b):
		return -1
	default:
		return 0
	}
}
