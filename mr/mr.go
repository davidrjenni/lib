// Copyright (c) 2017 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mr implements MapReduce using Go channels and Go routines.
package mr

import "strconv"

// Tuple holds two strings.
type Tuple struct {
	First, Second string
}

// Job represents a MapReduce job.
type Job interface {
	Map(key, value string, out chan<- Tuple)
	Reduce(key string, values []string, out chan<- Tuple)
}

// Run runs a MapReduce job on the given input.
func Run(j Job, input []string) chan Tuple {
	mappers := make(map[string]chan Tuple)
	for i, v := range input {
		name := strconv.Itoa(i)
		mappers[name] = make(chan Tuple)

		go func(key, value string, out chan<- Tuple) {
			j.Map(key, value, out)
			close(out)
		}(name, v, mappers[name])
	}

	data := make(map[string][]string)
	for m := range fanIn(mappers) {
		k := m.First
		v := m.Second
		values, ok := data[k]
		if !ok {
			values = make([]string, 0)
		}
		values = append(values, v)
		data[k] = values
	}

	reducers := make(map[string]chan Tuple)
	for k, v := range data {
		reducers[k] = make(chan Tuple)
		go func(key string, values []string, out chan<- Tuple) {
			j.Reduce(key, values, out)
			close(out)
		}(k, v, reducers[k])
	}

	return fanIn(reducers)
}

func fanIn(cm map[string]chan Tuple) chan Tuple {
	c := make(chan Tuple, 100)
	go func(c chan Tuple) {
		q := make(chan struct{})
		for _, v := range cm {
			go func(v chan Tuple) {
				for d := range v {
					c <- d
				}
				q <- struct{}{}
			}(v)
		}
		for i := 0; i < len(cm); i++ {
			<-q
		}
		close(c)
	}(c)
	return c
}
