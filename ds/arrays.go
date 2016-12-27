// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ds the contains implementations
// of the data structures from
//	http://opendatastructures.org.
// It is not recommended to use this package.
package ds

// V represents a value
// stored in a data structure.
type V interface{}

// --- Stack -------

// Stack implements a stack
// on top of a dynamic array.
type Stack struct {
	a Array // backing dynamic array
}

// Push pushes a value onto the stack.
//
// This operation has an amortized time
// complexity of O(1).
func (s *Stack) Push(v V) { s.a.Add(s.Len(), v) }

// Pop removes and returns a value from
// the top of the stack.
//
// This operation has an amortized time
// complexity of O(1).
func (s *Stack) Pop() (V, bool) { return s.a.Remove(s.Len() - 1) }

// Len returns the number of
// elements on the stack.
func (s *Stack) Len() int { return s.a.Len() }

// --- Dynamic Array -------

// Array implements a dynamic array, which
// grows and shrinks as needed.
type Array struct {
	s []V // backing slice
	n int // number of elements
}

// Len returns the number
// of elements in the array.
func (a *Array) Len() int { return a.n }

// Get returns the element at the
// given index.
//
// This operation has a time complexity of O(1).
func (a *Array) Get(i int) (V, bool) {
	if i < 0 || i > a.n-1 {
		return nil, false
	}
	return a.s[i], true
}

// Set sets the element at the given
// index and returns the old one.
//
// This operation has a time complexity of O(1).
func (a *Array) Set(i int, v V) (V, bool) {
	if i < 0 || i > a.n-1 {
		return nil, false
	}
	t := a.s[i]
	a.s[i] = v
	return t, true
}

// Add adds an element to the array at
// the given index and reports whether
// it was successful or not. The array is
// resized as needed.
//
// This operation has an amortized time
// complexity of O(n-i).
func (a *Array) Add(i int, v V) bool {
	if i < 0 || i > a.n {
		return false
	}
	if a.n+1 > len(a.s) {
		a.resize()
	}
	copy(a.s[i:], a.s[i+1:])
	a.s[i] = v
	a.n++
	return true
}

// Remove removes the element of the array
// at the given index and reports whether the
// operation was successful or not. The array
// is resized as needed.
//
// This operation has an amortized time
// complexity of O(n-i).
func (a *Array) Remove(i int) (V, bool) {
	if i < 0 || i > a.n {
		return nil, false
	}
	v := a.s[i]
	copy(a.s[i:], a.s[i+1:])
	a.n--
	if len(a.s) >= 3*a.n {
		a.resize()
	}
	return v, true
}

func (a *Array) resize() {
	s := make([]V, max(a.n*2, 1))
	copy(s, a.s)
	a.s = s
}

// --- Queue -------

// Queue implements a queue
// on top of a slice.
type Queue struct {
	s []V // backing slice
	r int // read index
	n int // number of elements
}

// Enqueue adds an element
// to the head of the queue.
//
// This operation has an amortized time
// complexity of O(1).
func (q *Queue) Enqueue(v V) {
	if q.n+1 > len(q.s) {
		q.resize()
	}
	q.s[(q.r+q.n)%len(q.s)] = v
	q.n++
}

// Dequeue removes the element
// from the tail of the queue.
//
// This operation has an amortized time
// complexity of O(1).
func (q *Queue) Dequeue() (V, bool) {
	if q.n == 0 {
		return nil, false
	}
	v := q.s[q.r]
	q.r = (q.r + 1) % len(q.s)
	q.n--
	if len(q.s) >= 3*q.n {
		q.resize()
	}
	return v, true
}

func (q *Queue) resize() {
	s := make([]V, max(q.n*2, 1))
	for i := 0; i < q.n; i++ {
		s[i] = q.s[(q.r+i)%len(q.s)]
	}
	q.s = s
	q.r = 0
}

// --- Utilities -------

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
