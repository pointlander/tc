// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
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

	fmt.Println(string(output))
}
