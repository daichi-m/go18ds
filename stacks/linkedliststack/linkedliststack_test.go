// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import (
	"strings"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := New[int]()
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if actualValue := stack.Values(); actualValue[0] != 3 || actualValue[1] != 2 || actualValue[2] != 1 {
		t.Errorf("Got %v expected %v", actualValue, "[3,2,1]")
	}
	if actualValue := stack.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := stack.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestStackPeek(t *testing.T) {
	stack := New[int]()
	if actualValue, ok := stack.Peek(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestStackPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	if actualValue, ok := stack.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := stack.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := stack.Pop(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, ok := stack.Pop(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := stack.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func TestStackIterator(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Iterator
	it := stack.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := index, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	stack.Clear()
	it = stack.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty stack")
	}
}

func TestStackIteratorBegin(t *testing.T) {
	stack := New[int]()
	it := stack.Iterator()
	it.Begin()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != 3 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, 3)
	}
}

func TestStackIteratorFirst(t *testing.T) {
	stack := New[int]()
	it := stack.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != 3 {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, 3)
	}
}

func TestStackSerialization(t *testing.T) {
	stack := New[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	var err error
	assert := func() {
		if actualValue, expectedValue := strings.Join(stack.Values(), ""), "cba"; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue, expectedValue := stack.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	json, err := stack.ToJSON()
	assert()

	err = stack.FromJSON(json)
	assert()
}

func benchmarkPush(b *testing.B, stack *Stack[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Push(n)
		}
	}
}

func benchmarkPop(b *testing.B, stack *Stack[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Pop()
		}
	}
}

func BenchmarkLinkedListStackPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, stack, size)
}

func BenchmarkLinkedListStackPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	stack := New[int]()
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkLinkedListStackPush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkLinkedListStackPush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}

func BenchmarkLinkedListStackPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	stack := New[int]()
	for n := 0; n < size; n++ {
		stack.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, stack, size)
}
