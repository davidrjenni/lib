// Copyright (c) 2017 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command wc implements word count as MapReduce job.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"github.com/davidrjenni/lib/mr"
)

type wordCount struct{}

func (w wordCount) Map(key, value string, out chan<- mr.Tuple) {
	s := bufio.NewScanner(strings.NewReader(value))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		out <- mr.Tuple{
			First:  s.Text(),
			Second: "1",
		}
	}
	if err := s.Err(); err != nil {
		log.Println("map error: ", err)
	}
}

func (w wordCount) Reduce(key string, values []string, out chan<- mr.Tuple) {
	c := 0
	for _, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Println("reduce error: ", err)
		} else {
			c += n
		}
	}
	out <- mr.Tuple{
		First:  key,
		Second: strconv.Itoa(c),
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("wc: ")

	var (
		input      = flag.String("input", "", "input text file")
		cpuprofile = flag.String("cpu", "", "cpu profile output")
	)

	flag.Parse()
	if *input == "" {
		flag.Usage()
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal("read input file: ", err)
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("create file: ", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	s := bufio.NewScanner(bytes.NewReader(b))
	s.Split(bufio.ScanLines)
	var values []string
	for s.Scan() {
		values = append(values, s.Text())
	}

	tuples := mr.Run(wordCount{}, values)
	for t := range tuples {
		fmt.Println(t.First, ":", t.Second)
	}
}
