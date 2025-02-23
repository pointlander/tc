// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	// FlagExperiment1
	FlagExperiment1 = flag.Bool("exp1", false, "experiment 1")
	// FlagExperiment2
	FlagExperiment2 = flag.Bool("exp2", false, "experiment 2")
)

// T is a node in a natural tree
type T struct {
	N int
	T []*T
}

// Parse parses a natural tree
func Parse(i int, input []byte) (int, *T) {
	var t *T
	for i < len(input) {
		switch input[i] {
		case '(':
			var tt *T
			i, tt = Parse(i+1, input)
			t.T = append(t.T, tt)
		case 't':
			if t == nil {
				t = &T{}
			} else {
				t.T = append(t.T, &T{})
			}
			i++
		case ')':
			return i + 1, t
		default:
			i++
		}
	}
	return i, t
}

// Label labels a natural tree
func (t *T) Label(n int) int {
	if n == 1 {
		if len(t.T) == 0 {
			t.N = n
			return n + 1
		}
		n = t.T[0].Label(n)
		t.N = n
		if len(t.T) == 2 {
			n = t.T[1].Label(n + 1)
		}
		return n + 1
	}
	t.N = n
	for _, v := range t.T {
		n = v.Label(n + 1)
	}
	return n
}

// String converts a natural tree to a string
func (t *T) String() string {
	var sb strings.Builder
	var str func(*T)
	str = func(t *T) {
		if len(t.T) == 0 {
			sb.WriteString("t")
			return
		}
		sb.WriteString("(t")
		for _, v := range t.T {
			sb.WriteString(" ")
			str(v)
		}
		sb.WriteString(")")
	}
	if t != nil {
		sb.WriteString("t")
	}
	for _, v := range t.T {
		sb.WriteString(" ")
		str(v)
	}
	return sb.String()
}

func main() {
	flag.Parse()

	if *FlagExperiment1 {
		Experiment1()
		return
	} else if *FlagExperiment2 {
		Experiment2()
		return
	}

	exe, err := filepath.Abs("../tricu/result/bin/tricu")
	if err != nil {
		panic(err)
	}
	f, err := filepath.Abs("../tricu/demos/size.tri")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(exe, "eval", "-f", f)

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	var labels func(int, *T)
	labels = func(depth int, t *T) {
		for i := 0; i < depth; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%d\n", t.N)
		for _, v := range t.T {
			labels(depth+1, v)
		}
	}
	fmt.Println(string(output))
	_, t := Parse(0, output)
	n := t.Label(1)
	fmt.Println("n", n)
	labels(0, t)
}
