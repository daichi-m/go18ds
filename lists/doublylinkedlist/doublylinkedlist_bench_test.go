package doublylinkedlist

import (
	"strconv"
	"testing"
)

func benchmarkGet(b *testing.B, list *List[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, list *List[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, list *List[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Remove(n)
		}
	}
}

var listSizes []int = []int{100, 1000, 10000, 100000}

func BenchmarkArrayListGet(b *testing.B) {
	for _, size := range listSizes {
		b.StopTimer()
		list := New[int]()
		for n := 0; n < size; n++ {
			list.Add(n)
		}
		b.StartTimer()
		b.Run("Get"+strconv.Itoa(size), func(b *testing.B) {
			benchmarkGet(b, list, size)
		})
	}
}

func BenchmarkArrayListAdd(b *testing.B) {
	for _, size := range listSizes {
		list := New[int]()
		b.Run("Add"+strconv.Itoa(size), func(b *testing.B) {
			benchmarkAdd(b, list, size)
		})
	}
}

func BenchmarkArrayListRemove(b *testing.B) {
	for _, size := range listSizes {
		b.StopTimer()
		list := New[int]()
		for n := 0; n < size; n++ {
			list.Add(n)
		}
		b.StartTimer()
		b.Run("Remove"+strconv.Itoa(size), func(b *testing.B) {
			benchmarkRemove(b, list, size)
		})
	}
}
