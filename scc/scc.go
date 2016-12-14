// Copyright (c) 2016 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scc provides a function to find the
// strongly connected components in a directed graph.
package scc

import "fmt"

// Node represents a node
// in a directed graph.
type Node struct {
	succs   []*Node
	value   interface{}
	index   int
	lowlink int
	onstack bool
}

// NewNode creates an initialized node
// with the given value and successors.
func NewNode(value interface{}, succs ...*Node) *Node {
	return &Node{
		value: value,
		succs: succs,
		index: -1,
	}
}

// Value returns the value of the node.
func (n *Node) Value() interface{} { return n.value }

// String returns the value of the node, printed as string.
func (n *Node) String() string { return fmt.Sprintf("%v", n.value) }

// AddSuccs adds successors of the node.
func (n *Node) AddSuccs(succs ...*Node) { n.succs = append(n.succs, succs...) }

// Find finds and returns all the strongly
// connected components using Tarjan's
// algorithm.
func Find(nodes ...*Node) [][]*Node {
	var c collector
	for _, n := range nodes {
		if n.index < 0 {
			c.strongconnect(n)
		}
	}
	return c.sccs
}

type collector struct {
	sccs  [][]*Node
	stack stack
	index int
}

func (c *collector) strongconnect(n *Node) {
	n.index, n.lowlink = c.index, c.index
	c.index++
	c.stack.push(n)
	for _, m := range n.succs {
		if m.index < 0 {
			c.strongconnect(m)
			n.lowlink = min(n.lowlink, m.lowlink)
		} else if m.onstack {
			n.lowlink = min(n.lowlink, m.index)
		}
	}
	if n.lowlink == n.index {
		var scc []*Node
		for {
			m := c.stack.pop()
			scc = append(scc, m)
			if m == n {
				break
			}
		}
		c.sccs = append(c.sccs, scc)
	}
}

// ------- Utilities -------

type stack []*Node

func (s *stack) push(n *Node) {
	n.onstack = true
	*s = append(*s, n)
}

func (s *stack) pop() *Node {
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	n.onstack = false
	return n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
