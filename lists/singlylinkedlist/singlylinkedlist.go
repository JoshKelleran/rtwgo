// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package singlylinkedlist implements the singly-linked list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package singlylinkedlist

import (
	"github.com/emirpasic/gods/v2/lists"
	"github.com/emirpasic/gods/v2/utils"
)

// Assert List implementation
var _ lists.List[int] = (*List[int])(nil)

// List holds the elements, where each element points to the next element
type List[T comparable] struct {
	first *element[T]
	last  *element[T]
	size  int
}

type element[T comparable] struct {
	value T
	next  *element[T]
}

// New instantiates a new list and adds the passed values, if any, to the list
func New[T comparable](values ...T) *List[T] {
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List[T]) Add(values ...T) {
}

// Append appends a value (one or more) at the end of the list (same as Add())
func (list *List[T]) Append(values ...T) {
}

// Prepend prepends a values (or more)
func (list *List[T]) Prepend(values ...T) {
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[T]) Get(index int) (T, bool) {
}

// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {

}

// Contains checks if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(values ...T) bool {

}

// Values returns all elements in the list.
func (list *List[T]) Values() []T {

}

// IndexOf returns index of provided element
func (list *List[T]) IndexOf(value T) int {

}

// Empty returns true if list does not contain any elements.
func (list *List[T]) Empty() bool {
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
}

// Sort sort values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {

}

// Swap swaps values of two elements at the given indices.
func (list *List[T]) Swap(i, j int) {
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {
}

// Set value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value T) {
}

// String returns a string representation of container
func (list *List[T]) String() string {
}

// Check that the index is within bounds of the list
func (list *List[T]) withinRange(index int) bool {
}
