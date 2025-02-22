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
		t := &T{}
		for i < len(input) {
			switch input[i] {
			case '(':
				var tt *T
				i, tt = parse(i+1, input)
				t.T = append(t.T, tt)
			case 't':
				t.T = append(t.T, &T{})
				i++
			case ')':
				return i + 1, t
			default:
				i++
			}
		}
		return i, t
	}
	var prnt func(*T, *strings.Builder)
	prnt = func(t *T, sb *strings.Builder) {
		if len(t.T) == 0 {
			sb.WriteString("t")
			return
		}
		sb.WriteString("(")
		for i, v := range t.T {
			if i > 0 {
				sb.WriteString(" ")
			}
			prnt(v, sb)
		}
		sb.WriteString(")")
	}
	show := func(t *T, sb *strings.Builder) {
		if len(t.T) == 0 {
			sb.WriteString("t ")
			return
		}
		for i, v := range t.T {
			if i > 0 {
				sb.WriteString(" ")
			}
			prnt(v, sb)
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
}
