// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scc_test

import (
	"fmt"

	"github.com/davidrjenni/lib/scc"
)

// ExampleFind shows how to find strongly connected components
// in a directed graph using the example from Wikipedia
// (https://en.wikipedia.org/wiki/Strongly_connected_component).
func ExampleFind() {
	a := scc.NewNode("a")
	b := scc.NewNode("b")
	c := scc.NewNode("c")
	d := scc.NewNode("d")
	e := scc.NewNode("e")
	f := scc.NewNode("f")
	g := scc.NewNode("g")
	h := scc.NewNode("h")

	a.AddSuccs(b)
	b.AddSuccs(c)
	c.AddSuccs(a)
	d.AddSuccs(b, c, e)
	e.AddSuccs(d, f)
	f.AddSuccs(c, g)
	g.AddSuccs(f)
	h.AddSuccs(e, g, h)

	for _, s := range scc.Find(a, b, c, d, e, f, g, h) {
		fmt.Println(s)
	}

	// Output:
	// [c b a]
	// [g f]
	// [e d]
	// [h]
}
