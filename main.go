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
func Parse(input []byte, ii ...int) (int, *T) {
	i := 0
	if len(ii) == 1 {
		i = ii[0]
	}
	var t *T
	for i < len(input) {
		switch input[i] {
		case '(':
			var tt *T
			i, tt = Parse(input, i+1)
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
func (t *T) Label(nn ...int) int {
	n := 1
	if len(nn) == 1 {
		n = nn[0]
	}
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

// Triangulation performs triangulation based on binary tree
func (t *T) Triangulation(size int) [][]int {
	size++
	polygon := make([][]int, size)
	polygon[0] = append(polygon[0], t.N)
	polygon[t.N] = append(polygon[t.N], 0)
	polygon[size-1] = append(polygon[size-1], t.N)
	polygon[t.N] = append(polygon[t.N], size-1)
	var tri func([][]int, *T, int)
	tri = func(polygon [][]int, t *T, n int) {
		for _, v := range t.T {
			if t.N < n {
				if v.N < t.N && t.N > 1 {
					polygon[t.N] = append(polygon[t.N], 0)
					polygon[0] = append(polygon[0], t.N)
				} else if n-t.N > 1 {
					polygon[t.N] = append(polygon[t.N], n)
					polygon[n] = append(polygon[n], t.N)
				}
			} else {
				if v.N > t.N && v.N < len(polygon)-1 {
					polygon[t.N] = append(polygon[t.N], len(polygon)-1)
					polygon[len(polygon)-1] = append(polygon[len(polygon)-1], t.N)
				} else if n-t.N > 1 {
					polygon[t.N] = append(polygon[t.N], n)
					polygon[n] = append(polygon[n], t.N)
				}
			}
			tri(polygon, v, n)
		}
	}
	for _, v := range t.T {
		tri(polygon, v, t.N)
	}
	return polygon
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
	_, t := Parse(output)
	n := t.Label()
	fmt.Println("n", n, t.N)
	labels(0, t)
	polygon := t.Triangulation(n)
	for i, v := range polygon {
		fmt.Println(i, v)
	}
}
