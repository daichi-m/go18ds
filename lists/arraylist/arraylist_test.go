// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"testing"

	"github.com/daichi-m/go18ds/utils"
)

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestListNew(t *testing.T) {

	tests := []struct {
		index       int
		expectValue string
		expectOk    bool
	}{
		{0, "a", true},
		{1, "b", true},
		{2, "c", true},
		{3, "", false},
	}

	list1 := New[string]()
	if !list1.Empty() {
		t.Errorf("Expected empty list got non-empty list")
	}

	list2 := New("a", "b", "c")
	if list2.Size() != 3 {
		t.Errorf("Expected list of size 3 got %v", list2.Size())
	}
	for _, tt := range tests {
		t.Run("New", func(t *testing.T) {
			if actVal, ok := list2.Get(tt.index); actVal != tt.expectValue || ok != tt.expectOk {
				t.Errorf("Got %v expected %v", actVal, tt.expectValue)
			}
		})
	}
}

func TestListAdd(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b", "c")

	tests := []struct {
		index       int
		expectValue string
		expectOk    bool
	}{
		{0, "a", true},
		{1, "b", true},
		{2, "c", true},
		{3, "", false},
	}

	for _, tt := range tests {
		t.Run("Add", func(t *testing.T) {
			if actVal, ok := list.Get(tt.index); actVal != tt.expectValue || ok != tt.expectOk {
				t.Errorf("Got %v expected %v", actVal, tt.expectValue)
			}
		})
	}
}

func TestListIndexOf(t *testing.T) {
	list := New("b", "c")
	tests := []struct {
		value     string
		expectIdx int
	}{
		{"a", -1},
		{"b", 0},
		{"c", 1},
	}

	for _, tt := range tests {
		t.Run("IndexOf", func(t *testing.T) {
			if actualIdx := list.IndexOf(tt.value); actualIdx != tt.expectIdx {
				t.Errorf("Got %v expected %v", actualIdx, tt.expectIdx)
			}
		})
	}
}

func TestListRemove(t *testing.T) {
	list := New("a", "b", "c")
	list.Remove(2)
	if actualValue, ok := list.Get(2); actualValue != "" || ok {
		t.Errorf("Got %v expected %v", actualValue, "")
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
	if !list.Empty() {
		t.Errorf("Got non-empty list expected empty list")
	}
}

func TestListGet(t *testing.T) {
	list := New("a", "b", "c")
	tests := []struct {
		index       int
		expectValue string
		expectOk    bool
	}{
		{0, "b", true},
		{1, "c", true},
		{2, "", false},
	}
	list.Remove(0)
	tests = tests[1:]
	for _, tt := range tests {
		t.Run("New", func(t *testing.T) {
			if actVal, ok := list.Get(tt.index); actVal != tt.expectValue ||
				ok != tt.expectOk {
				t.Errorf("Got %v expected %v", actVal, tt.expectValue)
			}
		})
	}
}

func TestListSwap(t *testing.T) {
	list := New("a", "b", "c")
	tests := []struct {
		index       int
		expectValue string
		expectOk    bool
	}{
		{1, "a", true},
		{0, "b", true},
		{2, "c", true},
		{3, "", false},
	}
	list.Swap(0, 1)
	for _, tt := range tests {
		t.Run("New", func(t *testing.T) {
			if actVal, ok := list.Get(tt.index); actVal != tt.expectValue ||
				ok != tt.expectOk {
				t.Errorf("Got %v expected %v", actVal, tt.expectValue)
			}
		})
	}
}

func TestListSort(t *testing.T) {
	list := New("e", "f", "g", "a", "b", "c", "d")
	list.Sort(utils.StringComparator)
	for i := range list.Values() {
		a1, ok1 := list.Get(i)
		a2, ok2 := list.Get(i + 1)
		if !ok1 || !ok2 {
			continue
		}
		if a1 > a2 {
			t.Errorf("List not sorted")
		}
	}
}

func TestListClear(t *testing.T) {
	list := New("e", "f", "g", "a", "b", "c", "d")
	list.Clear()
	if !list.Empty() {
		t.Errorf("List not cleared")
	}
}

func TestListContains(t *testing.T) {
	list1 := New("a", "b", "c")
	list2 := New[string]()

	tests := []struct {
		list     *List[string]
		val      []string
		expectOk bool
	}{
		{list1, []string{"a"}, true},
		{list1, []string{"a", "b", "c"}, true},
		{list1, []string{"d"}, false},
		{list1, []string{"a", "b", "c", "d"}, false},
		{list2, []string{"a"}, false},
		{list2, []string{"a", "b", "c"}, false},
	}

	for _, tt := range tests {
		t.Run("Contains", func(t *testing.T) {
			if actualOk := tt.list.Contains(tt.val...); actualOk != tt.expectOk {
				t.Errorf("Got %v expected %v", actualOk, tt.expectOk)
			}
		})
	}
}

func TestListValues(t *testing.T) {
	vals := []string{"a", "b", "c"}
	list := New(vals...)
	if actual, expected := list.Values(), vals; !equals(actual, expected) {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestListInsert(t *testing.T) {
	list := New[string]()
	list.Insert(0, "b", "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actual, expected := list.Values(), []string{"a", "b", "c", "d"}; !equals(actual, expected) {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}

func TestListSet(t *testing.T) {
	list := New[string]()
	list.Set(0, "a")
	list.Set(1, "b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.Set(2, "c") // append
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Set(4, "d")  // ignore
	list.Set(1, "bb") // update
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actual, expected := list.Values(), []string{"a", "bb", "c"}; !equals(actual, expected) {
		t.Errorf("Got %v expected %v", actual, expected)
	}
}
