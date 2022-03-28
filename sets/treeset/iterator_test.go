package treeset

import "testing"

func TestSetIteratorNextOnEmpty(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty set")
	}
}

func TestSetIteratorPrevOnEmpty(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty set")
	}
}

func TestSetIteratorNext(t *testing.T) {
	set := NewWithStringComparator("c", "a", "b")
	expect := []string{"a", "b", "c"}
	it := set.Iterator()
	count := 0
	for it.Next() {
		count++
		index, value := it.Index(), it.Value()
		if index >= len(expect) {
			t.Errorf("Expected index to be less than %v, got %v", len(expect), index)
		}
		if value != expect[index] {
			t.Errorf("Expected %v, got %v", expect[index], value)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSetIteratorPrev(t *testing.T) {
	set := NewWithStringComparator("c", "a", "b")
	expect := []string{"a", "b", "c"}
	it := set.Iterator()
	for it.Next() {
	}
	count := 0
	for it.Prev() {
		count++
		index, value := it.Index(), it.Value()
		if index >= len(expect) {
			t.Errorf("Expected index to be less than %v, got %v", len(expect), index)
		}
		if value != expect[index] {
			t.Errorf("Expected %v, got %v", expect[index], value)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestSetIteratorBegin(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()
	it.Begin()
	set.Add("a", "b", "c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestSetIteratorEnd(t *testing.T) {
	set := NewWithStringComparator()
	it := set.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	set.Add("a", "b", "c")
	it.End()
	if index := it.Index(); index != set.Size() {
		t.Errorf("Got %v expected %v", index, set.Size())
	}

	it.Prev()
	if index, value := it.Index(), it.Value(); index != set.Size()-1 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, set.Size()-1, "c")
	}
}

func TestSetIteratorFirst(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("a", "b", "c")
	it := set.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestSetIteratorLast(t *testing.T) {
	set := NewWithStringComparator()
	set.Add("a", "b", "c")
	it := set.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "c")
	}
}
