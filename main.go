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

	type T struct {
		N int
		T []*T
	}
	var parse func(int, []byte) (int, *T)
	parse = func(i int, input []byte) (int, *T) {
		var t *T
		for i < len(input) {
			switch input[i] {
			case '(':
				var tt *T
				i, tt = parse(i+1, input)
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
	var label func(int, *T) int
	label = func(n int, t *T) int {
		if n == 1 {
			if len(t.T) == 0 {
				t.N = n
				return n + 1
			}
			n = label(n, t.T[0])
			t.N = n
			if len(t.T) == 2 {
				n = label(n+1, t.T[1])
			}
			return n + 1
		}
		t.N = n
		for _, v := range t.T {
			n = label(n+1, v)
		}
		return n
	}
	var prnt func(*T, *strings.Builder)
	prnt = func(t *T, sb *strings.Builder) {
		if len(t.T) == 0 {
			sb.WriteString("t")
			return
		}
		sb.WriteString("(t")
		for _, v := range t.T {
			sb.WriteString(" ")
			prnt(v, sb)
		}
		sb.WriteString(")")
	}
	show := func(t *T, sb *strings.Builder) {
		if t != nil {
			sb.WriteString("t")
		}
		for _, v := range t.T {
			sb.WriteString(" ")
			prnt(v, sb)
		}
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
	_, t := parse(0, output)
	var sb strings.Builder
	show(t, &sb)
	fmt.Println(sb.String())
	if sb.String() != string(output) {
		panic("incorrect parsing")
	}
	n := label(1, t)
	fmt.Println("n", n)
	labels(0, t)

	tt := &T{
		T: []*T{
			&T{
				T: []*T{
					&T{},
					&T{
						T: []*T{
							&T{},
						},
					},
				},
			},
			&T{
				T: []*T{
					&T{},
				},
			},
		},
	}
	n = label(1, tt)
	fmt.Println("n", n)
	labels(0, tt)
}
