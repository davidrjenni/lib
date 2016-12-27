// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

import "testing"

func TestArray(t *testing.T) {
	const mincap, midcap, maxcap, n = 1, 84, 128, 65
	var a Array

	if _, ok := a.Get(0); ok {
		t.Errorf("no element at index 0")
	}
	if _, ok := a.Get(-1); ok {
		t.Errorf("no element at index -1")
	}

	if _, ok := a.Set(0, 1); ok {
		t.Errorf("no element to set at index 0")
	}
	if _, ok := a.Set(-1, 1); ok {
		t.Errorf("no element to set at index -1")
	}

	for i := 0; i < n; i++ {
		if ok := a.Add(i, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if a.Len() != n {
		t.Errorf("want %d, got %d", n, a.Len())
	}
	if len(a.s) != maxcap {
		t.Errorf("want %d, got %d", maxcap, len(a.s))
	}

	for i := 0; i < n; i++ {
		v, ok := a.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		a.Set(i, v.(int)*v.(int))
	}

	for i := n - 1; i >= 0; i -= 2 {
		r, ok := a.Remove(i)
		if !ok {
			t.Errorf("cannot remove: %d", i)
			continue
		}
		if r != i*i {
			t.Errorf("want %d, got %v", i*i, r)
		}
	}

	if a.Len() != n/2 {
		t.Errorf("want %d, got %d", n/2, a.Len())
	}
	if len(a.s) != midcap {
		t.Errorf("want %d, got %d", midcap, len(a.s))
	}
}

func TestStack(t *testing.T) {
	const mincap, maxcap, n = 1, 128, 65
	var s Stack

	for i := 0; i < n; i++ {
		s.Push(i)
	}
	if s.Len() != n {
		t.Errorf("want %d, got %d", n, s.Len())
	}
	if len(s.a.s) != maxcap {
		t.Errorf("want %d, got %d", maxcap, len(s.a.s))
	}

	for i := n - 1; i >= 0; i-- {
		r, ok := s.Pop()
		if !ok {
			t.Errorf("want %d, not found", i)
			continue
		}
		if r != i {
			t.Errorf("want %d, got %v", i, r)
		}
	}

	if _, ok := s.Pop(); ok {
		t.Errorf("no element in stack expected")
	}
	if len(s.a.s) != mincap {
		t.Errorf("want %d, got %d", mincap, len(s.a.s))
	}
}

func TestQueue(t *testing.T) {
	const mincap, maxcap, n = 1, 128, 65
	var q Queue

	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}
	if len(q.s) != maxcap {
		t.Errorf("want %d, got %d", maxcap, len(q.s))
	}

	for i := 0; i < n; i++ {
		r, ok := q.Dequeue()
		if !ok {
			t.Errorf("want %d, not found", i)
			continue
		}
		if r != i {
			t.Errorf("want %d, got %v", i, r)
		}
	}

	if _, ok := q.Dequeue(); ok {
		t.Errorf("no element in queue expected")
	}
	if len(q.s) != mincap {
		t.Errorf("want %d, got %d", mincap, len(q.s))
	}
}
