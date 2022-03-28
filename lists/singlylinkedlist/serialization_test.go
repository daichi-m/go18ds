package singlylinkedlist

import (
	"strings"
	"testing"
)

func TestListSerialization(t *testing.T) {
	list := New[string]()
	list.Add("a", "b", "c")

	var err error
	assert := func() {
		if actualValue, expectedValue := strings.Join(list.Values(), ""), "abc"; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := list.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	json, err := list.ToJSON()
	assert()

	err = list.FromJSON(json)
	assert()
}
