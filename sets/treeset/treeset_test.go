// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"fmt"
	"strings"
	"testing"
)

func concatenate[T any](args ...T) string {
	sb := strings.Builder{}
	for _, a := range args {
		sb.WriteString(fmt.Sprintf("%v", a))
	}
	return sb.String()
}

func sameElements[T int | string](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		vfound := false
		for _, u := range b {
			if u == v {
				vfound = true
				break
			}
		}
		if !vfound {
			return false
		}
	}
	return true
}

func TestSetNew(t *testing.T) {
	set := NewWithIntComparator(2, 1)
	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	values := set.Values()
	if !sameElements(values, []int{1, 2}) {
		t.Errorf("Got %v expected %v", values, []int{1, 2})
	}
}

func TestSetAdd(t *testing.T) {
	set := NewWithIntComparator()
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add() // Do nothing
	if set.Empty() {
		t.Errorf("Expected set to be non-empty, got %v", set.Empty())
	}
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if values, expect := set.Values(), []int{1, 2, 3}; !sameElements(values, expect) {
		t.Errorf("Got %v expected %v", values, expect)
	}
}

func TestSetContains(t *testing.T) {
	set := NewWithIntComparator(3, 1, 2, 3)

	tests := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
	}

	for _, tt := range tests {
		if actualValue := set.Contains(tt.value); actualValue != tt.expected {
			t.Errorf("Got %v expected %v", actualValue, tt.expected)
		}
	}
}

func TestSetRemove(t *testing.T) {
	set := NewWithIntComparator(3, 1, 2, 3)
	tests := []struct {
		value    int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
	}
	set.Remove(1)
	for _, tt := range tests {
		if actualValue := set.Contains(tt.value); actualValue != tt.expected {
			t.Errorf("Got %v expected %v", actualValue, tt.expected)
		}
	}
	set.Remove(3, 1)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	for _, tt := range tests {
		if actualValue := set.Contains(tt.value); actualValue != false {
			t.Errorf("Got %v expected %v", actualValue, false)
		}
	}
}
