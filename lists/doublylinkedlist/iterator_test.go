package doublylinkedlist

import "testing"

func TestListIteratorNextOnEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty list")
	}
}

func TestListIteratorNext(t *testing.T) {
	list := New("a", "b", "c")
	it := list.Iterator()
	expect := []string{"a", "b", "c"}
	count := 0
	for it.Next() {
		count++
		index, value := it.Index(), it.Value()
		if index >= len(expect) {
			t.Errorf("Index too big")
		}
		if actual := expect[index]; actual != value {
			t.Errorf("Got %v expected %v", value, actual)
		}
	}
	if actualCt, expectedCt := count, 3; actualCt != expectedCt {
		t.Errorf("Got %v expected %v", actualCt, expectedCt)
	}
}

func TestListIteratorPrevOnEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty list")
	}
}

func TestListIteratorPrev(t *testing.T) {
	list := New("a", "b", "c")
	list.Add()
	it := list.Iterator()
	for it.Next() {
	}
	expect := []string{"a", "b", "c"}
	count := 0
	for it.Prev() {
		count++
		index, value := it.Index(), it.Value()
		if index >= len(expect) {
			t.Errorf("Index too big")
		}
		if actual := expect[index]; actual != value {
			t.Errorf("Got %v expected %v", value, actual)
		}
	}
	if actualCt, expectedCt := count, 3; actualCt != expectedCt {
		t.Errorf("Got %v expected %v", actualCt, expectedCt)
	}
}

func TestListIteratorBegin(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	it.Begin()
	list.Add("a", "b", "c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestListIteratorEnd(t *testing.T) {
	list := New[string]()
	it := list.Iterator()

	if index := it.Index(); index != -1 {
		t.Errorf("Got %v expected %v", index, -1)
	}

	it.End()
	if index := it.Index(); index != 0 {
		t.Errorf("Got %v expected %v", index, 0)
	}

	list.Add("a", "b", "c")
	it.End()
	if index := it.Index(); index != list.Size() {
		t.Errorf("Got %v expected %v", index, list.Size())
	}

	it.Prev()
	if index, value := it.Index(), it.Value(); index != list.Size()-1 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, list.Size()-1, "c")
	}
}

func TestListIteratorFirst(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	list.Add("a", "b", "c")
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestListIteratorLast(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	if actualValue, expectedValue := it.Last(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	list.Add("a", "b", "c")
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 2 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "c")
	}
}
