// Copyright 2025 The tc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := []byte("t t (t (t t) (t (t t) (t t (t t (t t (t (t t) (t (t t) (t (t t) t))))))))")
	_, tt := Parse(0, input)
	if tt.String() != string(input) {
		t.Fatal("parsing failed")
	}
}

func TestLabel(t *testing.T) {
	tt := &T{
		N: 5,
		T: []*T{
			&T{
				N: 2,
				T: []*T{
					&T{N: 1},
					&T{
						N: 3,
						T: []*T{
							&T{N: 4},
						},
					},
				},
			},
			&T{
				N: 6,
				T: []*T{
					&T{N: 7},
				},
			},
		},
	}
	ttt := &T{
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
	n := ttt.Label(1)
	if n != 8 {
		t.Fatal("wrong number of nodes")
	}
	var compare func(a, b *T)
	compare = func(a, b *T) {
		if a.N != b.N {
			t.Fatal("node is mislabeled")
		}
		for i, v := range a.T {
			compare(v, b.T[i])
		}
	}
	compare(tt, ttt)
}
