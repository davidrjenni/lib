// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

import "testing"

func TestSListStack(t *testing.T) {
	const n = 65
	var l SList

	for i := 0; i < n; i++ {
		l.Push(i)
	}
	if l.Len() != n {
		t.Errorf("want %d, got %d", n, l.Len())
	}

	for i := n - 1; i >= 0; i-- {
		r, ok := l.Pop()
		if !ok {
			t.Errorf("want %d, not found", i)
			continue
		}
		if r != i {
			t.Errorf("want %d, got %v", i, r)
		}
	}

	if l.Len() != 0 {
		t.Errorf("want %d, got %d", 0, l.Len())
	}
	if _, ok := l.Pop(); ok {
		t.Errorf("no element in stack expected")
	}
}

func TestSListQueue(t *testing.T) {
	const n = 65
	var l SList

	for i := 0; i < n; i++ {
		l.Enqueue(i)
		if l.Len() != i+1 {
			t.Errorf("want %d, got %d", i, l.Len())
		}
	}
	if l.Len() != n {
		t.Errorf("want %d, got %d", n, l.Len())
	}

	for i := 0; i < n; i++ {
		r, ok := l.Dequeue()
		if !ok {
			t.Errorf("want %d, not found", i)
			continue
		}
		if r != i {
			t.Errorf("want %d, got %v", i, r)
		}
		if l.Len() != n-i-1 {
			t.Errorf("want %d, got %d", n-i-1, l.Len())
		}
	}

	if _, ok := l.Dequeue(); ok {
		t.Errorf("no element in list expected")
	}
	if l.Len() != 0 {
		t.Errorf("want %d, got %d", 0, l.Len())
	}
}

func TestDList(t *testing.T) {
	const n = 65
	var l DList

	if _, ok := l.Get(0); ok {
		t.Errorf("no element at index 0")
	}
	if _, ok := l.Get(-1); ok {
		t.Errorf("no element at index -1")
	}

	if _, ok := l.Set(0, 1); ok {
		t.Errorf("no element to set at index 0")
	}
	if _, ok := l.Set(-1, 1); ok {
		t.Errorf("no element to set at index -1")
	}

	if _, ok := l.Remove(0); ok {
		t.Errorf("no element to remove at index 0")
	}
	if _, ok := l.Remove(-1); ok {
		t.Errorf("no element to remove at index -1")
	}

	for i := 0; i < n; i++ {
		if ok := l.Add(i, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if l.Len() != n {
		t.Errorf("want %d, got %d", n, l.Len())
	}

	for i := 0; i < n; i++ {
		v, ok := l.Get(i)
		if !ok {
			t.Errorf("not found: %d", i)
			continue
		}
		l.Set(i, v.(int)*v.(int))
	}

	for i := n - 1; i >= 0; i -= 2 {
		r, ok := l.Remove(i)
		if !ok {
			t.Errorf("cannot remove: %d", i)
			continue
		}
		if r != i*i {
			t.Errorf("want %d, got %v", i*i, r)
		}
	}

	if l.Len() != n/2 {
		t.Errorf("want %d, got %d", n/2, l.Len())
	}

	for i := 0; i < n/2+1; i++ {
		if ok := l.Add(i*2, i); !ok {
			t.Errorf("cannot add: %d", i)
		}
	}
	if l.Len() != n {
		t.Errorf("want %d, got %d", n, l.Len())
	}
}
