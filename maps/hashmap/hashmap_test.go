// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"fmt"
	"testing"
)

func sameElements[K comparable](a []K, b []K) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestMapPut(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := m.Keys(), []int{1, 2, 3, 4, 5, 6, 7}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := m.Values(), []string{"a", "b", "c", "d", "e", "f", "g"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// key,expectedValue,expectedFound
	tests1 := []struct {
		key   int
		val   string
		found bool
	}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.val || actualFound != test.found {
			t.Errorf("Got %v expected %v", actualValue, test.val)
		}
	}
}

func TestMapRemove(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	if actualValue, expectedValue := m.Keys(), []int{1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := m.Values(), []string{"a", "b", "c", "d"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	tests2 := []struct {
		key   int
		val   string
		found bool
	}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "", false},
		{6, "", false},
		{7, "", false},
		{8, "", false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.val || actualFound != test.found {
			t.Errorf("Got %v expected %v", actualValue, test.val)
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%v", m.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", m.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapSerialization(t *testing.T) {
	m := New[string, float32]()
	m.Put("a", 1.0)
	m.Put("b", 2.0)
	m.Put("c", 3.0)

	var err error
	assert := func() {
		if actualValue, expectedValue := m.Keys(), []string{"a", "b", "c"}; !sameElements(actualValue, expectedValue) {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := m.Values(), []float32{1.0, 2.0, 3.0}; !sameElements(actualValue, expectedValue) {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := m.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	json, err := m.ToJSON()
	assert()

	err = m.FromJSON(json)
	assert()
}

func benchmarkGet(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkHashMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
