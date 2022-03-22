// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"testing"
	"time"
)

func TestIntComparator(t *testing.T) {
	// i1,i2,expected
	tests := []struct {
		i1, i2   int
		expected int
	}{
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{11, 22, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}

	for _, test := range tests {
		t.Run("IntCompare", func(tt *testing.T) {
			actual := NumberComparator(test.i1, test.i2)
			expected := test.expected
			if actual != expected {
				tt.Errorf("Got %v expected %v", actual, expected)
			}
		})
	}
}

func TestStringComparator(t *testing.T) {
	// s1,s2,expected
	tests := []struct {
		s1, s2   string
		expected int
	}{
		{"a", "a", 0},
		{"a", "b", -1},
		{"b", "a", 1},
		{"aa", "aab", -1},
		{"", "", 0},
		{"a", "", 1},
		{"", "a", -1},
		{"", "aaaaaaa", -1},
	}

	for _, test := range tests {
		t.Run("StringCompare", func(tt *testing.T) {
			actual := StringComparator(test.s1, test.s2)
			expected := test.expected
			if actual != expected {
				tt.Errorf("Got %v expected %v", actual, expected)
			}
		})
	}
}

func TestTimeComparator(t *testing.T) {
	now := time.Now()

	// i1,i2,expected
	tests := []struct {
		t1, t2   time.Time
		expected int
	}{
		{now, now, 0},
		{now.Add(24 * 7 * 2 * time.Hour), now, 1},
		{now, now.Add(24 * 7 * 2 * time.Hour), -1},
	}

	for _, test := range tests {
		t.Run("TimeCompare", func(tt *testing.T) {
			actual := TimeComparator(test.t1, test.t2)
			expected := test.expected
			if actual != expected {
				tt.Errorf("Got %v expected %v", actual, expected)
			}
		})
	}
}

func TestCustomComparator(t *testing.T) {
	type Custom struct {
		id   int
		name string
	}

	sortByID := func(c1, c2 Custom) int {
		switch {
		case c1.id > c2.id:
			return 1
		case c1.id < c2.id:
			return -1
		default:
			return 0
		}
	}

	// o1,o2,expected
	tests := []struct {
		c1, c2   Custom
		expected int
	}{
		{Custom{1, "a"}, Custom{1, "a"}, 0},
		{Custom{1, "a"}, Custom{2, "b"}, -1},
		{Custom{2, "b"}, Custom{1, "a"}, 1},
		{Custom{1, "a"}, Custom{1, "b"}, 0},
	}

	for _, test := range tests {
		t.Run("CustomCompare", func(tt *testing.T) {
			actual := sortByID(test.c1, test.c2)
			expected := test.expected
			if actual != expected {
				tt.Errorf("Got %v expected %v", actual, expected)
			}
		})
	}
}
