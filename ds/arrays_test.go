// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

import "testing"

func TestArray(t *testing.T) {
	const midcap, maxcap, n = 84, 128, 65
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

	if _, ok := a.Remove(0); ok {
		t.Errorf("no element to remove at index 0")
	}
	if _, ok := a.Remove(-1); ok {
		t.Errorf("no element to remove at index -1")
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

	a.reverse()
	for i := 0; i < n; i++ {
		v, ok := a.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		if e := (n - i - 1) * (n - i - 1); v != e {
			t.Errorf("want %d, got %v", e, v)
		}
	}
	a.reverse()

	var o Array
	o.addAll(a)
	for i := 0; i < n; i++ {
		v, ok := o.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		if v != i*i {
			t.Errorf("want %d, got %v", i*i, v)
		}
	}
	if o.Len() != a.Len() {
		t.Errorf("want %d, got %d", a.Len(), o.Len())
	}

	o = o.sub(0, o.Len()/2)
	if o.Len() != n/2 {
		t.Errorf("want %d, got %d", n/2, o.Len())
	}
	for i := 0; i < n/2; i++ {
		v, ok := o.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		if v != i*i {
			t.Errorf("want %d, got %v", i*i, v)
		}
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
		if q.Len() != i+1 {
			t.Errorf("want %d, got %d", i, q.Len())
		}
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
		if q.Len() != n-i-1 {
			t.Errorf("want %d, got %d", n-i-1, q.Len())
		}
	}

	if _, ok := q.Dequeue(); ok {
		t.Errorf("no element in queue expected")
	}
	if len(q.s) != mincap {
		t.Errorf("want %d, got %d", mincap, len(q.s))
	}
}

func TestDequeue(t *testing.T) {
	const midcap, maxcap, n = 84, 128, 65
	var d Dequeue

	if _, ok := d.Get(0); ok {
		t.Errorf("no element at index 0")
	}
	if _, ok := d.Get(-1); ok {
		t.Errorf("no element at index -1")
	}

	if _, ok := d.Set(0, 1); ok {
		t.Errorf("no element to set at index 0")
	}
	if _, ok := d.Set(-1, 1); ok {
		t.Errorf("no element to set at index -1")
	}

	if _, ok := d.Remove(0); ok {
		t.Errorf("no element to remove at index 0")
	}
	if _, ok := d.Remove(-1); ok {
		t.Errorf("no element to remove at index -1")
	}

	for i := 0; i < n; i++ {
		if ok := d.Add(i, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if d.Len() != n {
		t.Errorf("want %d, got %d", n, d.Len())
	}
	if len(d.s) != maxcap {
		t.Errorf("want %d, got %d", maxcap, len(d.s))
	}

	for i := 0; i < n; i++ {
		v, ok := d.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		d.Set(i, v.(int)*v.(int))
	}

	for i := n - 1; i >= 0; i -= 2 {
		r, ok := d.Remove(i)
		if !ok {
			t.Errorf("cannot remove: %d", i)
			continue
		}
		if r != i*i {
			t.Errorf("want %d, got %v", i*i, r)
		}
	}

	if d.Len() != n/2 {
		t.Errorf("want %d, got %d", n/2, d.Len())
	}
	if len(d.s) != midcap {
		t.Errorf("want %d, got %d", midcap, len(d.s))
	}

	for i := 0; i < n/2+1; i++ {
		if ok := d.Add(i*2, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if d.Len() != n {
		t.Errorf("want %d, got %d", n, d.Len())
	}
	if len(d.s) != midcap {
		t.Errorf("want %d, got %d", midcap, len(d.s))
	}
}

func TestRootishStack(t *testing.T) {
	const n = 65
	var r RootishStack

	if _, ok := r.Get(0); ok {
		t.Errorf("no element at index 0")
	}
	if _, ok := r.Get(-1); ok {
		t.Errorf("no element at index -1")
	}

	if _, ok := r.Set(0, 1); ok {
		t.Errorf("no element to set at index 0")
	}
	if _, ok := r.Set(-1, 1); ok {
		t.Errorf("no element to set at index -1")
	}

	if _, ok := r.Remove(0); ok {
		t.Errorf("no element to remove at index 0")
	}
	if _, ok := r.Remove(-1); ok {
		t.Errorf("no element to remove at index -1")
	}

	for i := 0; i < n; i++ {
		if ok := r.Add(i, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if r.Len() != n {
		t.Errorf("want %d, got %d", n, r.Len())
	}

	for i := 0; i < n; i++ {
		v, ok := r.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		r.Set(i, v.(int)*v.(int))
	}

	for i := n - 1; i >= 0; i -= 2 {
		x, ok := r.Remove(i)
		if !ok {
			t.Errorf("cannot remove: %d", i)
			continue
		}
		if x != i*i {
			t.Errorf("want %d, got %v", i*i, x)
		}
	}

	if r.Len() != n/2 {
		t.Errorf("want %d, got %d", n/2, r.Len())
	}

	for i := 0; i < n/2+1; i++ {
		if ok := r.Add(i*2, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if r.Len() != n {
		t.Errorf("want %d, got %d", n, r.Len())
	}
}

func TestDualDequeue(t *testing.T) {
	const n = 65
	var d DualDequeue

	if _, ok := d.Get(0); ok {
		t.Errorf("no element at index 0")
	}
	if _, ok := d.Get(-1); ok {
		t.Errorf("no element at index -1")
	}

	if _, ok := d.Set(0, 1); ok {
		t.Errorf("no element to set at index 0")
	}
	if _, ok := d.Set(-1, 1); ok {
		t.Errorf("no element to set at index -1")
	}

	if _, ok := d.Remove(0); ok {
		t.Errorf("no element to remove at index 0")
	}
	if _, ok := d.Remove(-1); ok {
		t.Errorf("no element to remove at index -1")
	}

	for i := 0; i < n; i++ {
		if ok := d.Add(i, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
		if d.Len() != i+1 {
			t.Errorf("want %d, got %d", i+1, d.Len())
		}
	}
	if d.Len() != n {
		t.Errorf("want %d, got %d", n, d.Len())
	}

	for i := 0; i < n; i++ {
		v, ok := d.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		d.Set(i, v.(int)*v.(int))
	}

	for i := n - 1; i >= 0; i -= 2 {
		x, ok := d.Remove(i)
		if !ok {
			t.Errorf("cannot remove: %d", i)
			continue
		}
		if x != i*i {
			t.Errorf("want %d, got %v", i*i, x)
		}
	}

	if d.Len() != n/2 {
		t.Errorf("want %d, got %d", n/2, d.Len())
	}

	for i := 0; i < n/2+1; i++ {
		if ok := d.Add(i*2, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if d.Len() != n {
		t.Errorf("want %d, got %d", n, d.Len())
	}
}
