package arraylist

import (
	"testing"
)

func TestListEach(t *testing.T) {
	list := New("a", "b", "c")
	expect := []string{"a", "b", "c"}
	list.Each(func(index int, value string) {
		if index >= len(expect) {
			t.Errorf("Index %v, supposed to be less than %v", index, len(expect)-1)
		}
		if actualValue, expectedValue := value, expect[index]; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	})
}

func TestListMap(t *testing.T) {
	list := New("a", "b", "c")
	tests := []struct {
		idx    int
		expect string
	}{
		{0, "mapped: a"},
		{1, "mapped: b"},
		{2, "mapped: c"},
	}
	mappedList := list.Map(func(index int, value string) string {
		return "mapped: " + value
	})
	for _, tt := range tests {
		if actual, _ := mappedList.Get(tt.idx); actual != tt.expect {
			t.Errorf("Got %v expected %v", actual, tt.expect)
		}
	}
}

func TestListSelect(t *testing.T) {
	list := New("a", "b", "c")
	tests := []struct {
		idx      int
		expect   string
		expectOk bool
	}{
		{0, "a", true},
		{1, "b", true},
		{2, "", false},
	}
	selectedList := list.Select(func(index int, value string) bool {
		return value >= "a" && value <= "b"
	})
	for _, tt := range tests {
		if actual, ok := selectedList.Get(tt.idx); actual != tt.expect || ok != tt.expectOk {
			t.Errorf("Got %v, %v expected %v, %v", actual, ok, tt.expect, tt.expectOk)
		}
	}
	if selectedList.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedList.Size(), 3)
	}
}

func TestListFilter(t *testing.T) {
	list := New[string]()
	list.Add("a", "b", "c", "d", "e", "f", "g")
	all := list.Filter(func(index int, value string) bool {
		return value >= "a" && value <= "c"
	})
	expect := []string{"a", "b", "c"}
	if !equals(all.Values(), expect) {
		t.Errorf("Got %v expected %v", all.Values(), expect)
	}
	all = list.Filter(func(index int, value string) bool {
		return len(value) > 1
	})
	if !equals(all.Values(), []string{}) {
		t.Errorf("Got %v expected %v", all, []string{})
	}
}

func TestListFind(t *testing.T) {
	list := New[string]()
	list.Add("a", "b", "c")
	foundIndex, foundValue := list.Find(func(index int, value string) bool {
		return value == "c"
	})
	if foundValue != "c" || foundIndex != 2 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, "c", 2)
	}
	foundIndex, foundValue = list.Find(func(index int, value string) bool {
		return value == "x"
	})
	if foundValue != "" || foundIndex != -1 {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundIndex, nil, nil)
	}
}
func TestListChaining(t *testing.T) {
	list := New[string]()
	list.Add("a", "b", "c")
	chainedList := list.Select(func(index int, value string) bool {
		return value > "a"
	}).Map(func(index int, value string) string {
		return value + value
	})
	if chainedList.Size() != 2 {
		t.Errorf("Got %v expected %v", chainedList.Size(), 2)
	}
	if actualValue, ok := chainedList.Get(0); actualValue != "bb" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := chainedList.Get(1); actualValue != "cc" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}
