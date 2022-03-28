package linkedhashset

import (
	"strconv"
	"testing"
)

func benchmarkContains(b *testing.B, set *Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

var setSize []int = []int{100, 1000, 10000, 100000}

func BenchmarkLinkedHashSetContains(b *testing.B) {

	for _, size := range setSize {
		b.StopTimer()
		set := New[int]()
		for n := 0; n < size; n++ {
			set.Add(n)
		}
		b.StartTimer()
		b.Run("Contains"+strconv.Itoa(size), func(b *testing.B) {
			benchmarkContains(b, set, size)
		})
	}
}

func BenchmarkLinkedHashSetAdd(b *testing.B) {
	for _, size := range setSize {
		b.StopTimer()
		set := New[int]()
		b.StartTimer()
		b.Run("Add"+strconv.Itoa(size), func(b *testing.B) {
			benchmarkAdd(b, set, size)
		})
	}
}

func BenchmarkLinkedHashSetRemove(b *testing.B) {
	for _, size := range setSize {
		b.StopTimer()
		set := New[int]()
		for n := 0; n < size; n++ {
			set.Add(n)
		}
		b.StartTimer()
		b.Run("Remove"+strconv.Itoa(size), func(b *testing.B) {
			benchmarkRemove(b, set, size)
		})
	}
}