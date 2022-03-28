package singlylinkedlist

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
