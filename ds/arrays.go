// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

import "math"

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

func (a *Array) addAll(o Array) {
	n := o.n + a.n
	if n > len(a.s) {
		s := make([]V, n)
		copy(s, a.s)
		a.s = s
	}
	copy(a.s[a.n:], o.s)
	a.n = n
}

func (a *Array) reverse() {
	for i, j := 0, a.n-1; i < j; i, j = i+1, j-1 {
		a.s[i], a.s[j] = a.s[j], a.s[i]
	}
}

func (a *Array) sub(f, t int) Array {
	var o Array
	/*
		if f < 0 || t > a.n || f > t {
			panic()
		}
	*/
	o.n = t - f
	o.s = make([]V, o.n)
	copy(o.s, a.s[f:t])
	return o
}

// Remove removes the element of the array
// at the given index and reports whether the
// operation was successful or not. The array
// is resized as needed.
//
// This operation has an amortized time
// complexity of O(n-i).
func (a *Array) Remove(i int) (V, bool) {
	if i < 0 || i > a.n-1 {
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
	r int // read offset
	n int // number of elements
}

// Len returns the number
// of elements in the queue.
func (q *Queue) Len() int { return q.n }

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

// --- Dequeue -------

// Dequeue is a queue which allows
// for efficient addition and removal
// at both ends of the queue.
type Dequeue struct {
	s []V // backing slice
	r int // read offset
	n int // number of elements
}

// Len returns the number
// of elements in the dequeue.
func (d *Dequeue) Len() int { return d.n }

// Get returns the element at the
// given index.
//
// This operation has a time complexity of O(1).
func (d *Dequeue) Get(i int) (V, bool) {
	if i < 0 || i > d.n-1 {
		return nil, false
	}
	return d.s[(d.r+i)%len(d.s)], true
}

// Set sets the element at the given
// index and returns the old one.
//
// This operation has a time complexity of O(1).
func (d *Dequeue) Set(i int, v V) (V, bool) {
	if i < 0 || i > d.n-1 {
		return nil, false
	}
	t := d.s[(d.r+i)%len(d.s)]
	d.s[(d.r+i)%len(d.s)] = v
	return t, true
}

// Add adds an element to the dequeue at
// the given index and reports whether it
// was successful or not. The dequeue is
// resized as needed.
//
// This operation has an amortized time
// complexity of O(min{i, n-i}).
func (d *Dequeue) Add(i int, v V) bool {
	if i < 0 || i > d.n {
		return false
	}
	if d.n+1 > len(d.s) {
		d.resize()
	}
	if i < d.n/2 {
		// shift left one position
		if d.r == 0 {
			d.r = len(d.s) - 1
		} else {
			d.r--
		}
		for j := 0; j <= i-1; j++ {
			d.s[(d.r+j)%len(d.s)] = d.s[(d.r+j+1)%len(d.s)]
		}
	} else {
		// shift right one position
		for j := d.n; j > i; j-- {
			d.s[(d.r+j)%len(d.s)] = d.s[(d.r+j-1)%len(d.s)]
		}
	}
	d.s[(d.r+i)%len(d.s)] = v
	d.n++
	return true
}

// Remove removes the element of the dequeue
// at the given index and reports whether the
// operation was successful or not. The dequeue
// is resized as needed.
//
// This operation has an amortized time
// complexity of O(min{i, n-i}).
func (d *Dequeue) Remove(i int) (V, bool) {
	if i < 0 || i > d.n-1 {
		return nil, false
	}
	t := d.s[(d.r+i)%len(d.s)]
	if i < d.n/2 {
		// shift right one position
		for j := i; j > 0; j-- {
			d.s[(d.r+j)%len(d.s)] = d.s[(d.r+j-1)%len(d.s)]
		}
		d.r = (d.r + 1) % len(d.s)
	} else {
		// shift left one position
		for j := i; j < d.n-1; j++ {
			d.s[(d.r+j)%len(d.s)] = d.s[(d.r+j+1)%len(d.s)]
		}
	}
	d.n--
	if 3*d.n < len(d.s) {
		d.resize()
	}
	return t, true
}

func (d *Dequeue) resize() {
	s := make([]V, max(d.n*2, 1))
	for i := 0; i < d.n; i++ {
		s[i] = d.s[(d.r+i)%len(d.s)]
	}
	d.s = s
	d.r = 0
}

// --- DualDequeue -------

// DualDequeue implements
// a dequeue by combining
// two dynamic arrays.
type DualDequeue struct {
	f, b Array // backing arrays
}

// Get returns the element at the
// given index.
//
// This operation has a time complexity of O(1).
func (d *DualDequeue) Get(i int) (V, bool) {
	l := d.f.Len()
	if i < l {
		return d.f.Get(l - i - 1)
	}
	return d.b.Get(i - l)
}

// Set sets the element at the given
// index and returns the old one.
//
// This operation has a time complexity of O(1).
func (d *DualDequeue) Set(i int, v V) (V, bool) {
	l := d.f.Len()
	if i < l {
		return d.f.Set(l-i-1, v)
	}
	return d.b.Set(i-l, v)
}

// Add adds an element to the dual dequeue at
// the given index and reports whether it was
// was successful or not. The dual dequeue is
// resized as needed.
//
// This operation has an amortized time
// complexity of O(min{i, n-i}).
func (d *DualDequeue) Add(i int, v V) bool {
	if l := d.f.Len(); i < l {
		d.f.Add(l-i, v)
	} else {
		d.b.Add(i-l, v)
	}
	d.balance()
	return true
}

// Remove removes the element of the dual dequeue
// at the given index and reports whether the operation
// was successful or not. The dual dequeue is resized
// as needed.
//
// This operation has an amortized time
// complexity of O(min{i, n-i}).
func (d *DualDequeue) Remove(i int) (V, bool) {
	var t V
	var ok bool
	if l := d.f.Len(); i < l {
		t, ok = d.f.Remove(l - i - 1)
		if !ok {
			return nil, false
		}
	} else {
		t, ok = d.b.Remove(i - l)
		if !ok {
			return nil, false
		}
	}
	d.balance()
	return t, true
}

func (d *DualDequeue) balance() {
	if 3*d.f.Len() < d.b.Len() {
		var f, b Array
		s := d.Len()/2 - d.f.Len()
		f.addAll(d.b.sub(0, s))
		f.reverse()
		f.addAll(d.f)
		b.addAll(d.b.sub(s, d.b.Len()))
		d.f, d.b = f, b
	} else if 3*d.b.Len() < d.f.Len() {
		var f, b Array
		s := d.f.Len() - d.Len()/2
		f.addAll(d.f.sub(s, d.f.Len()))
		b.addAll(d.f.sub(0, s))
		b.reverse()
		b.addAll(d.b)
		d.f, d.b = f, b
	}
}

// Len returns the number of
// elements in the dual dequeue.
func (d *DualDequeue) Len() int { return d.f.Len() + d.b.Len() }

// --- RootishStack -------

// RootishStack uses arrays in
// arrays to store elements and
// wastes at most O(sqrt(n)) space
// when storing n items.
type RootishStack struct {
	b Dequeue // backing blocks
	n int     // number of elements
}

// Get returns the element at the
// given index.
//
// This operation has a time complexity of O(1).
func (r *RootishStack) Get(i int) (V, bool) {
	if i < 0 || i > r.n-1 {
		return nil, false
	}
	b := i2b(i)
	a, _ := r.b.Get(b)
	return a.([]V)[i-b*(b+1)/2], true
}

// Set sets the element at the given
// index and returns the old one.
//
// This operation has a time complexity of O(1).
func (r *RootishStack) Set(i int, v V) (V, bool) {
	if i < 0 || i > r.n-1 {
		return nil, false
	}
	b := i2b(i)
	j := i - b*(b+1)/2
	a, _ := r.b.Get(b)
	t := a.([]V)[j]
	a.([]V)[j] = v
	return t, true
}

// Add adds an element to the stack at
// the given index and reports whether it
// was successful or not. The stack is
// resized as needed.
//
// This operation has an amortized time
// complexity of O(n-i).
func (r *RootishStack) Add(i int, v V) bool {
	if i < 0 || i > r.n {
		return false
	}
	if l := r.b.Len(); l*(l+1)/2 < r.n+1 {
		r.b.Add(r.b.Len(), V(make([]V, r.b.Len()+1)))
	}
	r.n++
	for j := r.n - 1; j > i; j-- {
		p, _ := r.Get(j - 1)
		r.Set(j, p)
	}
	r.Set(i, v)
	return true
}

// Remove removes the element of the stack
// at the given index and reports whether the
// operation was successful or not. The stack
// is resized as needed.
//
// This operation has an amortized time
// complexity of O(n-i).
func (r *RootishStack) Remove(i int) (V, bool) {
	if i < 0 || i > r.n-1 {
		return nil, false
	}
	t, _ := r.Get(i)
	for j := i; j < r.n-1; j++ {
		p, _ := r.Get(j + 1)
		r.Set(j, p)
	}
	r.n--
	for l := r.b.Len(); l > 0 && (l-2)*(l-1)/2 >= r.n; l-- {
		r.b.Remove(r.b.Len() - 1)
	}
	return t, true
}

// Len returns the number
// of elements in the stack.
func (r *RootishStack) Len() int { return r.n }

func i2b(i int) int {
	return int(math.Ceil((-3 + math.Sqrt(9+8*float64(i))) / 2.0))
}

// --- Utilities -------

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
