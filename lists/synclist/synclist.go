package synclist

import (
	"sync"

	"github.com/daichi-m/go18ds/lists"
	"github.com/daichi-m/go18ds/lists/arraylist"
	"github.com/daichi-m/go18ds/utils"
)

func assertListImplementation() {
	var _ lists.List[string] = (*List[string])(nil)
}

type List[T comparable] struct {
	list  lists.List[T]
	mutex sync.Mutex
}

// New instantiates a new list and adds the passed values, if any, to the list
func New[T comparable](values ...T) *List[T] {
	intList := arraylist.New(values...)
	return NewSyncList[T](intList)
}

// New instantiates a new list and adds the passed values, if any, to the list
func NewSyncList[T comparable](l lists.List[T]) *List[T] {
	return &List[T]{
		list:  l,
		mutex: sync.Mutex{},
	}
}

// Add appends a value at the end of the list
func (list *List[T]) Add(values ...T) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Add(values...)
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[T]) Get(index int) (T, bool) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.list.Get(index)
}

// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Remove(index)
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(values ...T) bool {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.list.Contains(values...)
}

// Values returns all elements in the list.
func (list *List[T]) Values() []T {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.list.Values()
}

// Empty returns true if list does not contain any elements.
func (list *List[T]) Empty() bool {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.list.Empty()
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	return list.list.Size()
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Clear()
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Sort(comparator)
}

// Swap swaps the two values at the specified positions.
func (list *List[T]) Swap(i, j int) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Swap(i, j)
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Insert(index, values...)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value T) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.list.Set(index, value)
}

// String returns a string representation of container
func (list *List[T]) String() string {
	return list.String()
}
