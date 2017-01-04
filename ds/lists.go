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
