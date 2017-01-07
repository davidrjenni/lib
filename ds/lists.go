// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

// --- SList -------

// snode represents a node
// in a singly-linked list.
type snode struct {
	n *snode // next pointer
	v V      // value
}

// SList is a singly-linked list with
// a head and tail pointer.
// It can be used as stack or queue.
type SList struct {
	h, t *snode // head and tail pointer
	n    int    // number of elements
}

// Push pushes a value onto the head of the list.
//
// This operation has a time complexity of O(1).
func (l *SList) Push(v V) {
	l.h = &snode{v: v, n: l.h}
	if l.n == 0 {
		l.t = l.h
	}
	l.n++
}

// Pop removes and returns a value from the
// head of the list.
//
// This operation has a time complexity of O(1).
func (l *SList) Pop() (V, bool) {
	if l.n == 0 {
		return nil, false
	}
	v := l.h.v
	l.h = l.h.n
	l.n--
	if l.n == 0 {
		l.t = nil
	}
	return v, true
}

// Enqueue adds an element to the tail of the list.
//
// This operation has a time complexity of O(1).
func (l *SList) Enqueue(v V) {
	n := &snode{v: v}
	if l.n == 0 {
		l.h = n
	} else {
		l.t.n = n
	}
	l.t = n
	l.n++
}

// Dequeue removes the element from the head of
// the list and is equivalent to Pop.
//
// This operation has a time complexity of O(1).
func (l *SList) Dequeue() (V, bool) { return l.Pop() }

// Len returns the number of elements in the list.
func (l *SList) Len() int { return l.n }

// --- DList -------

// dnode represents a node
// in a doubly-linked list.
type dnode struct {
	n *dnode // next pointer
	p *dnode // previous pointer
	v V      // value
}

// DList represents a doubly-linked list.
type DList struct {
	r *dnode // sentinel node, the root
	n int    // number of elements
}

// get returns the node at the
// given index or nil if not found.
func (l *DList) get(i int) *dnode {
	if i < l.n/2 {
		n := l.r.n
		for j := 0; j < i; j++ {
			n = n.n
		}
		return n
	}
	n := l.r
	for j := l.n; j > i; j-- {
		n = n.p
	}
	return n
}

// Len returns the number of elements in the list.
func (l *DList) Len() int { return l.n }

// Add adds an element to the list at the
// given index and reports whether it was
// successful or not.
//
// This operation has a time complexity
// of O(min{i, n-i}).
func (l *DList) Add(i int, v V) bool {
	if i < 0 || i > l.n {
		return false
	}
	// lazy initialization
	if l.r == nil {
		l.r = new(dnode)
		l.r.n = l.r
		l.r.p = l.r
	}
	p := l.get(i)
	n := &dnode{
		v: v,
		p: p.p,
		n: p,
	}
	n.n.p = n
	n.p.n = n
	l.n++
	return true
}

// Remove removes the element of the list at the
// given index and reports whether the operation was
// successful or not.
//
// This operation has a time complexity
// of O(min{i, n-i}).
func (l *DList) Remove(i int) (V, bool) {
	if i < 0 || i > l.n-1 {
		return nil, false
	}
	n := l.get(i)
	n.p.n = n.n
	n.n.p = n.p
	n.n = nil
	n.p = nil
	l.n--
	return n.v, true
}

// Get returns the element at the given index.
//
// This operation has a time complexity
// of O(min{i, n-i}).
func (l *DList) Get(i int) (V, bool) {
	if i < 0 || i > l.n-1 {
		return nil, false
	}
	return l.get(i).v, true
}

// Set sets the element at the given index
// and returns the old one.
//
// This operation has a time complexity
// of O(min{i, n-i}).
func (l *DList) Set(i int, v V) (V, bool) {
	if i < 0 || i > l.n-1 {
		return nil, false
	}
	n := l.get(i)
	t := n.v
	n.v = v
	return t, true
}
