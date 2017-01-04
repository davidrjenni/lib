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
