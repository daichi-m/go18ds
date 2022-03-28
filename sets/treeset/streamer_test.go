package treeset

import (
	"testing"
)

func TestSetEach(t *testing.T) {
	set := NewWithStringComparator("c", "a", "b")
	expect := []string{"a", "b", "c"}
	set.Each(func(index int, value string) {
		if index >= len(expect) {
			t.Errorf("Index %v is too many", index)
		}
		if value != expect[index] {
			t.Errorf("Got %v expected %v", value, expect[index])
		}
	})
}

func TestSetMap(t *testing.T) {
	set := NewWithStringComparator("c", "a", "b")
	expect := []string{"mapped: a", "mapped: b", "mapped: c"}
	mappedSet := set.Map(func(index int, value string) string {
		return "mapped: " + value
	})
	if !mappedSet.Contains(expect...) {
		t.Errorf("Expected to contain %v but set does not contain", expect)
	}
	if mappedSet.Contains("mapped: d") {
		t.Errorf("Did not expect set to contain %v", "mapped: d")
	}
	if mappedSet.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedSet.Size(), 3)
	}
}

func TestSetSelect(t *testing.T) {
	set := NewWithStringComparator("c", "a", "b")
	selectedSet := set.Select(func(index int, value string) bool {
		return value >= "a" && value <= "b"
	})
	if !selectedSet.Contains("a", "b") {
		t.Errorf("Got %v expected %v", selectedSet.Values(), []string{"a", "b"})
	}
	if selectedSet.Contains("c") {
		t.Errorf("Did not expect set to contain %v", "c")
	}
	if selectedSet.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedSet.Size(), 3)
	}
}

func TestSetFilter(t *testing.T) {
	vals := []string{"a", "b", "c"}
	set := NewWithStringComparator(vals...)
	all := set.Filter(func(index int, value string) bool {
		return value >= "a" && value <= "c"
	})
	if !sameElements(all.Values(), vals) {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = set.Filter(func(index int, value string) bool {
		return value >= "a" && value <= "b"
	})
	if !sameElements(all.Values(), vals[0:2]) {
		t.Errorf("Got %v expected %v", all, vals[0:1])
	}
}

func TestSetFind(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("c", "a", "b")
	foundIndex, foundValue := set.Find(func(index int, value string) bool {
		return value == "c"
	})
	if foundValue != "c" || foundIndex != 2 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, "c", 2)
	}
	foundIndex, foundValue = set.Find(func(index int, value string) bool {
		return value == "x"
	})
	if foundValue != "" || foundIndex != -1 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, nil, nil)
	}
}
