// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package avltree

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

type empty struct{}

var emptyVal empty = empty{}

func TestAVLTreePut(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	if actualValue := tree.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := concatenate(tree.Keys()...), "1234567"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := concatenate(tree.Values()...), "abcdefg"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	tests1 := []struct {
		key    int
		val    string
		expect bool
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
		actualValue, actualFound := tree.Get(test.key)
		if actualValue != test.val || actualFound != test.expect {
			t.Errorf("Got %v expected %v", actualValue, test.val)
		}
	}
}

func TestAVLTreeRemove(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite

	tree.Remove(5)
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

	if actualValue, expectedValue := concatenate(tree.Keys()...), "1234"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := concatenate(tree.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := concatenate(tree.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := tree.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	tests2 := []struct {
		key    int
		val    string
		expect bool
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
		actualValue, actualFound := tree.Get(test.key)
		if actualValue != test.val || actualFound != test.expect {
			t.Errorf("Got %v expected %v", actualValue, test.val)
		}
	}

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%v", tree.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if empty, size := tree.Empty(), tree.Size(); empty != true || size != -0 {
		t.Errorf("Got %v expected %v", empty, true)
	}
}

func TestAVLTreeLeftAndRight(t *testing.T) {
	tree := NewWithIntComparator[string]()

	if actualValue := tree.Left(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := tree.Right(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	tree.Put(1, "a")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x") // overwrite
	tree.Put(2, "b")

	if actualValue, expectedValue := fmt.Sprintf("%d", tree.Left().Key), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Left().Value), "x"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := fmt.Sprintf("%d", tree.Right().Key), "7"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", tree.Right().Value), "g"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeCeilingAndFloor(t *testing.T) {
	tree := NewWithIntComparator[string]()

	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
	if node, found := tree.Ceiling(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")

	if node, found := tree.Floor(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	if node, found := tree.Ceiling(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Ceiling(8); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
}

func TestAVLTreeIteratorNextOnEmpty(t *testing.T) {
	tree := NewWithIntComparator[string]()
	it := tree.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestAVLTreeIteratorPrevOnEmpty(t *testing.T) {
	tree := NewWithIntComparator[string]()
	it := tree.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestAVLTreeIterator1Next(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator1Prev(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator2Next(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator2Prev(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator3Next(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(1, "a")
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator3Prev(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(1, "a")
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.size
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator4Next(t *testing.T) {
	tree := NewWithIntComparator[int]()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		value := it.Value()
		switch value {
		case count:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIterator4Prev(t *testing.T) {
	tree := NewWithIntComparator[int]()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		value := it.Value()
		switch value {
		case count:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := value, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		count--
	}
	if actualValue, expectedValue := count, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestAVLTreeIteratorBegin(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.Begin()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	for it.Next() {
	}

	it.Begin()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestAVLTreeIteratorEnd(t *testing.T) {
	tree := NewWithIntComparator[string]()
	it := tree.Iterator()

	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.End()
	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it.End()
	if it.Key() != 0 {
		t.Errorf("Got %v expected %v", it.Key(), 0)
	}

	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestAVLTreeIteratorFirst(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestAVLTreeIteratorLast(t *testing.T) {
	tree := NewWithIntComparator[string]()
	tree.Put(3, "c")
	tree.Put(1, "a")
	tree.Put(2, "b")
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestAVLTreeSerialization(t *testing.T) {
	tree := NewWithStringComparator[string]()
	tree.Put("c", "3")
	tree.Put("b", "2")
	tree.Put("a", "1")

	var err error
	assert := func() {
		if actualValue, expectedValue := tree.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue := tree.Keys(); actualValue[0] != "a" || actualValue[1] != "b" || actualValue[2] != "c" {
			t.Errorf("Got %v expected %v", actualValue, "[a,b,c]")
		}
		if actualValue := tree.Values(); actualValue[0] != "1" || actualValue[1] != "2" || actualValue[2] != "3" {
			t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	json, err := tree.ToJSON()
	assert()

	err = tree.FromJSON(json)
	assert()
}

func benchmarkGet(b *testing.B, tree *Tree[int, empty], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, tree *Tree[int, empty], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n, emptyVal)
		}
	}
}

func benchmarkRemove(b *testing.B, tree *Tree[int, empty], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func BenchmarkAVLTreeGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreeGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreeGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreeGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkAVLTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator[empty]()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreePut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreePut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreePut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkAVLTreeRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkAVLTreeRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkAVLTreeRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkAVLTreeRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator[empty]()
	for n := 0; n < size; n++ {
		tree.Put(n, emptyVal)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
