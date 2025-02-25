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
func (t *T) Label() int {
	var left func(*T, int) int
	var right func(*T, int) int
	left = func(t *T, n int) int {
		switch len(t.T) {
		case 0:
			t.N = n
		case 1:
			n = left(t.T[0], n)
			t.N = n
		case 2:
			n = left(t.T[0], n)
			t.N = n
			n = right(t.T[1], n+1)
		}
		return n + 1
	}
	right = func(t *T, n int) int {
		switch len(t.T) {
		case 0:
			t.N = n
		case 1:
			t.N = n
			n = right(t.T[0], n+1)
		case 2:
			n = left(t.T[0], n)
			t.N = n
			n = right(t.T[1], n+1)
		}
		return n
	}
	return left(t, 1)
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
	polygon[size-1] = append(polygon[size-1], t.N)
	var left func([][]int, *T, int, int)
	left = func(polygon [][]int, t *T, a, b int) {
		for _, v := range t.T {
			if v.N < t.N {
				polygon[t.N] = append(polygon[t.N], a)
				left(polygon, v, a, t.N)
			} else {
				polygon[t.N] = append(polygon[t.N], b)
				left(polygon, v, t.N, b)
			}

		}
	}
	var right func([][]int, *T, int, int)
	right = func(polygon [][]int, t *T, a, b int) {
		for _, v := range t.T {
			if v.N > t.N {
				polygon[t.N] = append(polygon[t.N], a)
				right(polygon, v, a, t.N)
			} else {
				polygon[t.N] = append(polygon[t.N], b)
				right(polygon, v, t.N, b)
			}
		}
	}
	if len(t.T) == 2 {
		left(polygon, t.T[0], 0, t.N)
		right(polygon, t.T[1], len(polygon)-1, t.N)
	} else if len(t.T) == 1 {
		left(polygon, t.T[0], 0, t.N)
	}
	return polygon
}

func ITriangulation(polygon [][]int) *T {
	contains := func(list []int, i int) bool {
		for _, value := range list {
			if value == i {
				return true
			}
		}
		return false
	}
	root := 0
	for _, v := range polygon[0] {
		if contains(polygon[len(polygon)-1], v) {
			root = v
			break
		}
	}
	var build func([][]int, int, int) *T
	build = func(polygon [][]int, a, b int) *T {
		t := &T{}
		for i := a + 1; i < b; i++ {
			if contains(polygon[i], a) && contains(polygon[i], b) {
				t.T = append(t.T, build(polygon, a, i))
				t.T = append(t.T, build(polygon, i, b))
			}
		}
		if len(t.T) == 0 {
			for i := b - 1; i > a; i-- {
				if contains(polygon[i], a) {
					t.T = append(t.T, build(polygon, a, i))
				}
				if len(polygon[i]) > 0 {
					break
				}
			}
		}
		if len(t.T) == 0 {
			for i := a + 1; i < b; i++ {
				if contains(polygon[i], b) {
					t.T = append(t.T, build(polygon, i, b))
				}
				if len(polygon[i]) > 0 {
					break
				}
			}
		}
		return t
	}
	if root != 1 && root != len(polygon)-2 {
		return &T{
			T: []*T{
				build(polygon, 0, root),
				build(polygon, root, len(polygon)-1),
			},
		}
	} else if root != 1 {
		return &T{
			T: []*T{
				build(polygon, 0, root),
			},
		}
	} else {
		return &T{}
	}
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
